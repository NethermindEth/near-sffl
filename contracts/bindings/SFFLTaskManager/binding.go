// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contractSFFLTaskManager

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// BN254G1Point is an auto generated low-level Go binding around an user-defined struct.
type BN254G1Point struct {
	X *big.Int
	Y *big.Int
}

// BN254G2Point is an auto generated low-level Go binding around an user-defined struct.
type BN254G2Point struct {
	X [2]*big.Int
	Y [2]*big.Int
}

// CheckpointTask is an auto generated low-level Go binding around an user-defined struct.
type CheckpointTask struct {
	TaskCreatedBlock uint32
	FromTimestamp    uint64
	ToTimestamp      uint64
	QuorumThreshold  uint32
	QuorumNumbers    []byte
}

// CheckpointTaskResponse is an auto generated low-level Go binding around an user-defined struct.
type CheckpointTaskResponse struct {
	ReferenceTaskIndex     uint32
	StateRootUpdatesRoot   [32]byte
	OperatorSetUpdatesRoot [32]byte
}

// CheckpointTaskResponseMetadata is an auto generated low-level Go binding around an user-defined struct.
type CheckpointTaskResponseMetadata struct {
	TaskRespondedBlock uint32
	HashOfNonSigners   [32]byte
}

// IBLSSignatureCheckerNonSignerStakesAndSignature is an auto generated low-level Go binding around an user-defined struct.
type IBLSSignatureCheckerNonSignerStakesAndSignature struct {
	NonSignerQuorumBitmapIndices []uint32
	NonSignerPubkeys             []BN254G1Point
	QuorumApks                   []BN254G1Point
	ApkG2                        BN254G2Point
	Sigma                        BN254G1Point
	QuorumApkIndices             []uint32
	TotalStakeIndices            []uint32
	NonSignerStakeIndices        [][]uint32
}

// IBLSSignatureCheckerQuorumStakeTotals is an auto generated low-level Go binding around an user-defined struct.
type IBLSSignatureCheckerQuorumStakeTotals struct {
	SignedStakeForQuorum []*big.Int
	TotalStakeForQuorum  []*big.Int
}

// OperatorSetUpdateMessage is an auto generated low-level Go binding around an user-defined struct.
type OperatorSetUpdateMessage struct {
	Id        uint64
	Timestamp uint64
	Operators []RollupOperatorsOperator
}

// OperatorStateRetrieverCheckSignaturesIndices is an auto generated low-level Go binding around an user-defined struct.
type OperatorStateRetrieverCheckSignaturesIndices struct {
	NonSignerQuorumBitmapIndices []uint32
	QuorumApkIndices             []uint32
	TotalStakeIndices            []uint32
	NonSignerStakeIndices        [][]uint32
}

// OperatorStateRetrieverOperator is an auto generated low-level Go binding around an user-defined struct.
type OperatorStateRetrieverOperator struct {
	Operator   common.Address
	OperatorId [32]byte
	Stake      *big.Int
}

// RollupOperatorsOperator is an auto generated low-level Go binding around an user-defined struct.
type RollupOperatorsOperator struct {
	Pubkey BN254G1Point
	Weight *big.Int
}

// SparseMerkleTreeProof is an auto generated low-level Go binding around an user-defined struct.
type SparseMerkleTreeProof struct {
	Leaf      [32]byte
	Index     *big.Int
	Bitmask   *big.Int
	SideNodes [][32]byte
}

// StateRootUpdateMessage is an auto generated low-level Go binding around an user-defined struct.
type StateRootUpdateMessage struct {
	RollupId    uint32
	BlockHeight uint64
	Timestamp   uint64
	StateRoot   [32]byte
}

// ContractSFFLTaskManagerMetaData contains all meta data concerning the ContractSFFLTaskManager contract.
var ContractSFFLTaskManagerMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"taskResponseWindowBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"TASK_CHALLENGE_WINDOW_BLOCK\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"TASK_RESPONSE_WINDOW_BLOCK\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"THRESHOLD_DENOMINATOR\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"aggregator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"allCheckpointTaskHashes\",\"inputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"allCheckpointTaskResponses\",\"inputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"blsApkRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIBLSApkRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"checkQuorum\",\"inputs\":[{\"name\":\"messageHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"referenceBlockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"nonSignerStakesAndSignature\",\"type\":\"tuple\",\"internalType\":\"structIBLSSignatureChecker.NonSignerStakesAndSignature\",\"components\":[{\"name\":\"nonSignerQuorumBitmapIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerPubkeys\",\"type\":\"tuple[]\",\"internalType\":\"structBN254.G1Point[]\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"quorumApks\",\"type\":\"tuple[]\",\"internalType\":\"structBN254.G1Point[]\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"apkG2\",\"type\":\"tuple\",\"internalType\":\"structBN254.G2Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"Y\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}]},{\"name\":\"sigma\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"quorumApkIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"totalStakeIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerStakeIndices\",\"type\":\"uint32[][]\",\"internalType\":\"uint32[][]\"}]},{\"name\":\"quorumThreshold\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"checkSignatures\",\"inputs\":[{\"name\":\"msgHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"referenceBlockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"params\",\"type\":\"tuple\",\"internalType\":\"structIBLSSignatureChecker.NonSignerStakesAndSignature\",\"components\":[{\"name\":\"nonSignerQuorumBitmapIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerPubkeys\",\"type\":\"tuple[]\",\"internalType\":\"structBN254.G1Point[]\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"quorumApks\",\"type\":\"tuple[]\",\"internalType\":\"structBN254.G1Point[]\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"apkG2\",\"type\":\"tuple\",\"internalType\":\"structBN254.G2Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"Y\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}]},{\"name\":\"sigma\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"quorumApkIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"totalStakeIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerStakeIndices\",\"type\":\"uint32[][]\",\"internalType\":\"uint32[][]\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIBLSSignatureChecker.QuorumStakeTotals\",\"components\":[{\"name\":\"signedStakeForQuorum\",\"type\":\"uint96[]\",\"internalType\":\"uint96[]\"},{\"name\":\"totalStakeForQuorum\",\"type\":\"uint96[]\",\"internalType\":\"uint96[]\"}]},{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"checkpointTaskNumber\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"checkpointTaskSuccesfullyChallenged\",\"inputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"createCheckpointTask\",\"inputs\":[{\"name\":\"fromTimestamp\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"toTimestamp\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"quorumThreshold\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"delegation\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIDelegationManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"generator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCheckSignaturesIndices\",\"inputs\":[{\"name\":\"registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"referenceBlockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"nonSignerOperatorIds\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structOperatorStateRetriever.CheckSignaturesIndices\",\"components\":[{\"name\":\"nonSignerQuorumBitmapIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"quorumApkIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"totalStakeIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerStakeIndices\",\"type\":\"uint32[][]\",\"internalType\":\"uint32[][]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorState\",\"inputs\":[{\"name\":\"registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"blockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[][]\",\"internalType\":\"structOperatorStateRetriever.Operator[][]\",\"components\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"stake\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorState\",\"inputs\":[{\"name\":\"registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"operatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"blockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"tuple[][]\",\"internalType\":\"structOperatorStateRetriever.Operator[][]\",\"components\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"stake\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"_pauserRegistry\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"},{\"name\":\"initialOwner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_aggregator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_generator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"nextCheckpointTaskNum\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"pauseAll\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[{\"name\":\"index\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pauserRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"raiseAndResolveCheckpointChallenge\",\"inputs\":[{\"name\":\"task\",\"type\":\"tuple\",\"internalType\":\"structCheckpoint.Task\",\"components\":[{\"name\":\"taskCreatedBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"fromTimestamp\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"toTimestamp\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"quorumThreshold\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"taskResponse\",\"type\":\"tuple\",\"internalType\":\"structCheckpoint.TaskResponse\",\"components\":[{\"name\":\"referenceTaskIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"stateRootUpdatesRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"operatorSetUpdatesRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"taskResponseMetadata\",\"type\":\"tuple\",\"internalType\":\"structCheckpoint.TaskResponseMetadata\",\"components\":[{\"name\":\"taskRespondedBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"hashOfNonSigners\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"pubkeysOfNonSigningOperators\",\"type\":\"tuple[]\",\"internalType\":\"structBN254.G1Point[]\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registryCoordinator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"respondToCheckpointTask\",\"inputs\":[{\"name\":\"task\",\"type\":\"tuple\",\"internalType\":\"structCheckpoint.Task\",\"components\":[{\"name\":\"taskCreatedBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"fromTimestamp\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"toTimestamp\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"quorumThreshold\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"taskResponse\",\"type\":\"tuple\",\"internalType\":\"structCheckpoint.TaskResponse\",\"components\":[{\"name\":\"referenceTaskIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"stateRootUpdatesRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"operatorSetUpdatesRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"nonSignerStakesAndSignature\",\"type\":\"tuple\",\"internalType\":\"structIBLSSignatureChecker.NonSignerStakesAndSignature\",\"components\":[{\"name\":\"nonSignerQuorumBitmapIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerPubkeys\",\"type\":\"tuple[]\",\"internalType\":\"structBN254.G1Point[]\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"quorumApks\",\"type\":\"tuple[]\",\"internalType\":\"structBN254.G1Point[]\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"apkG2\",\"type\":\"tuple\",\"internalType\":\"structBN254.G2Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"Y\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}]},{\"name\":\"sigma\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"quorumApkIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"totalStakeIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerStakeIndices\",\"type\":\"uint32[][]\",\"internalType\":\"uint32[][]\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setPauserRegistry\",\"inputs\":[{\"name\":\"newPauserRegistry\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setStaleStakesForbidden\",\"inputs\":[{\"name\":\"value\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"stakeRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIStakeRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"staleStakesForbidden\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"trySignatureAndApkVerification\",\"inputs\":[{\"name\":\"msgHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"apk\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"apkG2\",\"type\":\"tuple\",\"internalType\":\"structBN254.G2Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"Y\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}]},{\"name\":\"sigma\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[{\"name\":\"pairingSuccessful\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"siganatureIsValid\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"unpause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"verifyMessageInclusionState\",\"inputs\":[{\"name\":\"message\",\"type\":\"tuple\",\"internalType\":\"structStateRootUpdate.Message\",\"components\":[{\"name\":\"rollupId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"blockHeight\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"timestamp\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"stateRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"taskResponse\",\"type\":\"tuple\",\"internalType\":\"structCheckpoint.TaskResponse\",\"components\":[{\"name\":\"referenceTaskIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"stateRootUpdatesRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"operatorSetUpdatesRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"proof\",\"type\":\"tuple\",\"internalType\":\"structSparseMerkleTree.Proof\",\"components\":[{\"name\":\"leaf\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"index\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"bitmask\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"sideNodes\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"verifyMessageInclusionState\",\"inputs\":[{\"name\":\"message\",\"type\":\"tuple\",\"internalType\":\"structOperatorSetUpdate.Message\",\"components\":[{\"name\":\"id\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"timestamp\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"operators\",\"type\":\"tuple[]\",\"internalType\":\"structRollupOperators.Operator[]\",\"components\":[{\"name\":\"pubkey\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"weight\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]},{\"name\":\"taskResponse\",\"type\":\"tuple\",\"internalType\":\"structCheckpoint.TaskResponse\",\"components\":[{\"name\":\"referenceTaskIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"stateRootUpdatesRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"operatorSetUpdatesRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"proof\",\"type\":\"tuple\",\"internalType\":\"structSparseMerkleTree.Proof\",\"components\":[{\"name\":\"leaf\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"index\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"bitmask\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"sideNodes\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"pure\"},{\"type\":\"event\",\"name\":\"CheckpointTaskChallengedSuccessfully\",\"inputs\":[{\"name\":\"taskIndex\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"},{\"name\":\"challenger\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CheckpointTaskChallengedUnsuccessfully\",\"inputs\":[{\"name\":\"taskIndex\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"},{\"name\":\"challenger\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CheckpointTaskCreated\",\"inputs\":[{\"name\":\"taskIndex\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"},{\"name\":\"task\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structCheckpoint.Task\",\"components\":[{\"name\":\"taskCreatedBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"fromTimestamp\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"toTimestamp\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"quorumThreshold\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CheckpointTaskResponded\",\"inputs\":[{\"name\":\"taskResponse\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structCheckpoint.TaskResponse\",\"components\":[{\"name\":\"referenceTaskIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"stateRootUpdatesRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"operatorSetUpdatesRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"taskResponseMetadata\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structCheckpoint.TaskResponseMetadata\",\"components\":[{\"name\":\"taskRespondedBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"hashOfNonSigners\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PauserRegistrySet\",\"inputs\":[{\"name\":\"pauserRegistry\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIPauserRegistry\"},{\"name\":\"newPauserRegistry\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIPauserRegistry\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StaleStakesForbiddenUpdate\",\"inputs\":[{\"name\":\"value\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false}]",
	Bin: "0x6101206040523480156200001257600080fd5b5060405162006169380380620061698339810160408190526200003591620001f7565b81806001600160a01b03166080816001600160a01b031681525050806001600160a01b031663683048356040518163ffffffff1660e01b8152600401602060405180830381865afa1580156200008f573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190620000b591906200023e565b6001600160a01b031660a0816001600160a01b031681525050806001600160a01b0316635df459466040518163ffffffff1660e01b8152600401602060405180830381865afa1580156200010d573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906200013391906200023e565b6001600160a01b031660c0816001600160a01b03168152505060a0516001600160a01b031663df5cf7236040518163ffffffff1660e01b8152600401602060405180830381865afa1580156200018d573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190620001b391906200023e565b6001600160a01b031660e052506097805460ff1916600117905563ffffffff16610100525062000265565b6001600160a01b0381168114620001f457600080fd5b50565b600080604083850312156200020b57600080fd5b82516200021881620001de565b602084015190925063ffffffff811681146200023357600080fd5b809150509250929050565b6000602082840312156200025157600080fd5b81516200025e81620001de565b9392505050565b60805160a05160c05160e05161010051615e80620002e9600039600081816102930152612dc401526000818161059801526121ac01526000818161041001526123960152600081816104370152818161256c015261272e01526000818161045e01528181610f0d01528181611ea90152818161202201526122500152615e806000f3fe608060405234801561001057600080fd5b50600436106102325760003560e01c8063715018a611610130578063cefdc1d4116100b8578063f294e43c1161007c578063f294e43c146105d5578063f2fde38b146105e8578063f63c5bab146105ba578063f8c8765e146105fb578063fabc1cbc1461060e57600080fd5b8063cefdc1d41461055f578063da16491f14610580578063df5cf72314610593578063ef024458146105ba578063efcf4edb146105c257600080fd5b80638da5cb5b116100ff5780638da5cb5b146104eb57806395eebee6146104fc578063a168e3c01461051f578063b98d09081461053f578063bdbf596c1461054c57600080fd5b8063715018a6146104a15780637afa1eed146104a9578063886f1195146104c55780638cbc379a146104d857600080fd5b80634f19ade7116101be5780635c975abb116101825780635c975abb146104035780635df459461461040b57806368304835146104325780636d14a987146104595780636efb46361461048057600080fd5b80634f19ade7146103675780634f739f7414610395578063595c6a67146103b55780635ac86ab7146103bd5780635ace2df7146103f057600080fd5b8063245a7bfc11610205578063245a7bfc146102ca578063292f7a4e146102f55780632e44b3491461031f5780633563b0d114610334578063416c7e5e1461035457600080fd5b806310d67a2f14610237578063136439dd1461024c578063171f1d5b1461025f5780631ad431891461028e575b600080fd5b61024a61024536600461455a565b610621565b005b61024a61025a366004614577565b6106dd565b61027261026d3660046146f5565b61081c565b6040805192151583529015156020830152015b60405180910390f35b6102b57f000000000000000000000000000000000000000000000000000000000000000081565b60405163ffffffff9091168152602001610285565b6098546102dd906001600160a01b031681565b6040516001600160a01b039091168152602001610285565b610308610303366004614a4f565b6109a6565b604080519215158352602083019190915201610285565b6097546102b590610100900463ffffffff1681565b610347610342366004614ae9565b610a73565b6040516102859190614c44565b61024a610362366004614c65565b610f0b565b610387610375366004614c82565b60996020526000908152604090205481565b604051908152602001610285565b6103a86103a3366004614c9f565b611080565b6040516102859190614da3565b61024a611704565b6103e06103cb366004614e6d565b606654600160ff9092169190911b9081161490565b6040519015158152602001610285565b61024a6103fe366004614eb4565b6117cb565b606654610387565b6102dd7f000000000000000000000000000000000000000000000000000000000000000081565b6102dd7f000000000000000000000000000000000000000000000000000000000000000081565b6102dd7f000000000000000000000000000000000000000000000000000000000000000081565b61049361048e366004614f45565b611af5565b604051610285929190615005565b61024a6129e3565b6097546102dd906501000000000090046001600160a01b031681565b6065546102dd906001600160a01b031681565b609754610100900463ffffffff166102b5565b6033546001600160a01b03166102dd565b6103e061050a366004614c82565b609b6020526000908152604090205460ff1681565b61038761052d366004614c82565b609a6020526000908152604090205481565b6097546103e09060ff1681565b6103e061055a366004615060565b6129f7565b61057261056d3660046150c0565b612ab2565b604051610285929190615102565b61024a61058e366004615123565b612c44565b6102dd7f000000000000000000000000000000000000000000000000000000000000000081565b6102b5606481565b61024a6105d03660046151a4565b612f45565b6103e06105e336600461521b565b613145565b61024a6105f636600461455a565b6131f3565b61024a610609366004615285565b613269565b61024a61061c366004614577565b6133d1565b606560009054906101000a90046001600160a01b03166001600160a01b031663eab66d7a6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610674573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061069891906152e1565b6001600160a01b0316336001600160a01b0316146106d15760405162461bcd60e51b81526004016106c8906152fe565b60405180910390fd5b6106da8161352d565b50565b60655460405163237dfb4760e11b81523360048201526001600160a01b03909116906346fbf68e90602401602060405180830381865afa158015610725573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906107499190615348565b6107655760405162461bcd60e51b81526004016106c890615365565b606654818116146107de5760405162461bcd60e51b815260206004820152603860248201527f5061757361626c652e70617573653a20696e76616c696420617474656d70742060448201527f746f20756e70617573652066756e6374696f6e616c697479000000000000000060648201526084016106c8565b606681905560405181815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d906020015b60405180910390a250565b60008060007f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f000000187876000015188602001518860000151600060028110610864576108646153ad565b60200201518951600160200201518a60200151600060028110610889576108896153ad565b60200201518b602001516001600281106108a5576108a56153ad565b602090810291909101518c518d8301516040516109029a99989796959401988952602089019790975260408801959095526060870193909352608086019190915260a085015260c084015260e08301526101008201526101200190565b6040516020818303038152906040528051906020012060001c61092591906153c3565b905061099861093e6109378884613624565b86906136bb565b61094661374f565b61098e61097f85610979604080518082018252600080825260209182015281518083019092526001825260029082015290565b90613624565b6109888c61380f565b906136bb565b886201d4c061389f565b909890975095505050505050565b6000806000806109b98a8a8a8a8a611af5565b9150915060005b88811015610a5f578563ffffffff16836020015182815181106109e5576109e56153ad565b60200260200101516109f791906153fb565b6001600160601b0316606463ffffffff1684600001518381518110610a1e57610a1e6153ad565b6020026020010151610a3091906153fb565b6001600160601b03161015610a4d5750600093509150610a689050565b80610a578161542a565b9150506109c0565b50600193509150505b965096945050505050565b60606000846001600160a01b031663683048356040518163ffffffff1660e01b8152600401602060405180830381865afa158015610ab5573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610ad991906152e1565b90506000856001600160a01b0316639e9923c26040518163ffffffff1660e01b8152600401602060405180830381865afa158015610b1b573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610b3f91906152e1565b90506000866001600160a01b0316635df459466040518163ffffffff1660e01b8152600401602060405180830381865afa158015610b81573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610ba591906152e1565b9050600086516001600160401b03811115610bc257610bc2614590565b604051908082528060200260200182016040528015610bf557816020015b6060815260200190600190039081610be05790505b50905060005b8751811015610efd576000888281518110610c1857610c186153ad565b0160200151604051638902624560e01b815260f89190911c6004820181905263ffffffff8a16602483015291506000906001600160a01b03871690638902624590604401600060405180830381865afa158015610c79573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052610ca19190810190615445565b905080516001600160401b03811115610cbc57610cbc614590565b604051908082528060200260200182016040528015610d0757816020015b6040805160608101825260008082526020808301829052928201528252600019909201910181610cda5790505b50848481518110610d1a57610d1a6153ad565b602002602001018190525060005b8151811015610ee7576040518060600160405280876001600160a01b03166347b314e8858581518110610d5d57610d5d6153ad565b60200260200101516040518263ffffffff1660e01b8152600401610d8391815260200190565b602060405180830381865afa158015610da0573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610dc491906152e1565b6001600160a01b03168152602001838381518110610de457610de46153ad565b60200260200101518152602001896001600160a01b031663fa28c627858581518110610e1257610e126153ad565b60209081029190910101516040516001600160e01b031960e084901b168152600481019190915260ff8816602482015263ffffffff8f166044820152606401602060405180830381865afa158015610e6e573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610e9291906154d5565b6001600160601b0316815250858581518110610eb057610eb06153ad565b60200260200101518281518110610ec957610ec96153ad565b60200260200101819052508080610edf9061542a565b915050610d28565b5050508080610ef59061542a565b915050610bfb565b5093505050505b9392505050565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316638da5cb5b6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610f69573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610f8d91906152e1565b6001600160a01b0316336001600160a01b0316146110395760405162461bcd60e51b815260206004820152605c60248201527f424c535369676e6174757265436865636b65722e6f6e6c79436f6f7264696e6160448201527f746f724f776e65723a2063616c6c6572206973206e6f7420746865206f776e6560648201527f72206f6620746865207265676973747279436f6f7264696e61746f7200000000608482015260a4016106c8565b6097805460ff19168215159081179091556040519081527f40e4ed880a29e0f6ddce307457fb75cddf4feef7d3ecb0301bfdf4976a0e2dfc9060200160405180910390a150565b6110ab6040518060800160405280606081526020016060815260200160608152602001606081525090565b6000876001600160a01b031663683048356040518163ffffffff1660e01b8152600401602060405180830381865afa1580156110eb573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061110f91906152e1565b905061113c6040518060800160405280606081526020016060815260200160608152602001606081525090565b6040516361c8a12f60e11b81526001600160a01b038a169063c391425e9061116c908b90899089906004016154fe565b600060405180830381865afa158015611189573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526111b19190810190615548565b81526040516340e03a8160e11b81526001600160a01b038316906381c07502906111e3908b908b908b906004016155ff565b600060405180830381865afa158015611200573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526112289190810190615548565b6040820152856001600160401b0381111561124557611245614590565b60405190808252806020026020018201604052801561127857816020015b60608152602001906001900390816112635790505b50606082015260005b60ff8116871115611615576000856001600160401b038111156112a6576112a6614590565b6040519080825280602002602001820160405280156112cf578160200160208202803683370190505b5083606001518360ff16815181106112e9576112e96153ad565b602002602001018190525060005b868110156115155760008c6001600160a01b03166304ec63518a8a85818110611322576113226153ad565b905060200201358e88600001518681518110611340576113406153ad565b60200260200101516040518463ffffffff1660e01b815260040161137d9392919092835263ffffffff918216602084015216604082015260600190565b602060405180830381865afa15801561139a573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906113be9190615628565b90508a8a8560ff168181106113d5576113d56153ad565b6001600160c01b03841692013560f81c9190911c60019081161415905061150257856001600160a01b031663dd9846b98a8a85818110611417576114176153ad565b905060200201358d8d8860ff16818110611433576114336153ad565b6040516001600160e01b031960e087901b1681526004810194909452919091013560f81c60248301525063ffffffff8f166044820152606401602060405180830381865afa158015611489573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906114ad9190615651565b85606001518560ff16815181106114c6576114c66153ad565b602002602001015184815181106114df576114df6153ad565b63ffffffff90921660209283029190910190910152826114fe8161542a565b9350505b508061150d8161542a565b9150506112f7565b506000816001600160401b0381111561153057611530614590565b604051908082528060200260200182016040528015611559578160200160208202803683370190505b50905060005b828110156115da5784606001518460ff1681518110611580576115806153ad565b60200260200101518181518110611599576115996153ad565b60200260200101518282815181106115b3576115b36153ad565b63ffffffff90921660209283029190910190910152806115d28161542a565b91505061155f565b508084606001518460ff16815181106115f5576115f56153ad565b60200260200101819052505050808061160d9061566e565b915050611281565b506000896001600160a01b0316635df459466040518163ffffffff1660e01b8152600401602060405180830381865afa158015611656573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061167a91906152e1565b60405163354952a360e21b81529091506001600160a01b0382169063d5254a8c906116ad908b908b908e9060040161568e565b600060405180830381865afa1580156116ca573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526116f29190810190615548565b60208301525098975050505050505050565b60655460405163237dfb4760e11b81523360048201526001600160a01b03909116906346fbf68e90602401602060405180830381865afa15801561174c573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906117709190615348565b61178c5760405162461bcd60e51b81526004016106c890615365565b600019606681905560405190815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d9060200160405180910390a2565b60006117da6020850185614c82565b63ffffffff81166000908152609a60205260409020549091506118345760405162461bcd60e51b815260206004820152601260248201527115185cdac81b9bdd081c995cdc1bdb99195960721b60448201526064016106c8565b61183e8484613ac3565b63ffffffff82166000908152609a6020526040902054146118975760405162461bcd60e51b815260206004820152601360248201527257726f6e67207461736b20726573706f6e736560681b60448201526064016106c8565b63ffffffff81166000908152609b602052604090205460ff16156118fd5760405162461bcd60e51b815260206004820152601760248201527f416c7265616479206265656e206368616c6c656e67656400000000000000000060448201526064016106c8565b606461190c6020850185614c82565b61191691906156b8565b63ffffffff164363ffffffff1611156119715760405162461bcd60e51b815260206004820152601860248201527f4368616c6c656e676520706572696f642065787069726564000000000000000060448201526064016106c8565b604051339063ffffffff8316907f0c6923c4a98292e75c5d677a1634527f87b6d19cf2c7d396aece99790c44a79590600090a350611aef565b8351811015611a16576119e78482815181106119c8576119c86153ad565b6020026020010151805160009081526020918201519091526040902090565b8282815181106119f9576119f96153ad565b602090810291909101015280611a0e8161542a565b9150506119aa565b506000611a266020880188614c82565b82604051602001611a389291906156e0565b60405160208183030381529060405280519060200120905084602001358114611aa35760405162461bcd60e51b815260206004820152601860248201527f57726f6e67206e6f6e2d7369676e6572207075626b657973000000000000000060448201526064016106c8565b63ffffffff83166000818152609b6020526040808220805460ff19166001179055513392917fff48388ad5e2a6d1845a7672040fba7d9b14b22b9e0eecd37046e5313d3aebc291a35050505b50505050565b6040805180820190915260608082526020820152600084611b6c5760405162461bcd60e51b81526020600482015260376024820152600080516020615e2b83398151915260448201527f7265733a20656d7074792071756f72756d20696e70757400000000000000000060648201526084016106c8565b60408301515185148015611b84575060a08301515185145b8015611b94575060c08301515185145b8015611ba4575060e08301515185145b611c0e5760405162461bcd60e51b81526020600482015260416024820152600080516020615e2b83398151915260448201527f7265733a20696e7075742071756f72756d206c656e677468206d69736d6174636064820152600d60fb1b608482015260a4016106c8565b82515160208401515114611c865760405162461bcd60e51b815260206004820152604460248201819052600080516020615e2b833981519152908201527f7265733a20696e707574206e6f6e7369676e6572206c656e677468206d69736d6064820152630c2e8c6d60e31b608482015260a4016106c8565b4363ffffffff168463ffffffff161115611cf65760405162461bcd60e51b815260206004820152603c6024820152600080516020615e2b83398151915260448201527f7265733a20696e76616c6964207265666572656e636520626c6f636b0000000060648201526084016106c8565b6040805180820182526000808252602080830191909152825180840190935260608084529083015290866001600160401b03811115611d3757611d37614590565b604051908082528060200260200182016040528015611d60578160200160208202803683370190505b506020820152866001600160401b03811115611d7e57611d7e614590565b604051908082528060200260200182016040528015611da7578160200160208202803683370190505b50815260408051808201909152606080825260208201528560200151516001600160401b03811115611ddb57611ddb614590565b604051908082528060200260200182016040528015611e04578160200160208202803683370190505b5081526020860151516001600160401b03811115611e2457611e24614590565b604051908082528060200260200182016040528015611e4d578160200160208202803683370190505b5081602001819052506000611f1f8a8a8080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152505060408051639aa1653d60e01b815290516001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169350639aa1653d925060048083019260209291908290030181865afa158015611ef6573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611f1a9190615728565b613af7565b905060005b87602001515181101561219b57611f4a886020015182815181106119c8576119c86153ad565b83602001518281518110611f6057611f606153ad565b60209081029190910101528015612020576020830151611f81600183615745565b81518110611f9157611f916153ad565b602002602001015160001c83602001518281518110611fb257611fb26153ad565b602002602001015160001c11612020576040805162461bcd60e51b8152602060048201526024810191909152600080516020615e2b83398151915260448201527f7265733a206e6f6e5369676e65725075626b657973206e6f7420736f7274656460648201526084016106c8565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166304ec635184602001518381518110612065576120656153ad565b60200260200101518b8b600001518581518110612084576120846153ad565b60200260200101516040518463ffffffff1660e01b81526004016120c19392919092835263ffffffff918216602084015216604082015260600190565b602060405180830381865afa1580156120de573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906121029190615628565b6001600160c01b031683600001518281518110612121576121216153ad565b60200260200101818152505061218761093761215b848660000151858151811061214d5761214d6153ad565b602002602001015116613ba9565b8a602001518481518110612171576121716153ad565b6020026020010151613bd490919063ffffffff16565b9450806121938161542a565b915050611f24565b50506121a683613cb8565b925060007f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166350f73e7c6040518163ffffffff1660e01b8152600401602060405180830381865afa158015612208573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061222c919061575c565b60975490915060ff1660005b8a8110156128b2578115612394578963ffffffff16837f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663249a0c428f8f8681811061228f5761228f6153ad565b60405160e085901b6001600160e01b031916815292013560f81c600483015250602401602060405180830381865afa1580156122cf573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906122f3919061575c565b6122fd9190615775565b10156123945760405162461bcd60e51b81526020600482015260666024820152600080516020615e2b83398151915260448201527f7265733a205374616b6552656769737472792075706461746573206d7573742060648201527f62652077697468696e207769746864726177616c44656c6179426c6f636b732060848201526577696e646f7760d01b60a482015260c4016106c8565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166368bccaac8d8d848181106123d5576123d56153ad565b9050013560f81c60f81b60f81c8c8c60a0015185815181106123f9576123f96153ad565b60209081029190910101516040516001600160e01b031960e086901b16815260ff909316600484015263ffffffff9182166024840152166044820152606401602060405180830381865afa158015612455573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612479919061578d565b6001600160401b03191661249c8a6040015183815181106119c8576119c86153ad565b67ffffffffffffffff1916146125385760405162461bcd60e51b81526020600482015260616024820152600080516020615e2b83398151915260448201527f7265733a2071756f72756d41706b206861736820696e2073746f72616765206460648201527f6f6573206e6f74206d617463682070726f76696465642071756f72756d2061706084820152606b60f81b60a482015260c4016106c8565b61256889604001518281518110612551576125516153ad565b6020026020010151876136bb90919063ffffffff16565b95507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663c8294c568d8d848181106125ab576125ab6153ad565b9050013560f81c60f81b60f81c8c8c60c0015185815181106125cf576125cf6153ad565b60209081029190910101516040516001600160e01b031960e086901b16815260ff909316600484015263ffffffff9182166024840152166044820152606401602060405180830381865afa15801561262b573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061264f91906154d5565b85602001518281518110612665576126656153ad565b6001600160601b03909216602092830291909101820152850151805182908110612691576126916153ad565b6020026020010151856000015182815181106126af576126af6153ad565b60200260200101906001600160601b031690816001600160601b0316815250506000805b8a602001515181101561289d57612727866000015182815181106126f9576126f96153ad565b60200260200101518f8f86818110612713576127136153ad565b600192013560f81c9290921c811614919050565b1561288b577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663f2be94ae8f8f8681811061276d5761276d6153ad565b9050013560f81c60f81b60f81c8e89602001518581518110612791576127916153ad565b60200260200101518f60e0015188815181106127af576127af6153ad565b602002602001015187815181106127c8576127c86153ad565b60209081029190910101516040516001600160e01b031960e087901b16815260ff909416600485015263ffffffff92831660248501526044840191909152166064820152608401602060405180830381865afa15801561282c573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061285091906154d5565b8751805185908110612864576128646153ad565b6020026020010181815161287891906157b8565b6001600160601b03169052506001909101905b806128958161542a565b9150506126d3565b505080806128aa9061542a565b915050612238565b5050506000806128cc8c868a606001518b6080015161081c565b915091508161293d5760405162461bcd60e51b81526020600482015260436024820152600080516020615e2b83398151915260448201527f7265733a2070616972696e6720707265636f6d70696c652063616c6c206661696064820152621b195960ea1b608482015260a4016106c8565b8061299e5760405162461bcd60e51b81526020600482015260396024820152600080516020615e2b83398151915260448201527f7265733a207369676e617475726520697320696e76616c69640000000000000060648201526084016106c8565b505060008782602001516040516020016129b99291906156e0565b60408051808303601f190181529190528051602090910120929b929a509198505050505050505050565b6129eb613d53565b6129f56000613dad565b565b6000612a0284613dff565b826020013514612a4a5760405162461bcd60e51b81526020600482015260136024820152720aee4dedcce40dacae6e6c2ceca40d2dcc8caf606b1b60448201526064016106c8565b612a5a8360200135606084613e38565b612a9a5760405162461bcd60e51b815260206004820152601160248201527024b73b30b634b21029a6aa10383937b7b360791b60448201526064016106c8565b6000612aa585613e4e565b8335149150509392505050565b6040805160018082528183019092526000916060918391602080830190803683370190505090508481600081518110612aed57612aed6153ad565b60209081029190910101526040516361c8a12f60e11b81526000906001600160a01b0388169063c391425e90612b2990889086906004016157e0565b600060405180830381865afa158015612b46573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052612b6e9190810190615548565b600081518110612b8057612b806153ad565b60209081029190910101516040516304ec635160e01b81526004810188905263ffffffff87811660248301529091166044820181905291506000906001600160a01b038916906304ec635190606401602060405180830381865afa158015612bec573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612c109190615628565b6001600160c01b031690506000612c2682613e7e565b905081612c348a838a610a73565b9550955050505050935093915050565b6098546001600160a01b03163314612c9e5760405162461bcd60e51b815260206004820152601d60248201527f41676772656761746f72206d757374206265207468652063616c6c657200000060448201526064016106c8565b6000612cad6020850185614c82565b9050366000612cbf6080870187615834565b90925090506000612cd66080880160608901614c82565b905060996000612ce96020890189614c82565b63ffffffff1663ffffffff16815260200190815260200160002054612d0d88613edb565b14612d4c5760405162461bcd60e51b815260206004820152600f60248201526e0aee4dedcce40e8c2e6d640d0c2e6d608b1b60448201526064016106c8565b6000609a81612d5e60208a018a614c82565b63ffffffff1663ffffffff1681526020019081526020016000205414612dbf5760405162461bcd60e51b815260206004820152601660248201527515185cdac8185b1c9958591e481c995cdc1bdb99195960521b60448201526064016106c8565b612de97f0000000000000000000000000000000000000000000000000000000000000000856156b8565b63ffffffff164363ffffffff161115612e3d5760405162461bcd60e51b815260206004820152601660248201527514995cdc1bdb9cd9481d1a5b5948195e18d95959195960521b60448201526064016106c8565b6000612e4887613eee565b9050600080612e5b8387878a8c896109a6565b9150915081612e9d5760405162461bcd60e51b815260206004820152600e60248201526d145d5bdc9d5b481b9bdd081b595d60921b60448201526064016106c8565b6040805180820190915263ffffffff4316815260208101829052612ed081612eca368d90038d018d61587a565b90613f01565b609a6000612ee160208e018e614c82565b63ffffffff1663ffffffff168152602001908152602001600020819055507f8016fcc5ad5dcf12fff2e128d239d9c6eb61f4041126bbac2c93fa8962627c1b8a82604051612f30929190615903565b60405180910390a15050505050505050505050565b6097546501000000000090046001600160a01b03163314612fb25760405162461bcd60e51b815260206004820152602160248201527f5461736b2067656e657261746f72206d757374206265207468652063616c6c656044820152603960f91b60648201526084016106c8565b606463ffffffff8416111561301b5760405162461bcd60e51b815260206004820152602960248201527f51756f72756d207468726573686f6c642067726561746572207468616e2064656044820152683737b6b4b730ba37b960b91b60648201526084016106c8565b60006040518060a001604052804363ffffffff168152602001876001600160401b03168152602001866001600160401b031681526020018563ffffffff16815260200184848080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250505091525090506130a181613f16565b6097805463ffffffff610100918290048116600090815260996020526040908190209490945591549251920416907f78aec7310ea6fd468e3d3bbd16a806fd4987515634d5b5bf4cf4f036d9c33225906130fc908490615959565b60405180910390a260975461311d90610100900463ffffffff1660016156b8565b609760016101000a81548163ffffffff021916908363ffffffff160217905550505050505050565b600061315084613f29565b8260200135146131985760405162461bcd60e51b81526020600482015260136024820152720aee4dedcce40dacae6e6c2ceca40d2dcc8caf606b1b60448201526064016106c8565b6131a88360400135604084613e38565b6131e85760405162461bcd60e51b815260206004820152601160248201527024b73b30b634b21029a6aa10383937b7b360791b60448201526064016106c8565b6000612aa585613f47565b6131fb613d53565b6001600160a01b0381166132605760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b60648201526084016106c8565b6106da81613dad565b600054610100900460ff16158080156132895750600054600160ff909116105b806132a35750303b1580156132a3575060005460ff166001145b6133065760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b60648201526084016106c8565b6000805460ff191660011790558015613329576000805461ff0019166101001790555b613334856000613f5a565b61333d84613dad565b609880546001600160a01b0319166001600160a01b03858116919091179091556097805465010000000000600160c81b031916650100000000009285169290920291909117905580156133ca576000805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b5050505050565b606560009054906101000a90046001600160a01b03166001600160a01b031663eab66d7a6040518163ffffffff1660e01b8152600401602060405180830381865afa158015613424573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061344891906152e1565b6001600160a01b0316336001600160a01b0316146134785760405162461bcd60e51b81526004016106c8906152fe565b6066541981196066541916146134f65760405162461bcd60e51b815260206004820152603860248201527f5061757361626c652e756e70617573653a20696e76616c696420617474656d7060448201527f7420746f2070617573652066756e6374696f6e616c697479000000000000000060648201526084016106c8565b606681905560405181815233907f3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c90602001610811565b6001600160a01b0381166135bb5760405162461bcd60e51b815260206004820152604960248201527f5061757361626c652e5f73657450617573657252656769737472793a206e657760448201527f50617573657252656769737472792063616e6e6f7420626520746865207a65726064820152686f206164647265737360b81b608482015260a4016106c8565b606554604080516001600160a01b03928316815291831660208301527f6e9fcd539896fca60e8b0f01dd580233e48a6b0f7df013b89ba7f565869acdb6910160405180910390a1606580546001600160a01b0319166001600160a01b0392909216919091179055565b604080518082019091526000808252602082015261364061446b565b835181526020808501519082015260408082018490526000908360608460076107d05a03fa905080801561367357613675565bfe5b50806136b35760405162461bcd60e51b815260206004820152600d60248201526c1958cb5b5d5b0b59985a5b1959609a1b60448201526064016106c8565b505092915050565b60408051808201909152600080825260208201526136d7614489565b835181526020808501518183015283516040808401919091529084015160608301526000908360808460066107d05a03fa90508080156136735750806136b35760405162461bcd60e51b815260206004820152600d60248201526c1958cb5859190b59985a5b1959609a1b60448201526064016106c8565b6137576144a7565b50604080516080810182527f198e9393920d483a7260bfb731fb5d25f1aa493335a9e71297e485b7aef312c28183019081527f1800deef121f1e76426a00665e5c4479674322d4f75edadd46debd5cd992f6ed6060830152815281518083019092527f275dc4a288d1afb3cbb1ac09187524c7db36395df7be3b99e673b13a075a65ec82527f1d9befcd05a5323e6da4d435f3b617cdb3af83285c2df711ef39c01571827f9d60208381019190915281019190915290565b60408051808201909152600080825260208201526000808061383f600080516020615e0b833981519152866153c3565b90505b61384b81614044565b9093509150600080516020615e0b833981519152828309831415613885576040805180820190915290815260208101919091529392505050565b600080516020615e0b833981519152600182089050613842565b6040805180820182528681526020808201869052825180840190935286835282018490526000918291906138d16144cc565b60005b6002811015613a965760006138ea8260066159d6565b90508482600281106138fe576138fe6153ad565b60200201515183613910836000615775565b600c8110613920576139206153ad565b6020020152848260028110613937576139376153ad565b6020020151602001518382600161394e9190615775565b600c811061395e5761395e6153ad565b6020020152838260028110613975576139756153ad565b6020020151515183613988836002615775565b600c8110613998576139986153ad565b60200201528382600281106139af576139af6153ad565b60200201515160016020020151836139c8836003615775565b600c81106139d8576139d86153ad565b60200201528382600281106139ef576139ef6153ad565b602002015160200151600060028110613a0a57613a0a6153ad565b602002015183613a1b836004615775565b600c8110613a2b57613a2b6153ad565b6020020152838260028110613a4257613a426153ad565b602002015160200151600160028110613a5d57613a5d6153ad565b602002015183613a6e836005615775565b600c8110613a7e57613a7e6153ad565b60200201525080613a8e8161542a565b9150506138d4565b50613a9f6144eb565b60006020826101808560088cfa9151919c9115159b50909950505050505050505050565b60008282604051602001613ad89291906159f5565b6040516020818303038152906040528051906020012090505b92915050565b600080613b03846140c6565b90508015610f04578260ff168460018651613b1e9190615745565b81518110613b2e57613b2e6153ad565b016020015160f81c10610f045760405162461bcd60e51b815260206004820152603f60248201527f4269746d61705574696c732e6f72646572656442797465734172726179546f4260448201527f69746d61703a206269746d61702065786365656473206d61782076616c75650060648201526084016106c8565b6000805b8215613af157613bbe600184615745565b9092169180613bcc81615a2b565b915050613bad565b60408051808201909152600080825260208201526102008261ffff1610613c305760405162461bcd60e51b815260206004820152601060248201526f7363616c61722d746f6f2d6c6172676560801b60448201526064016106c8565b8161ffff1660011415613c44575081613af1565b6040805180820190915260008082526020820181905284906001905b8161ffff168661ffff1610613cad57600161ffff871660ff83161c81161415613c9057613c8d84846136bb565b93505b613c9a83846136bb565b92506201fffe600192831b169101613c60565b509195945050505050565b60408051808201909152600080825260208201528151158015613cdd57506020820151155b15613cfb575050604080518082019091526000808252602082015290565b604051806040016040528083600001518152602001600080516020615e0b8339815191528460200151613d2e91906153c3565b613d4690600080516020615e0b833981519152615745565b905292915050565b919050565b6033546001600160a01b031633146129f55760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016106c8565b603380546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b60006040613e106020840184614c82565b63ffffffff16901b613e286040840160208501615a4d565b6001600160401b03161792915050565b600083613e458484614253565b14949350505050565b600081604051602001613e619190615a68565b604051602081830303815290604052805190602001209050919050565b60606000805b610100811015613ed4576001811b915083821615613ec457828160f81b604051602001613eb2929190615ac0565b60405160208183030381529060405292505b613ecd8161542a565b9050613e84565b5050919050565b600081604051602001613e619190615aef565b600081604051602001613e619190615bab565b60008282604051602001613ad8929190615bb9565b600081604051602001613e619190615959565b6000613f386020830183615a4d565b6001600160401b031692915050565b600081604051602001613e619190615bf8565b6065546001600160a01b0316158015613f7b57506001600160a01b03821615155b613ffd5760405162461bcd60e51b815260206004820152604760248201527f5061757361626c652e5f696e697469616c697a655061757365723a205f696e6960448201527f7469616c697a6550617573657228292063616e206f6e6c792062652063616c6c6064820152666564206f6e636560c81b608482015260a4016106c8565b606681905560405181815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d9060200160405180910390a26140408261352d565b5050565b60008080600080516020615e0b8339815191526003600080516020615e0b83398151915286600080516020615e0b8339815191528889090908905060006140ba827f0c19139cb84c680a6e14116da060561765e05aa45a1c72a34f082305b61f3f52600080516020615e0b83398151915261439f565b91959194509092505050565b60006101008251111561414f5760405162461bcd60e51b8152602060048201526044602482018190527f4269746d61705574696c732e6f72646572656442797465734172726179546f42908201527f69746d61703a206f7264657265644279746573417272617920697320746f6f206064820152636c6f6e6760e01b608482015260a4016106c8565b815161415d57506000919050565b60008083600081518110614173576141736153ad565b0160200151600160f89190911c81901b92505b845181101561424a578481815181106141a1576141a16153ad565b0160200151600160f89190911c1b91508282116142365760405162461bcd60e51b815260206004820152604760248201527f4269746d61705574696c732e6f72646572656442797465734172726179546f4260448201527f69746d61703a206f72646572656442797465734172726179206973206e6f74206064820152661bdc99195c995960ca1b608482015260a4016106c8565b918117916142438161542a565b9050614186565b50909392505050565b60006101008311156142a75760405162461bcd60e51b815260206004820152601b60248201527f44657074682073686f756c64206265206174206d6f737420323536000000000060448201526064016106c8565b82610100146142fd576142bb836002615db5565b8260200135106142fd5760405162461bcd60e51b815260206004820152600d60248201526c092dcecc2d8d2c840d2dcc8caf609b1b60448201526064016106c8565b81356000805b8581101561439557600060016040870135831c811614614324576000614354565b6143316060870187615dc1565b8461433b8161542a565b955081811061434c5761434c6153ad565b905060200201355b905060016020870135831c811614614375576143708482614447565b61437f565b61437f8185614447565b935050808061438d9061542a565b915050614303565b5090949350505050565b6000806143aa6144eb565b6143b2614509565b602080825281810181905260408201819052606082018890526080820187905260a082018690528260c08360056107d05a03fa925082801561367357508261443c5760405162461bcd60e51b815260206004820152601a60248201527f424e3235342e6578704d6f643a2063616c6c206661696c75726500000000000060448201526064016106c8565b505195945050505050565b60008215158061445657508115155b15613af1575060009182526020526040902090565b60405180606001604052806003906020820280368337509192915050565b60405180608001604052806004906020820280368337509192915050565b60405180604001604052806144ba614527565b81526020016144c7614527565b905290565b604051806101800160405280600c906020820280368337509192915050565b60405180602001604052806001906020820280368337509192915050565b6040518060c001604052806006906020820280368337509192915050565b60405180604001604052806002906020820280368337509192915050565b6001600160a01b03811681146106da57600080fd5b60006020828403121561456c57600080fd5b8135610f0481614545565b60006020828403121561458957600080fd5b5035919050565b634e487b7160e01b600052604160045260246000fd5b604080519081016001600160401b03811182821017156145c8576145c8614590565b60405290565b60405161010081016001600160401b03811182821017156145c8576145c8614590565b604051601f8201601f191681016001600160401b038111828210171561461957614619614590565b604052919050565b60006040828403121561463357600080fd5b61463b6145a6565b9050813581526020820135602082015292915050565b600082601f83011261466257600080fd5b604051604081018181106001600160401b038211171561468457614684614590565b806040525080604084018581111561469b57600080fd5b845b81811015613cad57803583526020928301920161469d565b6000608082840312156146c757600080fd5b6146cf6145a6565b90506146db8383614651565b81526146ea8360408401614651565b602082015292915050565b600080600080610120858703121561470c57600080fd5b8435935061471d8660208701614621565b925061472c86606087016146b5565b915061473b8660e08701614621565b905092959194509250565b60008083601f84011261475857600080fd5b5081356001600160401b0381111561476f57600080fd5b60208301915083602082850101111561478757600080fd5b9250929050565b63ffffffff811681146106da57600080fd5b8035613d4e8161478e565b60006001600160401b038211156147c4576147c4614590565b5060051b60200190565b600082601f8301126147df57600080fd5b813560206147f46147ef836147ab565b6145f1565b82815260059290921b8401810191818101908684111561481357600080fd5b8286015b8481101561483757803561482a8161478e565b8352918301918301614817565b509695505050505050565b600082601f83011261485357600080fd5b813560206148636147ef836147ab565b82815260069290921b8401810191818101908684111561488257600080fd5b8286015b84811015614837576148988882614621565b835291830191604001614886565b600082601f8301126148b757600080fd5b813560206148c76147ef836147ab565b82815260059290921b840181019181810190868411156148e657600080fd5b8286015b848110156148375780356001600160401b038111156149095760008081fd5b6149178986838b01016147ce565b8452509183019183016148ea565b6000610180828403121561493857600080fd5b6149406145ce565b905081356001600160401b038082111561495957600080fd5b614965858386016147ce565b8352602084013591508082111561497b57600080fd5b61498785838601614842565b602084015260408401359150808211156149a057600080fd5b6149ac85838601614842565b60408401526149be85606086016146b5565b60608401526149d08560e08601614621565b60808401526101208401359150808211156149ea57600080fd5b6149f6858386016147ce565b60a0840152610140840135915080821115614a1057600080fd5b614a1c858386016147ce565b60c0840152610160840135915080821115614a3657600080fd5b50614a43848285016148a6565b60e08301525092915050565b60008060008060008060a08789031215614a6857600080fd5b8635955060208701356001600160401b0380821115614a8657600080fd5b614a928a838b01614746565b909750955060408901359150614aa78261478e565b90935060608801359080821115614abd57600080fd5b50614aca89828a01614925565b9250506080870135614adb8161478e565b809150509295509295509295565b600080600060608486031215614afe57600080fd5b8335614b0981614545565b92506020848101356001600160401b0380821115614b2657600080fd5b818701915087601f830112614b3a57600080fd5b813581811115614b4c57614b4c614590565b614b5e601f8201601f191685016145f1565b91508082528884828501011115614b7457600080fd5b8084840185840137600084828401015250809450505050614b97604085016147a0565b90509250925092565b600081518084526020808501808196508360051b810191508286016000805b86811015614c36578385038a52825180518087529087019087870190845b81811015614c2157835180516001600160a01b031684528a8101518b8501526040908101516001600160601b03169084015292890192606090920191600101614bdd565b50509a87019a95505091850191600101614bbf565b509298975050505050505050565b602081526000610f046020830184614ba0565b80151581146106da57600080fd5b600060208284031215614c7757600080fd5b8135610f0481614c57565b600060208284031215614c9457600080fd5b8135610f048161478e565b60008060008060008060808789031215614cb857600080fd5b8635614cc381614545565b95506020870135614cd38161478e565b945060408701356001600160401b0380821115614cef57600080fd5b614cfb8a838b01614746565b90965094506060890135915080821115614d1457600080fd5b818901915089601f830112614d2857600080fd5b813581811115614d3757600080fd5b8a60208260051b8501011115614d4c57600080fd5b6020830194508093505050509295509295509295565b600081518084526020808501945080840160005b83811015614d9857815163ffffffff1687529582019590820190600101614d76565b509495945050505050565b600060208083528351608082850152614dbf60a0850182614d62565b905081850151601f1980868403016040870152614ddc8383614d62565b92506040870151915080868403016060870152614df98383614d62565b60608801518782038301608089015280518083529194508501925084840190600581901b8501860160005b82811015614e505784878303018452614e3e828751614d62565b95880195938801939150600101614e24565b509998505050505050505050565b60ff811681146106da57600080fd5b600060208284031215614e7f57600080fd5b8135610f0481614e5e565b600060a08284031215614e9c57600080fd5b50919050565b600060608284031215614e9c57600080fd5b60008060008084860360e0811215614ecb57600080fd5b85356001600160401b0380821115614ee257600080fd5b614eee89838a01614e8a565b9650614efd8960208a01614ea2565b95506040607f1984011215614f1157600080fd5b60808801945060c0880135925080831115614f2b57600080fd5b5050614f3987828801614842565b91505092959194509250565b600080600080600060808688031215614f5d57600080fd5b8535945060208601356001600160401b0380821115614f7b57600080fd5b614f8789838a01614746565b909650945060408801359150614f9c8261478e565b90925060608701359080821115614fb257600080fd5b50614fbf88828901614925565b9150509295509295909350565b600081518084526020808501945080840160005b83811015614d985781516001600160601b031687529582019590820190600101614fe0565b60408152600083516040808401526150206080840182614fcc565b90506020850151603f1984830301606085015261503d8282614fcc565b925050508260208301529392505050565b600060808284031215614e9c57600080fd5b6000806000610100848603121561507657600080fd5b615080858561504e565b925061508f8560808601614ea2565b915060e08401356001600160401b038111156150aa57600080fd5b6150b68682870161504e565b9150509250925092565b6000806000606084860312156150d557600080fd5b83356150e081614545565b92506020840135915060408401356150f78161478e565b809150509250925092565b82815260406020820152600061511b6040830184614ba0565b949350505050565b600080600060a0848603121561513857600080fd5b83356001600160401b038082111561514f57600080fd5b61515b87838801614e8a565b945061516a8760208801614ea2565b9350608086013591508082111561518057600080fd5b506150b686828701614925565b80356001600160401b0381168114613d4e57600080fd5b6000806000806000608086880312156151bc57600080fd5b6151c58661518d565b94506151d36020870161518d565b935060408601356151e38161478e565b925060608601356001600160401b038111156151fe57600080fd5b61520a88828901614746565b969995985093965092949392505050565b600080600060a0848603121561523057600080fd5b83356001600160401b038082111561524757600080fd5b61525387838801614ea2565b94506152628760208801614ea2565b9350608086013591508082111561527857600080fd5b506150b68682870161504e565b6000806000806080858703121561529b57600080fd5b84356152a681614545565b935060208501356152b681614545565b925060408501356152c681614545565b915060608501356152d681614545565b939692955090935050565b6000602082840312156152f357600080fd5b8151610f0481614545565b6020808252602a908201527f6d73672e73656e646572206973206e6f74207065726d697373696f6e6564206160408201526939903ab73830bab9b2b960b11b606082015260800190565b60006020828403121561535a57600080fd5b8151610f0481614c57565b60208082526028908201527f6d73672e73656e646572206973206e6f74207065726d697373696f6e6564206160408201526739903830bab9b2b960c11b606082015260800190565b634e487b7160e01b600052603260045260246000fd5b6000826153e057634e487b7160e01b600052601260045260246000fd5b500690565b634e487b7160e01b600052601160045260246000fd5b60006001600160601b0380831681851681830481118215151615615421576154216153e5565b02949350505050565b600060001982141561543e5761543e6153e5565b5060010190565b6000602080838503121561545857600080fd5b82516001600160401b0381111561546e57600080fd5b8301601f8101851361547f57600080fd5b805161548d6147ef826147ab565b81815260059190911b820183019083810190878311156154ac57600080fd5b928401925b828410156154ca578351825292840192908401906154b1565b979650505050505050565b6000602082840312156154e757600080fd5b81516001600160601b0381168114610f0457600080fd5b63ffffffff84168152604060208201819052810182905260006001600160fb1b0383111561552b57600080fd5b8260051b8085606085013760009201606001918252509392505050565b6000602080838503121561555b57600080fd5b82516001600160401b0381111561557157600080fd5b8301601f8101851361558257600080fd5b80516155906147ef826147ab565b81815260059190911b820183019083810190878311156155af57600080fd5b928401925b828410156154ca5783516155c78161478e565b825292840192908401906155b4565b81835281816020850137506000828201602090810191909152601f909101601f19169091010190565b63ffffffff8416815260406020820152600061561f6040830184866155d6565b95945050505050565b60006020828403121561563a57600080fd5b81516001600160c01b0381168114610f0457600080fd5b60006020828403121561566357600080fd5b8151610f048161478e565b600060ff821660ff811415615685576156856153e5565b60010192915050565b6040815260006156a26040830185876155d6565b905063ffffffff83166020830152949350505050565b600063ffffffff8083168185168083038211156156d7576156d76153e5565b01949350505050565b63ffffffff60e01b8360e01b1681526000600482018351602080860160005b8381101561571b578151855293820193908201906001016156ff565b5092979650505050505050565b60006020828403121561573a57600080fd5b8151610f0481614e5e565b600082821015615757576157576153e5565b500390565b60006020828403121561576e57600080fd5b5051919050565b60008219821115615788576157886153e5565b500190565b60006020828403121561579f57600080fd5b815167ffffffffffffffff1981168114610f0457600080fd5b60006001600160601b03838116908316818110156157d8576157d86153e5565b039392505050565b60006040820163ffffffff851683526020604081850152818551808452606086019150828701935060005b818110156158275784518352938301939183019160010161580b565b5090979650505050505050565b6000808335601e1984360301811261584b57600080fd5b8301803591506001600160401b0382111561586557600080fd5b60200191503681900382131561478757600080fd5b60006060828403121561588c57600080fd5b604051606081018181106001600160401b03821117156158ae576158ae614590565b60405282356158bc8161478e565b8152602083810135908201526040928301359281019290925250919050565b80356158e68161478e565b63ffffffff16825260208181013590830152604090810135910152565b60a0810161591182856158db565b825163ffffffff16606083015260208301516080830152610f04565b60005b83811015615948578181015183820152602001615930565b83811115611aef5750506000910152565b60208152600063ffffffff80845116602084015260208401516001600160401b038082166040860152806040870151166060860152505080606085015116608084015250608083015160a08084015280518060c08501526159c18160e086016020850161592d565b601f01601f19169290920160e0019392505050565b60008160001904831182151516156159f0576159f06153e5565b500290565b60a08101615a0382856158db565b8235615a0e8161478e565b63ffffffff16606083015260209290920135608090910152919050565b600061ffff80831681811415615a4357615a436153e5565b6001019392505050565b600060208284031215615a5f57600080fd5b610f048261518d565b608081018235615a778161478e565b63ffffffff168252615a8b6020840161518d565b6001600160401b03808216602085015280615aa86040870161518d565b16604085015250506060830135606083015292915050565b60008351615ad281846020880161592d565b6001600160f81b0319939093169190920190815260010192915050565b6020815260008235615b008161478e565b63ffffffff8082166020850152615b196020860161518d565b91506001600160401b03808316604086015280615b386040880161518d565b16606086015260608601359250615b4e8361478e565b818316608086015260808601359250601e19863603018312615b6f57600080fd5b918501918235915080821115615b8457600080fd5b50803603851315615b9457600080fd5b60a08085015261561f60c0850182602085016155d6565b60608101613af182846158db565b825163ffffffff168152602080840151908201526040808401519082015260a08101610f046060830184805163ffffffff168252602090810151910152565b60006020808352608083016001600160401b0380615c158761518d565b1683860152615c2583870161518d565b604082821681880152808801359150601e19883603018212615c4657600080fd5b90870190813583811115615c5957600080fd5b606093508381023603891315615c6e57600080fd5b87840184905293849052908401926000919060a08801835b82811015614e5057863582528787013588830152838701356fffffffffffffffffffffffffffffffff8116808214615cbc578687fd5b83860152509585019590850190600101615c86565b600181815b80851115615d0c578160001904821115615cf257615cf26153e5565b80851615615cff57918102915b93841c9390800290615cd6565b509250929050565b600082615d2357506001613af1565b81615d3057506000613af1565b8160018114615d465760028114615d5057615d6c565b6001915050613af1565b60ff841115615d6157615d616153e5565b50506001821b613af1565b5060208310610133831016604e8410600b8410161715615d8f575081810a613af1565b615d998383615cd1565b8060001904821115615dad57615dad6153e5565b029392505050565b6000610f048383615d14565b6000808335601e19843603018112615dd857600080fd5b8301803591506001600160401b03821115615df257600080fd5b6020019150600581901b360382131561478757600080fdfe30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd47424c535369676e6174757265436865636b65722e636865636b5369676e617475a2646970667358221220eab1d802095fdf4852f392980770803687d1d4bbe0b2c01edcddd5947ae9f17164736f6c634300080c0033",
}

// ContractSFFLTaskManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractSFFLTaskManagerMetaData.ABI instead.
var ContractSFFLTaskManagerABI = ContractSFFLTaskManagerMetaData.ABI

// ContractSFFLTaskManagerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ContractSFFLTaskManagerMetaData.Bin instead.
var ContractSFFLTaskManagerBin = ContractSFFLTaskManagerMetaData.Bin

// DeployContractSFFLTaskManager deploys a new Ethereum contract, binding an instance of ContractSFFLTaskManager to it.
func DeployContractSFFLTaskManager(auth *bind.TransactOpts, backend bind.ContractBackend, registryCoordinator common.Address, taskResponseWindowBlock uint32) (common.Address, *types.Transaction, *ContractSFFLTaskManager, error) {
	parsed, err := ContractSFFLTaskManagerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ContractSFFLTaskManagerBin), backend, registryCoordinator, taskResponseWindowBlock)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ContractSFFLTaskManager{ContractSFFLTaskManagerCaller: ContractSFFLTaskManagerCaller{contract: contract}, ContractSFFLTaskManagerTransactor: ContractSFFLTaskManagerTransactor{contract: contract}, ContractSFFLTaskManagerFilterer: ContractSFFLTaskManagerFilterer{contract: contract}}, nil
}

// ContractSFFLTaskManager is an auto generated Go binding around an Ethereum contract.
type ContractSFFLTaskManager struct {
	ContractSFFLTaskManagerCaller     // Read-only binding to the contract
	ContractSFFLTaskManagerTransactor // Write-only binding to the contract
	ContractSFFLTaskManagerFilterer   // Log filterer for contract events
}

// ContractSFFLTaskManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractSFFLTaskManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractSFFLTaskManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractSFFLTaskManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractSFFLTaskManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractSFFLTaskManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractSFFLTaskManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractSFFLTaskManagerSession struct {
	Contract     *ContractSFFLTaskManager // Generic contract binding to set the session for
	CallOpts     bind.CallOpts            // Call options to use throughout this session
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// ContractSFFLTaskManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractSFFLTaskManagerCallerSession struct {
	Contract *ContractSFFLTaskManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                  // Call options to use throughout this session
}

// ContractSFFLTaskManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractSFFLTaskManagerTransactorSession struct {
	Contract     *ContractSFFLTaskManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                  // Transaction auth options to use throughout this session
}

// ContractSFFLTaskManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractSFFLTaskManagerRaw struct {
	Contract *ContractSFFLTaskManager // Generic contract binding to access the raw methods on
}

// ContractSFFLTaskManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractSFFLTaskManagerCallerRaw struct {
	Contract *ContractSFFLTaskManagerCaller // Generic read-only contract binding to access the raw methods on
}

// ContractSFFLTaskManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractSFFLTaskManagerTransactorRaw struct {
	Contract *ContractSFFLTaskManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContractSFFLTaskManager creates a new instance of ContractSFFLTaskManager, bound to a specific deployed contract.
func NewContractSFFLTaskManager(address common.Address, backend bind.ContractBackend) (*ContractSFFLTaskManager, error) {
	contract, err := bindContractSFFLTaskManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ContractSFFLTaskManager{ContractSFFLTaskManagerCaller: ContractSFFLTaskManagerCaller{contract: contract}, ContractSFFLTaskManagerTransactor: ContractSFFLTaskManagerTransactor{contract: contract}, ContractSFFLTaskManagerFilterer: ContractSFFLTaskManagerFilterer{contract: contract}}, nil
}

// NewContractSFFLTaskManagerCaller creates a new read-only instance of ContractSFFLTaskManager, bound to a specific deployed contract.
func NewContractSFFLTaskManagerCaller(address common.Address, caller bind.ContractCaller) (*ContractSFFLTaskManagerCaller, error) {
	contract, err := bindContractSFFLTaskManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractSFFLTaskManagerCaller{contract: contract}, nil
}

// NewContractSFFLTaskManagerTransactor creates a new write-only instance of ContractSFFLTaskManager, bound to a specific deployed contract.
func NewContractSFFLTaskManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractSFFLTaskManagerTransactor, error) {
	contract, err := bindContractSFFLTaskManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractSFFLTaskManagerTransactor{contract: contract}, nil
}

// NewContractSFFLTaskManagerFilterer creates a new log filterer instance of ContractSFFLTaskManager, bound to a specific deployed contract.
func NewContractSFFLTaskManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractSFFLTaskManagerFilterer, error) {
	contract, err := bindContractSFFLTaskManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractSFFLTaskManagerFilterer{contract: contract}, nil
}

// bindContractSFFLTaskManager binds a generic wrapper to an already deployed contract.
func bindContractSFFLTaskManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContractSFFLTaskManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractSFFLTaskManager.Contract.ContractSFFLTaskManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractSFFLTaskManager.Contract.ContractSFFLTaskManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractSFFLTaskManager.Contract.ContractSFFLTaskManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractSFFLTaskManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractSFFLTaskManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractSFFLTaskManager.Contract.contract.Transact(opts, method, params...)
}

// TASKCHALLENGEWINDOWBLOCK is a free data retrieval call binding the contract method 0xf63c5bab.
//
// Solidity: function TASK_CHALLENGE_WINDOW_BLOCK() view returns(uint32)
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerCaller) TASKCHALLENGEWINDOWBLOCK(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _ContractSFFLTaskManager.contract.Call(opts, &out, "TASK_CHALLENGE_WINDOW_BLOCK")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// TASKCHALLENGEWINDOWBLOCK is a free data retrieval call binding the contract method 0xf63c5bab.
//
// Solidity: function TASK_CHALLENGE_WINDOW_BLOCK() view returns(uint32)
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerSession) TASKCHALLENGEWINDOWBLOCK() (uint32, error) {
	return _ContractSFFLTaskManager.Contract.TASKCHALLENGEWINDOWBLOCK(&_ContractSFFLTaskManager.CallOpts)
}

// TASKCHALLENGEWINDOWBLOCK is a free data retrieval call binding the contract method 0xf63c5bab.
//
// Solidity: function TASK_CHALLENGE_WINDOW_BLOCK() view returns(uint32)
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerCallerSession) TASKCHALLENGEWINDOWBLOCK() (uint32, error) {
	return _ContractSFFLTaskManager.Contract.TASKCHALLENGEWINDOWBLOCK(&_ContractSFFLTaskManager.CallOpts)
}

// TASKRESPONSEWINDOWBLOCK is a free data retrieval call binding the contract method 0x1ad43189.
//
// Solidity: function TASK_RESPONSE_WINDOW_BLOCK() view returns(uint32)
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerCaller) TASKRESPONSEWINDOWBLOCK(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _ContractSFFLTaskManager.contract.Call(opts, &out, "TASK_RESPONSE_WINDOW_BLOCK")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// TASKRESPONSEWINDOWBLOCK is a free data retrieval call binding the contract method 0x1ad43189.
//
// Solidity: function TASK_RESPONSE_WINDOW_BLOCK() view returns(uint32)
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerSession) TASKRESPONSEWINDOWBLOCK() (uint32, error) {
	return _ContractSFFLTaskManager.Contract.TASKRESPONSEWINDOWBLOCK(&_ContractSFFLTaskManager.CallOpts)
}

// TASKRESPONSEWINDOWBLOCK is a free data retrieval call binding the contract method 0x1ad43189.
//
// Solidity: function TASK_RESPONSE_WINDOW_BLOCK() view returns(uint32)
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerCallerSession) TASKRESPONSEWINDOWBLOCK() (uint32, error) {
	return _ContractSFFLTaskManager.Contract.TASKRESPONSEWINDOWBLOCK(&_ContractSFFLTaskManager.CallOpts)
}

// THRESHOLDDENOMINATOR is a free data retrieval call binding the contract method 0xef024458.
//
// Solidity: function THRESHOLD_DENOMINATOR() view returns(uint32)
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerCaller) THRESHOLDDENOMINATOR(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _ContractSFFLTaskManager.contract.Call(opts, &out, "THRESHOLD_DENOMINATOR")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// THRESHOLDDENOMINATOR is a free data retrieval call binding the contract method 0xef024458.
//
// Solidity: function THRESHOLD_DENOMINATOR() view returns(uint32)
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerSession) THRESHOLDDENOMINATOR() (uint32, error) {
	return _ContractSFFLTaskManager.Contract.THRESHOLDDENOMINATOR(&_ContractSFFLTaskManager.CallOpts)
}

// THRESHOLDDENOMINATOR is a free data retrieval call binding the contract method 0xef024458.
//
// Solidity: function THRESHOLD_DENOMINATOR() view returns(uint32)
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerCallerSession) THRESHOLDDENOMINATOR() (uint32, error) {
	return _ContractSFFLTaskManager.Contract.THRESHOLDDENOMINATOR(&_ContractSFFLTaskManager.CallOpts)
}

// Aggregator is a free data retrieval call binding the contract method 0x245a7bfc.
//
// Solidity: function aggregator() view returns(address)
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerCaller) Aggregator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractSFFLTaskManager.contract.Call(opts, &out, "aggregator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Aggregator is a free data retrieval call binding the contract method 0x245a7bfc.
//
// Solidity: function aggregator() view returns(address)
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerSession) Aggregator() (common.Address, error) {
	return _ContractSFFLTaskManager.Contract.Aggregator(&_ContractSFFLTaskManager.CallOpts)
}

// Aggregator is a free data retrieval call binding the contract method 0x245a7bfc.
//
// Solidity: function aggregator() view returns(address)
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerCallerSession) Aggregator() (common.Address, error) {
	return _ContractSFFLTaskManager.Contract.Aggregator(&_ContractSFFLTaskManager.CallOpts)
}

// AllCheckpointTaskHashes is a free data retrieval call binding the contract method 0x4f19ade7.
//
// Solidity: function allCheckpointTaskHashes(uint32 ) view returns(bytes32)
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerCaller) AllCheckpointTaskHashes(opts *bind.CallOpts, arg0 uint32) ([32]byte, error) {
	var out []interface{}
	err := _ContractSFFLTaskManager.contract.Call(opts, &out, "allCheckpointTaskHashes", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// AllCheckpointTaskHashes is a free data retrieval call binding the contract method 0x4f19ade7.
//
// Solidity: function allCheckpointTaskHashes(uint32 ) view returns(bytes32)
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerSession) AllCheckpointTaskHashes(arg0 uint32) ([32]byte, error) {
	return _ContractSFFLTaskManager.Contract.AllCheckpointTaskHashes(&_ContractSFFLTaskManager.CallOpts, arg0)
}

// AllCheckpointTaskHashes is a free data retrieval call binding the contract method 0x4f19ade7.
//
// Solidity: function allCheckpointTaskHashes(uint32 ) view returns(bytes32)
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerCallerSession) AllCheckpointTaskHashes(arg0 uint32) ([32]byte, error) {
	return _ContractSFFLTaskManager.Contract.AllCheckpointTaskHashes(&_ContractSFFLTaskManager.CallOpts, arg0)
}

// AllCheckpointTaskResponses is a free data retrieval call binding the contract method 0xa168e3c0.
//
// Solidity: function allCheckpointTaskResponses(uint32 ) view returns(bytes32)
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerCaller) AllCheckpointTaskResponses(opts *bind.CallOpts, arg0 uint32) ([32]byte, error) {
	var out []interface{}
	err := _ContractSFFLTaskManager.contract.Call(opts, &out, "allCheckpointTaskResponses", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// AllCheckpointTaskResponses is a free data retrieval call binding the contract method 0xa168e3c0.
//
// Solidity: function allCheckpointTaskResponses(uint32 ) view returns(bytes32)
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerSession) AllCheckpointTaskResponses(arg0 uint32) ([32]byte, error) {
	return _ContractSFFLTaskManager.Contract.AllCheckpointTaskResponses(&_ContractSFFLTaskManager.CallOpts, arg0)
}

// AllCheckpointTaskResponses is a free data retrieval call binding the contract method 0xa168e3c0.
//
// Solidity: function allCheckpointTaskResponses(uint32 ) view returns(bytes32)
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerCallerSession) AllCheckpointTaskResponses(arg0 uint32) ([32]byte, error) {
	return _ContractSFFLTaskManager.Contract.AllCheckpointTaskResponses(&_ContractSFFLTaskManager.CallOpts, arg0)
}

// BlsApkRegistry is a free data retrieval call binding the contract method 0x5df45946.
//
// Solidity: function blsApkRegistry() view returns(address)
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerCaller) BlsApkRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractSFFLTaskManager.contract.Call(opts, &out, "blsApkRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BlsApkRegistry is a free data retrieval call binding the contract method 0x5df45946.
//
// Solidity: function blsApkRegistry() view returns(address)
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerSession) BlsApkRegistry() (common.Address, error) {
	return _ContractSFFLTaskManager.Contract.BlsApkRegistry(&_ContractSFFLTaskManager.CallOpts)
}

// BlsApkRegistry is a free data retrieval call binding the contract method 0x5df45946.
//
// Solidity: function blsApkRegistry() view returns(address)
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerCallerSession) BlsApkRegistry() (common.Address, error) {
	return _ContractSFFLTaskManager.Contract.BlsApkRegistry(&_ContractSFFLTaskManager.CallOpts)
}

// CheckQuorum is a free data retrieval call binding the contract method 0x292f7a4e.
//
// Solidity: function checkQuorum(bytes32 messageHash, bytes quorumNumbers, uint32 referenceBlockNumber, (uint32[],(uint256,uint256)[],(uint256,uint256)[],(uint256[2],uint256[2]),(uint256,uint256),uint32[],uint32[],uint32[][]) nonSignerStakesAndSignature, uint32 quorumThreshold) view returns(bool, bytes32)
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerCaller) CheckQuorum(opts *bind.CallOpts, messageHash [32]byte, quorumNumbers []byte, referenceBlockNumber uint32, nonSignerStakesAndSignature IBLSSignatureCheckerNonSignerStakesAndSignature, quorumThreshold uint32) (bool, [32]byte, error) {
	var out []interface{}
	err := _ContractSFFLTaskManager.contract.Call(opts, &out, "checkQuorum", messageHash, quorumNumbers, referenceBlockNumber, nonSignerStakesAndSignature, quorumThreshold)

	if err != nil {
		return *new(bool), *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	out1 := *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)

	return out0, out1, err

}

// CheckQuorum is a free data retrieval call binding the contract method 0x292f7a4e.
//
// Solidity: function checkQuorum(bytes32 messageHash, bytes quorumNumbers, uint32 referenceBlockNumber, (uint32[],(uint256,uint256)[],(uint256,uint256)[],(uint256[2],uint256[2]),(uint256,uint256),uint32[],uint32[],uint32[][]) nonSignerStakesAndSignature, uint32 quorumThreshold) view returns(bool, bytes32)
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerSession) CheckQuorum(messageHash [32]byte, quorumNumbers []byte, referenceBlockNumber uint32, nonSignerStakesAndSignature IBLSSignatureCheckerNonSignerStakesAndSignature, quorumThreshold uint32) (bool, [32]byte, error) {
	return _ContractSFFLTaskManager.Contract.CheckQuorum(&_ContractSFFLTaskManager.CallOpts, messageHash, quorumNumbers, referenceBlockNumber, nonSignerStakesAndSignature, quorumThreshold)
}

// CheckQuorum is a free data retrieval call binding the contract method 0x292f7a4e.
//
// Solidity: function checkQuorum(bytes32 messageHash, bytes quorumNumbers, uint32 referenceBlockNumber, (uint32[],(uint256,uint256)[],(uint256,uint256)[],(uint256[2],uint256[2]),(uint256,uint256),uint32[],uint32[],uint32[][]) nonSignerStakesAndSignature, uint32 quorumThreshold) view returns(bool, bytes32)
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerCallerSession) CheckQuorum(messageHash [32]byte, quorumNumbers []byte, referenceBlockNumber uint32, nonSignerStakesAndSignature IBLSSignatureCheckerNonSignerStakesAndSignature, quorumThreshold uint32) (bool, [32]byte, error) {
	return _ContractSFFLTaskManager.Contract.CheckQuorum(&_ContractSFFLTaskManager.CallOpts, messageHash, quorumNumbers, referenceBlockNumber, nonSignerStakesAndSignature, quorumThreshold)
}

// CheckSignatures is a free data retrieval call binding the contract method 0x6efb4636.
//
// Solidity: function checkSignatures(bytes32 msgHash, bytes quorumNumbers, uint32 referenceBlockNumber, (uint32[],(uint256,uint256)[],(uint256,uint256)[],(uint256[2],uint256[2]),(uint256,uint256),uint32[],uint32[],uint32[][]) params) view returns((uint96[],uint96[]), bytes32)
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerCaller) CheckSignatures(opts *bind.CallOpts, msgHash [32]byte, quorumNumbers []byte, referenceBlockNumber uint32, params IBLSSignatureCheckerNonSignerStakesAndSignature) (IBLSSignatureCheckerQuorumStakeTotals, [32]byte, error) {
	var out []interface{}
	err := _ContractSFFLTaskManager.contract.Call(opts, &out, "checkSignatures", msgHash, quorumNumbers, referenceBlockNumber, params)

	if err != nil {
		return *new(IBLSSignatureCheckerQuorumStakeTotals), *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new(IBLSSignatureCheckerQuorumStakeTotals)).(*IBLSSignatureCheckerQuorumStakeTotals)
	out1 := *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)

	return out0, out1, err

}

// CheckSignatures is a free data retrieval call binding the contract method 0x6efb4636.
//
// Solidity: function checkSignatures(bytes32 msgHash, bytes quorumNumbers, uint32 referenceBlockNumber, (uint32[],(uint256,uint256)[],(uint256,uint256)[],(uint256[2],uint256[2]),(uint256,uint256),uint32[],uint32[],uint32[][]) params) view returns((uint96[],uint96[]), bytes32)
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerSession) CheckSignatures(msgHash [32]byte, quorumNumbers []byte, referenceBlockNumber uint32, params IBLSSignatureCheckerNonSignerStakesAndSignature) (IBLSSignatureCheckerQuorumStakeTotals, [32]byte, error) {
	return _ContractSFFLTaskManager.Contract.CheckSignatures(&_ContractSFFLTaskManager.CallOpts, msgHash, quorumNumbers, referenceBlockNumber, params)
}

// CheckSignatures is a free data retrieval call binding the contract method 0x6efb4636.
//
// Solidity: function checkSignatures(bytes32 msgHash, bytes quorumNumbers, uint32 referenceBlockNumber, (uint32[],(uint256,uint256)[],(uint256,uint256)[],(uint256[2],uint256[2]),(uint256,uint256),uint32[],uint32[],uint32[][]) params) view returns((uint96[],uint96[]), bytes32)
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerCallerSession) CheckSignatures(msgHash [32]byte, quorumNumbers []byte, referenceBlockNumber uint32, params IBLSSignatureCheckerNonSignerStakesAndSignature) (IBLSSignatureCheckerQuorumStakeTotals, [32]byte, error) {
	return _ContractSFFLTaskManager.Contract.CheckSignatures(&_ContractSFFLTaskManager.CallOpts, msgHash, quorumNumbers, referenceBlockNumber, params)
}

// CheckpointTaskNumber is a free data retrieval call binding the contract method 0x8cbc379a.
//
// Solidity: function checkpointTaskNumber() view returns(uint32)
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerCaller) CheckpointTaskNumber(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _ContractSFFLTaskManager.contract.Call(opts, &out, "checkpointTaskNumber")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// CheckpointTaskNumber is a free data retrieval call binding the contract method 0x8cbc379a.
//
// Solidity: function checkpointTaskNumber() view returns(uint32)
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerSession) CheckpointTaskNumber() (uint32, error) {
	return _ContractSFFLTaskManager.Contract.CheckpointTaskNumber(&_ContractSFFLTaskManager.CallOpts)
}

// CheckpointTaskNumber is a free data retrieval call binding the contract method 0x8cbc379a.
//
// Solidity: function checkpointTaskNumber() view returns(uint32)
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerCallerSession) CheckpointTaskNumber() (uint32, error) {
	return _ContractSFFLTaskManager.Contract.CheckpointTaskNumber(&_ContractSFFLTaskManager.CallOpts)
}

// CheckpointTaskSuccesfullyChallenged is a free data retrieval call binding the contract method 0x95eebee6.
//
// Solidity: function checkpointTaskSuccesfullyChallenged(uint32 ) view returns(bool)
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerCaller) CheckpointTaskSuccesfullyChallenged(opts *bind.CallOpts, arg0 uint32) (bool, error) {
	var out []interface{}
	err := _ContractSFFLTaskManager.contract.Call(opts, &out, "checkpointTaskSuccesfullyChallenged", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CheckpointTaskSuccesfullyChallenged is a free data retrieval call binding the contract method 0x95eebee6.
//
// Solidity: function checkpointTaskSuccesfullyChallenged(uint32 ) view returns(bool)
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerSession) CheckpointTaskSuccesfullyChallenged(arg0 uint32) (bool, error) {
	return _ContractSFFLTaskManager.Contract.CheckpointTaskSuccesfullyChallenged(&_ContractSFFLTaskManager.CallOpts, arg0)
}

// CheckpointTaskSuccesfullyChallenged is a free data retrieval call binding the contract method 0x95eebee6.
//
// Solidity: function checkpointTaskSuccesfullyChallenged(uint32 ) view returns(bool)
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerCallerSession) CheckpointTaskSuccesfullyChallenged(arg0 uint32) (bool, error) {
	return _ContractSFFLTaskManager.Contract.CheckpointTaskSuccesfullyChallenged(&_ContractSFFLTaskManager.CallOpts, arg0)
}

// Delegation is a free data retrieval call binding the contract method 0xdf5cf723.
//
// Solidity: function delegation() view returns(address)
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerCaller) Delegation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractSFFLTaskManager.contract.Call(opts, &out, "delegation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Delegation is a free data retrieval call binding the contract method 0xdf5cf723.
//
// Solidity: function delegation() view returns(address)
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerSession) Delegation() (common.Address, error) {
	return _ContractSFFLTaskManager.Contract.Delegation(&_ContractSFFLTaskManager.CallOpts)
}

// Delegation is a free data retrieval call binding the contract method 0xdf5cf723.
//
// Solidity: function delegation() view returns(address)
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerCallerSession) Delegation() (common.Address, error) {
	return _ContractSFFLTaskManager.Contract.Delegation(&_ContractSFFLTaskManager.CallOpts)
}

// Generator is a free data retrieval call binding the contract method 0x7afa1eed.
//
// Solidity: function generator() view returns(address)
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerCaller) Generator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractSFFLTaskManager.contract.Call(opts, &out, "generator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Generator is a free data retrieval call binding the contract method 0x7afa1eed.
//
// Solidity: function generator() view returns(address)
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerSession) Generator() (common.Address, error) {
	return _ContractSFFLTaskManager.Contract.Generator(&_ContractSFFLTaskManager.CallOpts)
}

// Generator is a free data retrieval call binding the contract method 0x7afa1eed.
//
// Solidity: function generator() view returns(address)
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerCallerSession) Generator() (common.Address, error) {
	return _ContractSFFLTaskManager.Contract.Generator(&_ContractSFFLTaskManager.CallOpts)
}

// GetCheckSignaturesIndices is a free data retrieval call binding the contract method 0x4f739f74.
//
// Solidity: function getCheckSignaturesIndices(address registryCoordinator, uint32 referenceBlockNumber, bytes quorumNumbers, bytes32[] nonSignerOperatorIds) view returns((uint32[],uint32[],uint32[],uint32[][]))
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerCaller) GetCheckSignaturesIndices(opts *bind.CallOpts, registryCoordinator common.Address, referenceBlockNumber uint32, quorumNumbers []byte, nonSignerOperatorIds [][32]byte) (OperatorStateRetrieverCheckSignaturesIndices, error) {
	var out []interface{}
	err := _ContractSFFLTaskManager.contract.Call(opts, &out, "getCheckSignaturesIndices", registryCoordinator, referenceBlockNumber, quorumNumbers, nonSignerOperatorIds)

	if err != nil {
		return *new(OperatorStateRetrieverCheckSignaturesIndices), err
	}

	out0 := *abi.ConvertType(out[0], new(OperatorStateRetrieverCheckSignaturesIndices)).(*OperatorStateRetrieverCheckSignaturesIndices)

	return out0, err

}

// GetCheckSignaturesIndices is a free data retrieval call binding the contract method 0x4f739f74.
//
// Solidity: function getCheckSignaturesIndices(address registryCoordinator, uint32 referenceBlockNumber, bytes quorumNumbers, bytes32[] nonSignerOperatorIds) view returns((uint32[],uint32[],uint32[],uint32[][]))
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerSession) GetCheckSignaturesIndices(registryCoordinator common.Address, referenceBlockNumber uint32, quorumNumbers []byte, nonSignerOperatorIds [][32]byte) (OperatorStateRetrieverCheckSignaturesIndices, error) {
	return _ContractSFFLTaskManager.Contract.GetCheckSignaturesIndices(&_ContractSFFLTaskManager.CallOpts, registryCoordinator, referenceBlockNumber, quorumNumbers, nonSignerOperatorIds)
}

// GetCheckSignaturesIndices is a free data retrieval call binding the contract method 0x4f739f74.
//
// Solidity: function getCheckSignaturesIndices(address registryCoordinator, uint32 referenceBlockNumber, bytes quorumNumbers, bytes32[] nonSignerOperatorIds) view returns((uint32[],uint32[],uint32[],uint32[][]))
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerCallerSession) GetCheckSignaturesIndices(registryCoordinator common.Address, referenceBlockNumber uint32, quorumNumbers []byte, nonSignerOperatorIds [][32]byte) (OperatorStateRetrieverCheckSignaturesIndices, error) {
	return _ContractSFFLTaskManager.Contract.GetCheckSignaturesIndices(&_ContractSFFLTaskManager.CallOpts, registryCoordinator, referenceBlockNumber, quorumNumbers, nonSignerOperatorIds)
}

// GetOperatorState is a free data retrieval call binding the contract method 0x3563b0d1.
//
// Solidity: function getOperatorState(address registryCoordinator, bytes quorumNumbers, uint32 blockNumber) view returns((address,bytes32,uint96)[][])
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerCaller) GetOperatorState(opts *bind.CallOpts, registryCoordinator common.Address, quorumNumbers []byte, blockNumber uint32) ([][]OperatorStateRetrieverOperator, error) {
	var out []interface{}
	err := _ContractSFFLTaskManager.contract.Call(opts, &out, "getOperatorState", registryCoordinator, quorumNumbers, blockNumber)

	if err != nil {
		return *new([][]OperatorStateRetrieverOperator), err
	}

	out0 := *abi.ConvertType(out[0], new([][]OperatorStateRetrieverOperator)).(*[][]OperatorStateRetrieverOperator)

	return out0, err

}

// GetOperatorState is a free data retrieval call binding the contract method 0x3563b0d1.
//
// Solidity: function getOperatorState(address registryCoordinator, bytes quorumNumbers, uint32 blockNumber) view returns((address,bytes32,uint96)[][])
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerSession) GetOperatorState(registryCoordinator common.Address, quorumNumbers []byte, blockNumber uint32) ([][]OperatorStateRetrieverOperator, error) {
	return _ContractSFFLTaskManager.Contract.GetOperatorState(&_ContractSFFLTaskManager.CallOpts, registryCoordinator, quorumNumbers, blockNumber)
}

// GetOperatorState is a free data retrieval call binding the contract method 0x3563b0d1.
//
// Solidity: function getOperatorState(address registryCoordinator, bytes quorumNumbers, uint32 blockNumber) view returns((address,bytes32,uint96)[][])
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerCallerSession) GetOperatorState(registryCoordinator common.Address, quorumNumbers []byte, blockNumber uint32) ([][]OperatorStateRetrieverOperator, error) {
	return _ContractSFFLTaskManager.Contract.GetOperatorState(&_ContractSFFLTaskManager.CallOpts, registryCoordinator, quorumNumbers, blockNumber)
}

// GetOperatorState0 is a free data retrieval call binding the contract method 0xcefdc1d4.
//
// Solidity: function getOperatorState(address registryCoordinator, bytes32 operatorId, uint32 blockNumber) view returns(uint256, (address,bytes32,uint96)[][])
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerCaller) GetOperatorState0(opts *bind.CallOpts, registryCoordinator common.Address, operatorId [32]byte, blockNumber uint32) (*big.Int, [][]OperatorStateRetrieverOperator, error) {
	var out []interface{}
	err := _ContractSFFLTaskManager.contract.Call(opts, &out, "getOperatorState0", registryCoordinator, operatorId, blockNumber)

	if err != nil {
		return *new(*big.Int), *new([][]OperatorStateRetrieverOperator), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new([][]OperatorStateRetrieverOperator)).(*[][]OperatorStateRetrieverOperator)

	return out0, out1, err

}

// GetOperatorState0 is a free data retrieval call binding the contract method 0xcefdc1d4.
//
// Solidity: function getOperatorState(address registryCoordinator, bytes32 operatorId, uint32 blockNumber) view returns(uint256, (address,bytes32,uint96)[][])
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerSession) GetOperatorState0(registryCoordinator common.Address, operatorId [32]byte, blockNumber uint32) (*big.Int, [][]OperatorStateRetrieverOperator, error) {
	return _ContractSFFLTaskManager.Contract.GetOperatorState0(&_ContractSFFLTaskManager.CallOpts, registryCoordinator, operatorId, blockNumber)
}

// GetOperatorState0 is a free data retrieval call binding the contract method 0xcefdc1d4.
//
// Solidity: function getOperatorState(address registryCoordinator, bytes32 operatorId, uint32 blockNumber) view returns(uint256, (address,bytes32,uint96)[][])
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerCallerSession) GetOperatorState0(registryCoordinator common.Address, operatorId [32]byte, blockNumber uint32) (*big.Int, [][]OperatorStateRetrieverOperator, error) {
	return _ContractSFFLTaskManager.Contract.GetOperatorState0(&_ContractSFFLTaskManager.CallOpts, registryCoordinator, operatorId, blockNumber)
}

// NextCheckpointTaskNum is a free data retrieval call binding the contract method 0x2e44b349.
//
// Solidity: function nextCheckpointTaskNum() view returns(uint32)
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerCaller) NextCheckpointTaskNum(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _ContractSFFLTaskManager.contract.Call(opts, &out, "nextCheckpointTaskNum")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// NextCheckpointTaskNum is a free data retrieval call binding the contract method 0x2e44b349.
//
// Solidity: function nextCheckpointTaskNum() view returns(uint32)
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerSession) NextCheckpointTaskNum() (uint32, error) {
	return _ContractSFFLTaskManager.Contract.NextCheckpointTaskNum(&_ContractSFFLTaskManager.CallOpts)
}

// NextCheckpointTaskNum is a free data retrieval call binding the contract method 0x2e44b349.
//
// Solidity: function nextCheckpointTaskNum() view returns(uint32)
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerCallerSession) NextCheckpointTaskNum() (uint32, error) {
	return _ContractSFFLTaskManager.Contract.NextCheckpointTaskNum(&_ContractSFFLTaskManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractSFFLTaskManager.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerSession) Owner() (common.Address, error) {
	return _ContractSFFLTaskManager.Contract.Owner(&_ContractSFFLTaskManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerCallerSession) Owner() (common.Address, error) {
	return _ContractSFFLTaskManager.Contract.Owner(&_ContractSFFLTaskManager.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerCaller) Paused(opts *bind.CallOpts, index uint8) (bool, error) {
	var out []interface{}
	err := _ContractSFFLTaskManager.contract.Call(opts, &out, "paused", index)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerSession) Paused(index uint8) (bool, error) {
	return _ContractSFFLTaskManager.Contract.Paused(&_ContractSFFLTaskManager.CallOpts, index)
}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerCallerSession) Paused(index uint8) (bool, error) {
	return _ContractSFFLTaskManager.Contract.Paused(&_ContractSFFLTaskManager.CallOpts, index)
}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerCaller) Paused0(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ContractSFFLTaskManager.contract.Call(opts, &out, "paused0")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerSession) Paused0() (*big.Int, error) {
	return _ContractSFFLTaskManager.Contract.Paused0(&_ContractSFFLTaskManager.CallOpts)
}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerCallerSession) Paused0() (*big.Int, error) {
	return _ContractSFFLTaskManager.Contract.Paused0(&_ContractSFFLTaskManager.CallOpts)
}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerCaller) PauserRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractSFFLTaskManager.contract.Call(opts, &out, "pauserRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerSession) PauserRegistry() (common.Address, error) {
	return _ContractSFFLTaskManager.Contract.PauserRegistry(&_ContractSFFLTaskManager.CallOpts)
}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerCallerSession) PauserRegistry() (common.Address, error) {
	return _ContractSFFLTaskManager.Contract.PauserRegistry(&_ContractSFFLTaskManager.CallOpts)
}

// RegistryCoordinator is a free data retrieval call binding the contract method 0x6d14a987.
//
// Solidity: function registryCoordinator() view returns(address)
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerCaller) RegistryCoordinator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractSFFLTaskManager.contract.Call(opts, &out, "registryCoordinator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RegistryCoordinator is a free data retrieval call binding the contract method 0x6d14a987.
//
// Solidity: function registryCoordinator() view returns(address)
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerSession) RegistryCoordinator() (common.Address, error) {
	return _ContractSFFLTaskManager.Contract.RegistryCoordinator(&_ContractSFFLTaskManager.CallOpts)
}

// RegistryCoordinator is a free data retrieval call binding the contract method 0x6d14a987.
//
// Solidity: function registryCoordinator() view returns(address)
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerCallerSession) RegistryCoordinator() (common.Address, error) {
	return _ContractSFFLTaskManager.Contract.RegistryCoordinator(&_ContractSFFLTaskManager.CallOpts)
}

// StakeRegistry is a free data retrieval call binding the contract method 0x68304835.
//
// Solidity: function stakeRegistry() view returns(address)
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerCaller) StakeRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractSFFLTaskManager.contract.Call(opts, &out, "stakeRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StakeRegistry is a free data retrieval call binding the contract method 0x68304835.
//
// Solidity: function stakeRegistry() view returns(address)
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerSession) StakeRegistry() (common.Address, error) {
	return _ContractSFFLTaskManager.Contract.StakeRegistry(&_ContractSFFLTaskManager.CallOpts)
}

// StakeRegistry is a free data retrieval call binding the contract method 0x68304835.
//
// Solidity: function stakeRegistry() view returns(address)
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerCallerSession) StakeRegistry() (common.Address, error) {
	return _ContractSFFLTaskManager.Contract.StakeRegistry(&_ContractSFFLTaskManager.CallOpts)
}

// StaleStakesForbidden is a free data retrieval call binding the contract method 0xb98d0908.
//
// Solidity: function staleStakesForbidden() view returns(bool)
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerCaller) StaleStakesForbidden(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ContractSFFLTaskManager.contract.Call(opts, &out, "staleStakesForbidden")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// StaleStakesForbidden is a free data retrieval call binding the contract method 0xb98d0908.
//
// Solidity: function staleStakesForbidden() view returns(bool)
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerSession) StaleStakesForbidden() (bool, error) {
	return _ContractSFFLTaskManager.Contract.StaleStakesForbidden(&_ContractSFFLTaskManager.CallOpts)
}

// StaleStakesForbidden is a free data retrieval call binding the contract method 0xb98d0908.
//
// Solidity: function staleStakesForbidden() view returns(bool)
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerCallerSession) StaleStakesForbidden() (bool, error) {
	return _ContractSFFLTaskManager.Contract.StaleStakesForbidden(&_ContractSFFLTaskManager.CallOpts)
}

// TrySignatureAndApkVerification is a free data retrieval call binding the contract method 0x171f1d5b.
//
// Solidity: function trySignatureAndApkVerification(bytes32 msgHash, (uint256,uint256) apk, (uint256[2],uint256[2]) apkG2, (uint256,uint256) sigma) view returns(bool pairingSuccessful, bool siganatureIsValid)
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerCaller) TrySignatureAndApkVerification(opts *bind.CallOpts, msgHash [32]byte, apk BN254G1Point, apkG2 BN254G2Point, sigma BN254G1Point) (struct {
	PairingSuccessful bool
	SiganatureIsValid bool
}, error) {
	var out []interface{}
	err := _ContractSFFLTaskManager.contract.Call(opts, &out, "trySignatureAndApkVerification", msgHash, apk, apkG2, sigma)

	outstruct := new(struct {
		PairingSuccessful bool
		SiganatureIsValid bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.PairingSuccessful = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.SiganatureIsValid = *abi.ConvertType(out[1], new(bool)).(*bool)

	return *outstruct, err

}

// TrySignatureAndApkVerification is a free data retrieval call binding the contract method 0x171f1d5b.
//
// Solidity: function trySignatureAndApkVerification(bytes32 msgHash, (uint256,uint256) apk, (uint256[2],uint256[2]) apkG2, (uint256,uint256) sigma) view returns(bool pairingSuccessful, bool siganatureIsValid)
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerSession) TrySignatureAndApkVerification(msgHash [32]byte, apk BN254G1Point, apkG2 BN254G2Point, sigma BN254G1Point) (struct {
	PairingSuccessful bool
	SiganatureIsValid bool
}, error) {
	return _ContractSFFLTaskManager.Contract.TrySignatureAndApkVerification(&_ContractSFFLTaskManager.CallOpts, msgHash, apk, apkG2, sigma)
}

// TrySignatureAndApkVerification is a free data retrieval call binding the contract method 0x171f1d5b.
//
// Solidity: function trySignatureAndApkVerification(bytes32 msgHash, (uint256,uint256) apk, (uint256[2],uint256[2]) apkG2, (uint256,uint256) sigma) view returns(bool pairingSuccessful, bool siganatureIsValid)
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerCallerSession) TrySignatureAndApkVerification(msgHash [32]byte, apk BN254G1Point, apkG2 BN254G2Point, sigma BN254G1Point) (struct {
	PairingSuccessful bool
	SiganatureIsValid bool
}, error) {
	return _ContractSFFLTaskManager.Contract.TrySignatureAndApkVerification(&_ContractSFFLTaskManager.CallOpts, msgHash, apk, apkG2, sigma)
}

// VerifyMessageInclusionState is a free data retrieval call binding the contract method 0xbdbf596c.
//
// Solidity: function verifyMessageInclusionState((uint32,uint64,uint64,bytes32) message, (uint32,bytes32,bytes32) taskResponse, (bytes32,uint256,uint256,bytes32[]) proof) pure returns(bool)
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerCaller) VerifyMessageInclusionState(opts *bind.CallOpts, message StateRootUpdateMessage, taskResponse CheckpointTaskResponse, proof SparseMerkleTreeProof) (bool, error) {
	var out []interface{}
	err := _ContractSFFLTaskManager.contract.Call(opts, &out, "verifyMessageInclusionState", message, taskResponse, proof)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyMessageInclusionState is a free data retrieval call binding the contract method 0xbdbf596c.
//
// Solidity: function verifyMessageInclusionState((uint32,uint64,uint64,bytes32) message, (uint32,bytes32,bytes32) taskResponse, (bytes32,uint256,uint256,bytes32[]) proof) pure returns(bool)
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerSession) VerifyMessageInclusionState(message StateRootUpdateMessage, taskResponse CheckpointTaskResponse, proof SparseMerkleTreeProof) (bool, error) {
	return _ContractSFFLTaskManager.Contract.VerifyMessageInclusionState(&_ContractSFFLTaskManager.CallOpts, message, taskResponse, proof)
}

// VerifyMessageInclusionState is a free data retrieval call binding the contract method 0xbdbf596c.
//
// Solidity: function verifyMessageInclusionState((uint32,uint64,uint64,bytes32) message, (uint32,bytes32,bytes32) taskResponse, (bytes32,uint256,uint256,bytes32[]) proof) pure returns(bool)
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerCallerSession) VerifyMessageInclusionState(message StateRootUpdateMessage, taskResponse CheckpointTaskResponse, proof SparseMerkleTreeProof) (bool, error) {
	return _ContractSFFLTaskManager.Contract.VerifyMessageInclusionState(&_ContractSFFLTaskManager.CallOpts, message, taskResponse, proof)
}

// VerifyMessageInclusionState0 is a free data retrieval call binding the contract method 0xf294e43c.
//
// Solidity: function verifyMessageInclusionState((uint64,uint64,((uint256,uint256),uint128)[]) message, (uint32,bytes32,bytes32) taskResponse, (bytes32,uint256,uint256,bytes32[]) proof) pure returns(bool)
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerCaller) VerifyMessageInclusionState0(opts *bind.CallOpts, message OperatorSetUpdateMessage, taskResponse CheckpointTaskResponse, proof SparseMerkleTreeProof) (bool, error) {
	var out []interface{}
	err := _ContractSFFLTaskManager.contract.Call(opts, &out, "verifyMessageInclusionState0", message, taskResponse, proof)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyMessageInclusionState0 is a free data retrieval call binding the contract method 0xf294e43c.
//
// Solidity: function verifyMessageInclusionState((uint64,uint64,((uint256,uint256),uint128)[]) message, (uint32,bytes32,bytes32) taskResponse, (bytes32,uint256,uint256,bytes32[]) proof) pure returns(bool)
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerSession) VerifyMessageInclusionState0(message OperatorSetUpdateMessage, taskResponse CheckpointTaskResponse, proof SparseMerkleTreeProof) (bool, error) {
	return _ContractSFFLTaskManager.Contract.VerifyMessageInclusionState0(&_ContractSFFLTaskManager.CallOpts, message, taskResponse, proof)
}

// VerifyMessageInclusionState0 is a free data retrieval call binding the contract method 0xf294e43c.
//
// Solidity: function verifyMessageInclusionState((uint64,uint64,((uint256,uint256),uint128)[]) message, (uint32,bytes32,bytes32) taskResponse, (bytes32,uint256,uint256,bytes32[]) proof) pure returns(bool)
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerCallerSession) VerifyMessageInclusionState0(message OperatorSetUpdateMessage, taskResponse CheckpointTaskResponse, proof SparseMerkleTreeProof) (bool, error) {
	return _ContractSFFLTaskManager.Contract.VerifyMessageInclusionState0(&_ContractSFFLTaskManager.CallOpts, message, taskResponse, proof)
}

// CreateCheckpointTask is a paid mutator transaction binding the contract method 0xefcf4edb.
//
// Solidity: function createCheckpointTask(uint64 fromTimestamp, uint64 toTimestamp, uint32 quorumThreshold, bytes quorumNumbers) returns()
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerTransactor) CreateCheckpointTask(opts *bind.TransactOpts, fromTimestamp uint64, toTimestamp uint64, quorumThreshold uint32, quorumNumbers []byte) (*types.Transaction, error) {
	return _ContractSFFLTaskManager.contract.Transact(opts, "createCheckpointTask", fromTimestamp, toTimestamp, quorumThreshold, quorumNumbers)
}

// CreateCheckpointTask is a paid mutator transaction binding the contract method 0xefcf4edb.
//
// Solidity: function createCheckpointTask(uint64 fromTimestamp, uint64 toTimestamp, uint32 quorumThreshold, bytes quorumNumbers) returns()
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerSession) CreateCheckpointTask(fromTimestamp uint64, toTimestamp uint64, quorumThreshold uint32, quorumNumbers []byte) (*types.Transaction, error) {
	return _ContractSFFLTaskManager.Contract.CreateCheckpointTask(&_ContractSFFLTaskManager.TransactOpts, fromTimestamp, toTimestamp, quorumThreshold, quorumNumbers)
}

// CreateCheckpointTask is a paid mutator transaction binding the contract method 0xefcf4edb.
//
// Solidity: function createCheckpointTask(uint64 fromTimestamp, uint64 toTimestamp, uint32 quorumThreshold, bytes quorumNumbers) returns()
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerTransactorSession) CreateCheckpointTask(fromTimestamp uint64, toTimestamp uint64, quorumThreshold uint32, quorumNumbers []byte) (*types.Transaction, error) {
	return _ContractSFFLTaskManager.Contract.CreateCheckpointTask(&_ContractSFFLTaskManager.TransactOpts, fromTimestamp, toTimestamp, quorumThreshold, quorumNumbers)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address _pauserRegistry, address initialOwner, address _aggregator, address _generator) returns()
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerTransactor) Initialize(opts *bind.TransactOpts, _pauserRegistry common.Address, initialOwner common.Address, _aggregator common.Address, _generator common.Address) (*types.Transaction, error) {
	return _ContractSFFLTaskManager.contract.Transact(opts, "initialize", _pauserRegistry, initialOwner, _aggregator, _generator)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address _pauserRegistry, address initialOwner, address _aggregator, address _generator) returns()
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerSession) Initialize(_pauserRegistry common.Address, initialOwner common.Address, _aggregator common.Address, _generator common.Address) (*types.Transaction, error) {
	return _ContractSFFLTaskManager.Contract.Initialize(&_ContractSFFLTaskManager.TransactOpts, _pauserRegistry, initialOwner, _aggregator, _generator)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address _pauserRegistry, address initialOwner, address _aggregator, address _generator) returns()
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerTransactorSession) Initialize(_pauserRegistry common.Address, initialOwner common.Address, _aggregator common.Address, _generator common.Address) (*types.Transaction, error) {
	return _ContractSFFLTaskManager.Contract.Initialize(&_ContractSFFLTaskManager.TransactOpts, _pauserRegistry, initialOwner, _aggregator, _generator)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerTransactor) Pause(opts *bind.TransactOpts, newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractSFFLTaskManager.contract.Transact(opts, "pause", newPausedStatus)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerSession) Pause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractSFFLTaskManager.Contract.Pause(&_ContractSFFLTaskManager.TransactOpts, newPausedStatus)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerTransactorSession) Pause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractSFFLTaskManager.Contract.Pause(&_ContractSFFLTaskManager.TransactOpts, newPausedStatus)
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerTransactor) PauseAll(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractSFFLTaskManager.contract.Transact(opts, "pauseAll")
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerSession) PauseAll() (*types.Transaction, error) {
	return _ContractSFFLTaskManager.Contract.PauseAll(&_ContractSFFLTaskManager.TransactOpts)
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerTransactorSession) PauseAll() (*types.Transaction, error) {
	return _ContractSFFLTaskManager.Contract.PauseAll(&_ContractSFFLTaskManager.TransactOpts)
}

// RaiseAndResolveCheckpointChallenge is a paid mutator transaction binding the contract method 0x5ace2df7.
//
// Solidity: function raiseAndResolveCheckpointChallenge((uint32,uint64,uint64,uint32,bytes) task, (uint32,bytes32,bytes32) taskResponse, (uint32,bytes32) taskResponseMetadata, (uint256,uint256)[] pubkeysOfNonSigningOperators) returns()
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerTransactor) RaiseAndResolveCheckpointChallenge(opts *bind.TransactOpts, task CheckpointTask, taskResponse CheckpointTaskResponse, taskResponseMetadata CheckpointTaskResponseMetadata, pubkeysOfNonSigningOperators []BN254G1Point) (*types.Transaction, error) {
	return _ContractSFFLTaskManager.contract.Transact(opts, "raiseAndResolveCheckpointChallenge", task, taskResponse, taskResponseMetadata, pubkeysOfNonSigningOperators)
}

// RaiseAndResolveCheckpointChallenge is a paid mutator transaction binding the contract method 0x5ace2df7.
//
// Solidity: function raiseAndResolveCheckpointChallenge((uint32,uint64,uint64,uint32,bytes) task, (uint32,bytes32,bytes32) taskResponse, (uint32,bytes32) taskResponseMetadata, (uint256,uint256)[] pubkeysOfNonSigningOperators) returns()
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerSession) RaiseAndResolveCheckpointChallenge(task CheckpointTask, taskResponse CheckpointTaskResponse, taskResponseMetadata CheckpointTaskResponseMetadata, pubkeysOfNonSigningOperators []BN254G1Point) (*types.Transaction, error) {
	return _ContractSFFLTaskManager.Contract.RaiseAndResolveCheckpointChallenge(&_ContractSFFLTaskManager.TransactOpts, task, taskResponse, taskResponseMetadata, pubkeysOfNonSigningOperators)
}

// RaiseAndResolveCheckpointChallenge is a paid mutator transaction binding the contract method 0x5ace2df7.
//
// Solidity: function raiseAndResolveCheckpointChallenge((uint32,uint64,uint64,uint32,bytes) task, (uint32,bytes32,bytes32) taskResponse, (uint32,bytes32) taskResponseMetadata, (uint256,uint256)[] pubkeysOfNonSigningOperators) returns()
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerTransactorSession) RaiseAndResolveCheckpointChallenge(task CheckpointTask, taskResponse CheckpointTaskResponse, taskResponseMetadata CheckpointTaskResponseMetadata, pubkeysOfNonSigningOperators []BN254G1Point) (*types.Transaction, error) {
	return _ContractSFFLTaskManager.Contract.RaiseAndResolveCheckpointChallenge(&_ContractSFFLTaskManager.TransactOpts, task, taskResponse, taskResponseMetadata, pubkeysOfNonSigningOperators)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractSFFLTaskManager.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerSession) RenounceOwnership() (*types.Transaction, error) {
	return _ContractSFFLTaskManager.Contract.RenounceOwnership(&_ContractSFFLTaskManager.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ContractSFFLTaskManager.Contract.RenounceOwnership(&_ContractSFFLTaskManager.TransactOpts)
}

// RespondToCheckpointTask is a paid mutator transaction binding the contract method 0xda16491f.
//
// Solidity: function respondToCheckpointTask((uint32,uint64,uint64,uint32,bytes) task, (uint32,bytes32,bytes32) taskResponse, (uint32[],(uint256,uint256)[],(uint256,uint256)[],(uint256[2],uint256[2]),(uint256,uint256),uint32[],uint32[],uint32[][]) nonSignerStakesAndSignature) returns()
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerTransactor) RespondToCheckpointTask(opts *bind.TransactOpts, task CheckpointTask, taskResponse CheckpointTaskResponse, nonSignerStakesAndSignature IBLSSignatureCheckerNonSignerStakesAndSignature) (*types.Transaction, error) {
	return _ContractSFFLTaskManager.contract.Transact(opts, "respondToCheckpointTask", task, taskResponse, nonSignerStakesAndSignature)
}

// RespondToCheckpointTask is a paid mutator transaction binding the contract method 0xda16491f.
//
// Solidity: function respondToCheckpointTask((uint32,uint64,uint64,uint32,bytes) task, (uint32,bytes32,bytes32) taskResponse, (uint32[],(uint256,uint256)[],(uint256,uint256)[],(uint256[2],uint256[2]),(uint256,uint256),uint32[],uint32[],uint32[][]) nonSignerStakesAndSignature) returns()
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerSession) RespondToCheckpointTask(task CheckpointTask, taskResponse CheckpointTaskResponse, nonSignerStakesAndSignature IBLSSignatureCheckerNonSignerStakesAndSignature) (*types.Transaction, error) {
	return _ContractSFFLTaskManager.Contract.RespondToCheckpointTask(&_ContractSFFLTaskManager.TransactOpts, task, taskResponse, nonSignerStakesAndSignature)
}

// RespondToCheckpointTask is a paid mutator transaction binding the contract method 0xda16491f.
//
// Solidity: function respondToCheckpointTask((uint32,uint64,uint64,uint32,bytes) task, (uint32,bytes32,bytes32) taskResponse, (uint32[],(uint256,uint256)[],(uint256,uint256)[],(uint256[2],uint256[2]),(uint256,uint256),uint32[],uint32[],uint32[][]) nonSignerStakesAndSignature) returns()
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerTransactorSession) RespondToCheckpointTask(task CheckpointTask, taskResponse CheckpointTaskResponse, nonSignerStakesAndSignature IBLSSignatureCheckerNonSignerStakesAndSignature) (*types.Transaction, error) {
	return _ContractSFFLTaskManager.Contract.RespondToCheckpointTask(&_ContractSFFLTaskManager.TransactOpts, task, taskResponse, nonSignerStakesAndSignature)
}

// SetPauserRegistry is a paid mutator transaction binding the contract method 0x10d67a2f.
//
// Solidity: function setPauserRegistry(address newPauserRegistry) returns()
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerTransactor) SetPauserRegistry(opts *bind.TransactOpts, newPauserRegistry common.Address) (*types.Transaction, error) {
	return _ContractSFFLTaskManager.contract.Transact(opts, "setPauserRegistry", newPauserRegistry)
}

// SetPauserRegistry is a paid mutator transaction binding the contract method 0x10d67a2f.
//
// Solidity: function setPauserRegistry(address newPauserRegistry) returns()
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerSession) SetPauserRegistry(newPauserRegistry common.Address) (*types.Transaction, error) {
	return _ContractSFFLTaskManager.Contract.SetPauserRegistry(&_ContractSFFLTaskManager.TransactOpts, newPauserRegistry)
}

// SetPauserRegistry is a paid mutator transaction binding the contract method 0x10d67a2f.
//
// Solidity: function setPauserRegistry(address newPauserRegistry) returns()
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerTransactorSession) SetPauserRegistry(newPauserRegistry common.Address) (*types.Transaction, error) {
	return _ContractSFFLTaskManager.Contract.SetPauserRegistry(&_ContractSFFLTaskManager.TransactOpts, newPauserRegistry)
}

// SetStaleStakesForbidden is a paid mutator transaction binding the contract method 0x416c7e5e.
//
// Solidity: function setStaleStakesForbidden(bool value) returns()
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerTransactor) SetStaleStakesForbidden(opts *bind.TransactOpts, value bool) (*types.Transaction, error) {
	return _ContractSFFLTaskManager.contract.Transact(opts, "setStaleStakesForbidden", value)
}

// SetStaleStakesForbidden is a paid mutator transaction binding the contract method 0x416c7e5e.
//
// Solidity: function setStaleStakesForbidden(bool value) returns()
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerSession) SetStaleStakesForbidden(value bool) (*types.Transaction, error) {
	return _ContractSFFLTaskManager.Contract.SetStaleStakesForbidden(&_ContractSFFLTaskManager.TransactOpts, value)
}

// SetStaleStakesForbidden is a paid mutator transaction binding the contract method 0x416c7e5e.
//
// Solidity: function setStaleStakesForbidden(bool value) returns()
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerTransactorSession) SetStaleStakesForbidden(value bool) (*types.Transaction, error) {
	return _ContractSFFLTaskManager.Contract.SetStaleStakesForbidden(&_ContractSFFLTaskManager.TransactOpts, value)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ContractSFFLTaskManager.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ContractSFFLTaskManager.Contract.TransferOwnership(&_ContractSFFLTaskManager.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ContractSFFLTaskManager.Contract.TransferOwnership(&_ContractSFFLTaskManager.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerTransactor) Unpause(opts *bind.TransactOpts, newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractSFFLTaskManager.contract.Transact(opts, "unpause", newPausedStatus)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerSession) Unpause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractSFFLTaskManager.Contract.Unpause(&_ContractSFFLTaskManager.TransactOpts, newPausedStatus)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerTransactorSession) Unpause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractSFFLTaskManager.Contract.Unpause(&_ContractSFFLTaskManager.TransactOpts, newPausedStatus)
}

// ContractSFFLTaskManagerCheckpointTaskChallengedSuccessfullyIterator is returned from FilterCheckpointTaskChallengedSuccessfully and is used to iterate over the raw logs and unpacked data for CheckpointTaskChallengedSuccessfully events raised by the ContractSFFLTaskManager contract.
type ContractSFFLTaskManagerCheckpointTaskChallengedSuccessfullyIterator struct {
	Event *ContractSFFLTaskManagerCheckpointTaskChallengedSuccessfully // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractSFFLTaskManagerCheckpointTaskChallengedSuccessfullyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractSFFLTaskManagerCheckpointTaskChallengedSuccessfully)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractSFFLTaskManagerCheckpointTaskChallengedSuccessfully)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractSFFLTaskManagerCheckpointTaskChallengedSuccessfullyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractSFFLTaskManagerCheckpointTaskChallengedSuccessfullyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractSFFLTaskManagerCheckpointTaskChallengedSuccessfully represents a CheckpointTaskChallengedSuccessfully event raised by the ContractSFFLTaskManager contract.
type ContractSFFLTaskManagerCheckpointTaskChallengedSuccessfully struct {
	TaskIndex  uint32
	Challenger common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterCheckpointTaskChallengedSuccessfully is a free log retrieval operation binding the contract event 0xff48388ad5e2a6d1845a7672040fba7d9b14b22b9e0eecd37046e5313d3aebc2.
//
// Solidity: event CheckpointTaskChallengedSuccessfully(uint32 indexed taskIndex, address indexed challenger)
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerFilterer) FilterCheckpointTaskChallengedSuccessfully(opts *bind.FilterOpts, taskIndex []uint32, challenger []common.Address) (*ContractSFFLTaskManagerCheckpointTaskChallengedSuccessfullyIterator, error) {

	var taskIndexRule []interface{}
	for _, taskIndexItem := range taskIndex {
		taskIndexRule = append(taskIndexRule, taskIndexItem)
	}
	var challengerRule []interface{}
	for _, challengerItem := range challenger {
		challengerRule = append(challengerRule, challengerItem)
	}

	logs, sub, err := _ContractSFFLTaskManager.contract.FilterLogs(opts, "CheckpointTaskChallengedSuccessfully", taskIndexRule, challengerRule)
	if err != nil {
		return nil, err
	}
	return &ContractSFFLTaskManagerCheckpointTaskChallengedSuccessfullyIterator{contract: _ContractSFFLTaskManager.contract, event: "CheckpointTaskChallengedSuccessfully", logs: logs, sub: sub}, nil
}

// WatchCheckpointTaskChallengedSuccessfully is a free log subscription operation binding the contract event 0xff48388ad5e2a6d1845a7672040fba7d9b14b22b9e0eecd37046e5313d3aebc2.
//
// Solidity: event CheckpointTaskChallengedSuccessfully(uint32 indexed taskIndex, address indexed challenger)
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerFilterer) WatchCheckpointTaskChallengedSuccessfully(opts *bind.WatchOpts, sink chan<- *ContractSFFLTaskManagerCheckpointTaskChallengedSuccessfully, taskIndex []uint32, challenger []common.Address) (event.Subscription, error) {

	var taskIndexRule []interface{}
	for _, taskIndexItem := range taskIndex {
		taskIndexRule = append(taskIndexRule, taskIndexItem)
	}
	var challengerRule []interface{}
	for _, challengerItem := range challenger {
		challengerRule = append(challengerRule, challengerItem)
	}

	logs, sub, err := _ContractSFFLTaskManager.contract.WatchLogs(opts, "CheckpointTaskChallengedSuccessfully", taskIndexRule, challengerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractSFFLTaskManagerCheckpointTaskChallengedSuccessfully)
				if err := _ContractSFFLTaskManager.contract.UnpackLog(event, "CheckpointTaskChallengedSuccessfully", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseCheckpointTaskChallengedSuccessfully is a log parse operation binding the contract event 0xff48388ad5e2a6d1845a7672040fba7d9b14b22b9e0eecd37046e5313d3aebc2.
//
// Solidity: event CheckpointTaskChallengedSuccessfully(uint32 indexed taskIndex, address indexed challenger)
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerFilterer) ParseCheckpointTaskChallengedSuccessfully(log types.Log) (*ContractSFFLTaskManagerCheckpointTaskChallengedSuccessfully, error) {
	event := new(ContractSFFLTaskManagerCheckpointTaskChallengedSuccessfully)
	if err := _ContractSFFLTaskManager.contract.UnpackLog(event, "CheckpointTaskChallengedSuccessfully", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractSFFLTaskManagerCheckpointTaskChallengedUnsuccessfullyIterator is returned from FilterCheckpointTaskChallengedUnsuccessfully and is used to iterate over the raw logs and unpacked data for CheckpointTaskChallengedUnsuccessfully events raised by the ContractSFFLTaskManager contract.
type ContractSFFLTaskManagerCheckpointTaskChallengedUnsuccessfullyIterator struct {
	Event *ContractSFFLTaskManagerCheckpointTaskChallengedUnsuccessfully // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractSFFLTaskManagerCheckpointTaskChallengedUnsuccessfullyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractSFFLTaskManagerCheckpointTaskChallengedUnsuccessfully)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractSFFLTaskManagerCheckpointTaskChallengedUnsuccessfully)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractSFFLTaskManagerCheckpointTaskChallengedUnsuccessfullyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractSFFLTaskManagerCheckpointTaskChallengedUnsuccessfullyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractSFFLTaskManagerCheckpointTaskChallengedUnsuccessfully represents a CheckpointTaskChallengedUnsuccessfully event raised by the ContractSFFLTaskManager contract.
type ContractSFFLTaskManagerCheckpointTaskChallengedUnsuccessfully struct {
	TaskIndex  uint32
	Challenger common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterCheckpointTaskChallengedUnsuccessfully is a free log retrieval operation binding the contract event 0x0c6923c4a98292e75c5d677a1634527f87b6d19cf2c7d396aece99790c44a795.
//
// Solidity: event CheckpointTaskChallengedUnsuccessfully(uint32 indexed taskIndex, address indexed challenger)
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerFilterer) FilterCheckpointTaskChallengedUnsuccessfully(opts *bind.FilterOpts, taskIndex []uint32, challenger []common.Address) (*ContractSFFLTaskManagerCheckpointTaskChallengedUnsuccessfullyIterator, error) {

	var taskIndexRule []interface{}
	for _, taskIndexItem := range taskIndex {
		taskIndexRule = append(taskIndexRule, taskIndexItem)
	}
	var challengerRule []interface{}
	for _, challengerItem := range challenger {
		challengerRule = append(challengerRule, challengerItem)
	}

	logs, sub, err := _ContractSFFLTaskManager.contract.FilterLogs(opts, "CheckpointTaskChallengedUnsuccessfully", taskIndexRule, challengerRule)
	if err != nil {
		return nil, err
	}
	return &ContractSFFLTaskManagerCheckpointTaskChallengedUnsuccessfullyIterator{contract: _ContractSFFLTaskManager.contract, event: "CheckpointTaskChallengedUnsuccessfully", logs: logs, sub: sub}, nil
}

// WatchCheckpointTaskChallengedUnsuccessfully is a free log subscription operation binding the contract event 0x0c6923c4a98292e75c5d677a1634527f87b6d19cf2c7d396aece99790c44a795.
//
// Solidity: event CheckpointTaskChallengedUnsuccessfully(uint32 indexed taskIndex, address indexed challenger)
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerFilterer) WatchCheckpointTaskChallengedUnsuccessfully(opts *bind.WatchOpts, sink chan<- *ContractSFFLTaskManagerCheckpointTaskChallengedUnsuccessfully, taskIndex []uint32, challenger []common.Address) (event.Subscription, error) {

	var taskIndexRule []interface{}
	for _, taskIndexItem := range taskIndex {
		taskIndexRule = append(taskIndexRule, taskIndexItem)
	}
	var challengerRule []interface{}
	for _, challengerItem := range challenger {
		challengerRule = append(challengerRule, challengerItem)
	}

	logs, sub, err := _ContractSFFLTaskManager.contract.WatchLogs(opts, "CheckpointTaskChallengedUnsuccessfully", taskIndexRule, challengerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractSFFLTaskManagerCheckpointTaskChallengedUnsuccessfully)
				if err := _ContractSFFLTaskManager.contract.UnpackLog(event, "CheckpointTaskChallengedUnsuccessfully", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseCheckpointTaskChallengedUnsuccessfully is a log parse operation binding the contract event 0x0c6923c4a98292e75c5d677a1634527f87b6d19cf2c7d396aece99790c44a795.
//
// Solidity: event CheckpointTaskChallengedUnsuccessfully(uint32 indexed taskIndex, address indexed challenger)
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerFilterer) ParseCheckpointTaskChallengedUnsuccessfully(log types.Log) (*ContractSFFLTaskManagerCheckpointTaskChallengedUnsuccessfully, error) {
	event := new(ContractSFFLTaskManagerCheckpointTaskChallengedUnsuccessfully)
	if err := _ContractSFFLTaskManager.contract.UnpackLog(event, "CheckpointTaskChallengedUnsuccessfully", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractSFFLTaskManagerCheckpointTaskCreatedIterator is returned from FilterCheckpointTaskCreated and is used to iterate over the raw logs and unpacked data for CheckpointTaskCreated events raised by the ContractSFFLTaskManager contract.
type ContractSFFLTaskManagerCheckpointTaskCreatedIterator struct {
	Event *ContractSFFLTaskManagerCheckpointTaskCreated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractSFFLTaskManagerCheckpointTaskCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractSFFLTaskManagerCheckpointTaskCreated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractSFFLTaskManagerCheckpointTaskCreated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractSFFLTaskManagerCheckpointTaskCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractSFFLTaskManagerCheckpointTaskCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractSFFLTaskManagerCheckpointTaskCreated represents a CheckpointTaskCreated event raised by the ContractSFFLTaskManager contract.
type ContractSFFLTaskManagerCheckpointTaskCreated struct {
	TaskIndex uint32
	Task      CheckpointTask
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterCheckpointTaskCreated is a free log retrieval operation binding the contract event 0x78aec7310ea6fd468e3d3bbd16a806fd4987515634d5b5bf4cf4f036d9c33225.
//
// Solidity: event CheckpointTaskCreated(uint32 indexed taskIndex, (uint32,uint64,uint64,uint32,bytes) task)
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerFilterer) FilterCheckpointTaskCreated(opts *bind.FilterOpts, taskIndex []uint32) (*ContractSFFLTaskManagerCheckpointTaskCreatedIterator, error) {

	var taskIndexRule []interface{}
	for _, taskIndexItem := range taskIndex {
		taskIndexRule = append(taskIndexRule, taskIndexItem)
	}

	logs, sub, err := _ContractSFFLTaskManager.contract.FilterLogs(opts, "CheckpointTaskCreated", taskIndexRule)
	if err != nil {
		return nil, err
	}
	return &ContractSFFLTaskManagerCheckpointTaskCreatedIterator{contract: _ContractSFFLTaskManager.contract, event: "CheckpointTaskCreated", logs: logs, sub: sub}, nil
}

// WatchCheckpointTaskCreated is a free log subscription operation binding the contract event 0x78aec7310ea6fd468e3d3bbd16a806fd4987515634d5b5bf4cf4f036d9c33225.
//
// Solidity: event CheckpointTaskCreated(uint32 indexed taskIndex, (uint32,uint64,uint64,uint32,bytes) task)
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerFilterer) WatchCheckpointTaskCreated(opts *bind.WatchOpts, sink chan<- *ContractSFFLTaskManagerCheckpointTaskCreated, taskIndex []uint32) (event.Subscription, error) {

	var taskIndexRule []interface{}
	for _, taskIndexItem := range taskIndex {
		taskIndexRule = append(taskIndexRule, taskIndexItem)
	}

	logs, sub, err := _ContractSFFLTaskManager.contract.WatchLogs(opts, "CheckpointTaskCreated", taskIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractSFFLTaskManagerCheckpointTaskCreated)
				if err := _ContractSFFLTaskManager.contract.UnpackLog(event, "CheckpointTaskCreated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseCheckpointTaskCreated is a log parse operation binding the contract event 0x78aec7310ea6fd468e3d3bbd16a806fd4987515634d5b5bf4cf4f036d9c33225.
//
// Solidity: event CheckpointTaskCreated(uint32 indexed taskIndex, (uint32,uint64,uint64,uint32,bytes) task)
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerFilterer) ParseCheckpointTaskCreated(log types.Log) (*ContractSFFLTaskManagerCheckpointTaskCreated, error) {
	event := new(ContractSFFLTaskManagerCheckpointTaskCreated)
	if err := _ContractSFFLTaskManager.contract.UnpackLog(event, "CheckpointTaskCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractSFFLTaskManagerCheckpointTaskRespondedIterator is returned from FilterCheckpointTaskResponded and is used to iterate over the raw logs and unpacked data for CheckpointTaskResponded events raised by the ContractSFFLTaskManager contract.
type ContractSFFLTaskManagerCheckpointTaskRespondedIterator struct {
	Event *ContractSFFLTaskManagerCheckpointTaskResponded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractSFFLTaskManagerCheckpointTaskRespondedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractSFFLTaskManagerCheckpointTaskResponded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractSFFLTaskManagerCheckpointTaskResponded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractSFFLTaskManagerCheckpointTaskRespondedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractSFFLTaskManagerCheckpointTaskRespondedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractSFFLTaskManagerCheckpointTaskResponded represents a CheckpointTaskResponded event raised by the ContractSFFLTaskManager contract.
type ContractSFFLTaskManagerCheckpointTaskResponded struct {
	TaskResponse         CheckpointTaskResponse
	TaskResponseMetadata CheckpointTaskResponseMetadata
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterCheckpointTaskResponded is a free log retrieval operation binding the contract event 0x8016fcc5ad5dcf12fff2e128d239d9c6eb61f4041126bbac2c93fa8962627c1b.
//
// Solidity: event CheckpointTaskResponded((uint32,bytes32,bytes32) taskResponse, (uint32,bytes32) taskResponseMetadata)
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerFilterer) FilterCheckpointTaskResponded(opts *bind.FilterOpts) (*ContractSFFLTaskManagerCheckpointTaskRespondedIterator, error) {

	logs, sub, err := _ContractSFFLTaskManager.contract.FilterLogs(opts, "CheckpointTaskResponded")
	if err != nil {
		return nil, err
	}
	return &ContractSFFLTaskManagerCheckpointTaskRespondedIterator{contract: _ContractSFFLTaskManager.contract, event: "CheckpointTaskResponded", logs: logs, sub: sub}, nil
}

// WatchCheckpointTaskResponded is a free log subscription operation binding the contract event 0x8016fcc5ad5dcf12fff2e128d239d9c6eb61f4041126bbac2c93fa8962627c1b.
//
// Solidity: event CheckpointTaskResponded((uint32,bytes32,bytes32) taskResponse, (uint32,bytes32) taskResponseMetadata)
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerFilterer) WatchCheckpointTaskResponded(opts *bind.WatchOpts, sink chan<- *ContractSFFLTaskManagerCheckpointTaskResponded) (event.Subscription, error) {

	logs, sub, err := _ContractSFFLTaskManager.contract.WatchLogs(opts, "CheckpointTaskResponded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractSFFLTaskManagerCheckpointTaskResponded)
				if err := _ContractSFFLTaskManager.contract.UnpackLog(event, "CheckpointTaskResponded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseCheckpointTaskResponded is a log parse operation binding the contract event 0x8016fcc5ad5dcf12fff2e128d239d9c6eb61f4041126bbac2c93fa8962627c1b.
//
// Solidity: event CheckpointTaskResponded((uint32,bytes32,bytes32) taskResponse, (uint32,bytes32) taskResponseMetadata)
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerFilterer) ParseCheckpointTaskResponded(log types.Log) (*ContractSFFLTaskManagerCheckpointTaskResponded, error) {
	event := new(ContractSFFLTaskManagerCheckpointTaskResponded)
	if err := _ContractSFFLTaskManager.contract.UnpackLog(event, "CheckpointTaskResponded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractSFFLTaskManagerInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the ContractSFFLTaskManager contract.
type ContractSFFLTaskManagerInitializedIterator struct {
	Event *ContractSFFLTaskManagerInitialized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractSFFLTaskManagerInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractSFFLTaskManagerInitialized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractSFFLTaskManagerInitialized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractSFFLTaskManagerInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractSFFLTaskManagerInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractSFFLTaskManagerInitialized represents a Initialized event raised by the ContractSFFLTaskManager contract.
type ContractSFFLTaskManagerInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerFilterer) FilterInitialized(opts *bind.FilterOpts) (*ContractSFFLTaskManagerInitializedIterator, error) {

	logs, sub, err := _ContractSFFLTaskManager.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ContractSFFLTaskManagerInitializedIterator{contract: _ContractSFFLTaskManager.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ContractSFFLTaskManagerInitialized) (event.Subscription, error) {

	logs, sub, err := _ContractSFFLTaskManager.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractSFFLTaskManagerInitialized)
				if err := _ContractSFFLTaskManager.contract.UnpackLog(event, "Initialized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerFilterer) ParseInitialized(log types.Log) (*ContractSFFLTaskManagerInitialized, error) {
	event := new(ContractSFFLTaskManagerInitialized)
	if err := _ContractSFFLTaskManager.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractSFFLTaskManagerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ContractSFFLTaskManager contract.
type ContractSFFLTaskManagerOwnershipTransferredIterator struct {
	Event *ContractSFFLTaskManagerOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractSFFLTaskManagerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractSFFLTaskManagerOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractSFFLTaskManagerOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractSFFLTaskManagerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractSFFLTaskManagerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractSFFLTaskManagerOwnershipTransferred represents a OwnershipTransferred event raised by the ContractSFFLTaskManager contract.
type ContractSFFLTaskManagerOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ContractSFFLTaskManagerOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ContractSFFLTaskManager.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ContractSFFLTaskManagerOwnershipTransferredIterator{contract: _ContractSFFLTaskManager.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ContractSFFLTaskManagerOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ContractSFFLTaskManager.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractSFFLTaskManagerOwnershipTransferred)
				if err := _ContractSFFLTaskManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerFilterer) ParseOwnershipTransferred(log types.Log) (*ContractSFFLTaskManagerOwnershipTransferred, error) {
	event := new(ContractSFFLTaskManagerOwnershipTransferred)
	if err := _ContractSFFLTaskManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractSFFLTaskManagerPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the ContractSFFLTaskManager contract.
type ContractSFFLTaskManagerPausedIterator struct {
	Event *ContractSFFLTaskManagerPaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractSFFLTaskManagerPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractSFFLTaskManagerPaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractSFFLTaskManagerPaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractSFFLTaskManagerPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractSFFLTaskManagerPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractSFFLTaskManagerPaused represents a Paused event raised by the ContractSFFLTaskManager contract.
type ContractSFFLTaskManagerPaused struct {
	Account         common.Address
	NewPausedStatus *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0xab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d.
//
// Solidity: event Paused(address indexed account, uint256 newPausedStatus)
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerFilterer) FilterPaused(opts *bind.FilterOpts, account []common.Address) (*ContractSFFLTaskManagerPausedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ContractSFFLTaskManager.contract.FilterLogs(opts, "Paused", accountRule)
	if err != nil {
		return nil, err
	}
	return &ContractSFFLTaskManagerPausedIterator{contract: _ContractSFFLTaskManager.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0xab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d.
//
// Solidity: event Paused(address indexed account, uint256 newPausedStatus)
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *ContractSFFLTaskManagerPaused, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ContractSFFLTaskManager.contract.WatchLogs(opts, "Paused", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractSFFLTaskManagerPaused)
				if err := _ContractSFFLTaskManager.contract.UnpackLog(event, "Paused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePaused is a log parse operation binding the contract event 0xab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d.
//
// Solidity: event Paused(address indexed account, uint256 newPausedStatus)
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerFilterer) ParsePaused(log types.Log) (*ContractSFFLTaskManagerPaused, error) {
	event := new(ContractSFFLTaskManagerPaused)
	if err := _ContractSFFLTaskManager.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractSFFLTaskManagerPauserRegistrySetIterator is returned from FilterPauserRegistrySet and is used to iterate over the raw logs and unpacked data for PauserRegistrySet events raised by the ContractSFFLTaskManager contract.
type ContractSFFLTaskManagerPauserRegistrySetIterator struct {
	Event *ContractSFFLTaskManagerPauserRegistrySet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractSFFLTaskManagerPauserRegistrySetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractSFFLTaskManagerPauserRegistrySet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractSFFLTaskManagerPauserRegistrySet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractSFFLTaskManagerPauserRegistrySetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractSFFLTaskManagerPauserRegistrySetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractSFFLTaskManagerPauserRegistrySet represents a PauserRegistrySet event raised by the ContractSFFLTaskManager contract.
type ContractSFFLTaskManagerPauserRegistrySet struct {
	PauserRegistry    common.Address
	NewPauserRegistry common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterPauserRegistrySet is a free log retrieval operation binding the contract event 0x6e9fcd539896fca60e8b0f01dd580233e48a6b0f7df013b89ba7f565869acdb6.
//
// Solidity: event PauserRegistrySet(address pauserRegistry, address newPauserRegistry)
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerFilterer) FilterPauserRegistrySet(opts *bind.FilterOpts) (*ContractSFFLTaskManagerPauserRegistrySetIterator, error) {

	logs, sub, err := _ContractSFFLTaskManager.contract.FilterLogs(opts, "PauserRegistrySet")
	if err != nil {
		return nil, err
	}
	return &ContractSFFLTaskManagerPauserRegistrySetIterator{contract: _ContractSFFLTaskManager.contract, event: "PauserRegistrySet", logs: logs, sub: sub}, nil
}

// WatchPauserRegistrySet is a free log subscription operation binding the contract event 0x6e9fcd539896fca60e8b0f01dd580233e48a6b0f7df013b89ba7f565869acdb6.
//
// Solidity: event PauserRegistrySet(address pauserRegistry, address newPauserRegistry)
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerFilterer) WatchPauserRegistrySet(opts *bind.WatchOpts, sink chan<- *ContractSFFLTaskManagerPauserRegistrySet) (event.Subscription, error) {

	logs, sub, err := _ContractSFFLTaskManager.contract.WatchLogs(opts, "PauserRegistrySet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractSFFLTaskManagerPauserRegistrySet)
				if err := _ContractSFFLTaskManager.contract.UnpackLog(event, "PauserRegistrySet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePauserRegistrySet is a log parse operation binding the contract event 0x6e9fcd539896fca60e8b0f01dd580233e48a6b0f7df013b89ba7f565869acdb6.
//
// Solidity: event PauserRegistrySet(address pauserRegistry, address newPauserRegistry)
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerFilterer) ParsePauserRegistrySet(log types.Log) (*ContractSFFLTaskManagerPauserRegistrySet, error) {
	event := new(ContractSFFLTaskManagerPauserRegistrySet)
	if err := _ContractSFFLTaskManager.contract.UnpackLog(event, "PauserRegistrySet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractSFFLTaskManagerStaleStakesForbiddenUpdateIterator is returned from FilterStaleStakesForbiddenUpdate and is used to iterate over the raw logs and unpacked data for StaleStakesForbiddenUpdate events raised by the ContractSFFLTaskManager contract.
type ContractSFFLTaskManagerStaleStakesForbiddenUpdateIterator struct {
	Event *ContractSFFLTaskManagerStaleStakesForbiddenUpdate // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractSFFLTaskManagerStaleStakesForbiddenUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractSFFLTaskManagerStaleStakesForbiddenUpdate)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractSFFLTaskManagerStaleStakesForbiddenUpdate)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractSFFLTaskManagerStaleStakesForbiddenUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractSFFLTaskManagerStaleStakesForbiddenUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractSFFLTaskManagerStaleStakesForbiddenUpdate represents a StaleStakesForbiddenUpdate event raised by the ContractSFFLTaskManager contract.
type ContractSFFLTaskManagerStaleStakesForbiddenUpdate struct {
	Value bool
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterStaleStakesForbiddenUpdate is a free log retrieval operation binding the contract event 0x40e4ed880a29e0f6ddce307457fb75cddf4feef7d3ecb0301bfdf4976a0e2dfc.
//
// Solidity: event StaleStakesForbiddenUpdate(bool value)
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerFilterer) FilterStaleStakesForbiddenUpdate(opts *bind.FilterOpts) (*ContractSFFLTaskManagerStaleStakesForbiddenUpdateIterator, error) {

	logs, sub, err := _ContractSFFLTaskManager.contract.FilterLogs(opts, "StaleStakesForbiddenUpdate")
	if err != nil {
		return nil, err
	}
	return &ContractSFFLTaskManagerStaleStakesForbiddenUpdateIterator{contract: _ContractSFFLTaskManager.contract, event: "StaleStakesForbiddenUpdate", logs: logs, sub: sub}, nil
}

// WatchStaleStakesForbiddenUpdate is a free log subscription operation binding the contract event 0x40e4ed880a29e0f6ddce307457fb75cddf4feef7d3ecb0301bfdf4976a0e2dfc.
//
// Solidity: event StaleStakesForbiddenUpdate(bool value)
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerFilterer) WatchStaleStakesForbiddenUpdate(opts *bind.WatchOpts, sink chan<- *ContractSFFLTaskManagerStaleStakesForbiddenUpdate) (event.Subscription, error) {

	logs, sub, err := _ContractSFFLTaskManager.contract.WatchLogs(opts, "StaleStakesForbiddenUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractSFFLTaskManagerStaleStakesForbiddenUpdate)
				if err := _ContractSFFLTaskManager.contract.UnpackLog(event, "StaleStakesForbiddenUpdate", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseStaleStakesForbiddenUpdate is a log parse operation binding the contract event 0x40e4ed880a29e0f6ddce307457fb75cddf4feef7d3ecb0301bfdf4976a0e2dfc.
//
// Solidity: event StaleStakesForbiddenUpdate(bool value)
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerFilterer) ParseStaleStakesForbiddenUpdate(log types.Log) (*ContractSFFLTaskManagerStaleStakesForbiddenUpdate, error) {
	event := new(ContractSFFLTaskManagerStaleStakesForbiddenUpdate)
	if err := _ContractSFFLTaskManager.contract.UnpackLog(event, "StaleStakesForbiddenUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractSFFLTaskManagerUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the ContractSFFLTaskManager contract.
type ContractSFFLTaskManagerUnpausedIterator struct {
	Event *ContractSFFLTaskManagerUnpaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractSFFLTaskManagerUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractSFFLTaskManagerUnpaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractSFFLTaskManagerUnpaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractSFFLTaskManagerUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractSFFLTaskManagerUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractSFFLTaskManagerUnpaused represents a Unpaused event raised by the ContractSFFLTaskManager contract.
type ContractSFFLTaskManagerUnpaused struct {
	Account         common.Address
	NewPausedStatus *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c.
//
// Solidity: event Unpaused(address indexed account, uint256 newPausedStatus)
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerFilterer) FilterUnpaused(opts *bind.FilterOpts, account []common.Address) (*ContractSFFLTaskManagerUnpausedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ContractSFFLTaskManager.contract.FilterLogs(opts, "Unpaused", accountRule)
	if err != nil {
		return nil, err
	}
	return &ContractSFFLTaskManagerUnpausedIterator{contract: _ContractSFFLTaskManager.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c.
//
// Solidity: event Unpaused(address indexed account, uint256 newPausedStatus)
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *ContractSFFLTaskManagerUnpaused, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ContractSFFLTaskManager.contract.WatchLogs(opts, "Unpaused", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractSFFLTaskManagerUnpaused)
				if err := _ContractSFFLTaskManager.contract.UnpackLog(event, "Unpaused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUnpaused is a log parse operation binding the contract event 0x3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c.
//
// Solidity: event Unpaused(address indexed account, uint256 newPausedStatus)
func (_ContractSFFLTaskManager *ContractSFFLTaskManagerFilterer) ParseUnpaused(log types.Log) (*ContractSFFLTaskManagerUnpaused, error) {
	event := new(ContractSFFLTaskManagerUnpaused)
	if err := _ContractSFFLTaskManager.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
