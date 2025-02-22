// Code generated by mockery v2.51.1. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// EnvironmentGetter is an autogenerated mock type for the EnvironmentGetter type
type EnvironmentGetter struct {
	mock.Mock
}

// Getenv provides a mock function with given fields: key
func (_m *EnvironmentGetter) Getenv(key string) string {
	ret := _m.Called(key)

	if len(ret) == 0 {
		panic("no return value specified for Getenv")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(key)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// NewEnvironmentGetter creates a new instance of EnvironmentGetter. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewEnvironmentGetter(t interface {
	mock.TestingT
	Cleanup(func())
}) *EnvironmentGetter {
	mock := &EnvironmentGetter{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
