// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/NethermindEth/near-sffl/aggregator (interfaces: MessageBlsAggregationService)
//
// Generated by this command:
//
//	mockgen -destination=./mocks/message_blsagg.go -package=mocks github.com/NethermindEth/near-sffl/aggregator MessageBlsAggregationService
//
// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"
	time "time"

	bls "github.com/Layr-Labs/eigensdk-go/crypto/bls"
	types "github.com/NethermindEth/near-sffl/aggregator/types"
	gomock "go.uber.org/mock/gomock"
)

// MockMessageBlsAggregationService is a mock of MessageBlsAggregationService interface.
type MockMessageBlsAggregationService struct {
	ctrl     *gomock.Controller
	recorder *MockMessageBlsAggregationServiceMockRecorder
}

// MockMessageBlsAggregationServiceMockRecorder is the mock recorder for MockMessageBlsAggregationService.
type MockMessageBlsAggregationServiceMockRecorder struct {
	mock *MockMessageBlsAggregationService
}

// NewMockMessageBlsAggregationService creates a new mock instance.
func NewMockMessageBlsAggregationService(ctrl *gomock.Controller) *MockMessageBlsAggregationService {
	mock := &MockMessageBlsAggregationService{ctrl: ctrl}
	mock.recorder = &MockMessageBlsAggregationServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMessageBlsAggregationService) EXPECT() *MockMessageBlsAggregationServiceMockRecorder {
	return m.recorder
}

// GetResponseChannel mocks base method.
func (m *MockMessageBlsAggregationService) GetResponseChannel() <-chan types.MessageBlsAggregationServiceResponse {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetResponseChannel")
	ret0, _ := ret[0].(<-chan types.MessageBlsAggregationServiceResponse)
	return ret0
}

// GetResponseChannel indicates an expected call of GetResponseChannel.
func (mr *MockMessageBlsAggregationServiceMockRecorder) GetResponseChannel() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetResponseChannel", reflect.TypeOf((*MockMessageBlsAggregationService)(nil).GetResponseChannel))
}

// InitializeNewMessage mocks base method.
func (m *MockMessageBlsAggregationService) InitializeNewMessage(arg0 [32]byte, arg1 []byte, arg2 []uint32, arg3 time.Duration) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InitializeNewMessage", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// InitializeNewMessage indicates an expected call of InitializeNewMessage.
func (mr *MockMessageBlsAggregationServiceMockRecorder) InitializeNewMessage(arg0, arg1, arg2, arg3 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InitializeNewMessage", reflect.TypeOf((*MockMessageBlsAggregationService)(nil).InitializeNewMessage), arg0, arg1, arg2, arg3)
}

// ProcessNewSignature mocks base method.
func (m *MockMessageBlsAggregationService) ProcessNewSignature(arg0 context.Context, arg1 [32]byte, arg2 *bls.Signature, arg3 [32]byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProcessNewSignature", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// ProcessNewSignature indicates an expected call of ProcessNewSignature.
func (mr *MockMessageBlsAggregationServiceMockRecorder) ProcessNewSignature(arg0, arg1, arg2, arg3 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProcessNewSignature", reflect.TypeOf((*MockMessageBlsAggregationService)(nil).ProcessNewSignature), arg0, arg1, arg2, arg3)
}
