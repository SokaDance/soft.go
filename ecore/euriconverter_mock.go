// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package ecore

import (
	io "io"
	url "net/url"

	mock "github.com/stretchr/testify/mock"
)

// MockEURIConverter is an autogenerated mock type for the EURIConverter type
type MockEURIConverter struct {
	mock.Mock
}

// CreateReader provides a mock function with given fields: uri
func (_m *MockEURIConverter) CreateReader(uri *url.URL) io.ReadCloser {
	ret := _m.Called(uri)

	var r0 io.ReadCloser
	if rf, ok := ret.Get(0).(func(*url.URL) io.ReadCloser); ok {
		r0 = rf(uri)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(io.ReadCloser)
		}
	}

	return r0
}

// CreateWriter provides a mock function with given fields: uri
func (_m *MockEURIConverter) CreateWriter(uri *url.URL) io.WriteCloser {
	ret := _m.Called(uri)

	var r0 io.WriteCloser
	if rf, ok := ret.Get(0).(func(*url.URL) io.WriteCloser); ok {
		r0 = rf(uri)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(io.WriteCloser)
		}
	}

	return r0
}

// GetURIHandler provides a mock function with given fields: uri
func (_m *MockEURIConverter) GetURIHandler(uri *url.URL) EURIHandler {
	ret := _m.Called(uri)

	var r0 EURIHandler
	if rf, ok := ret.Get(0).(func(*url.URL) EURIHandler); ok {
		r0 = rf(uri)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(EURIHandler)
		}
	}

	return r0
}

// GetURIHandlers provides a mock function with given fields:
func (_m *MockEURIConverter) GetURIHandlers() EList {
	ret := _m.Called()

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

// Normalize provides a mock function with given fields: uri
func (_m *MockEURIConverter) Normalize(uri *url.URL) *url.URL {
	ret := _m.Called(uri)

	var r0 *url.URL
	if rf, ok := ret.Get(0).(func(*url.URL) *url.URL); ok {
		r0 = rf(uri)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*url.URL)
		}
	}

	return r0
}
