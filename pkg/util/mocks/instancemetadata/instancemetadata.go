// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/Azure/ARO-RP/pkg/util/instancemetadata (interfaces: InstanceMetadata)
//
// Generated by this command:
//
//	mockgen -destination=../mocks/instancemetadata/instancemetadata.go github.com/Azure/ARO-RP/pkg/util/instancemetadata InstanceMetadata
//

// Package mock_instancemetadata is a generated GoMock package.
package mock_instancemetadata

import (
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"

	azureclient "github.com/Azure/ARO-RP/pkg/util/azureclient"
)

// MockInstanceMetadata is a mock of InstanceMetadata interface.
type MockInstanceMetadata struct {
	ctrl     *gomock.Controller
	recorder *MockInstanceMetadataMockRecorder
	isgomock struct{}
}

// MockInstanceMetadataMockRecorder is the mock recorder for MockInstanceMetadata.
type MockInstanceMetadataMockRecorder struct {
	mock *MockInstanceMetadata
}

// NewMockInstanceMetadata creates a new mock instance.
func NewMockInstanceMetadata(ctrl *gomock.Controller) *MockInstanceMetadata {
	mock := &MockInstanceMetadata{ctrl: ctrl}
	mock.recorder = &MockInstanceMetadataMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockInstanceMetadata) EXPECT() *MockInstanceMetadataMockRecorder {
	return m.recorder
}

// Environment mocks base method.
func (m *MockInstanceMetadata) Environment() *azureclient.AROEnvironment {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Environment")
	ret0, _ := ret[0].(*azureclient.AROEnvironment)
	return ret0
}

// Environment indicates an expected call of Environment.
func (mr *MockInstanceMetadataMockRecorder) Environment() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Environment", reflect.TypeOf((*MockInstanceMetadata)(nil).Environment))
}

// Hostname mocks base method.
func (m *MockInstanceMetadata) Hostname() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Hostname")
	ret0, _ := ret[0].(string)
	return ret0
}

// Hostname indicates an expected call of Hostname.
func (mr *MockInstanceMetadataMockRecorder) Hostname() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Hostname", reflect.TypeOf((*MockInstanceMetadata)(nil).Hostname))
}

// Location mocks base method.
func (m *MockInstanceMetadata) Location() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Location")
	ret0, _ := ret[0].(string)
	return ret0
}

// Location indicates an expected call of Location.
func (mr *MockInstanceMetadataMockRecorder) Location() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Location", reflect.TypeOf((*MockInstanceMetadata)(nil).Location))
}

// ResourceGroup mocks base method.
func (m *MockInstanceMetadata) ResourceGroup() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ResourceGroup")
	ret0, _ := ret[0].(string)
	return ret0
}

// ResourceGroup indicates an expected call of ResourceGroup.
func (mr *MockInstanceMetadataMockRecorder) ResourceGroup() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResourceGroup", reflect.TypeOf((*MockInstanceMetadata)(nil).ResourceGroup))
}

// SubscriptionID mocks base method.
func (m *MockInstanceMetadata) SubscriptionID() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SubscriptionID")
	ret0, _ := ret[0].(string)
	return ret0
}

// SubscriptionID indicates an expected call of SubscriptionID.
func (mr *MockInstanceMetadataMockRecorder) SubscriptionID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubscriptionID", reflect.TypeOf((*MockInstanceMetadata)(nil).SubscriptionID))
}

// TenantID mocks base method.
func (m *MockInstanceMetadata) TenantID() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TenantID")
	ret0, _ := ret[0].(string)
	return ret0
}

// TenantID indicates an expected call of TenantID.
func (mr *MockInstanceMetadataMockRecorder) TenantID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TenantID", reflect.TypeOf((*MockInstanceMetadata)(nil).TenantID))
}
