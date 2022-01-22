// *****************************************************************************
// Copyright(c) 2021 MASA Group
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// *****************************************************************************

package ecore

// MockEObjectList is an autogenerated mock type for the EObjectList type
type MockEObjectList[T any] struct {
	MockEList[T]
}

// GetUnResolvedList provides a mock function with given fields:
func (_m *MockEObjectList[T]) GetUnResolvedList() EObjectList[T] {
	ret := _m.Called()

	var r0 EObjectList[T]
	if rf, ok := ret.Get(0).(func() EObjectList[T]); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(EObjectList[T])
		}
	}

	return r0
}
