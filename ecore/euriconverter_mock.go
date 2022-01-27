// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package ecore

import (
	mock "github.com/stretchr/testify/mock"
	io "io"
)

// MockEURIConverter is an autogenerated mock type for the EURIConverter type
type MockEURIConverter struct {
	mock.Mock
}

// CreateReader provides a mock function with given fields: uri
func (_m *MockEURIConverter) CreateReader(uri *URI) (io.ReadCloser, error) {
	ret := _m.Called(uri)

	var r0 io.ReadCloser
	var r1 error
	if rf, ok := ret.Get(0).(func(*URI) (io.ReadCloser, error)); ok {
		r0, r1 = rf(uri)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(io.ReadCloser)
		}
		r1 = ret.Error(1)
	}
	return r0, r1
}

// CreateWriter provides a mock function with given fields: uri
func (_m *MockEURIConverter) CreateWriter(uri *URI) (io.WriteCloser, error) {
	ret := _m.Called(uri)

	var r0 io.WriteCloser
	var r1 error
	if rf, ok := ret.Get(0).(func(*URI) (io.WriteCloser, error)); ok {
		r0, r1 = rf(uri)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(io.WriteCloser)
		}
		r1 = ret.Error(1)
	}
	return r0, r1
}

// GetURIHandler provides a mock function with given fields: uri
func (_m *MockEURIConverter) GetURIHandler(uri *URI) EURIHandler {
	ret := _m.Called(uri)

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

// GetURIHandlers provides a mock function with given fields:
func (_m *MockEURIConverter) GetURIHandlers() EList[EURIHandler] {
	ret := _m.Called()

	var r0 EList[EURIHandler]
	if rf, ok := ret.Get(0).(func() EList[EURIHandler]); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(EList[EURIHandler])
		}
	}

	return r0
}

// Normalize provides a mock function with given fields: uri
func (_m *MockEURIConverter) Normalize(uri *URI) *URI {
	ret := _m.Called(uri)

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

// GetURIMap provides a mock function with given fields:
func (_m *MockEURIConverter) GetURIMap() map[URI]URI {
	ret := _m.Called()

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
