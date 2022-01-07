// Code generated by soft.generator.go. DO NOT EDIT.



package ecore

import "reflect"

type MockEStructuralFeature struct {
	MockETypedElement
}

// IsChangeable get the value of isChangeable
func (eStructuralFeature *MockEStructuralFeature) IsChangeable() bool {
	ret := eStructuralFeature.Called()

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

// SetChangeable provides mock implementation for setting the value of isChangeable
func (eStructuralFeature *MockEStructuralFeature) SetChangeable( newIsChangeable bool ) {
	eStructuralFeature.Called(newIsChangeable)
}

// GetDefaultValue get the value of defaultValue
func (eStructuralFeature *MockEStructuralFeature) GetDefaultValue() any {
	ret := eStructuralFeature.Called()

	var r any
	if rf, ok := ret.Get(0).(func() any); ok {
		r = rf()
	} else {
		if ret.Get(0) != nil {
			r = ret.Get(0).(any)
		}
	}

	return r
}

// SetDefaultValue provides mock implementation for setting the value of defaultValue
func (eStructuralFeature *MockEStructuralFeature) SetDefaultValue( newDefaultValue any ) {
	eStructuralFeature.Called(newDefaultValue)
}

// GetDefaultValueLiteral get the value of defaultValueLiteral
func (eStructuralFeature *MockEStructuralFeature) GetDefaultValueLiteral() string {
	ret := eStructuralFeature.Called()

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

// SetDefaultValueLiteral provides mock implementation for setting the value of defaultValueLiteral
func (eStructuralFeature *MockEStructuralFeature) SetDefaultValueLiteral( newDefaultValueLiteral string ) {
	eStructuralFeature.Called(newDefaultValueLiteral)
}

// IsDerived get the value of isDerived
func (eStructuralFeature *MockEStructuralFeature) IsDerived() bool {
	ret := eStructuralFeature.Called()

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

// SetDerived provides mock implementation for setting the value of isDerived
func (eStructuralFeature *MockEStructuralFeature) SetDerived( newIsDerived bool ) {
	eStructuralFeature.Called(newIsDerived)
}

// GetEContainingClass get the value of eContainingClass
func (eStructuralFeature *MockEStructuralFeature) GetEContainingClass() EClass {
	ret := eStructuralFeature.Called()

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

// GetFeatureID get the value of featureID
func (eStructuralFeature *MockEStructuralFeature) GetFeatureID() int {
	ret := eStructuralFeature.Called()

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

// SetFeatureID provides mock implementation for setting the value of featureID
func (eStructuralFeature *MockEStructuralFeature) SetFeatureID( newFeatureID int ) {
	eStructuralFeature.Called(newFeatureID)
}

// IsTransient get the value of isTransient
func (eStructuralFeature *MockEStructuralFeature) IsTransient() bool {
	ret := eStructuralFeature.Called()

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

// SetTransient provides mock implementation for setting the value of isTransient
func (eStructuralFeature *MockEStructuralFeature) SetTransient( newIsTransient bool ) {
	eStructuralFeature.Called(newIsTransient)
}

// IsUnsettable get the value of isUnsettable
func (eStructuralFeature *MockEStructuralFeature) IsUnsettable() bool {
	ret := eStructuralFeature.Called()

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

// SetUnsettable provides mock implementation for setting the value of isUnsettable
func (eStructuralFeature *MockEStructuralFeature) SetUnsettable( newIsUnsettable bool ) {
	eStructuralFeature.Called(newIsUnsettable)
}

// IsVolatile get the value of isVolatile
func (eStructuralFeature *MockEStructuralFeature) IsVolatile() bool {
	ret := eStructuralFeature.Called()

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

// SetVolatile provides mock implementation for setting the value of isVolatile
func (eStructuralFeature *MockEStructuralFeature) SetVolatile( newIsVolatile bool ) {
	eStructuralFeature.Called(newIsVolatile)
}


// GetContainerClass provides mock implementation
func (eStructuralFeature *MockEStructuralFeature) GetContainerClass() reflect.Type {
	ret := eStructuralFeature.Called()

	var r reflect.Type
	if rf, ok := ret.Get(0).(func() reflect.Type); ok {
		r = rf()
	} else {
		if ret.Get(0) != nil {
			r = ret.Get(0).(reflect.Type)
		}
	}

	return r
}



