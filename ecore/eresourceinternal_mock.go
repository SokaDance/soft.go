package ecore

import (
	io "io"
)

// MockEResourceInternal is an autogenerated mock type for the EResourceInternal type
type MockEResourceInternal struct {
	MockEResource
}

// DoLoad provides a mock function with given fields: rd
func (_m *MockEResourceInternal) DoLoad(rd io.Reader, options map[string]interface{}) {
	_m.Called(rd, options)
}

// DoSave provides a mock function with given fields: rd
func (_m *MockEResourceInternal) DoSave(rd io.Writer, options map[string]interface{}) {
	_m.Called(rd, options)
}

// DoUnload provides a mock function with given fields:
func (_m *MockEResourceInternal) DoUnload() {
	_m.Called()
}

// basicSetLoaded provides a mock function with given fields: _a0, _a1
func (_m *MockEResourceInternal) basicSetLoaded(_a0 bool, _a1 ENotificationChain) ENotificationChain {
	ret := _m.Called(_a0, _a1)

	var r0 ENotificationChain
	if rf, ok := ret.Get(0).(func(bool, ENotificationChain) ENotificationChain); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(ENotificationChain)
		}
	}

	return r0
}

// basicSetResourceSet provides a mock function with given fields: _a0, _a1
func (_m *MockEResourceInternal) basicSetResourceSet(_a0 EResourceSet, _a1 ENotificationChain) ENotificationChain {
	ret := _m.Called(_a0, _a1)

	var r0 ENotificationChain
	if rf, ok := ret.Get(0).(func(EResourceSet, ENotificationChain) ENotificationChain); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(ENotificationChain)
		}
	}

	return r0
}
