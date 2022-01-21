// Code generated by soft.generator.go. DO NOT EDIT.

package ecore

type MockEStringToStringMapEntry struct {
	MockEObjectInternal
}

// GetStringKey get the value of key
func (eStringToStringMapEntry *MockEStringToStringMapEntry) GetStringKey() string {
	ret := eStringToStringMapEntry.Called()

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

// SetStringKey provides mock implementation for setting the value of key
func (eStringToStringMapEntry *MockEStringToStringMapEntry) SetStringKey(newKey string) {
	eStringToStringMapEntry.Called(newKey)
}

// GetStringValue get the value of value
func (eStringToStringMapEntry *MockEStringToStringMapEntry) GetStringValue() string {
	ret := eStringToStringMapEntry.Called()

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

// SetStringValue provides mock implementation for setting the value of value
func (eStringToStringMapEntry *MockEStringToStringMapEntry) SetStringValue(newValue string) {
	eStringToStringMapEntry.Called(newValue)
}
