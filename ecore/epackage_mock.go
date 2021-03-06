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

type MockEPackage struct {
	MockENamedElement
}

// GetEClassifiers get the value of eClassifiers
func (ePackage *MockEPackage) GetEClassifiers() EList {
	ret := ePackage.Called()

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

// GetEFactoryInstance get the value of eFactoryInstance
func (ePackage *MockEPackage) GetEFactoryInstance() EFactory {
	ret := ePackage.Called()

	var r EFactory
	if rf, ok := ret.Get(0).(func() EFactory); ok {
		r = rf()
	} else {
		if ret.Get(0) != nil {
			r = ret.Get(0).(EFactory)
		}
	}

	return r
}

// SetEFactoryInstance provides mock implementation for setting the value of eFactoryInstance
func (ePackage *MockEPackage) SetEFactoryInstance(newEFactoryInstance EFactory) {
	ePackage.Called(newEFactoryInstance)
}

// GetESubPackages get the value of eSubPackages
func (ePackage *MockEPackage) GetESubPackages() EList {
	ret := ePackage.Called()

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

// GetESuperPackage get the value of eSuperPackage
func (ePackage *MockEPackage) GetESuperPackage() EPackage {
	ret := ePackage.Called()

	var r EPackage
	if rf, ok := ret.Get(0).(func() EPackage); ok {
		r = rf()
	} else {
		if ret.Get(0) != nil {
			r = ret.Get(0).(EPackage)
		}
	}

	return r
}

// GetNsPrefix get the value of nsPrefix
func (ePackage *MockEPackage) GetNsPrefix() string {
	ret := ePackage.Called()

	var r string
	if rf, ok := ret.Get(0).(func() string); ok {
		r = rf()
	} else {
		if ret.Get(0) != nil {
			r = ret.Get(0).(string)
		}
	}

	return r
}

// SetNsPrefix provides mock implementation for setting the value of nsPrefix
func (ePackage *MockEPackage) SetNsPrefix(newNsPrefix string) {
	ePackage.Called(newNsPrefix)
}

// GetNsURI get the value of nsURI
func (ePackage *MockEPackage) GetNsURI() string {
	ret := ePackage.Called()

	var r string
	if rf, ok := ret.Get(0).(func() string); ok {
		r = rf()
	} else {
		if ret.Get(0) != nil {
			r = ret.Get(0).(string)
		}
	}

	return r
}

// SetNsURI provides mock implementation for setting the value of nsURI
func (ePackage *MockEPackage) SetNsURI(newNsURI string) {
	ePackage.Called(newNsURI)
}

// GetEClassifier provides mock implementation
func (ePackage *MockEPackage) GetEClassifier(name string) EClassifier {
	ret := ePackage.Called(name)

	var r EClassifier
	if rf, ok := ret.Get(0).(func() EClassifier); ok {
		r = rf()
	} else {
		if ret.Get(0) != nil {
			r = ret.Get(0).(EClassifier)
		}
	}

	return r
}
