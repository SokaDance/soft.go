// Code generated by soft.generator.go. DO NOT EDIT.



package ecore


type MockEEnum struct {
	MockEDataType
}

// GetELiterals get the value of eLiterals
func (eEnum *MockEEnum) GetELiterals() EList[EEnumLiteral] {
	ret := eEnum.Called()

	var r EList[EEnumLiteral]
	if rf, ok := ret.Get(0).(func() EList[EEnumLiteral]); ok {
		r = rf()
	} else {
		if ret.Get(0) != nil {
			r = ret.Get(0).(EList[EEnumLiteral])
		}
	}

	return r
}


// GetEEnumLiteralByLiteral provides mock implementation
func (eEnum *MockEEnum) GetEEnumLiteralByLiteral(literal string) EEnumLiteral {
	ret := eEnum.Called(literal)

	var r EEnumLiteral
	if rf, ok := ret.Get(0).(func() EEnumLiteral); ok {
		r = rf()
	} else {
		if ret.Get(0) != nil {
			r = ret.Get(0).(EEnumLiteral)
		}
	}

	return r
}

// GetEEnumLiteralByName provides mock implementation
func (eEnum *MockEEnum) GetEEnumLiteralByName(name string) EEnumLiteral {
	ret := eEnum.Called(name)

	var r EEnumLiteral
	if rf, ok := ret.Get(0).(func() EEnumLiteral); ok {
		r = rf()
	} else {
		if ret.Get(0) != nil {
			r = ret.Get(0).(EEnumLiteral)
		}
	}

	return r
}

// GetEEnumLiteralByValue provides mock implementation
func (eEnum *MockEEnum) GetEEnumLiteralByValue(value int) EEnumLiteral {
	ret := eEnum.Called(value)

	var r EEnumLiteral
	if rf, ok := ret.Get(0).(func() EEnumLiteral); ok {
		r = rf()
	} else {
		if ret.Get(0) != nil {
			r = ret.Get(0).(EEnumLiteral)
		}
	}

	return r
}



