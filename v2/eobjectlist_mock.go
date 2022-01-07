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
type MockEObjectList[T EObject] struct {
	MockEList[T]
}

// GetUnResolvedList provides a mock function with given fields:
func (_m *MockEObjectList[T]) GetUnResolvedList() EList[T] {
	ret := _m.Called()

	var r0 EList[T]
	if rf, ok := ret.Get(0).(func() EList[T]); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(EList[T])
		}
	}

	return r0
}
