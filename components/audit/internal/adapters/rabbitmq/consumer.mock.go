// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/LerianStudio/midaz/components/audit/internal/adapters/rabbitmq (interfaces: ConsumerRepository)
//
// Generated by this command:
//
//	mockgen --destination=consumer.mock.go --package=rabbitmq . ConsumerRepository
//

// Package rabbitmq is a generated GoMock package.
package rabbitmq

import (
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockConsumerRepository is a mock of ConsumerRepository interface.
type MockConsumerRepository struct {
	ctrl     *gomock.Controller
	recorder *MockConsumerRepositoryMockRecorder
}

// MockConsumerRepositoryMockRecorder is the mock recorder for MockConsumerRepository.
type MockConsumerRepositoryMockRecorder struct {
	mock *MockConsumerRepository
}

// NewMockConsumerRepository creates a new mock instance.
func NewMockConsumerRepository(ctrl *gomock.Controller) *MockConsumerRepository {
	mock := &MockConsumerRepository{ctrl: ctrl}
	mock.recorder = &MockConsumerRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockConsumerRepository) EXPECT() *MockConsumerRepositoryMockRecorder {
	return m.recorder
}

// ConsumerAudit mocks base method.
func (m *MockConsumerRepository) ConsumerAudit() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ConsumerAudit")
}

// ConsumerAudit indicates an expected call of ConsumerAudit.
func (mr *MockConsumerRepositoryMockRecorder) ConsumerAudit() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConsumerAudit", reflect.TypeOf((*MockConsumerRepository)(nil).ConsumerAudit))
}
