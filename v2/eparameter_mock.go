// Code generated by soft.generator.go. DO NOT EDIT.



package ecore


type MockEParameter struct {
	MockETypedElement
}

// GetEOperation get the value of eOperation
func (eParameter *MockEParameter) GetEOperation() EOperation {
	ret := eParameter.Called()

	var r EOperation
	if rf, ok := ret.Get(0).(func() EOperation); ok {
		r = rf()
	} else {
		if ret.Get(0) != nil {
			r = ret.Get(0).(EOperation)
		}
	}

	return r
}




