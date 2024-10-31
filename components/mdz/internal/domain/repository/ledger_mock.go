// Code generated by MockGen. DO NOT EDIT.
// Source: /Users/maxwelbm/Workspace/midaz/components/mdz/internal/domain/repository/ledger.go
//
// Generated by this command:
//
//	mockgen -source=/Users/maxwelbm/Workspace/midaz/components/mdz/internal/domain/repository/ledger.go -destination=/Users/maxwelbm/Workspace/midaz/components/mdz/internal/domain/repository/ledger_mock.go -package repository
//

// Package repository is a generated GoMock package.
package repository

import (
	reflect "reflect"

	model "github.com/LerianStudio/midaz/components/mdz/internal/model"
	gomock "go.uber.org/mock/gomock"
)

// MockLedger is a mock of Ledger interface.
type MockLedger struct {
	ctrl     *gomock.Controller
	recorder *MockLedgerMockRecorder
	isgomock struct{}
}

// MockLedgerMockRecorder is the mock recorder for MockLedger.
type MockLedgerMockRecorder struct {
	mock *MockLedger
}

// NewMockLedger creates a new mock instance.
func NewMockLedger(ctrl *gomock.Controller) *MockLedger {
	mock := &MockLedger{ctrl: ctrl}
	mock.recorder = &MockLedgerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLedger) EXPECT() *MockLedgerMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockLedger) Create(organizationID string, inp model.LedgerInput) (*model.LedgerCreate, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", organizationID, inp)
	ret0, _ := ret[0].(*model.LedgerCreate)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockLedgerMockRecorder) Create(organizationID, inp any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockLedger)(nil).Create), organizationID, inp)
}
