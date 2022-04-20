// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"

	testing "testing"
)

// RequesterArray is an autogenerated mock type for the RequesterArray type
type RequesterArray struct {
	mock.Mock
}

// Get provides a mock function with given fields: path
func (_m *RequesterArray) Get(path string) ([2]string, error) {
	ret := _m.Called(path)

	var r0 [2]string
	if rf, ok := ret.Get(0).(func(string) [2]string); ok {
		r0 = rf(path)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([2]string)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(path)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewRequesterArray creates a new instance of RequesterArray. It also registers a cleanup function to assert the mocks expectations.
func NewRequesterArray(t testing.TB) *RequesterArray {
	mock := &RequesterArray{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
