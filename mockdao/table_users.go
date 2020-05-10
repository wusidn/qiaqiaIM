// Code generated by MockGen. DO NOT EDIT.
// Source: dao/table_users.go

// Package mockdao is a generated GoMock package.
package mockdao

import (
	gomock "github.com/golang/mock/gomock"
	dao "github.com/wusidn/qiaqia/dao"
	reflect "reflect"
)

// MockTableUsers is a mock of TableUsers interface
type MockTableUsers struct {
	ctrl     *gomock.Controller
	recorder *MockTableUsersMockRecorder
}

// MockTableUsersMockRecorder is the mock recorder for MockTableUsers
type MockTableUsersMockRecorder struct {
	mock *MockTableUsers
}

// NewMockTableUsers creates a new mock instance
func NewMockTableUsers(ctrl *gomock.Controller) *MockTableUsers {
	mock := &MockTableUsers{ctrl: ctrl}
	mock.recorder = &MockTableUsersMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockTableUsers) EXPECT() *MockTableUsersMockRecorder {
	return m.recorder
}

// InsertUserInfo mocks base method
func (m *MockTableUsers) InsertUserInfo(userInfo dao.UserInfo) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertUserInfo", userInfo)
	ret0, _ := ret[0].(error)
	return ret0
}

// InsertUserInfo indicates an expected call of InsertUserInfo
func (mr *MockTableUsersMockRecorder) InsertUserInfo(userInfo interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertUserInfo", reflect.TypeOf((*MockTableUsers)(nil).InsertUserInfo), userInfo)
}

// QueryUserInfo mocks base method
func (m *MockTableUsers) QueryUserInfo(userId int) (dao.UserInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QueryUserInfo", userId)
	ret0, _ := ret[0].(dao.UserInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QueryUserInfo indicates an expected call of QueryUserInfo
func (mr *MockTableUsersMockRecorder) QueryUserInfo(userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryUserInfo", reflect.TypeOf((*MockTableUsers)(nil).QueryUserInfo), userId)
}

// QueryUserInfoByEmail mocks base method
func (m *MockTableUsers) QueryUserInfoByEmail(email string) (dao.UserInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QueryUserInfoByEmail", email)
	ret0, _ := ret[0].(dao.UserInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QueryUserInfoByEmail indicates an expected call of QueryUserInfoByEmail
func (mr *MockTableUsersMockRecorder) QueryUserInfoByEmail(email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryUserInfoByEmail", reflect.TypeOf((*MockTableUsers)(nil).QueryUserInfoByEmail), email)
}

// QueryUserInfoByPhoneNumber mocks base method
func (m *MockTableUsers) QueryUserInfoByPhoneNumber(phoneNumber string) (dao.UserInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QueryUserInfoByPhoneNumber", phoneNumber)
	ret0, _ := ret[0].(dao.UserInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QueryUserInfoByPhoneNumber indicates an expected call of QueryUserInfoByPhoneNumber
func (mr *MockTableUsersMockRecorder) QueryUserInfoByPhoneNumber(phoneNumber interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryUserInfoByPhoneNumber", reflect.TypeOf((*MockTableUsers)(nil).QueryUserInfoByPhoneNumber), phoneNumber)
}

// QueryUserInfoByEmailOrPhoneNumber mocks base method
func (m *MockTableUsers) QueryUserInfoByEmailOrPhoneNumber(email, phoneNumber string) (dao.UserInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QueryUserInfoByEmailOrPhoneNumber", email, phoneNumber)
	ret0, _ := ret[0].(dao.UserInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QueryUserInfoByEmailOrPhoneNumber indicates an expected call of QueryUserInfoByEmailOrPhoneNumber
func (mr *MockTableUsersMockRecorder) QueryUserInfoByEmailOrPhoneNumber(email, phoneNumber interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryUserInfoByEmailOrPhoneNumber", reflect.TypeOf((*MockTableUsers)(nil).QueryUserInfoByEmailOrPhoneNumber), email, phoneNumber)
}

// UpdateUserInfo mocks base method
func (m *MockTableUsers) UpdateUserInfo(userInfo dao.UserInfo) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUserInfo", userInfo)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateUserInfo indicates an expected call of UpdateUserInfo
func (mr *MockTableUsersMockRecorder) UpdateUserInfo(userInfo interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUserInfo", reflect.TypeOf((*MockTableUsers)(nil).UpdateUserInfo), userInfo)
}