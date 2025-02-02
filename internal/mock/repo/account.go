// Code generated by MockGen. DO NOT EDIT.
// Source: account.go

// Package mock_repo is a generated GoMock package.
package mock_repo

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	aggregate "github.com/illusory-server/accounts/internal/domain/aggregate"
	vo "github.com/illusory-server/accounts/internal/domain/vo"
)

// MockAccountCommand is a mock of AccountCommand interface.
type MockAccountCommand struct {
	ctrl     *gomock.Controller
	recorder *MockAccountCommandMockRecorder
}

// MockAccountCommandMockRecorder is the mock recorder for MockAccountCommand.
type MockAccountCommandMockRecorder struct {
	mock *MockAccountCommand
}

// NewMockAccountCommand creates a new mock instance.
func NewMockAccountCommand(ctrl *gomock.Controller) *MockAccountCommand {
	mock := &MockAccountCommand{ctrl: ctrl}
	mock.recorder = &MockAccountCommandMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAccountCommand) EXPECT() *MockAccountCommandMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockAccountCommand) Create(ctx context.Context, account *aggregate.Account) (*aggregate.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, account)
	ret0, _ := ret[0].(*aggregate.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockAccountCommandMockRecorder) Create(ctx, account interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockAccountCommand)(nil).Create), ctx, account)
}

// CreateMany mocks base method.
func (m *MockAccountCommand) CreateMany(ctx context.Context, accounts []*aggregate.Account) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateMany", ctx, accounts)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateMany indicates an expected call of CreateMany.
func (mr *MockAccountCommandMockRecorder) CreateMany(ctx, accounts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateMany", reflect.TypeOf((*MockAccountCommand)(nil).CreateMany), ctx, accounts)
}

// DeleteByEmail mocks base method.
func (m *MockAccountCommand) DeleteByEmail(ctx context.Context, email string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteByEmail", ctx, email)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteByEmail indicates an expected call of DeleteByEmail.
func (mr *MockAccountCommandMockRecorder) DeleteByEmail(ctx, email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteByEmail", reflect.TypeOf((*MockAccountCommand)(nil).DeleteByEmail), ctx, email)
}

// DeleteById mocks base method.
func (m *MockAccountCommand) DeleteById(ctx context.Context, id string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteById", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteById indicates an expected call of DeleteById.
func (mr *MockAccountCommandMockRecorder) DeleteById(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteById", reflect.TypeOf((*MockAccountCommand)(nil).DeleteById), ctx, id)
}

// DeleteByNickname mocks base method.
func (m *MockAccountCommand) DeleteByNickname(ctx context.Context, nickname string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteByNickname", ctx, nickname)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteByNickname indicates an expected call of DeleteByNickname.
func (mr *MockAccountCommandMockRecorder) DeleteByNickname(ctx, nickname interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteByNickname", reflect.TypeOf((*MockAccountCommand)(nil).DeleteByNickname), ctx, nickname)
}

// UpdateAvatarLinkById mocks base method.
func (m *MockAccountCommand) UpdateAvatarLinkById(ctx context.Context, id vo.ID, link vo.Link) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAvatarLinkById", ctx, id, link)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateAvatarLinkById indicates an expected call of UpdateAvatarLinkById.
func (mr *MockAccountCommandMockRecorder) UpdateAvatarLinkById(ctx, id, link interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAvatarLinkById", reflect.TypeOf((*MockAccountCommand)(nil).UpdateAvatarLinkById), ctx, id, link)
}

// UpdateInfoById mocks base method.
func (m *MockAccountCommand) UpdateInfoById(ctx context.Context, id vo.ID, info vo.AccountInfo) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateInfoById", ctx, id, info)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateInfoById indicates an expected call of UpdateInfoById.
func (mr *MockAccountCommandMockRecorder) UpdateInfoById(ctx, id, info interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateInfoById", reflect.TypeOf((*MockAccountCommand)(nil).UpdateInfoById), ctx, id, info)
}

// UpdateNicknameById mocks base method.
func (m *MockAccountCommand) UpdateNicknameById(ctx context.Context, id vo.ID, nickname string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateNicknameById", ctx, id, nickname)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateNicknameById indicates an expected call of UpdateNicknameById.
func (mr *MockAccountCommandMockRecorder) UpdateNicknameById(ctx, id, nickname interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateNicknameById", reflect.TypeOf((*MockAccountCommand)(nil).UpdateNicknameById), ctx, id, nickname)
}

// UpdatePasswordById mocks base method.
func (m *MockAccountCommand) UpdatePasswordById(ctx context.Context, id vo.ID, newPassword vo.Password) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdatePasswordById", ctx, id, newPassword)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdatePasswordById indicates an expected call of UpdatePasswordById.
func (mr *MockAccountCommandMockRecorder) UpdatePasswordById(ctx, id, newPassword interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePasswordById", reflect.TypeOf((*MockAccountCommand)(nil).UpdatePasswordById), ctx, id, newPassword)
}

// UpdateRoleById mocks base method.
func (m *MockAccountCommand) UpdateRoleById(ctx context.Context, id vo.ID, role vo.Role) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateRoleById", ctx, id, role)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateRoleById indicates an expected call of UpdateRoleById.
func (mr *MockAccountCommandMockRecorder) UpdateRoleById(ctx, id, role interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateRoleById", reflect.TypeOf((*MockAccountCommand)(nil).UpdateRoleById), ctx, id, role)
}

// MockAccountQuery is a mock of AccountQuery interface.
type MockAccountQuery struct {
	ctrl     *gomock.Controller
	recorder *MockAccountQueryMockRecorder
}

// MockAccountQueryMockRecorder is the mock recorder for MockAccountQuery.
type MockAccountQueryMockRecorder struct {
	mock *MockAccountQuery
}

// NewMockAccountQuery creates a new mock instance.
func NewMockAccountQuery(ctrl *gomock.Controller) *MockAccountQuery {
	mock := &MockAccountQuery{ctrl: ctrl}
	mock.recorder = &MockAccountQueryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAccountQuery) EXPECT() *MockAccountQueryMockRecorder {
	return m.recorder
}

// CheckAccountRoleById mocks base method.
func (m *MockAccountQuery) CheckAccountRoleById(ctx context.Context, id string, expectedRole vo.AccountRoleType) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckAccountRoleById", ctx, id, expectedRole)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckAccountRoleById indicates an expected call of CheckAccountRoleById.
func (mr *MockAccountQueryMockRecorder) CheckAccountRoleById(ctx, id, expectedRole interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckAccountRoleById", reflect.TypeOf((*MockAccountQuery)(nil).CheckAccountRoleById), ctx, id, expectedRole)
}

// GetByEmail mocks base method.
func (m *MockAccountQuery) GetByEmail(ctx context.Context, email string) (*aggregate.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByEmail", ctx, email)
	ret0, _ := ret[0].(*aggregate.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByEmail indicates an expected call of GetByEmail.
func (mr *MockAccountQueryMockRecorder) GetByEmail(ctx, email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByEmail", reflect.TypeOf((*MockAccountQuery)(nil).GetByEmail), ctx, email)
}

// GetById mocks base method.
func (m *MockAccountQuery) GetById(ctx context.Context, id string) (*aggregate.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetById", ctx, id)
	ret0, _ := ret[0].(*aggregate.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetById indicates an expected call of GetById.
func (mr *MockAccountQueryMockRecorder) GetById(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetById", reflect.TypeOf((*MockAccountQuery)(nil).GetById), ctx, id)
}

// GetByIds mocks base method.
func (m *MockAccountQuery) GetByIds(ctx context.Context, ids []string) ([]*aggregate.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByIds", ctx, ids)
	ret0, _ := ret[0].([]*aggregate.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByIds indicates an expected call of GetByIds.
func (mr *MockAccountQueryMockRecorder) GetByIds(ctx, ids interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByIds", reflect.TypeOf((*MockAccountQuery)(nil).GetByIds), ctx, ids)
}

// GetByNickname mocks base method.
func (m *MockAccountQuery) GetByNickname(ctx context.Context, nickname string) (*aggregate.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByNickname", ctx, nickname)
	ret0, _ := ret[0].(*aggregate.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByNickname indicates an expected call of GetByNickname.
func (mr *MockAccountQueryMockRecorder) GetByNickname(ctx, nickname interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByNickname", reflect.TypeOf((*MockAccountQuery)(nil).GetByNickname), ctx, nickname)
}

// GetByQuery mocks base method.
func (m *MockAccountQuery) GetByQuery(ctx context.Context, query vo.Query) ([]*aggregate.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByQuery", ctx, query)
	ret0, _ := ret[0].([]*aggregate.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByQuery indicates an expected call of GetByQuery.
func (mr *MockAccountQueryMockRecorder) GetByQuery(ctx, query interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByQuery", reflect.TypeOf((*MockAccountQuery)(nil).GetByQuery), ctx, query)
}

// GetPageCountByLimit mocks base method.
func (m *MockAccountQuery) GetPageCountByLimit(ctx context.Context, limit uint64) (uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPageCountByLimit", ctx, limit)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPageCountByLimit indicates an expected call of GetPageCountByLimit.
func (mr *MockAccountQueryMockRecorder) GetPageCountByLimit(ctx, limit interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPageCountByLimit", reflect.TypeOf((*MockAccountQuery)(nil).GetPageCountByLimit), ctx, limit)
}

// HasByEmail mocks base method.
func (m *MockAccountQuery) HasByEmail(ctx context.Context, email string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HasByEmail", ctx, email)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// HasByEmail indicates an expected call of HasByEmail.
func (mr *MockAccountQueryMockRecorder) HasByEmail(ctx, email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HasByEmail", reflect.TypeOf((*MockAccountQuery)(nil).HasByEmail), ctx, email)
}

// HasById mocks base method.
func (m *MockAccountQuery) HasById(ctx context.Context, id string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HasById", ctx, id)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// HasById indicates an expected call of HasById.
func (mr *MockAccountQueryMockRecorder) HasById(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HasById", reflect.TypeOf((*MockAccountQuery)(nil).HasById), ctx, id)
}

// HasByNickname mocks base method.
func (m *MockAccountQuery) HasByNickname(ctx context.Context, nickname string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HasByNickname", ctx, nickname)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// HasByNickname indicates an expected call of HasByNickname.
func (mr *MockAccountQueryMockRecorder) HasByNickname(ctx, nickname interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HasByNickname", reflect.TypeOf((*MockAccountQuery)(nil).HasByNickname), ctx, nickname)
}
