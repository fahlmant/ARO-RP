// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/Azure/ARO-RP/pkg/util/azureclient/azuresdk/armstorage (interfaces: AccountsClient)

// Package mock_armstorage is a generated GoMock package.
package mock_armstorage

import (
	context "context"
	reflect "reflect"

	armstorage "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storage/armstorage"
	gomock "github.com/golang/mock/gomock"
)

// MockAccountsClient is a mock of AccountsClient interface.
type MockAccountsClient struct {
	ctrl     *gomock.Controller
	recorder *MockAccountsClientMockRecorder
}

// MockAccountsClientMockRecorder is the mock recorder for MockAccountsClient.
type MockAccountsClientMockRecorder struct {
	mock *MockAccountsClient
}

// NewMockAccountsClient creates a new mock instance.
func NewMockAccountsClient(ctrl *gomock.Controller) *MockAccountsClient {
	mock := &MockAccountsClient{ctrl: ctrl}
	mock.recorder = &MockAccountsClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAccountsClient) EXPECT() *MockAccountsClientMockRecorder {
	return m.recorder
}

// GetProperties mocks base method.
func (m *MockAccountsClient) GetProperties(arg0 context.Context, arg1, arg2 string, arg3 *armstorage.AccountsClientGetPropertiesOptions) (armstorage.AccountsClientGetPropertiesResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProperties", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(armstorage.AccountsClientGetPropertiesResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProperties indicates an expected call of GetProperties.
func (mr *MockAccountsClientMockRecorder) GetProperties(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProperties", reflect.TypeOf((*MockAccountsClient)(nil).GetProperties), arg0, arg1, arg2, arg3)
}