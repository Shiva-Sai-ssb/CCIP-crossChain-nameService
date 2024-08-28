// Code generated by mockery v2.38.0. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	types "github.com/smartcontractkit/chainlink/v2/common/types"
)

// ForwarderManager is an autogenerated mock type for the ForwarderManager type
type ForwarderManager[ADDR types.Hashable] struct {
	mock.Mock
}

// Close provides a mock function with given fields:
func (_m *ForwarderManager[ADDR]) Close() error {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Close")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ConvertPayload provides a mock function with given fields: dest, origPayload
func (_m *ForwarderManager[ADDR]) ConvertPayload(dest ADDR, origPayload []byte) ([]byte, error) {
	ret := _m.Called(dest, origPayload)

	if len(ret) == 0 {
		panic("no return value specified for ConvertPayload")
	}

	var r0 []byte
	var r1 error
	if rf, ok := ret.Get(0).(func(ADDR, []byte) ([]byte, error)); ok {
		return rf(dest, origPayload)
	}
	if rf, ok := ret.Get(0).(func(ADDR, []byte) []byte); ok {
		r0 = rf(dest, origPayload)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	if rf, ok := ret.Get(1).(func(ADDR, []byte) error); ok {
		r1 = rf(dest, origPayload)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ForwarderFor provides a mock function with given fields: addr
func (_m *ForwarderManager[ADDR]) ForwarderFor(addr ADDR) (ADDR, error) {
	ret := _m.Called(addr)

	if len(ret) == 0 {
		panic("no return value specified for ForwarderFor")
	}

	var r0 ADDR
	var r1 error
	if rf, ok := ret.Get(0).(func(ADDR) (ADDR, error)); ok {
		return rf(addr)
	}
	if rf, ok := ret.Get(0).(func(ADDR) ADDR); ok {
		r0 = rf(addr)
	} else {
		r0 = ret.Get(0).(ADDR)
	}

	if rf, ok := ret.Get(1).(func(ADDR) error); ok {
		r1 = rf(addr)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// HealthReport provides a mock function with given fields:
func (_m *ForwarderManager[ADDR]) HealthReport() map[string]error {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for HealthReport")
	}

	var r0 map[string]error
	if rf, ok := ret.Get(0).(func() map[string]error); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]error)
		}
	}

	return r0
}

// Name provides a mock function with given fields:
func (_m *ForwarderManager[ADDR]) Name() string {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Name")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Ready provides a mock function with given fields:
func (_m *ForwarderManager[ADDR]) Ready() error {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Ready")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Start provides a mock function with given fields: _a0
func (_m *ForwarderManager[ADDR]) Start(_a0 context.Context) error {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for Start")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewForwarderManager creates a new instance of ForwarderManager. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewForwarderManager[ADDR types.Hashable](t interface {
	mock.TestingT
	Cleanup(func())
}) *ForwarderManager[ADDR] {
	mock := &ForwarderManager[ADDR]{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
