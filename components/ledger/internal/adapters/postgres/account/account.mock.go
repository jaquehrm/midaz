// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/LerianStudio/midaz/components/ledger/internal/adapters/postgres/account (interfaces: Repository)
//
// Generated by this command:
//
//	mockgen --destination=account.mock.go --package=account . Repository
//

// Package account is a generated GoMock package.
package account

import (
	context "context"
	reflect "reflect"

	mmodel "github.com/LerianStudio/midaz/pkg/mmodel"
	http "github.com/LerianStudio/midaz/pkg/net/http"
	uuid "github.com/google/uuid"
	gomock "go.uber.org/mock/gomock"
)

// MockRepository is a mock of Repository interface.
type MockRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryMockRecorder
}

// MockRepositoryMockRecorder is the mock recorder for MockRepository.
type MockRepositoryMockRecorder struct {
	mock *MockRepository
}

// NewMockRepository creates a new mock instance.
func NewMockRepository(ctrl *gomock.Controller) *MockRepository {
	mock := &MockRepository{ctrl: ctrl}
	mock.recorder = &MockRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepository) EXPECT() *MockRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockRepository) Create(arg0 context.Context, arg1 *mmodel.Account) (*mmodel.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1)
	ret0, _ := ret[0].(*mmodel.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockRepositoryMockRecorder) Create(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockRepository)(nil).Create), arg0, arg1)
}

// Delete mocks base method.
func (m *MockRepository) Delete(arg0 context.Context, arg1, arg2 uuid.UUID, arg3 *uuid.UUID, arg4 uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockRepositoryMockRecorder) Delete(arg0, arg1, arg2, arg3, arg4 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockRepository)(nil).Delete), arg0, arg1, arg2, arg3, arg4)
}

// Find mocks base method.
func (m *MockRepository) Find(arg0 context.Context, arg1, arg2 uuid.UUID, arg3 *uuid.UUID, arg4 uuid.UUID) (*mmodel.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Find", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(*mmodel.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Find indicates an expected call of Find.
func (mr *MockRepositoryMockRecorder) Find(arg0, arg1, arg2, arg3, arg4 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockRepository)(nil).Find), arg0, arg1, arg2, arg3, arg4)
}

// FindAlias mocks base method.
func (m *MockRepository) FindAlias(arg0 context.Context, arg1, arg2 uuid.UUID, arg3 *uuid.UUID, arg4 string) (*mmodel.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAlias", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(*mmodel.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindAlias indicates an expected call of FindAlias.
func (mr *MockRepositoryMockRecorder) FindAlias(arg0, arg1, arg2, arg3, arg4 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAlias", reflect.TypeOf((*MockRepository)(nil).FindAlias), arg0, arg1, arg2, arg3, arg4)
}

// FindAll mocks base method.
func (m *MockRepository) FindAll(arg0 context.Context, arg1, arg2 uuid.UUID, arg3 *uuid.UUID, arg4 http.Pagination) ([]*mmodel.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAll", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].([]*mmodel.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindAll indicates an expected call of FindAll.
func (mr *MockRepositoryMockRecorder) FindAll(arg0, arg1, arg2, arg3, arg4 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAll", reflect.TypeOf((*MockRepository)(nil).FindAll), arg0, arg1, arg2, arg3, arg4)
}

// FindByAlias mocks base method.
func (m *MockRepository) FindByAlias(arg0 context.Context, arg1, arg2 uuid.UUID, arg3 string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByAlias", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByAlias indicates an expected call of FindByAlias.
func (mr *MockRepositoryMockRecorder) FindByAlias(arg0, arg1, arg2, arg3 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByAlias", reflect.TypeOf((*MockRepository)(nil).FindByAlias), arg0, arg1, arg2, arg3)
}

// FindWithDeleted mocks base method.
func (m *MockRepository) FindWithDeleted(arg0 context.Context, arg1, arg2 uuid.UUID, arg3 *uuid.UUID, arg4 uuid.UUID) (*mmodel.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindWithDeleted", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(*mmodel.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindWithDeleted indicates an expected call of FindWithDeleted.
func (mr *MockRepositoryMockRecorder) FindWithDeleted(arg0, arg1, arg2, arg3, arg4 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindWithDeleted", reflect.TypeOf((*MockRepository)(nil).FindWithDeleted), arg0, arg1, arg2, arg3, arg4)
}

// ListAccountsByAlias mocks base method.
func (m *MockRepository) ListAccountsByAlias(arg0 context.Context, arg1, arg2 uuid.UUID, arg3 []string) ([]*mmodel.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAccountsByAlias", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].([]*mmodel.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAccountsByAlias indicates an expected call of ListAccountsByAlias.
func (mr *MockRepositoryMockRecorder) ListAccountsByAlias(arg0, arg1, arg2, arg3 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAccountsByAlias", reflect.TypeOf((*MockRepository)(nil).ListAccountsByAlias), arg0, arg1, arg2, arg3)
}

// ListAccountsByIDs mocks base method.
func (m *MockRepository) ListAccountsByIDs(arg0 context.Context, arg1, arg2 uuid.UUID, arg3 []uuid.UUID) ([]*mmodel.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAccountsByIDs", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].([]*mmodel.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAccountsByIDs indicates an expected call of ListAccountsByIDs.
func (mr *MockRepositoryMockRecorder) ListAccountsByIDs(arg0, arg1, arg2, arg3 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAccountsByIDs", reflect.TypeOf((*MockRepository)(nil).ListAccountsByIDs), arg0, arg1, arg2, arg3)
}

// ListByAlias mocks base method.
func (m *MockRepository) ListByAlias(arg0 context.Context, arg1, arg2, arg3 uuid.UUID, arg4 []string) ([]*mmodel.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListByAlias", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].([]*mmodel.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListByAlias indicates an expected call of ListByAlias.
func (mr *MockRepositoryMockRecorder) ListByAlias(arg0, arg1, arg2, arg3, arg4 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListByAlias", reflect.TypeOf((*MockRepository)(nil).ListByAlias), arg0, arg1, arg2, arg3, arg4)
}

// ListByIDs mocks base method.
func (m *MockRepository) ListByIDs(arg0 context.Context, arg1, arg2 uuid.UUID, arg3 *uuid.UUID, arg4 []uuid.UUID) ([]*mmodel.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListByIDs", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].([]*mmodel.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListByIDs indicates an expected call of ListByIDs.
func (mr *MockRepositoryMockRecorder) ListByIDs(arg0, arg1, arg2, arg3, arg4 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListByIDs", reflect.TypeOf((*MockRepository)(nil).ListByIDs), arg0, arg1, arg2, arg3, arg4)
}

// Update mocks base method.
func (m *MockRepository) Update(arg0 context.Context, arg1, arg2 uuid.UUID, arg3 *uuid.UUID, arg4 uuid.UUID, arg5 *mmodel.Account) (*mmodel.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0, arg1, arg2, arg3, arg4, arg5)
	ret0, _ := ret[0].(*mmodel.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockRepositoryMockRecorder) Update(arg0, arg1, arg2, arg3, arg4, arg5 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockRepository)(nil).Update), arg0, arg1, arg2, arg3, arg4, arg5)
}
