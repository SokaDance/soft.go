// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

// *****************************************************************************
// Copyright(c) 2021 MASA Group
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// *****************************************************************************

package ecore

import mock "github.com/stretchr/testify/mock"

type MockECollection[T any] struct {
	mock.Mock
}

// Add provides a mock function with given fields: _a0
func (_m *MockECollection[T]) Add(_a0 T) bool {
	ret := _m.Called(_a0)

	var r0 bool
	if rf, ok := ret.Get(0).(func(T) bool); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// AddAll provides a mock function with given fields: _a0
func (_m *MockECollection[T]) AddAll(_a0 ECollection[T]) bool {
	ret := _m.Called(_a0)

	var r0 bool
	if rf, ok := ret.Get(0).(func(ECollection[T]) bool); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// Clear provides a mock function with given fields:
func (_m *MockECollection[T]) Clear() {
	_m.Called()
}

// Contains provides a mock function with given fields: _a0
func (_m *MockECollection[T]) Contains(_a0 T) bool {
	ret := _m.Called(_a0)

	var r0 bool
	if rf, ok := ret.Get(0).(func(T) bool); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// Empty provides a mock function with given fields:
func (_m *MockECollection[T]) Empty() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// Iterator provides a mock function with given fields:
func (_m *MockECollection[T]) Iterator() EIterator[T] {
	ret := _m.Called()

	var r0 EIterator[T]
	if rf, ok := ret.Get(0).(func() EIterator[T]); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(EIterator[T])
		}
	}

	return r0
}


// Remove provides a mock function with given fields: _a0
func (_m *MockECollection[T]) Remove(_a0 T) bool {
	ret := _m.Called(_a0)

	var r0 bool
	if rf, ok := ret.Get(0).(func(T) bool); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// RemoveAll provides a mock function with given fields: _a0
func (_m *MockECollection[T]) RemoveAll(_a0 ECollection[T]) bool {
	ret := _m.Called(_a0)

	var r0 bool
	if rf, ok := ret.Get(0).(func(ECollection[T]) bool); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// RemoveAll provides a mock function with given fields: _a0
func (_m *MockECollection[T]) RetainAll(_a0 ECollection[T]) bool {
	ret := _m.Called(_a0)

	var r0 bool
	if rf, ok := ret.Get(0).(func(ECollection[T]) bool); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}


// Size provides a mock function with given fields:
func (_m *MockECollection[T]) Size() int {
	ret := _m.Called()

	var r0 int
	if rf, ok := ret.Get(0).(func() int); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

// ToArray provides a mock function with given fields:
func (_m *MockECollection[T]) ToArray() []T {
	ret := _m.Called()

	var r0 []T
	if rf, ok := ret.Get(0).(func() []T); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]T)
		}
	}

	return r0
}