// *****************************************************************************
// Copyright(c) 2021 MASA Group
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// *****************************************************************************

package ecore

type MockEMap[K comparable,V any] struct {
	MockEList[EMapEntry[K,V]]
}

// ContainsKey provides a mock function with given fields: key
func (_m *MockEMap[K,V]) ContainsKey(key K) bool {
	ret := _m.Called(key)

	var r0 bool
	if rf, ok := ret.Get(0).(func(K) bool); ok {
		r0 = rf(key)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// ContainsValue provides a mock function with given fields: value
func (_m *MockEMap[K,V]) ContainsValue(value V) bool {
	ret := _m.Called(value)

	var r0 bool
	if rf, ok := ret.Get(0).(func(V) bool); ok {
		r0 = rf(value)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// GetValue provides a mock function with given fields: value
func (_m *MockEMap[K,V]) GetValue(value K) V {
	ret := _m.Called(value)

	var r0 V
	if rf, ok := ret.Get(0).(func(K) V); ok {
		r0 = rf(value)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(V)
		}
	}

	return r0
}

// Put provides a mock function with given fields: key, value
func (_m *MockEMap[K,V]) Put(key K, value V) {
	_m.Called(key, value)
}

// RemoveKey provides a mock function with given fields: key
func (_m *MockEMap[K,V]) RemoveKey(key K) V {
	ret := _m.Called(key)

	var r0 V
	if rf, ok := ret.Get(0).(func(K) V); ok {
		r0 = rf(key)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(V)
		}
	}

	return r0
}

// ToMap provides a mock function with given fields:
func (_m *MockEMap[K,V]) ToMap() map[K]V {
	ret := _m.Called()

	var r0 map[K]V
	if rf, ok := ret.Get(0).(func() map[K]V); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[K]V)
		}
	}

	return r0
}
