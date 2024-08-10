/* Autogenerated file. Do not edit manually. */
/* tslint:disable */
/* eslint-disable */
import type {
  BaseContract,
  BigNumberish,
  BytesLike,
  FunctionFragment,
  Result,
  Interface,
  EventFragment,
  AddressLike,
  ContractRunner,
  ContractMethod,
  Listener,
} from "ethers";
import type {
  TypedContractEvent,
  TypedDeferredTopicFilter,
  TypedEventLog,
  TypedLogDescription,
  TypedListener,
  TypedContractMethod,
} from "./common";

export declare namespace BN254 {
  export type G1PointStruct = { X: BigNumberish; Y: BigNumberish };

  export type G1PointStructOutput = [X: bigint, Y: bigint] & {
    X: bigint;
    Y: bigint;
  };

  export type G2PointStruct = {
    X: [BigNumberish, BigNumberish];
    Y: [BigNumberish, BigNumberish];
  };

  export type G2PointStructOutput = [
    X: [bigint, bigint],
    Y: [bigint, bigint]
  ] & { X: [bigint, bigint]; Y: [bigint, bigint] };
}

export declare namespace RollupOperators {
  export type OperatorStruct = {
    pubkey: BN254.G1PointStruct;
    weight: BigNumberish;
  };

  export type OperatorStructOutput = [
    pubkey: BN254.G1PointStructOutput,
    weight: bigint
  ] & { pubkey: BN254.G1PointStructOutput; weight: bigint };

  export type SignatureInfoStruct = {
    nonSignerPubkeys: BN254.G1PointStruct[];
    apkG2: BN254.G2PointStruct;
    sigma: BN254.G1PointStruct;
  };

  export type SignatureInfoStructOutput = [
    nonSignerPubkeys: BN254.G1PointStructOutput[],
    apkG2: BN254.G2PointStructOutput,
    sigma: BN254.G1PointStructOutput
  ] & {
    nonSignerPubkeys: BN254.G1PointStructOutput[];
    apkG2: BN254.G2PointStructOutput;
    sigma: BN254.G1PointStructOutput;
  };
}

export declare namespace OperatorSetUpdate {
  export type MessageStruct = {
    id: BigNumberish;
    timestamp: BigNumberish;
    operators: RollupOperators.OperatorStruct[];
  };

  export type MessageStructOutput = [
    id: bigint,
    timestamp: bigint,
    operators: RollupOperators.OperatorStructOutput[]
  ] & {
    id: bigint;
    timestamp: bigint;
    operators: RollupOperators.OperatorStructOutput[];
  };
}

export declare namespace StateRootUpdate {
  export type MessageStruct = {
    rollupId: BigNumberish;
    blockHeight: BigNumberish;
    timestamp: BigNumberish;
    nearDaTransactionId: BytesLike;
    nearDaCommitment: BytesLike;
    stateRoot: BytesLike;
  };

  export type MessageStructOutput = [
    rollupId: bigint,
    blockHeight: bigint,
    timestamp: bigint,
    nearDaTransactionId: string,
    nearDaCommitment: string,
    stateRoot: string
  ] & {
    rollupId: bigint;
    blockHeight: bigint;
    timestamp: bigint;
    nearDaTransactionId: string;
    nearDaCommitment: string;
    stateRoot: string;
  };
}

export declare namespace SFFLRegistryBase {
  export type ProofParamsStruct = {
    target: AddressLike;
    storageKey: BytesLike;
    stateTrieWitness: BytesLike;
    storageTrieWitness: BytesLike;
  };

  export type ProofParamsStructOutput = [
    target: string,
    storageKey: string,
    stateTrieWitness: string,
    storageTrieWitness: string
  ] & {
    target: string;
    storageKey: string;
    stateTrieWitness: string;
    storageTrieWitness: string;
  };
}

export interface SFFLRegistryRollupInterface extends Interface {
  getFunction(
    nameOrSignature:
      | "PAUSED_UPDATE_OPERATOR_SET"
      | "PAUSED_UPDATE_STATE_ROOT"
      | "THRESHOLD_DENOMINATOR"
      | "aggregator"
      | "forceOperatorSetUpdate"
      | "getApk"
      | "getOperatorWeight"
      | "getQuorumThreshold"
      | "getStateRoot"
      | "getStorageValue"
      | "getTotalWeight"
      | "initialize"
      | "nextOperatorUpdateId"
      | "owner"
      | "pause"
      | "pauseAll"
      | "paused(uint8)"
      | "paused()"
      | "pauserRegistry"
      | "renounceOwnership"
      | "setInitialOperatorSet"
      | "setPauserRegistry"
      | "setQuorumThreshold"
      | "transferOwnership"
      | "unpause"
      | "updateAndGetStorageValue"
      | "updateOperatorSet"
      | "updateStateRoot"
  ): FunctionFragment;

  getEvent(
    nameOrSignatureOrTopic:
      | "Initialized"
      | "OwnershipTransferred"
      | "Paused"
      | "PauserRegistrySet"
      | "StateRootUpdated"
      | "Unpaused"
  ): EventFragment;

  encodeFunctionData(
    functionFragment: "PAUSED_UPDATE_OPERATOR_SET",
    values?: undefined
  ): string;
  encodeFunctionData(
    functionFragment: "PAUSED_UPDATE_STATE_ROOT",
    values?: undefined
  ): string;
  encodeFunctionData(
    functionFragment: "THRESHOLD_DENOMINATOR",
    values?: undefined
  ): string;
  encodeFunctionData(
    functionFragment: "aggregator",
    values?: undefined
  ): string;
  encodeFunctionData(
    functionFragment: "forceOperatorSetUpdate",
    values: [OperatorSetUpdate.MessageStruct]
  ): string;
  encodeFunctionData(functionFragment: "getApk", values?: undefined): string;
  encodeFunctionData(
    functionFragment: "getOperatorWeight",
    values: [BytesLike]
  ): string;
  encodeFunctionData(
    functionFragment: "getQuorumThreshold",
    values?: undefined
  ): string;
  encodeFunctionData(
    functionFragment: "getStateRoot",
    values: [BigNumberish, BigNumberish]
  ): string;
  encodeFunctionData(
    functionFragment: "getStorageValue",
    values: [StateRootUpdate.MessageStruct, SFFLRegistryBase.ProofParamsStruct]
  ): string;
  encodeFunctionData(
    functionFragment: "getTotalWeight",
    values?: undefined
  ): string;
  encodeFunctionData(
    functionFragment: "initialize",
    values: [BigNumberish, AddressLike, AddressLike, AddressLike]
  ): string;
  encodeFunctionData(
    functionFragment: "nextOperatorUpdateId",
    values?: undefined
  ): string;
  encodeFunctionData(functionFragment: "owner", values?: undefined): string;
  encodeFunctionData(functionFragment: "pause", values: [BigNumberish]): string;
  encodeFunctionData(functionFragment: "pauseAll", values?: undefined): string;
  encodeFunctionData(
    functionFragment: "paused(uint8)",
    values: [BigNumberish]
  ): string;
  encodeFunctionData(functionFragment: "paused()", values?: undefined): string;
  encodeFunctionData(
    functionFragment: "pauserRegistry",
    values?: undefined
  ): string;
  encodeFunctionData(
    functionFragment: "renounceOwnership",
    values?: undefined
  ): string;
  encodeFunctionData(
    functionFragment: "setInitialOperatorSet",
    values: [RollupOperators.OperatorStruct[], BigNumberish]
  ): string;
  encodeFunctionData(
    functionFragment: "setPauserRegistry",
    values: [AddressLike]
  ): string;
  encodeFunctionData(
    functionFragment: "setQuorumThreshold",
    values: [BigNumberish]
  ): string;
  encodeFunctionData(
    functionFragment: "transferOwnership",
    values: [AddressLike]
  ): string;
  encodeFunctionData(
    functionFragment: "unpause",
    values: [BigNumberish]
  ): string;
  encodeFunctionData(
    functionFragment: "updateAndGetStorageValue",
    values: [
      StateRootUpdate.MessageStruct,
      SFFLRegistryBase.ProofParamsStruct,
      BytesLike
    ]
  ): string;
  encodeFunctionData(
    functionFragment: "updateOperatorSet",
    values: [
      OperatorSetUpdate.MessageStruct,
      RollupOperators.SignatureInfoStruct
    ]
  ): string;
  encodeFunctionData(
    functionFragment: "updateStateRoot",
    values: [StateRootUpdate.MessageStruct, RollupOperators.SignatureInfoStruct]
  ): string;

  decodeFunctionResult(
    functionFragment: "PAUSED_UPDATE_OPERATOR_SET",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "PAUSED_UPDATE_STATE_ROOT",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "THRESHOLD_DENOMINATOR",
    data: BytesLike
  ): Result;
  decodeFunctionResult(functionFragment: "aggregator", data: BytesLike): Result;
  decodeFunctionResult(
    functionFragment: "forceOperatorSetUpdate",
    data: BytesLike
  ): Result;
  decodeFunctionResult(functionFragment: "getApk", data: BytesLike): Result;
  decodeFunctionResult(
    functionFragment: "getOperatorWeight",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "getQuorumThreshold",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "getStateRoot",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "getStorageValue",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "getTotalWeight",
    data: BytesLike
  ): Result;
  decodeFunctionResult(functionFragment: "initialize", data: BytesLike): Result;
  decodeFunctionResult(
    functionFragment: "nextOperatorUpdateId",
    data: BytesLike
  ): Result;
  decodeFunctionResult(functionFragment: "owner", data: BytesLike): Result;
  decodeFunctionResult(functionFragment: "pause", data: BytesLike): Result;
  decodeFunctionResult(functionFragment: "pauseAll", data: BytesLike): Result;
  decodeFunctionResult(
    functionFragment: "paused(uint8)",
    data: BytesLike
  ): Result;
  decodeFunctionResult(functionFragment: "paused()", data: BytesLike): Result;
  decodeFunctionResult(
    functionFragment: "pauserRegistry",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "renounceOwnership",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "setInitialOperatorSet",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "setPauserRegistry",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "setQuorumThreshold",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "transferOwnership",
    data: BytesLike
  ): Result;
  decodeFunctionResult(functionFragment: "unpause", data: BytesLike): Result;
  decodeFunctionResult(
    functionFragment: "updateAndGetStorageValue",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "updateOperatorSet",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "updateStateRoot",
    data: BytesLike
  ): Result;
}

export namespace InitializedEvent {
  export type InputTuple = [version: BigNumberish];
  export type OutputTuple = [version: bigint];
  export interface OutputObject {
    version: bigint;
  }
  export type Event = TypedContractEvent<InputTuple, OutputTuple, OutputObject>;
  export type Filter = TypedDeferredTopicFilter<Event>;
  export type Log = TypedEventLog<Event>;
  export type LogDescription = TypedLogDescription<Event>;
}

export namespace OwnershipTransferredEvent {
  export type InputTuple = [previousOwner: AddressLike, newOwner: AddressLike];
  export type OutputTuple = [previousOwner: string, newOwner: string];
  export interface OutputObject {
    previousOwner: string;
    newOwner: string;
  }
  export type Event = TypedContractEvent<InputTuple, OutputTuple, OutputObject>;
  export type Filter = TypedDeferredTopicFilter<Event>;
  export type Log = TypedEventLog<Event>;
  export type LogDescription = TypedLogDescription<Event>;
}

export namespace PausedEvent {
  export type InputTuple = [
    account: AddressLike,
    newPausedStatus: BigNumberish
  ];
  export type OutputTuple = [account: string, newPausedStatus: bigint];
  export interface OutputObject {
    account: string;
    newPausedStatus: bigint;
  }
  export type Event = TypedContractEvent<InputTuple, OutputTuple, OutputObject>;
  export type Filter = TypedDeferredTopicFilter<Event>;
  export type Log = TypedEventLog<Event>;
  export type LogDescription = TypedLogDescription<Event>;
}

export namespace PauserRegistrySetEvent {
  export type InputTuple = [
    pauserRegistry: AddressLike,
    newPauserRegistry: AddressLike
  ];
  export type OutputTuple = [pauserRegistry: string, newPauserRegistry: string];
  export interface OutputObject {
    pauserRegistry: string;
    newPauserRegistry: string;
  }
  export type Event = TypedContractEvent<InputTuple, OutputTuple, OutputObject>;
  export type Filter = TypedDeferredTopicFilter<Event>;
  export type Log = TypedEventLog<Event>;
  export type LogDescription = TypedLogDescription<Event>;
}

export namespace StateRootUpdatedEvent {
  export type InputTuple = [
    rollupId: BigNumberish,
    blockHeight: BigNumberish,
    stateRoot: BytesLike
  ];
  export type OutputTuple = [
    rollupId: bigint,
    blockHeight: bigint,
    stateRoot: string
  ];
  export interface OutputObject {
    rollupId: bigint;
    blockHeight: bigint;
    stateRoot: string;
  }
  export type Event = TypedContractEvent<InputTuple, OutputTuple, OutputObject>;
  export type Filter = TypedDeferredTopicFilter<Event>;
  export type Log = TypedEventLog<Event>;
  export type LogDescription = TypedLogDescription<Event>;
}

export namespace UnpausedEvent {
  export type InputTuple = [
    account: AddressLike,
    newPausedStatus: BigNumberish
  ];
  export type OutputTuple = [account: string, newPausedStatus: bigint];
  export interface OutputObject {
    account: string;
    newPausedStatus: bigint;
  }
  export type Event = TypedContractEvent<InputTuple, OutputTuple, OutputObject>;
  export type Filter = TypedDeferredTopicFilter<Event>;
  export type Log = TypedEventLog<Event>;
  export type LogDescription = TypedLogDescription<Event>;
}

export interface SFFLRegistryRollup extends BaseContract {
  connect(runner?: ContractRunner | null): SFFLRegistryRollup;
  waitForDeployment(): Promise<this>;

  interface: SFFLRegistryRollupInterface;

  queryFilter<TCEvent extends TypedContractEvent>(
    event: TCEvent,
    fromBlockOrBlockhash?: string | number | undefined,
    toBlock?: string | number | undefined
  ): Promise<Array<TypedEventLog<TCEvent>>>;
  queryFilter<TCEvent extends TypedContractEvent>(
    filter: TypedDeferredTopicFilter<TCEvent>,
    fromBlockOrBlockhash?: string | number | undefined,
    toBlock?: string | number | undefined
  ): Promise<Array<TypedEventLog<TCEvent>>>;

  on<TCEvent extends TypedContractEvent>(
    event: TCEvent,
    listener: TypedListener<TCEvent>
  ): Promise<this>;
  on<TCEvent extends TypedContractEvent>(
    filter: TypedDeferredTopicFilter<TCEvent>,
    listener: TypedListener<TCEvent>
  ): Promise<this>;

  once<TCEvent extends TypedContractEvent>(
    event: TCEvent,
    listener: TypedListener<TCEvent>
  ): Promise<this>;
  once<TCEvent extends TypedContractEvent>(
    filter: TypedDeferredTopicFilter<TCEvent>,
    listener: TypedListener<TCEvent>
  ): Promise<this>;

  listeners<TCEvent extends TypedContractEvent>(
    event: TCEvent
  ): Promise<Array<TypedListener<TCEvent>>>;
  listeners(eventName?: string): Promise<Array<Listener>>;
  removeAllListeners<TCEvent extends TypedContractEvent>(
    event?: TCEvent
  ): Promise<this>;

  PAUSED_UPDATE_OPERATOR_SET: TypedContractMethod<[], [bigint], "view">;

  PAUSED_UPDATE_STATE_ROOT: TypedContractMethod<[], [bigint], "view">;

  THRESHOLD_DENOMINATOR: TypedContractMethod<[], [bigint], "view">;

  aggregator: TypedContractMethod<[], [string], "view">;

  forceOperatorSetUpdate: TypedContractMethod<
    [message: OperatorSetUpdate.MessageStruct],
    [void],
    "nonpayable"
  >;

  getApk: TypedContractMethod<[], [BN254.G1PointStructOutput], "view">;

  getOperatorWeight: TypedContractMethod<
    [pubkeyHash: BytesLike],
    [bigint],
    "view"
  >;

  getQuorumThreshold: TypedContractMethod<[], [bigint], "view">;

  getStateRoot: TypedContractMethod<
    [rollupId: BigNumberish, blockHeight: BigNumberish],
    [string],
    "view"
  >;

  getStorageValue: TypedContractMethod<
    [
      message: StateRootUpdate.MessageStruct,
      proofParams: SFFLRegistryBase.ProofParamsStruct
    ],
    [string],
    "view"
  >;

  getTotalWeight: TypedContractMethod<[], [bigint], "view">;

  initialize: TypedContractMethod<
    [
      quorumThreshold: BigNumberish,
      initialOwner: AddressLike,
      _aggregator: AddressLike,
      _pauserRegistry: AddressLike
    ],
    [void],
    "nonpayable"
  >;

  nextOperatorUpdateId: TypedContractMethod<[], [bigint], "view">;

  owner: TypedContractMethod<[], [string], "view">;

  pause: TypedContractMethod<
    [newPausedStatus: BigNumberish],
    [void],
    "nonpayable"
  >;

  pauseAll: TypedContractMethod<[], [void], "nonpayable">;

  "paused(uint8)": TypedContractMethod<
    [index: BigNumberish],
    [boolean],
    "view"
  >;

  "paused()": TypedContractMethod<[], [bigint], "view">;

  pauserRegistry: TypedContractMethod<[], [string], "view">;

  renounceOwnership: TypedContractMethod<[], [void], "nonpayable">;

  setInitialOperatorSet: TypedContractMethod<
    [
      operators: RollupOperators.OperatorStruct[],
      _nextOperatorUpdateId: BigNumberish
    ],
    [void],
    "nonpayable"
  >;

  setPauserRegistry: TypedContractMethod<
    [newPauserRegistry: AddressLike],
    [void],
    "nonpayable"
  >;

  setQuorumThreshold: TypedContractMethod<
    [newQuorumThreshold: BigNumberish],
    [void],
    "nonpayable"
  >;

  transferOwnership: TypedContractMethod<
    [newOwner: AddressLike],
    [void],
    "nonpayable"
  >;

  unpause: TypedContractMethod<
    [newPausedStatus: BigNumberish],
    [void],
    "nonpayable"
  >;

  updateAndGetStorageValue: TypedContractMethod<
    [
      message: StateRootUpdate.MessageStruct,
      proofParams: SFFLRegistryBase.ProofParamsStruct,
      agreement: BytesLike
    ],
    [string],
    "nonpayable"
  >;

  updateOperatorSet: TypedContractMethod<
    [
      message: OperatorSetUpdate.MessageStruct,
      signatureInfo: RollupOperators.SignatureInfoStruct
    ],
    [void],
    "nonpayable"
  >;

  updateStateRoot: TypedContractMethod<
    [
      message: StateRootUpdate.MessageStruct,
      signatureInfo: RollupOperators.SignatureInfoStruct
    ],
    [void],
    "nonpayable"
  >;

  getFunction<T extends ContractMethod = ContractMethod>(
    key: string | FunctionFragment
  ): T;

  getFunction(
    nameOrSignature: "PAUSED_UPDATE_OPERATOR_SET"
  ): TypedContractMethod<[], [bigint], "view">;
  getFunction(
    nameOrSignature: "PAUSED_UPDATE_STATE_ROOT"
  ): TypedContractMethod<[], [bigint], "view">;
  getFunction(
    nameOrSignature: "THRESHOLD_DENOMINATOR"
  ): TypedContractMethod<[], [bigint], "view">;
  getFunction(
    nameOrSignature: "aggregator"
  ): TypedContractMethod<[], [string], "view">;
  getFunction(
    nameOrSignature: "forceOperatorSetUpdate"
  ): TypedContractMethod<
    [message: OperatorSetUpdate.MessageStruct],
    [void],
    "nonpayable"
  >;
  getFunction(
    nameOrSignature: "getApk"
  ): TypedContractMethod<[], [BN254.G1PointStructOutput], "view">;
  getFunction(
    nameOrSignature: "getOperatorWeight"
  ): TypedContractMethod<[pubkeyHash: BytesLike], [bigint], "view">;
  getFunction(
    nameOrSignature: "getQuorumThreshold"
  ): TypedContractMethod<[], [bigint], "view">;
  getFunction(
    nameOrSignature: "getStateRoot"
  ): TypedContractMethod<
    [rollupId: BigNumberish, blockHeight: BigNumberish],
    [string],
    "view"
  >;
  getFunction(
    nameOrSignature: "getStorageValue"
  ): TypedContractMethod<
    [
      message: StateRootUpdate.MessageStruct,
      proofParams: SFFLRegistryBase.ProofParamsStruct
    ],
    [string],
    "view"
  >;
  getFunction(
    nameOrSignature: "getTotalWeight"
  ): TypedContractMethod<[], [bigint], "view">;
  getFunction(
    nameOrSignature: "initialize"
  ): TypedContractMethod<
    [
      quorumThreshold: BigNumberish,
      initialOwner: AddressLike,
      _aggregator: AddressLike,
      _pauserRegistry: AddressLike
    ],
    [void],
    "nonpayable"
  >;
  getFunction(
    nameOrSignature: "nextOperatorUpdateId"
  ): TypedContractMethod<[], [bigint], "view">;
  getFunction(
    nameOrSignature: "owner"
  ): TypedContractMethod<[], [string], "view">;
  getFunction(
    nameOrSignature: "pause"
  ): TypedContractMethod<[newPausedStatus: BigNumberish], [void], "nonpayable">;
  getFunction(
    nameOrSignature: "pauseAll"
  ): TypedContractMethod<[], [void], "nonpayable">;
  getFunction(
    nameOrSignature: "paused(uint8)"
  ): TypedContractMethod<[index: BigNumberish], [boolean], "view">;
  getFunction(
    nameOrSignature: "paused()"
  ): TypedContractMethod<[], [bigint], "view">;
  getFunction(
    nameOrSignature: "pauserRegistry"
  ): TypedContractMethod<[], [string], "view">;
  getFunction(
    nameOrSignature: "renounceOwnership"
  ): TypedContractMethod<[], [void], "nonpayable">;
  getFunction(
    nameOrSignature: "setInitialOperatorSet"
  ): TypedContractMethod<
    [
      operators: RollupOperators.OperatorStruct[],
      _nextOperatorUpdateId: BigNumberish
    ],
    [void],
    "nonpayable"
  >;
  getFunction(
    nameOrSignature: "setPauserRegistry"
  ): TypedContractMethod<
    [newPauserRegistry: AddressLike],
    [void],
    "nonpayable"
  >;
  getFunction(
    nameOrSignature: "setQuorumThreshold"
  ): TypedContractMethod<
    [newQuorumThreshold: BigNumberish],
    [void],
    "nonpayable"
  >;
  getFunction(
    nameOrSignature: "transferOwnership"
  ): TypedContractMethod<[newOwner: AddressLike], [void], "nonpayable">;
  getFunction(
    nameOrSignature: "unpause"
  ): TypedContractMethod<[newPausedStatus: BigNumberish], [void], "nonpayable">;
  getFunction(
    nameOrSignature: "updateAndGetStorageValue"
  ): TypedContractMethod<
    [
      message: StateRootUpdate.MessageStruct,
      proofParams: SFFLRegistryBase.ProofParamsStruct,
      agreement: BytesLike
    ],
    [string],
    "nonpayable"
  >;
  getFunction(
    nameOrSignature: "updateOperatorSet"
  ): TypedContractMethod<
    [
      message: OperatorSetUpdate.MessageStruct,
      signatureInfo: RollupOperators.SignatureInfoStruct
    ],
    [void],
    "nonpayable"
  >;
  getFunction(
    nameOrSignature: "updateStateRoot"
  ): TypedContractMethod<
    [
      message: StateRootUpdate.MessageStruct,
      signatureInfo: RollupOperators.SignatureInfoStruct
    ],
    [void],
    "nonpayable"
  >;

  getEvent(
    key: "Initialized"
  ): TypedContractEvent<
    InitializedEvent.InputTuple,
    InitializedEvent.OutputTuple,
    InitializedEvent.OutputObject
  >;
  getEvent(
    key: "OwnershipTransferred"
  ): TypedContractEvent<
    OwnershipTransferredEvent.InputTuple,
    OwnershipTransferredEvent.OutputTuple,
    OwnershipTransferredEvent.OutputObject
  >;
  getEvent(
    key: "Paused"
  ): TypedContractEvent<
    PausedEvent.InputTuple,
    PausedEvent.OutputTuple,
    PausedEvent.OutputObject
  >;
  getEvent(
    key: "PauserRegistrySet"
  ): TypedContractEvent<
    PauserRegistrySetEvent.InputTuple,
    PauserRegistrySetEvent.OutputTuple,
    PauserRegistrySetEvent.OutputObject
  >;
  getEvent(
    key: "StateRootUpdated"
  ): TypedContractEvent<
    StateRootUpdatedEvent.InputTuple,
    StateRootUpdatedEvent.OutputTuple,
    StateRootUpdatedEvent.OutputObject
  >;
  getEvent(
    key: "Unpaused"
  ): TypedContractEvent<
    UnpausedEvent.InputTuple,
    UnpausedEvent.OutputTuple,
    UnpausedEvent.OutputObject
  >;

  filters: {
    "Initialized(uint8)": TypedContractEvent<
      InitializedEvent.InputTuple,
      InitializedEvent.OutputTuple,
      InitializedEvent.OutputObject
    >;
    Initialized: TypedContractEvent<
      InitializedEvent.InputTuple,
      InitializedEvent.OutputTuple,
      InitializedEvent.OutputObject
    >;

    "OwnershipTransferred(address,address)": TypedContractEvent<
      OwnershipTransferredEvent.InputTuple,
      OwnershipTransferredEvent.OutputTuple,
      OwnershipTransferredEvent.OutputObject
    >;
    OwnershipTransferred: TypedContractEvent<
      OwnershipTransferredEvent.InputTuple,
      OwnershipTransferredEvent.OutputTuple,
      OwnershipTransferredEvent.OutputObject
    >;

    "Paused(address,uint256)": TypedContractEvent<
      PausedEvent.InputTuple,
      PausedEvent.OutputTuple,
      PausedEvent.OutputObject
    >;
    Paused: TypedContractEvent<
      PausedEvent.InputTuple,
      PausedEvent.OutputTuple,
      PausedEvent.OutputObject
    >;

    "PauserRegistrySet(address,address)": TypedContractEvent<
      PauserRegistrySetEvent.InputTuple,
      PauserRegistrySetEvent.OutputTuple,
      PauserRegistrySetEvent.OutputObject
    >;
    PauserRegistrySet: TypedContractEvent<
      PauserRegistrySetEvent.InputTuple,
      PauserRegistrySetEvent.OutputTuple,
      PauserRegistrySetEvent.OutputObject
    >;

    "StateRootUpdated(uint32,uint64,bytes32)": TypedContractEvent<
      StateRootUpdatedEvent.InputTuple,
      StateRootUpdatedEvent.OutputTuple,
      StateRootUpdatedEvent.OutputObject
    >;
    StateRootUpdated: TypedContractEvent<
      StateRootUpdatedEvent.InputTuple,
      StateRootUpdatedEvent.OutputTuple,
      StateRootUpdatedEvent.OutputObject
    >;

    "Unpaused(address,uint256)": TypedContractEvent<
      UnpausedEvent.InputTuple,
      UnpausedEvent.OutputTuple,
      UnpausedEvent.OutputObject
    >;
    Unpaused: TypedContractEvent<
      UnpausedEvent.InputTuple,
      UnpausedEvent.OutputTuple,
      UnpausedEvent.OutputObject
    >;
  };
}
