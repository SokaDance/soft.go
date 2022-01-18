// Code generated by soft.generator.go. DO NOT EDIT.

package ecore

type MockETypedElement struct {
	MockENamedElement
}

// GetEType get the value of eType
func (eTypedElement *MockETypedElement) GetEType() EClassifier {
	ret := eTypedElement.Called()

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

// SetEType provides mock implementation for setting the value of eType
func (eTypedElement *MockETypedElement) SetEType(newEType EClassifier) {
	eTypedElement.Called(newEType)
}

// UnsetEType provides mock implementation for unset the value of eType
func (eTypedElement *MockETypedElement) UnsetEType() {
	eTypedElement.Called()
}

// GetLowerBound get the value of lowerBound
func (eTypedElement *MockETypedElement) GetLowerBound() int {
	ret := eTypedElement.Called()

	var r int
	if rf, ok := ret.Get(0).(func() int); ok {
		r = rf()
	} else {
		if ret.Get(0) != nil {
			r = ret.Get(0).(int)
		}
	}

	return r
}

// SetLowerBound provides mock implementation for setting the value of lowerBound
func (eTypedElement *MockETypedElement) SetLowerBound(newLowerBound int) {
	eTypedElement.Called(newLowerBound)
}

// IsMany get the value of isMany
func (eTypedElement *MockETypedElement) IsMany() bool {
	ret := eTypedElement.Called()

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

// IsOrdered get the value of isOrdered
func (eTypedElement *MockETypedElement) IsOrdered() bool {
	ret := eTypedElement.Called()

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

// SetOrdered provides mock implementation for setting the value of isOrdered
func (eTypedElement *MockETypedElement) SetOrdered(newIsOrdered bool) {
	eTypedElement.Called(newIsOrdered)
}

// IsRequired get the value of isRequired
func (eTypedElement *MockETypedElement) IsRequired() bool {
	ret := eTypedElement.Called()

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

// IsUnique get the value of isUnique
func (eTypedElement *MockETypedElement) IsUnique() bool {
	ret := eTypedElement.Called()

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

// SetUnique provides mock implementation for setting the value of isUnique
func (eTypedElement *MockETypedElement) SetUnique(newIsUnique bool) {
	eTypedElement.Called(newIsUnique)
}

// GetUpperBound get the value of upperBound
func (eTypedElement *MockETypedElement) GetUpperBound() int {
	ret := eTypedElement.Called()

	var r int
	if rf, ok := ret.Get(0).(func() int); ok {
		r = rf()
	} else {
		if ret.Get(0) != nil {
			r = ret.Get(0).(int)
		}
	}

	return r
}

// SetUpperBound provides mock implementation for setting the value of upperBound
func (eTypedElement *MockETypedElement) SetUpperBound(newUpperBound int) {
	eTypedElement.Called(newUpperBound)
}
