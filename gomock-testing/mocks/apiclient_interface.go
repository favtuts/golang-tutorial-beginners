// Code generated by MockGen. DO NOT EDIT.
// Source: apiclient_interface.go

// Package mock_main is a generated GoMock package.
package mock_main

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockApiClient is a mock of ApiClient interface.
type MockApiClient struct {
	ctrl     *gomock.Controller
	recorder *MockApiClientMockRecorder
}

// MockApiClientMockRecorder is the mock recorder for MockApiClient.
type MockApiClientMockRecorder struct {
	mock *MockApiClient
}

// NewMockApiClient creates a new mock instance.
func NewMockApiClient(ctrl *gomock.Controller) *MockApiClient {
	mock := &MockApiClient{ctrl: ctrl}
	mock.recorder = &MockApiClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockApiClient) EXPECT() *MockApiClientMockRecorder {
	return m.recorder
}

// GetData mocks base method.
func (m *MockApiClient) GetData() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetData")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetData indicates an expected call of GetData.
func (mr *MockApiClientMockRecorder) GetData() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetData", reflect.TypeOf((*MockApiClient)(nil).GetData))
}
