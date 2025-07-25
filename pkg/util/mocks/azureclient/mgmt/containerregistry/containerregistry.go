// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/Azure/ARO-RP/pkg/util/azureclient/mgmt/containerregistry (interfaces: TokensClient,RegistriesClient)
//
// Generated by this command:
//
//	mockgen -destination=../../../../util/mocks/azureclient/mgmt/containerregistry/containerregistry.go github.com/Azure/ARO-RP/pkg/util/azureclient/mgmt/containerregistry TokensClient,RegistriesClient
//

// Package mock_containerregistry is a generated GoMock package.
package mock_containerregistry

import (
	context "context"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"

	containerregistry "github.com/Azure/azure-sdk-for-go/services/preview/containerregistry/mgmt/2020-11-01-preview/containerregistry"
)

// MockTokensClient is a mock of TokensClient interface.
type MockTokensClient struct {
	ctrl     *gomock.Controller
	recorder *MockTokensClientMockRecorder
	isgomock struct{}
}

// MockTokensClientMockRecorder is the mock recorder for MockTokensClient.
type MockTokensClientMockRecorder struct {
	mock *MockTokensClient
}

// NewMockTokensClient creates a new mock instance.
func NewMockTokensClient(ctrl *gomock.Controller) *MockTokensClient {
	mock := &MockTokensClient{ctrl: ctrl}
	mock.recorder = &MockTokensClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTokensClient) EXPECT() *MockTokensClientMockRecorder {
	return m.recorder
}

// CreateAndWait mocks base method.
func (m *MockTokensClient) CreateAndWait(ctx context.Context, resourceGroupName, registryName, tokenName string, tokenCreateParameters containerregistry.Token) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateAndWait", ctx, resourceGroupName, registryName, tokenName, tokenCreateParameters)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateAndWait indicates an expected call of CreateAndWait.
func (mr *MockTokensClientMockRecorder) CreateAndWait(ctx, resourceGroupName, registryName, tokenName, tokenCreateParameters any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAndWait", reflect.TypeOf((*MockTokensClient)(nil).CreateAndWait), ctx, resourceGroupName, registryName, tokenName, tokenCreateParameters)
}

// DeleteAndWait mocks base method.
func (m *MockTokensClient) DeleteAndWait(ctx context.Context, resourceGroupName, registryName, tokenName string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAndWait", ctx, resourceGroupName, registryName, tokenName)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteAndWait indicates an expected call of DeleteAndWait.
func (mr *MockTokensClientMockRecorder) DeleteAndWait(ctx, resourceGroupName, registryName, tokenName any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAndWait", reflect.TypeOf((*MockTokensClient)(nil).DeleteAndWait), ctx, resourceGroupName, registryName, tokenName)
}

// GetTokenProperties mocks base method.
func (m *MockTokensClient) GetTokenProperties(ctx context.Context, resourceGroupName, registryName, tokenName string) (containerregistry.TokenProperties, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTokenProperties", ctx, resourceGroupName, registryName, tokenName)
	ret0, _ := ret[0].(containerregistry.TokenProperties)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTokenProperties indicates an expected call of GetTokenProperties.
func (mr *MockTokensClientMockRecorder) GetTokenProperties(ctx, resourceGroupName, registryName, tokenName any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTokenProperties", reflect.TypeOf((*MockTokensClient)(nil).GetTokenProperties), ctx, resourceGroupName, registryName, tokenName)
}

// MockRegistriesClient is a mock of RegistriesClient interface.
type MockRegistriesClient struct {
	ctrl     *gomock.Controller
	recorder *MockRegistriesClientMockRecorder
	isgomock struct{}
}

// MockRegistriesClientMockRecorder is the mock recorder for MockRegistriesClient.
type MockRegistriesClientMockRecorder struct {
	mock *MockRegistriesClient
}

// NewMockRegistriesClient creates a new mock instance.
func NewMockRegistriesClient(ctrl *gomock.Controller) *MockRegistriesClient {
	mock := &MockRegistriesClient{ctrl: ctrl}
	mock.recorder = &MockRegistriesClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRegistriesClient) EXPECT() *MockRegistriesClientMockRecorder {
	return m.recorder
}

// GenerateCredentials mocks base method.
func (m *MockRegistriesClient) GenerateCredentials(ctx context.Context, resourceGroupName, registryName string, generateCredentialsParameters containerregistry.GenerateCredentialsParameters) (containerregistry.GenerateCredentialsResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateCredentials", ctx, resourceGroupName, registryName, generateCredentialsParameters)
	ret0, _ := ret[0].(containerregistry.GenerateCredentialsResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GenerateCredentials indicates an expected call of GenerateCredentials.
func (mr *MockRegistriesClientMockRecorder) GenerateCredentials(ctx, resourceGroupName, registryName, generateCredentialsParameters any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateCredentials", reflect.TypeOf((*MockRegistriesClient)(nil).GenerateCredentials), ctx, resourceGroupName, registryName, generateCredentialsParameters)
}
