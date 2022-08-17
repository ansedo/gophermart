// Code generated by MockGen. DO NOT EDIT.
// Source: internal/pkg/repository/postgres.go

// Package mock_repository is a generated GoMock package.
package mock_repository

import (
	context "context"
	reflect "reflect"

	models "github.com/ansedo/gophermart/internal/models"
	gomock "github.com/golang/mock/gomock"
	sqlx "github.com/jmoiron/sqlx"
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

// Transaction mocks base method.
func (m *MockRepository) Transaction(ctx context.Context, f func(context.Context, *sqlx.Tx) error) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Transaction", ctx, f)
	ret0, _ := ret[0].(error)
	return ret0
}

// Transaction indicates an expected call of Transaction.
func (mr *MockRepositoryMockRecorder) Transaction(ctx, f interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Transaction", reflect.TypeOf((*MockRepository)(nil).Transaction), ctx, f)
}

// GetAllBalance mocks base method.
func (m *MockRepository) GetBalanceByUID(ctx context.Context, UID string) (models.Balance, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBalanceByUID", ctx, UID)
	ret0, _ := ret[0].(models.Balance)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllBalance indicates an expected call of GetAllBalance.
func (mr *MockRepositoryMockRecorder) GetAllBalance(ctx, UID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBalanceByUID", reflect.TypeOf((*MockRepository)(nil).GetBalanceByUID), ctx, UID)
}

// GetCurrentBalance mocks base method.
func (m *MockRepository) GetCurrentBalanceByUID(ctx context.Context, UID string, tx *sqlx.Tx) (float64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCurrentBalanceByUID", ctx, UID, tx)
	ret0, _ := ret[0].(float64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCurrentBalance indicates an expected call of GetCurrentBalance.
func (mr *MockRepositoryMockRecorder) GetCurrentBalance(ctx, UID, tx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCurrentBalanceByUID", reflect.TypeOf((*MockRepository)(nil).GetCurrentBalanceByUID), ctx, UID, tx)
}

// GetOrders mocks base method.
func (m *MockRepository) GetOrdersByUID(ctx context.Context, UID string) ([]models.Order, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOrdersByUID", ctx, UID)
	ret0, _ := ret[0].([]models.Order)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOrders indicates an expected call of GetOrders.
func (mr *MockRepositoryMockRecorder) GetOrders(ctx, UID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrdersByUID", reflect.TypeOf((*MockRepository)(nil).GetOrdersByUID), ctx, UID)
}

// GetUnprocessedOrders mocks base method.
func (m *MockRepository) GetUnfinishedOrders(ctx context.Context) ([]models.Order, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUnfinishedOrders", ctx)
	ret0, _ := ret[0].([]models.Order)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUnprocessedOrders indicates an expected call of GetUnprocessedOrders.
func (mr *MockRepositoryMockRecorder) GetUnprocessedOrders(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUnfinishedOrders", reflect.TypeOf((*MockRepository)(nil).GetUnfinishedOrders), ctx)
}

// GetUser mocks base method.
func (m *MockRepository) GetUserByLogin(ctx context.Context, user models.User) (models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByLogin", ctx, user)
	ret0, _ := ret[0].(models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUser indicates an expected call of GetUser.
func (mr *MockRepositoryMockRecorder) GetUser(ctx, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByLogin", reflect.TypeOf((*MockRepository)(nil).GetUserByLogin), ctx, user)
}

// GetWithdrawals mocks base method.
func (m *MockRepository) GetWithdrawalsByUID(ctx context.Context, UID string) ([]models.Withdrawal, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetWithdrawalsByUID", ctx, UID)
	ret0, _ := ret[0].([]models.Withdrawal)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetWithdrawals indicates an expected call of GetWithdrawals.
func (mr *MockRepositoryMockRecorder) GetWithdrawals(ctx, UID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWithdrawalsByUID", reflect.TypeOf((*MockRepository)(nil).GetWithdrawalsByUID), ctx, UID)
}

// SaveBalance mocks base method.
func (m *MockRepository) SaveBalance(ctx context.Context, UID string, value *float64, tx *sqlx.Tx) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveBalance", ctx, UID, value, tx)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveBalance indicates an expected call of SaveBalance.
func (mr *MockRepositoryMockRecorder) SaveBalance(ctx, UID, value, tx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveBalance", reflect.TypeOf((*MockRepository)(nil).SaveBalance), ctx, UID, value, tx)
}

// SaveOrder mocks base method.
func (m *MockRepository) SaveOrder(ctx context.Context, OrderID models.Order, tx *sqlx.Tx) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddOrder", ctx, OrderID, tx)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveOrder indicates an expected call of SaveOrder.
func (mr *MockRepositoryMockRecorder) SaveOrder(ctx, OrderID, tx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddOrder", reflect.TypeOf((*MockRepository)(nil).SaveOrder), ctx, OrderID, tx)
}

// SaveUser mocks base method.
func (m *MockRepository) SaveUser(ctx context.Context, user models.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddUser", ctx, user)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveUser indicates an expected call of SaveUser.
func (mr *MockRepositoryMockRecorder) SaveUser(ctx, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddUser", reflect.TypeOf((*MockRepository)(nil).SaveUser), ctx, user)
}

// SaveWithdrawal mocks base method.
func (m *MockRepository) SaveWithdrawal(ctx context.Context, wth models.Withdrawal, tx *sqlx.Tx) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddWithdrawal", ctx, wth, tx)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveWithdrawal indicates an expected call of SaveWithdrawal.
func (mr *MockRepositoryMockRecorder) SaveWithdrawal(ctx, wth, tx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddWithdrawal", reflect.TypeOf((*MockRepository)(nil).SaveWithdrawal), ctx, wth, tx)
}

// UpdateOrder mocks base method.
func (m *MockRepository) UpdateOrder(ctx context.Context, order models.Order, tx *sqlx.Tx) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateOrder", ctx, order, tx)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateOrder indicates an expected call of UpdateOrder.
func (mr *MockRepositoryMockRecorder) UpdateOrder(ctx, order, tx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateOrder", reflect.TypeOf((*MockRepository)(nil).UpdateOrder), ctx, order, tx)
}

// WithdrawBalance mocks base method.
func (m *MockRepository) WithdrawBalance(ctx context.Context, UID string, value float64, tx *sqlx.Tx) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IncrBalanceWithdrawnByUID", ctx, UID, value, tx)
	ret0, _ := ret[0].(error)
	return ret0
}

// WithdrawBalance indicates an expected call of WithdrawBalance.
func (mr *MockRepositoryMockRecorder) WithdrawBalance(ctx, UID, value, tx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IncrBalanceWithdrawnByUID", reflect.TypeOf((*MockRepository)(nil).WithdrawBalance), ctx, UID, value, tx)
}
