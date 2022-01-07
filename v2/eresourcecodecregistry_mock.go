// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package ecore

import mock "github.com/stretchr/testify/mock"

// MockEResourceCodecRegistry is an autogenerated mock type for the EResourceCodecRegistry type
type MockEResourceCodecRegistry struct {
	mock.Mock
}

// GetCodec provides a mock function with given fields: uri
func (_m *MockEResourceCodecRegistry) GetCodec(uri *URI) EResourceCodec {
	ret := _m.Called(uri)

	var r0 EResourceCodec
	if rf, ok := ret.Get(0).(func(*URI) EResourceCodec); ok {
		r0 = rf(uri)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(EResourceCodec)
		}
	}

	return r0
}

// GetExtensionToCodecMap provides a mock function with given fields:
func (_m *MockEResourceCodecRegistry) GetExtensionToCodecMap() map[string]EResourceCodec {
	ret := _m.Called()

	var r0 map[string]EResourceCodec
	if rf, ok := ret.Get(0).(func() map[string]EResourceCodec); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]EResourceCodec)
		}
	}

	return r0
}

// GetProtocolToCodecMap provides a mock function with given fields:
func (_m *MockEResourceCodecRegistry) GetProtocolToCodecMap() map[string]EResourceCodec {
	ret := _m.Called()

	var r0 map[string]EResourceCodec
	if rf, ok := ret.Get(0).(func() map[string]EResourceCodec); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]EResourceCodec)
		}
	}

	return r0
}
