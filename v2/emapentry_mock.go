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

// MockEMapEntry is an autogenerated mock type for the EMapEntry type
type MockEMapEntry[K comparable,V any] struct {
	mock.Mock
}

// GetKey provides a mock function with given fields:
func (_m *MockEMapEntry[K,V]) GetKey() K {
	ret := _m.Called()

	var r0 K
	if rf, ok := ret.Get(0).(func() K); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(K)
		}
	}

	return r0
}

// GetValue provides a mock function with given fields:
func (_m *MockEMapEntry[K,V]) GetValue() V {
	ret := _m.Called()

	var r0 V
	if rf, ok := ret.Get(0).(func() V); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(V)
		}
	}

	return r0
}

// SetKey provides a mock function with given fields: _a0
func (_m *MockEMapEntry[K,V]) SetKey(_a0 K) {
	_m.Called(_a0)
}

// SetValue provides a mock function with given fields: _a0
func (_m *MockEMapEntry[K,V]) SetValue(_a0 V) {
	_m.Called(_a0)
}