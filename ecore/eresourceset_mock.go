// *****************************************************************************
// Copyright(c) 2021 MASA Group
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// *****************************************************************************

package ecore

// MockEResourceSet is an autogenerated mock type for the EResourceSet type
type MockEResourceSet struct {
	MockENotifier
}

// CreateResource provides a mock function with given fields: uri
func (_m *MockEResourceSet) CreateResource(uri *URI) EResource {
	ret := _m.Called(uri)

	var r0 EResource
	if rf, ok := ret.Get(0).(func(*URI) EResource); ok {
		r0 = rf(uri)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(EResource)
		}
	}

	return r0
}

// GetEObject provides a mock function with given fields: uri, loadOnDemand
func (_m *MockEResourceSet) GetEObject(uri *URI, loadOnDemand bool) EObject {
	ret := _m.Called(uri, loadOnDemand)

	var r0 EObject
	if rf, ok := ret.Get(0).(func(*URI, bool) EObject); ok {
		r0 = rf(uri, loadOnDemand)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(EObject)
		}
	}

	return r0
}

// GetPackageRegistry provides a mock function with given fields:
func (_m *MockEResourceSet) GetPackageRegistry() EPackageRegistry {
	ret := _m.Called()

	var r0 EPackageRegistry
	if rf, ok := ret.Get(0).(func() EPackageRegistry); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(EPackageRegistry)
		}
	}

	return r0
}

// GetResource provides a mock function with given fields: uri, loadOnDemand
func (_m *MockEResourceSet) GetResource(uri *URI, loadOnDemand bool) EResource {
	ret := _m.Called(uri, loadOnDemand)

	var r0 EResource
	if rf, ok := ret.Get(0).(func(*URI, bool) EResource); ok {
		r0 = rf(uri, loadOnDemand)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(EResource)
		}
	}

	return r0
}

// GetResourceCodecRegistry provides a mock function with given fields:
func (_m *MockEResourceSet) GetResourceCodecRegistry() EResourceCodecRegistry {
	ret := _m.Called()

	var r0 EResourceCodecRegistry
	if rf, ok := ret.Get(0).(func() EResourceCodecRegistry); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(EResourceCodecRegistry)
		}
	}

	return r0
}

// GetResources provides a mock function with given fields:
func (_m *MockEResourceSet) GetResources() EList {
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

// GetURIConverter provides a mock function with given fields:
func (_m *MockEResourceSet) GetURIConverter() EURIConverter {
	ret := _m.Called()

	var r0 EURIConverter
	if rf, ok := ret.Get(0).(func() EURIConverter); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(EURIConverter)
		}
	}

	return r0
}

// GetURIResourceMap provides a mock function with given fields:
func (_m *MockEResourceSet) GetURIResourceMap() map[*URI]EResource {
	ret := _m.Called()

	var r0 map[*URI]EResource
	if rf, ok := ret.Get(0).(func() map[*URI]EResource); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[*URI]EResource)
		}
	}

	return r0
}

// SetPackageRegistry provides a mock function with given fields: packageregistry
func (_m *MockEResourceSet) SetPackageRegistry(packageregistry EPackageRegistry) {
	_m.Called(packageregistry)
}

// SetResourceCodecRegistry provides a mock function with given fields: resourceCodecRegistry
func (_m *MockEResourceSet) SetResourceCodecRegistry(resourceCodecRegistry EResourceCodecRegistry) {
	_m.Called(resourceCodecRegistry)
}

// SetURIConverter provides a mock function with given fields: uriConverter
func (_m *MockEResourceSet) SetURIConverter(uriConverter EURIConverter) {
	_m.Called(uriConverter)
}

// SetURIResourceMap provides a mock function with given fields: uriMap
func (_m *MockEResourceSet) SetURIResourceMap(uriMap map[*URI]EResource) {
	_m.Called(uriMap)
}
