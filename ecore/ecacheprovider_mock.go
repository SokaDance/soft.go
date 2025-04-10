// Code generated by mockery v2.53.2. DO NOT EDIT.

package ecore

import mock "github.com/stretchr/testify/mock"

// MockECacheProvider_Prototype is an autogenerated mock type for the ECacheProvider type
type MockECacheProvider struct {
	mock.Mock
	MockECacheProvider_Prototype
}

type MockECacheProvider_Prototype struct {
	mock *mock.Mock
}

func (_mp *MockECacheProvider_Prototype) SetMock(mock *mock.Mock) {
	_mp.mock = mock
}

type MockECacheProvider_Expecter struct {
	mock *mock.Mock
}

func (_me *MockECacheProvider_Expecter) SetMock(mock *mock.Mock) {
	_me.mock = mock
}

func (_m *MockECacheProvider_Prototype) EXPECT() *MockECacheProvider_Expecter {
	return &MockECacheProvider_Expecter{mock: _m.mock}
}

// IsCache provides a mock function with no fields
func (_m *MockECacheProvider_Prototype) IsCache() bool {
	ret := _m.mock.Called()

	if len(ret) == 0 {
		panic("no return value specified for IsCache")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// MockECacheProvider_IsCache_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'IsCache'
type MockECacheProvider_IsCache_Call struct {
	*mock.Call
}

// IsCache is a helper method to define mock.On call
func (_e *MockECacheProvider_Expecter) IsCache() *MockECacheProvider_IsCache_Call {
	return &MockECacheProvider_IsCache_Call{Call: _e.mock.On("IsCache")}
}

func (_c *MockECacheProvider_IsCache_Call) Run(run func()) *MockECacheProvider_IsCache_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockECacheProvider_IsCache_Call) Return(_a0 bool) *MockECacheProvider_IsCache_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockECacheProvider_IsCache_Call) RunAndReturn(run func() bool) *MockECacheProvider_IsCache_Call {
	_c.Call.Return(run)
	return _c
}

// SetCache provides a mock function with given fields: _a0
func (_m *MockECacheProvider_Prototype) SetCache(_a0 bool) {
	_m.mock.Called(_a0)
}

// MockECacheProvider_SetCache_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetCache'
type MockECacheProvider_SetCache_Call struct {
	*mock.Call
}

// SetCache is a helper method to define mock.On call
//   - _a0 bool
func (_e *MockECacheProvider_Expecter) SetCache(_a0 interface{}) *MockECacheProvider_SetCache_Call {
	return &MockECacheProvider_SetCache_Call{Call: _e.mock.On("SetCache", _a0)}
}

func (_c *MockECacheProvider_SetCache_Call) Run(run func(_a0 bool)) *MockECacheProvider_SetCache_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(bool))
	})
	return _c
}

func (_c *MockECacheProvider_SetCache_Call) Return() *MockECacheProvider_SetCache_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockECacheProvider_SetCache_Call) RunAndReturn(run func(bool)) *MockECacheProvider_SetCache_Call {
	_c.Run(run)
	return _c
}

// NewMockECacheProvider creates a new instance of MockECacheProvider. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockECacheProvider(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockECacheProvider {
	mock := &MockECacheProvider{}
	mock.SetMock(&mock.Mock)
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
