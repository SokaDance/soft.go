// Code generated by soft.generator.go. DO NOT EDIT.



package ecore


type MockEOperation struct {
	MockETypedElement
}

// GetEContainingClass get the value of eContainingClass
func (eOperation *MockEOperation) GetEContainingClass() EClass {
	ret := eOperation.Called()

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

// GetEExceptions get the value of eExceptions
func (eOperation *MockEOperation) GetEExceptions() EList[EClassifier] {
	ret := eOperation.Called()

	var r EList[EClassifier]
	if rf, ok := ret.Get(0).(func() EList[EClassifier]); ok {
		r = rf()
	} else {
		if ret.Get(0) != nil {
			r = ret.Get(0).(EList[EClassifier])
		}
	}

	return r
}

// UnsetEExceptions provides mock implementation for unset the value of eExceptions
func (eOperation *MockEOperation) UnsetEExceptions() {
	eOperation.Called()
}

// GetEParameters get the value of eParameters
func (eOperation *MockEOperation) GetEParameters() EList[EParameter] {
	ret := eOperation.Called()

	var r EList[EParameter]
	if rf, ok := ret.Get(0).(func() EList[EParameter]); ok {
		r = rf()
	} else {
		if ret.Get(0) != nil {
			r = ret.Get(0).(EList[EParameter])
		}
	}

	return r
}

// GetOperationID get the value of operationID
func (eOperation *MockEOperation) GetOperationID() int {
	ret := eOperation.Called()

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

// SetOperationID provides mock implementation for setting the value of operationID
func (eOperation *MockEOperation) SetOperationID( newOperationID int ) {
	eOperation.Called(newOperationID)
}


// IsOverrideOf provides mock implementation
func (eOperation *MockEOperation) IsOverrideOf(someOperation EOperation) bool {
	ret := eOperation.Called(someOperation)

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


