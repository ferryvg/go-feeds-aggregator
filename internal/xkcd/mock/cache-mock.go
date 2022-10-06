// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/ferryvg/go-feeds-aggregator/internal/xkcd (interfaces: Cache)

// Package mock_xkcd is a generated GoMock package.
package mock_xkcd

import (
	reflect "reflect"

	xkcd "github.com/ferryvg/go-feeds-aggregator/internal/xkcd"
	gomock "github.com/golang/mock/gomock"
)

// MockCache is a mock of Cache interface.
type MockCache struct {
	ctrl     *gomock.Controller
	recorder *MockCacheMockRecorder
}

// MockCacheMockRecorder is the mock recorder for MockCache.
type MockCacheMockRecorder struct {
	mock *MockCache
}

// NewMockCache creates a new mock instance.
func NewMockCache(ctrl *gomock.Controller) *MockCache {
	mock := &MockCache{ctrl: ctrl}
	mock.recorder = &MockCacheMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCache) EXPECT() *MockCacheMockRecorder {
	return m.recorder
}

// Get mocks base method.
func (m *MockCache) Get(arg0 int) (*xkcd.SourceItem, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0)
	ret0, _ := ret[0].(*xkcd.SourceItem)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockCacheMockRecorder) Get(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockCache)(nil).Get), arg0)
}

// Set mocks base method.
func (m *MockCache) Set(arg0 *xkcd.SourceItem) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Set", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Set indicates an expected call of Set.
func (mr *MockCacheMockRecorder) Set(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Set", reflect.TypeOf((*MockCache)(nil).Set), arg0)
}
