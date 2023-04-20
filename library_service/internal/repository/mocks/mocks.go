// Code generated by MockGen. DO NOT EDIT.
// Source: manager.go

// Package mock_repository is a generated GoMock package.
package mock_repository

import (
	context "context"
	model "onelab/internal/model"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockIUserRepository is a mock of IUserRepository interface.
type MockIUserRepository struct {
	ctrl     *gomock.Controller
	recorder *MockIUserRepositoryMockRecorder
}

// MockIUserRepositoryMockRecorder is the mock recorder for MockIUserRepository.
type MockIUserRepositoryMockRecorder struct {
	mock *MockIUserRepository
}

// NewMockIUserRepository creates a new mock instance.
func NewMockIUserRepository(ctrl *gomock.Controller) *MockIUserRepository {
	mock := &MockIUserRepository{ctrl: ctrl}
	mock.recorder = &MockIUserRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIUserRepository) EXPECT() *MockIUserRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockIUserRepository) Create(ctx context.Context, user model.User) (uint, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, user)
	ret0, _ := ret[0].(uint)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockIUserRepositoryMockRecorder) Create(ctx, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockIUserRepository)(nil).Create), ctx, user)
}

// Get mocks base method.
func (m *MockIUserRepository) Get(ctx context.Context, id int) (model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, id)
	ret0, _ := ret[0].(model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockIUserRepositoryMockRecorder) Get(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockIUserRepository)(nil).Get), ctx, id)
}

// GetByEmail mocks base method.
func (m *MockIUserRepository) GetByEmail(ctx context.Context, email string) (model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByEmail", ctx, email)
	ret0, _ := ret[0].(model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByEmail indicates an expected call of GetByEmail.
func (mr *MockIUserRepositoryMockRecorder) GetByEmail(ctx, email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByEmail", reflect.TypeOf((*MockIUserRepository)(nil).GetByEmail), ctx, email)
}

// Update mocks base method.
func (m *MockIUserRepository) Update(ctx context.Context, user model.Login) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, user)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockIUserRepositoryMockRecorder) Update(ctx, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockIUserRepository)(nil).Update), ctx, user)
}

// MockIBookRepository is a mock of IBookRepository interface.
type MockIBookRepository struct {
	ctrl     *gomock.Controller
	recorder *MockIBookRepositoryMockRecorder
}

// MockIBookRepositoryMockRecorder is the mock recorder for MockIBookRepository.
type MockIBookRepositoryMockRecorder struct {
	mock *MockIBookRepository
}

// NewMockIBookRepository creates a new mock instance.
func NewMockIBookRepository(ctrl *gomock.Controller) *MockIBookRepository {
	mock := &MockIBookRepository{ctrl: ctrl}
	mock.recorder = &MockIBookRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIBookRepository) EXPECT() *MockIBookRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockIBookRepository) Create(ctx context.Context, book model.Book) (uint, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, book)
	ret0, _ := ret[0].(uint)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockIBookRepositoryMockRecorder) Create(ctx, book interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockIBookRepository)(nil).Create), ctx, book)
}

// Get mocks base method.
func (m *MockIBookRepository) Get(ctx context.Context, id int) (model.Book, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, id)
	ret0, _ := ret[0].(model.Book)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockIBookRepositoryMockRecorder) Get(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockIBookRepository)(nil).Get), ctx, id)
}

// GetAll mocks base method.
func (m *MockIBookRepository) GetAll(ctx context.Context) ([]model.Book, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll", ctx)
	ret0, _ := ret[0].([]model.Book)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll.
func (mr *MockIBookRepositoryMockRecorder) GetAll(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockIBookRepository)(nil).GetAll), ctx)
}

// MockILibraryRepository is a mock of ILibraryRepository interface.
type MockILibraryRepository struct {
	ctrl     *gomock.Controller
	recorder *MockILibraryRepositoryMockRecorder
}

// MockILibraryRepositoryMockRecorder is the mock recorder for MockILibraryRepository.
type MockILibraryRepositoryMockRecorder struct {
	mock *MockILibraryRepository
}

// NewMockILibraryRepository creates a new mock instance.
func NewMockILibraryRepository(ctrl *gomock.Controller) *MockILibraryRepository {
	mock := &MockILibraryRepository{ctrl: ctrl}
	mock.recorder = &MockILibraryRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockILibraryRepository) EXPECT() *MockILibraryRepositoryMockRecorder {
	return m.recorder
}

// GiveBook mocks base method.
func (m *MockILibraryRepository) GiveBook(ctx context.Context, item model.CreateBookReader) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GiveBook", ctx, item)
	ret0, _ := ret[0].(error)
	return ret0
}

// GiveBook indicates an expected call of GiveBook.
func (mr *MockILibraryRepositoryMockRecorder) GiveBook(ctx, item interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GiveBook", reflect.TypeOf((*MockILibraryRepository)(nil).GiveBook), ctx, item)
}

// List mocks base method.
func (m *MockILibraryRepository) List(ctx context.Context) ([]model.BookReader, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", ctx)
	ret0, _ := ret[0].([]model.BookReader)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockILibraryRepositoryMockRecorder) List(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockILibraryRepository)(nil).List), ctx)
}

// ListMonth mocks base method.
func (m *MockILibraryRepository) ListMonth(ctx context.Context) ([]model.BookReader, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListMonth", ctx)
	ret0, _ := ret[0].([]model.BookReader)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListMonth indicates an expected call of ListMonth.
func (mr *MockILibraryRepositoryMockRecorder) ListMonth(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListMonth", reflect.TypeOf((*MockILibraryRepository)(nil).ListMonth), ctx)
}

// ReturnBook mocks base method.
func (m *MockILibraryRepository) ReturnBook(ctx context.Context, bookReader model.BookReader) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReturnBook", ctx, bookReader)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReturnBook indicates an expected call of ReturnBook.
func (mr *MockILibraryRepositoryMockRecorder) ReturnBook(ctx, bookReader interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReturnBook", reflect.TypeOf((*MockILibraryRepository)(nil).ReturnBook), ctx, bookReader)
}