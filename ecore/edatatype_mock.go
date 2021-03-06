// Code generated by soft.generator.go. DO NOT EDIT.

// *****************************************************************************
// Copyright(c) 2021 MASA Group
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// *****************************************************************************

package ecore

type MockEDataType struct {
	MockEClassifier
}

// IsSerializable get the value of isSerializable
func (eDataType *MockEDataType) IsSerializable() bool {
	ret := eDataType.Called()

	var r bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r = rf()
	} else {
		if ret.Get(0) != nil {
			r = ret.Get(0).(bool)
		}
	}

	return r
}

// SetSerializable provides mock implementation for setting the value of isSerializable
func (eDataType *MockEDataType) SetSerializable(newIsSerializable bool) {
	eDataType.Called(newIsSerializable)
}
