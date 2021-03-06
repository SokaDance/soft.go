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

type MockEReference struct {
	MockEStructuralFeature
}

// IsContainer get the value of isContainer
func (eReference *MockEReference) IsContainer() bool {
	ret := eReference.Called()

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

// IsContainment get the value of isContainment
func (eReference *MockEReference) IsContainment() bool {
	ret := eReference.Called()

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

// SetContainment provides mock implementation for setting the value of isContainment
func (eReference *MockEReference) SetContainment(newIsContainment bool) {
	eReference.Called(newIsContainment)
}

// GetEKeys get the value of eKeys
func (eReference *MockEReference) GetEKeys() EList {
	ret := eReference.Called()

	var r EList
	if rf, ok := ret.Get(0).(func() EList); ok {
		r = rf()
	} else {
		if ret.Get(0) != nil {
			r = ret.Get(0).(EList)
		}
	}

	return r
}

// GetEOpposite get the value of eOpposite
func (eReference *MockEReference) GetEOpposite() EReference {
	ret := eReference.Called()

	var r EReference
	if rf, ok := ret.Get(0).(func() EReference); ok {
		r = rf()
	} else {
		if ret.Get(0) != nil {
			r = ret.Get(0).(EReference)
		}
	}

	return r
}

// SetEOpposite provides mock implementation for setting the value of eOpposite
func (eReference *MockEReference) SetEOpposite(newEOpposite EReference) {
	eReference.Called(newEOpposite)
}

// GetEReferenceType get the value of eReferenceType
func (eReference *MockEReference) GetEReferenceType() EClass {
	ret := eReference.Called()

	var r EClass
	if rf, ok := ret.Get(0).(func() EClass); ok {
		r = rf()
	} else {
		if ret.Get(0) != nil {
			r = ret.Get(0).(EClass)
		}
	}

	return r
}

// IsResolveProxies get the value of isResolveProxies
func (eReference *MockEReference) IsResolveProxies() bool {
	ret := eReference.Called()

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

// SetResolveProxies provides mock implementation for setting the value of isResolveProxies
func (eReference *MockEReference) SetResolveProxies(newIsResolveProxies bool) {
	eReference.Called(newIsResolveProxies)
}
