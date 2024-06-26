// Code generated by mockery v2.42.0. DO NOT EDIT.

package ecore

import (
	io "io"

	mock "github.com/stretchr/testify/mock"
)

// MockEURIConverter is an autogenerated mock type for the EURIConverter type
type MockEURIConverter struct {
	mock.Mock
}

type MockEURIConverter_Expecter struct {
	mock *mock.Mock
}

func (_m *MockEURIConverter) EXPECT() *MockEURIConverter_Expecter {
	return &MockEURIConverter_Expecter{mock: &_m.Mock}
}

// CreateReader provides a mock function with given fields: uri
func (_m *MockEURIConverter) CreateReader(uri *URI) (io.ReadCloser, error) {
	ret := _m.Called(uri)

	if len(ret) == 0 {
		panic("no return value specified for CreateReader")
	}

	var r0 io.ReadCloser
	var r1 error
	if rf, ok := ret.Get(0).(func(*URI) (io.ReadCloser, error)); ok {
		return rf(uri)
	}
	if rf, ok := ret.Get(0).(func(*URI) io.ReadCloser); ok {
		r0 = rf(uri)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(io.ReadCloser)
		}
	}

	if rf, ok := ret.Get(1).(func(*URI) error); ok {
		r1 = rf(uri)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockEURIConverter_CreateReader_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateReader'
type MockEURIConverter_CreateReader_Call struct {
	*mock.Call
}

// CreateReader is a helper method to define mock.On call
//   - uri *URI
func (_e *MockEURIConverter_Expecter) CreateReader(uri interface{}) *MockEURIConverter_CreateReader_Call {
	return &MockEURIConverter_CreateReader_Call{Call: _e.mock.On("CreateReader", uri)}
}

func (_c *MockEURIConverter_CreateReader_Call) Run(run func(uri *URI)) *MockEURIConverter_CreateReader_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*URI))
	})
	return _c
}

func (_c *MockEURIConverter_CreateReader_Call) Return(_a0 io.ReadCloser, _a1 error) *MockEURIConverter_CreateReader_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockEURIConverter_CreateReader_Call) RunAndReturn(run func(*URI) (io.ReadCloser, error)) *MockEURIConverter_CreateReader_Call {
	_c.Call.Return(run)
	return _c
}

// CreateWriter provides a mock function with given fields: uri
func (_m *MockEURIConverter) CreateWriter(uri *URI) (io.WriteCloser, error) {
	ret := _m.Called(uri)

	if len(ret) == 0 {
		panic("no return value specified for CreateWriter")
	}

	var r0 io.WriteCloser
	var r1 error
	if rf, ok := ret.Get(0).(func(*URI) (io.WriteCloser, error)); ok {
		return rf(uri)
	}
	if rf, ok := ret.Get(0).(func(*URI) io.WriteCloser); ok {
		r0 = rf(uri)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(io.WriteCloser)
		}
	}

	if rf, ok := ret.Get(1).(func(*URI) error); ok {
		r1 = rf(uri)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockEURIConverter_CreateWriter_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateWriter'
type MockEURIConverter_CreateWriter_Call struct {
	*mock.Call
}

// CreateWriter is a helper method to define mock.On call
//   - uri *URI
func (_e *MockEURIConverter_Expecter) CreateWriter(uri interface{}) *MockEURIConverter_CreateWriter_Call {
	return &MockEURIConverter_CreateWriter_Call{Call: _e.mock.On("CreateWriter", uri)}
}

func (_c *MockEURIConverter_CreateWriter_Call) Run(run func(uri *URI)) *MockEURIConverter_CreateWriter_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*URI))
	})
	return _c
}

func (_c *MockEURIConverter_CreateWriter_Call) Return(_a0 io.WriteCloser, _a1 error) *MockEURIConverter_CreateWriter_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockEURIConverter_CreateWriter_Call) RunAndReturn(run func(*URI) (io.WriteCloser, error)) *MockEURIConverter_CreateWriter_Call {
	_c.Call.Return(run)
	return _c
}

// GetURIHandler provides a mock function with given fields: uri
func (_m *MockEURIConverter) GetURIHandler(uri *URI) EURIHandler {
	ret := _m.Called(uri)

	if len(ret) == 0 {
		panic("no return value specified for GetURIHandler")
	}

	var r0 EURIHandler
	if rf, ok := ret.Get(0).(func(*URI) EURIHandler); ok {
		r0 = rf(uri)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(EURIHandler)
		}
	}

	return r0
}

// MockEURIConverter_GetURIHandler_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetURIHandler'
type MockEURIConverter_GetURIHandler_Call struct {
	*mock.Call
}

// GetURIHandler is a helper method to define mock.On call
//   - uri *URI
func (_e *MockEURIConverter_Expecter) GetURIHandler(uri interface{}) *MockEURIConverter_GetURIHandler_Call {
	return &MockEURIConverter_GetURIHandler_Call{Call: _e.mock.On("GetURIHandler", uri)}
}

func (_c *MockEURIConverter_GetURIHandler_Call) Run(run func(uri *URI)) *MockEURIConverter_GetURIHandler_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*URI))
	})
	return _c
}

func (_c *MockEURIConverter_GetURIHandler_Call) Return(_a0 EURIHandler) *MockEURIConverter_GetURIHandler_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockEURIConverter_GetURIHandler_Call) RunAndReturn(run func(*URI) EURIHandler) *MockEURIConverter_GetURIHandler_Call {
	_c.Call.Return(run)
	return _c
}

// GetURIHandlers provides a mock function with given fields:
func (_m *MockEURIConverter) GetURIHandlers() EList {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetURIHandlers")
	}

	var r0 EList
	if rf, ok := ret.Get(0).(func() EList); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(EList)
		}
	}

	return r0
}

// MockEURIConverter_GetURIHandlers_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetURIHandlers'
type MockEURIConverter_GetURIHandlers_Call struct {
	*mock.Call
}

// GetURIHandlers is a helper method to define mock.On call
func (_e *MockEURIConverter_Expecter) GetURIHandlers() *MockEURIConverter_GetURIHandlers_Call {
	return &MockEURIConverter_GetURIHandlers_Call{Call: _e.mock.On("GetURIHandlers")}
}

func (_c *MockEURIConverter_GetURIHandlers_Call) Run(run func()) *MockEURIConverter_GetURIHandlers_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockEURIConverter_GetURIHandlers_Call) Return(_a0 EList) *MockEURIConverter_GetURIHandlers_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockEURIConverter_GetURIHandlers_Call) RunAndReturn(run func() EList) *MockEURIConverter_GetURIHandlers_Call {
	_c.Call.Return(run)
	return _c
}

// GetURIMap provides a mock function with given fields:
func (_m *MockEURIConverter) GetURIMap() map[URI]URI {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetURIMap")
	}

	var r0 map[URI]URI
	if rf, ok := ret.Get(0).(func() map[URI]URI); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[URI]URI)
		}
	}

	return r0
}

// MockEURIConverter_GetURIMap_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetURIMap'
type MockEURIConverter_GetURIMap_Call struct {
	*mock.Call
}

// GetURIMap is a helper method to define mock.On call
func (_e *MockEURIConverter_Expecter) GetURIMap() *MockEURIConverter_GetURIMap_Call {
	return &MockEURIConverter_GetURIMap_Call{Call: _e.mock.On("GetURIMap")}
}

func (_c *MockEURIConverter_GetURIMap_Call) Run(run func()) *MockEURIConverter_GetURIMap_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockEURIConverter_GetURIMap_Call) Return(_a0 map[URI]URI) *MockEURIConverter_GetURIMap_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockEURIConverter_GetURIMap_Call) RunAndReturn(run func() map[URI]URI) *MockEURIConverter_GetURIMap_Call {
	_c.Call.Return(run)
	return _c
}

// Normalize provides a mock function with given fields: uri
func (_m *MockEURIConverter) Normalize(uri *URI) *URI {
	ret := _m.Called(uri)

	if len(ret) == 0 {
		panic("no return value specified for Normalize")
	}

	var r0 *URI
	if rf, ok := ret.Get(0).(func(*URI) *URI); ok {
		r0 = rf(uri)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*URI)
		}
	}

	return r0
}

// MockEURIConverter_Normalize_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Normalize'
type MockEURIConverter_Normalize_Call struct {
	*mock.Call
}

// Normalize is a helper method to define mock.On call
//   - uri *URI
func (_e *MockEURIConverter_Expecter) Normalize(uri interface{}) *MockEURIConverter_Normalize_Call {
	return &MockEURIConverter_Normalize_Call{Call: _e.mock.On("Normalize", uri)}
}

func (_c *MockEURIConverter_Normalize_Call) Run(run func(uri *URI)) *MockEURIConverter_Normalize_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*URI))
	})
	return _c
}

func (_c *MockEURIConverter_Normalize_Call) Return(_a0 *URI) *MockEURIConverter_Normalize_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockEURIConverter_Normalize_Call) RunAndReturn(run func(*URI) *URI) *MockEURIConverter_Normalize_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockEURIConverter creates a new instance of MockEURIConverter. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockEURIConverter(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockEURIConverter {
	mock := &MockEURIConverter{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
