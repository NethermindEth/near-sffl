// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/NethermindEth/near-sffl/aggregator (interfaces: RestAggregatorer)
//
// Generated by this command:
//
//	mockgen -destination=./mocks/rest_aggregator.go -package=mocks github.com/NethermindEth/near-sffl/aggregator RestAggregatorer
//
// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	types "github.com/NethermindEth/near-sffl/aggregator/types"
	gomock "go.uber.org/mock/gomock"
)

// MockRestAggregatorer is a mock of RestAggregatorer interface.
type MockRestAggregatorer struct {
	ctrl     *gomock.Controller
	recorder *MockRestAggregatorerMockRecorder
}

// MockRestAggregatorerMockRecorder is the mock recorder for MockRestAggregatorer.
type MockRestAggregatorerMockRecorder struct {
	mock *MockRestAggregatorer
}

// NewMockRestAggregatorer creates a new mock instance.
func NewMockRestAggregatorer(ctrl *gomock.Controller) *MockRestAggregatorer {
	mock := &MockRestAggregatorer{ctrl: ctrl}
	mock.recorder = &MockRestAggregatorerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRestAggregatorer) EXPECT() *MockRestAggregatorerMockRecorder {
	return m.recorder
}

// GetCheckpointMessages mocks base method.
func (m *MockRestAggregatorer) GetCheckpointMessages(arg0, arg1 uint64) (*types.GetCheckpointMessagesResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCheckpointMessages", arg0, arg1)
	ret0, _ := ret[0].(*types.GetCheckpointMessagesResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCheckpointMessages indicates an expected call of GetCheckpointMessages.
func (mr *MockRestAggregatorerMockRecorder) GetCheckpointMessages(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCheckpointMessages", reflect.TypeOf((*MockRestAggregatorer)(nil).GetCheckpointMessages), arg0, arg1)
}

// GetOperatorSetUpdateAggregation mocks base method.
func (m *MockRestAggregatorer) GetOperatorSetUpdateAggregation(arg0 uint64) (*types.GetOperatorSetUpdateAggregationResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOperatorSetUpdateAggregation", arg0)
	ret0, _ := ret[0].(*types.GetOperatorSetUpdateAggregationResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOperatorSetUpdateAggregation indicates an expected call of GetOperatorSetUpdateAggregation.
func (mr *MockRestAggregatorerMockRecorder) GetOperatorSetUpdateAggregation(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOperatorSetUpdateAggregation", reflect.TypeOf((*MockRestAggregatorer)(nil).GetOperatorSetUpdateAggregation), arg0)
}

// GetStateRootUpdateAggregation mocks base method.
func (m *MockRestAggregatorer) GetStateRootUpdateAggregation(arg0 uint32, arg1 uint64) (*types.GetStateRootUpdateAggregationResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStateRootUpdateAggregation", arg0, arg1)
	ret0, _ := ret[0].(*types.GetStateRootUpdateAggregationResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetStateRootUpdateAggregation indicates an expected call of GetStateRootUpdateAggregation.
func (mr *MockRestAggregatorerMockRecorder) GetStateRootUpdateAggregation(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStateRootUpdateAggregation", reflect.TypeOf((*MockRestAggregatorer)(nil).GetStateRootUpdateAggregation), arg0, arg1)
}