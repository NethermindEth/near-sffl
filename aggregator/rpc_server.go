package aggregator

import (
	"context"
	"errors"
	"net/http"
	"net/rpc"
	"time"

	"github.com/NethermindEth/near-sffl/aggregator/types"
	servicemanager "github.com/NethermindEth/near-sffl/contracts/bindings/SFFLServiceManager"
	taskmanager "github.com/NethermindEth/near-sffl/contracts/bindings/SFFLTaskManager"
	"github.com/NethermindEth/near-sffl/core"

	"github.com/Layr-Labs/eigensdk-go/crypto/bls"
	sdktypes "github.com/Layr-Labs/eigensdk-go/types"
)

var (
	TaskNotFoundError400                     = errors.New("400. Task not found")
	OperatorNotPartOfTaskQuorum400           = errors.New("400. Operator not part of quorum")
	TaskResponseDigestNotFoundError500       = errors.New("500. Failed to get task response digest")
	MessageDigestNotFoundError500            = errors.New("500. Failed to get message digest")
	UnknownErrorWhileVerifyingSignature400   = errors.New("400. Failed to verify signature")
	SignatureVerificationFailed400           = errors.New("400. Signature verification failed")
	CallToGetCheckSignaturesIndicesFailed500 = errors.New("500. Failed to get check signatures indices")
)

func (agg *Aggregator) startServer(ctx context.Context) error {

	err := rpc.Register(agg)
	if err != nil {
		agg.logger.Fatal("Format of service TaskManager isn't correct. ", "err", err)
	}
	rpc.HandleHTTP()
	err = http.ListenAndServe(agg.serverIpPortAddr, nil)
	if err != nil {
		agg.logger.Fatal("ListenAndServe", "err", err)
	}

	return nil
}

type SignedCheckpointTaskResponse struct {
	TaskResponse taskmanager.CheckpointTaskResponse
	BlsSignature bls.Signature
	OperatorId   bls.OperatorId
}

// rpc endpoint which is called by operator
// reply doesn't need to be checked. If there are no errors, the task response is accepted
// rpc framework forces a reply type to exist, so we put bool as a placeholder
func (agg *Aggregator) ProcessSignedCheckpointTaskResponse(signedCheckpointTaskResponse *SignedCheckpointTaskResponse, reply *bool) error {
	agg.logger.Infof("Received signed task response: %#v", signedCheckpointTaskResponse)
	taskIndex := signedCheckpointTaskResponse.TaskResponse.ReferenceTaskIndex
	taskResponseDigest, err := core.GetCheckpointTaskResponseDigest(&signedCheckpointTaskResponse.TaskResponse)
	if err != nil {
		agg.logger.Error("Failed to get task response digest", "err", err)
		return TaskResponseDigestNotFoundError500
	}
	agg.taskResponsesMu.Lock()
	if _, ok := agg.taskResponses[taskIndex]; !ok {
		agg.taskResponses[taskIndex] = make(map[sdktypes.TaskResponseDigest]taskmanager.CheckpointTaskResponse)
	}
	if _, ok := agg.taskResponses[taskIndex][taskResponseDigest]; !ok {
		agg.taskResponses[taskIndex][taskResponseDigest] = signedCheckpointTaskResponse.TaskResponse
	}
	agg.taskResponsesMu.Unlock()

	err = agg.taskBlsAggregationService.ProcessNewSignature(
		context.Background(), taskIndex, taskResponseDigest,
		&signedCheckpointTaskResponse.BlsSignature, signedCheckpointTaskResponse.OperatorId,
	)
	return err
}

type SignedStateRootUpdateMessage struct {
	Message      servicemanager.StateRootUpdateMessage
	BlsSignature bls.Signature
	OperatorId   bls.OperatorId
}

func (agg *Aggregator) ProcessSignedStateRootUpdateMessage(signedStateRootUpdateMessage *SignedStateRootUpdateMessage, reply *bool) error {
	agg.logger.Infof("Received signed state root update message: %#v", signedStateRootUpdateMessage)
	messageDigest, err := core.GetStateRootUpdateMessageDigest(&signedStateRootUpdateMessage.Message)
	if err != nil {
		agg.logger.Error("Failed to get message digest", "err", err)
		return TaskResponseDigestNotFoundError500
	}

	agg.stateRootUpdatesMu.Lock()
	agg.stateRootUpdates[messageDigest] = signedStateRootUpdateMessage.Message
	agg.stateRootUpdatesMu.Unlock()

	agg.messageBlsAggregationService.InitializeNewMessage(messageDigest, types.QUORUM_NUMBERS, []uint32{types.QUORUM_THRESHOLD_NUMERATOR}, 1*time.Hour)

	err = agg.messageBlsAggregationService.ProcessNewSignature(
		context.Background(), messageDigest,
		&signedStateRootUpdateMessage.BlsSignature, signedStateRootUpdateMessage.OperatorId,
	)
	return err
}
