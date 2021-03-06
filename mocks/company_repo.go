// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/wgarunap/xm-rest-api/domain/repository (interfaces: Company)

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	domain "github.com/wgarunap/xm-rest-api/domain"
	repository "github.com/wgarunap/xm-rest-api/domain/repository"
	reflect "reflect"
)

// MockCompany is a mock of Company interface
type MockCompany struct {
	ctrl     *gomock.Controller
	recorder *MockCompanyMockRecorder
}

// MockCompanyMockRecorder is the mock recorder for MockCompany
type MockCompanyMockRecorder struct {
	mock *MockCompany
}

// NewMockCompany creates a new mock instance
func NewMockCompany(ctrl *gomock.Controller) *MockCompany {
	mock := &MockCompany{ctrl: ctrl}
	mock.recorder = &MockCompanyMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCompany) EXPECT() *MockCompanyMockRecorder {
	return m.recorder
}

// Create mocks base method
func (m *MockCompany) Create(arg0 domain.Company) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create
func (mr *MockCompanyMockRecorder) Create(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockCompany)(nil).Create), arg0)
}

// Delete mocks base method
func (m *MockCompany) Delete(arg0 ...repository.Filter) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range arg0 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Delete", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *MockCompanyMockRecorder) Delete(arg0 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockCompany)(nil).Delete), arg0...)
}

// Get mocks base method
func (m *MockCompany) Get(arg0 ...repository.Filter) ([]domain.Company, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range arg0 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Get", varargs...)
	ret0, _ := ret[0].([]domain.Company)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockCompanyMockRecorder) Get(arg0 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockCompany)(nil).Get), arg0...)
}

// Update mocks base method
func (m *MockCompany) Update(arg0 domain.Company) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update
func (mr *MockCompanyMockRecorder) Update(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockCompany)(nil).Update), arg0)
}
