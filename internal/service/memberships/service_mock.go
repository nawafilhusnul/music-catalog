// Code generated by MockGen. DO NOT EDIT.
// Source: service.go
//
// Generated by this command:
//
//	mockgen -source=service.go -destination=service_mock.go -package=memberships
//

// Package memberships is a generated GoMock package.
package memberships

import (
	context "context"
	reflect "reflect"

	memberships "github.com/nawafilhusnul/music-catalog/internal/models/memberships"
	gomock "go.uber.org/mock/gomock"
)

// Mockrepository is a mock of repository interface.
type Mockrepository struct {
	ctrl     *gomock.Controller
	recorder *MockrepositoryMockRecorder
	isgomock struct{}
}

// MockrepositoryMockRecorder is the mock recorder for Mockrepository.
type MockrepositoryMockRecorder struct {
	mock *Mockrepository
}

// NewMockrepository creates a new mock instance.
func NewMockrepository(ctrl *gomock.Controller) *Mockrepository {
	mock := &Mockrepository{ctrl: ctrl}
	mock.recorder = &MockrepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *Mockrepository) EXPECT() *MockrepositoryMockRecorder {
	return m.recorder
}

// CreateUser mocks base method.
func (m *Mockrepository) CreateUser(ctx context.Context, model *memberships.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", ctx, model)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockrepositoryMockRecorder) CreateUser(ctx, model any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*Mockrepository)(nil).CreateUser), ctx, model)
}

// GetUser mocks base method.
func (m *Mockrepository) GetUser(ctx context.Context, email, username string, id uint) (*memberships.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUser", ctx, email, username, id)
	ret0, _ := ret[0].(*memberships.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUser indicates an expected call of GetUser.
func (mr *MockrepositoryMockRecorder) GetUser(ctx, email, username, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUser", reflect.TypeOf((*Mockrepository)(nil).GetUser), ctx, email, username, id)
}
