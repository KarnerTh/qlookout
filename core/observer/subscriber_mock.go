// Code generated by mockery v2.21.4. DO NOT EDIT.

package observer

import mock "github.com/stretchr/testify/mock"

// MockSubscriber is an autogenerated mock type for the Subscriber type
type MockSubscriber[T interface{}] struct {
	mock.Mock
}

// Subscribe provides a mock function with given fields:
func (_m *MockSubscriber[T]) Subscribe() <-chan T {
	ret := _m.Called()

	var r0 <-chan T
	if rf, ok := ret.Get(0).(func() <-chan T); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(<-chan T)
		}
	}

	return r0
}

type mockConstructorTestingTNewMockSubscriber interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockSubscriber creates a new instance of MockSubscriber. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockSubscriber[T interface{}](t mockConstructorTestingTNewMockSubscriber) *MockSubscriber[T] {
	mock := &MockSubscriber[T]{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}