// Code generated by mockery v2.21.4. DO NOT EDIT.

package config

import mock "github.com/stretchr/testify/mock"

// MockConfig is an autogenerated mock type for the Config type
type MockConfig struct {
	mock.Mock
}

// BaseUrl provides a mock function with given fields:
func (_m *MockConfig) BaseUrl() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// DataSource provides a mock function with given fields:
func (_m *MockConfig) DataSource() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// DatabaseFile provides a mock function with given fields:
func (_m *MockConfig) DatabaseFile() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// LogLevel provides a mock function with given fields:
func (_m *MockConfig) LogLevel() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// MailFromAddress provides a mock function with given fields:
func (_m *MockConfig) MailFromAddress() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// MailPassword provides a mock function with given fields:
func (_m *MockConfig) MailPassword() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// MailSmtpHost provides a mock function with given fields:
func (_m *MockConfig) MailSmtpHost() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// MailSmtpPort provides a mock function with given fields:
func (_m *MockConfig) MailSmtpPort() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// MailToAddress provides a mock function with given fields:
func (_m *MockConfig) MailToAddress() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// MailUsername provides a mock function with given fields:
func (_m *MockConfig) MailUsername() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

type mockConstructorTestingTNewMockConfig interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockConfig creates a new instance of MockConfig. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockConfig(t mockConstructorTestingTNewMockConfig) *MockConfig {
	mock := &MockConfig{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
