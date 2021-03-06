// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/jim-minter/rp/pkg/util/azureclient/authorization (interfaces: RoleAssignmentsClient)

// Package mock_authorization is a generated GoMock package.
package mock_authorization

import (
	context "context"
	reflect "reflect"

	authorization "github.com/Azure/azure-sdk-for-go/services/authorization/mgmt/2015-07-01/authorization"
	gomock "github.com/golang/mock/gomock"
)

// MockRoleAssignmentsClient is a mock of RoleAssignmentsClient interface
type MockRoleAssignmentsClient struct {
	ctrl     *gomock.Controller
	recorder *MockRoleAssignmentsClientMockRecorder
}

// MockRoleAssignmentsClientMockRecorder is the mock recorder for MockRoleAssignmentsClient
type MockRoleAssignmentsClientMockRecorder struct {
	mock *MockRoleAssignmentsClient
}

// NewMockRoleAssignmentsClient creates a new mock instance
func NewMockRoleAssignmentsClient(ctrl *gomock.Controller) *MockRoleAssignmentsClient {
	mock := &MockRoleAssignmentsClient{ctrl: ctrl}
	mock.recorder = &MockRoleAssignmentsClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRoleAssignmentsClient) EXPECT() *MockRoleAssignmentsClientMockRecorder {
	return m.recorder
}

// Create mocks base method
func (m *MockRoleAssignmentsClient) Create(arg0 context.Context, arg1, arg2 string, arg3 authorization.RoleAssignmentCreateParameters) (authorization.RoleAssignment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(authorization.RoleAssignment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create
func (mr *MockRoleAssignmentsClientMockRecorder) Create(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockRoleAssignmentsClient)(nil).Create), arg0, arg1, arg2, arg3)
}
