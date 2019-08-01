// Code generated by mockery v1.0.0. DO NOT EDIT.

package ecore

import mock "github.com/stretchr/testify/mock"

// MockEResource is an autogenerated mock type for the EResource type
type MockEResource struct {
	mock.Mock
}

// Attached provides a mock function with given fields: object
func (_m *MockEResource) Attached(object EObject) {
	_m.Called(object)
}

// Detached provides a mock function with given fields: object
func (_m *MockEResource) Detached(object EObject) {
	_m.Called(object)
}

// GetContents provides a mock function with given fields:
func (_m *MockEResource) GetContents() EList {
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