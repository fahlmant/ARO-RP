// Code generated by MockGen. DO NOT EDIT.
// Source: ./client.go
//
// Generated by this command:
//
//	mockgen -typed -source ./client.go -destination=../../../mocks/azureclient/azuresdk/azsecrets/azsecrets.go github.com/Azure/ARO-RP/pkg/util/azureclient/azuresdk/azsecrets Client
//

// Package mock_azsecrets is a generated GoMock package.
package mock_azsecrets

import (
	context "context"
	reflect "reflect"

	runtime "github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	azsecrets "github.com/Azure/azure-sdk-for-go/sdk/security/keyvault/azsecrets"
	gomock "go.uber.org/mock/gomock"
)

// MockClient is a mock of Client interface.
type MockClient struct {
	ctrl     *gomock.Controller
	recorder *MockClientMockRecorder
}

// MockClientMockRecorder is the mock recorder for MockClient.
type MockClientMockRecorder struct {
	mock *MockClient
}

// NewMockClient creates a new mock instance.
func NewMockClient(ctrl *gomock.Controller) *MockClient {
	mock := &MockClient{ctrl: ctrl}
	mock.recorder = &MockClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClient) EXPECT() *MockClientMockRecorder {
	return m.recorder
}

// DeleteSecret mocks base method.
func (m *MockClient) DeleteSecret(ctx context.Context, name string, options *azsecrets.DeleteSecretOptions) (azsecrets.DeleteSecretResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteSecret", ctx, name, options)
	ret0, _ := ret[0].(azsecrets.DeleteSecretResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteSecret indicates an expected call of DeleteSecret.
func (mr *MockClientMockRecorder) DeleteSecret(ctx, name, options any) *MockClientDeleteSecretCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteSecret", reflect.TypeOf((*MockClient)(nil).DeleteSecret), ctx, name, options)
	return &MockClientDeleteSecretCall{Call: call}
}

// MockClientDeleteSecretCall wrap *gomock.Call
type MockClientDeleteSecretCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockClientDeleteSecretCall) Return(arg0 azsecrets.DeleteSecretResponse, arg1 error) *MockClientDeleteSecretCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockClientDeleteSecretCall) Do(f func(context.Context, string, *azsecrets.DeleteSecretOptions) (azsecrets.DeleteSecretResponse, error)) *MockClientDeleteSecretCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockClientDeleteSecretCall) DoAndReturn(f func(context.Context, string, *azsecrets.DeleteSecretOptions) (azsecrets.DeleteSecretResponse, error)) *MockClientDeleteSecretCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// GetSecret mocks base method.
func (m *MockClient) GetSecret(ctx context.Context, name, version string, options *azsecrets.GetSecretOptions) (azsecrets.GetSecretResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSecret", ctx, name, version, options)
	ret0, _ := ret[0].(azsecrets.GetSecretResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSecret indicates an expected call of GetSecret.
func (mr *MockClientMockRecorder) GetSecret(ctx, name, version, options any) *MockClientGetSecretCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSecret", reflect.TypeOf((*MockClient)(nil).GetSecret), ctx, name, version, options)
	return &MockClientGetSecretCall{Call: call}
}

// MockClientGetSecretCall wrap *gomock.Call
type MockClientGetSecretCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockClientGetSecretCall) Return(arg0 azsecrets.GetSecretResponse, arg1 error) *MockClientGetSecretCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockClientGetSecretCall) Do(f func(context.Context, string, string, *azsecrets.GetSecretOptions) (azsecrets.GetSecretResponse, error)) *MockClientGetSecretCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockClientGetSecretCall) DoAndReturn(f func(context.Context, string, string, *azsecrets.GetSecretOptions) (azsecrets.GetSecretResponse, error)) *MockClientGetSecretCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// NewListDeletedSecretPropertiesPager mocks base method.
func (m *MockClient) NewListDeletedSecretPropertiesPager(options *azsecrets.ListDeletedSecretPropertiesOptions) *runtime.Pager[azsecrets.ListDeletedSecretPropertiesResponse] {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewListDeletedSecretPropertiesPager", options)
	ret0, _ := ret[0].(*runtime.Pager[azsecrets.ListDeletedSecretPropertiesResponse])
	return ret0
}

// NewListDeletedSecretPropertiesPager indicates an expected call of NewListDeletedSecretPropertiesPager.
func (mr *MockClientMockRecorder) NewListDeletedSecretPropertiesPager(options any) *MockClientNewListDeletedSecretPropertiesPagerCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewListDeletedSecretPropertiesPager", reflect.TypeOf((*MockClient)(nil).NewListDeletedSecretPropertiesPager), options)
	return &MockClientNewListDeletedSecretPropertiesPagerCall{Call: call}
}

// MockClientNewListDeletedSecretPropertiesPagerCall wrap *gomock.Call
type MockClientNewListDeletedSecretPropertiesPagerCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockClientNewListDeletedSecretPropertiesPagerCall) Return(arg0 *runtime.Pager[azsecrets.ListDeletedSecretPropertiesResponse]) *MockClientNewListDeletedSecretPropertiesPagerCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockClientNewListDeletedSecretPropertiesPagerCall) Do(f func(*azsecrets.ListDeletedSecretPropertiesOptions) *runtime.Pager[azsecrets.ListDeletedSecretPropertiesResponse]) *MockClientNewListDeletedSecretPropertiesPagerCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockClientNewListDeletedSecretPropertiesPagerCall) DoAndReturn(f func(*azsecrets.ListDeletedSecretPropertiesOptions) *runtime.Pager[azsecrets.ListDeletedSecretPropertiesResponse]) *MockClientNewListDeletedSecretPropertiesPagerCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// NewListSecretPropertiesPager mocks base method.
func (m *MockClient) NewListSecretPropertiesPager(options *azsecrets.ListSecretPropertiesOptions) *runtime.Pager[azsecrets.ListSecretPropertiesResponse] {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewListSecretPropertiesPager", options)
	ret0, _ := ret[0].(*runtime.Pager[azsecrets.ListSecretPropertiesResponse])
	return ret0
}

// NewListSecretPropertiesPager indicates an expected call of NewListSecretPropertiesPager.
func (mr *MockClientMockRecorder) NewListSecretPropertiesPager(options any) *MockClientNewListSecretPropertiesPagerCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewListSecretPropertiesPager", reflect.TypeOf((*MockClient)(nil).NewListSecretPropertiesPager), options)
	return &MockClientNewListSecretPropertiesPagerCall{Call: call}
}

// MockClientNewListSecretPropertiesPagerCall wrap *gomock.Call
type MockClientNewListSecretPropertiesPagerCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockClientNewListSecretPropertiesPagerCall) Return(arg0 *runtime.Pager[azsecrets.ListSecretPropertiesResponse]) *MockClientNewListSecretPropertiesPagerCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockClientNewListSecretPropertiesPagerCall) Do(f func(*azsecrets.ListSecretPropertiesOptions) *runtime.Pager[azsecrets.ListSecretPropertiesResponse]) *MockClientNewListSecretPropertiesPagerCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockClientNewListSecretPropertiesPagerCall) DoAndReturn(f func(*azsecrets.ListSecretPropertiesOptions) *runtime.Pager[azsecrets.ListSecretPropertiesResponse]) *MockClientNewListSecretPropertiesPagerCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// PurgeDeletedSecret mocks base method.
func (m *MockClient) PurgeDeletedSecret(ctx context.Context, name string, options *azsecrets.PurgeDeletedSecretOptions) (azsecrets.PurgeDeletedSecretResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PurgeDeletedSecret", ctx, name, options)
	ret0, _ := ret[0].(azsecrets.PurgeDeletedSecretResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PurgeDeletedSecret indicates an expected call of PurgeDeletedSecret.
func (mr *MockClientMockRecorder) PurgeDeletedSecret(ctx, name, options any) *MockClientPurgeDeletedSecretCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PurgeDeletedSecret", reflect.TypeOf((*MockClient)(nil).PurgeDeletedSecret), ctx, name, options)
	return &MockClientPurgeDeletedSecretCall{Call: call}
}

// MockClientPurgeDeletedSecretCall wrap *gomock.Call
type MockClientPurgeDeletedSecretCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockClientPurgeDeletedSecretCall) Return(arg0 azsecrets.PurgeDeletedSecretResponse, arg1 error) *MockClientPurgeDeletedSecretCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockClientPurgeDeletedSecretCall) Do(f func(context.Context, string, *azsecrets.PurgeDeletedSecretOptions) (azsecrets.PurgeDeletedSecretResponse, error)) *MockClientPurgeDeletedSecretCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockClientPurgeDeletedSecretCall) DoAndReturn(f func(context.Context, string, *azsecrets.PurgeDeletedSecretOptions) (azsecrets.PurgeDeletedSecretResponse, error)) *MockClientPurgeDeletedSecretCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// SetSecret mocks base method.
func (m *MockClient) SetSecret(ctx context.Context, name string, parameters azsecrets.SetSecretParameters, options *azsecrets.SetSecretOptions) (azsecrets.SetSecretResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetSecret", ctx, name, parameters, options)
	ret0, _ := ret[0].(azsecrets.SetSecretResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SetSecret indicates an expected call of SetSecret.
func (mr *MockClientMockRecorder) SetSecret(ctx, name, parameters, options any) *MockClientSetSecretCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetSecret", reflect.TypeOf((*MockClient)(nil).SetSecret), ctx, name, parameters, options)
	return &MockClientSetSecretCall{Call: call}
}

// MockClientSetSecretCall wrap *gomock.Call
type MockClientSetSecretCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockClientSetSecretCall) Return(arg0 azsecrets.SetSecretResponse, arg1 error) *MockClientSetSecretCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockClientSetSecretCall) Do(f func(context.Context, string, azsecrets.SetSecretParameters, *azsecrets.SetSecretOptions) (azsecrets.SetSecretResponse, error)) *MockClientSetSecretCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockClientSetSecretCall) DoAndReturn(f func(context.Context, string, azsecrets.SetSecretParameters, *azsecrets.SetSecretOptions) (azsecrets.SetSecretResponse, error)) *MockClientSetSecretCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
