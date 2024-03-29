// Code generated by mockery v2.21.4. DO NOT EDIT.

package watch

import (
	cron "github.com/robfig/cron/v3"
	mock "github.com/stretchr/testify/mock"
)

// MockWatcher is an autogenerated mock type for the Watcher type
type MockWatcher struct {
	mock.Mock
}

// StopWatching provides a mock function with given fields: id
func (_m *MockWatcher) StopWatching(id cron.EntryID) {
	_m.Called(id)
}

// Watch provides a mock function with given fields: config
func (_m *MockWatcher) Watch(config WatchConfig) cron.EntryID {
	ret := _m.Called(config)

	var r0 cron.EntryID
	if rf, ok := ret.Get(0).(func(WatchConfig) cron.EntryID); ok {
		r0 = rf(config)
	} else {
		r0 = ret.Get(0).(cron.EntryID)
	}

	return r0
}

type mockConstructorTestingTNewMockWatcher interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockWatcher creates a new instance of MockWatcher. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockWatcher(t mockConstructorTestingTNewMockWatcher) *MockWatcher {
	mock := &MockWatcher{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
