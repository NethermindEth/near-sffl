package messages

import (
	"github.com/Layr-Labs/eigensdk-go/crypto/bls"
	"github.com/ethereum/go-ethereum/accounts/abi"

	taskmanager "github.com/NethermindEth/near-sffl/contracts/bindings/SFFLTaskManager"
	"github.com/NethermindEth/near-sffl/core"
	coretypes "github.com/NethermindEth/near-sffl/core/types"
)

type CheckpointTaskResponse struct {
	ReferenceTaskIndex     coretypes.TaskIndex
	StateRootUpdatesRoot   [32]byte
	OperatorSetUpdatesRoot [32]byte
}

type SignedCheckpointTaskResponse struct {
	TaskResponse CheckpointTaskResponse
	BlsSignature bls.Signature
	OperatorId   bls.OperatorId
}

type CheckpointMessages struct {
	StateRootUpdateMessages              []StateRootUpdateMessage
	StateRootUpdateMessageAggregations   []MessageBlsAggregation
	OperatorSetUpdateMessages            []OperatorSetUpdateMessage
	OperatorSetUpdateMessageAggregations []MessageBlsAggregation
}

func NewCheckpointTaskResponseFromBinding(binding taskmanager.CheckpointTaskResponse) CheckpointTaskResponse {
	return CheckpointTaskResponse(binding)
}

func (msg CheckpointTaskResponse) ToBinding() taskmanager.CheckpointTaskResponse {
	return taskmanager.CheckpointTaskResponse(msg)
}

func (msg CheckpointTaskResponse) AbiEncode() ([]byte, error) {
	typ, err := abi.NewType("tuple", "", []abi.ArgumentMarshaling{
		{Name: "referenceTaskIndex", Type: "uint32"},
		{Name: "stateRootUpdatesRoot", Type: "bytes32"},
		{Name: "operatorSetUpdatesRoot", Type: "bytes32"},
	})
	if err != nil {
		return nil, err
	}

	arguments := abi.Arguments{{Type: typ}}

	bytes, err := arguments.Pack(msg.ToBinding())
	if err != nil {
		return nil, err
	}

	return bytes, nil
}

func (msg CheckpointTaskResponse) Digest() ([32]byte, error) {
	data, err := msg.AbiEncode()
	if err != nil {
		return [32]byte{}, err
	}

	digest, err := core.Keccak256(data)
	if err != nil {
		return [32]byte{}, err
	}

	return digest, nil
}
