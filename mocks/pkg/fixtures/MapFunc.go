// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"

	testing "testing"
)

// MapFunc is an autogenerated mock type for the MapFunc type
type MapFunc struct {
	mock.Mock
}

// Get provides a mock function with given fields: m
func (_m *MapFunc) Get(m map[string]func(string) string) error {
	ret := _m.Called(m)

	var r0 error
	if rf, ok := ret.Get(0).(func(map[string]func(string) string) error); ok {
		r0 = rf(m)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewMapFunc creates a new instance of MapFunc. It also registers a cleanup function to assert the mocks expectations.
func NewMapFunc(t testing.TB) *MapFunc {
	mock := &MapFunc{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
