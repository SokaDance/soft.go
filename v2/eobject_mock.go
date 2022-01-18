// Code generated by soft.generator.go. DO NOT EDIT.

package ecore

type MockEObject struct {
	MockENotifier
}

// EAllContents provides mock implementation
func (eObject *MockEObject) EAllContents() EIterator[EObject] {
	ret := eObject.Called()

	var r EIterator[EObject]
	if rf, ok := ret.Get(0).(func() EIterator[EObject]); ok {
		r = rf()
	} else {
		if ret.Get(0) != nil {
			r = ret.Get(0).(EIterator[EObject])
		}
	}

	return r
}

// EClass provides mock implementation
func (eObject *MockEObject) EClass() EClass {
	ret := eObject.Called()

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

// EContainer provides mock implementation
func (eObject *MockEObject) EContainer() EObject {
	ret := eObject.Called()

	var r EObject
	if rf, ok := ret.Get(0).(func() EObject); ok {
		r = rf()
	} else {
		if ret.Get(0) != nil {
			r = ret.Get(0).(EObject)
		}
	}

	return r
}

// EContainingFeature provides mock implementation
func (eObject *MockEObject) EContainingFeature() EStructuralFeature {
	ret := eObject.Called()

	var r EStructuralFeature
	if rf, ok := ret.Get(0).(func() EStructuralFeature); ok {
		r = rf()
	} else {
		if ret.Get(0) != nil {
			r = ret.Get(0).(EStructuralFeature)
		}
	}

	return r
}

// EContainmentFeature provides mock implementation
func (eObject *MockEObject) EContainmentFeature() EReference {
	ret := eObject.Called()

	var r EReference
	if rf, ok := ret.Get(0).(func() EReference); ok {
		r = rf()
	} else {
		if ret.Get(0) != nil {
			r = ret.Get(0).(EReference)
		}
	}

	return r
}

// EContents provides mock implementation
func (eObject *MockEObject) EContents() EList[EObject] {
	ret := eObject.Called()

	var r EList[EObject]
	if rf, ok := ret.Get(0).(func() EList[EObject]); ok {
		r = rf()
	} else {
		if ret.Get(0) != nil {
			r = ret.Get(0).(EList[EObject])
		}
	}

	return r
}

// ECrossReferences provides mock implementation
func (eObject *MockEObject) ECrossReferences() EList[EObject] {
	ret := eObject.Called()

	var r EList[EObject]
	if rf, ok := ret.Get(0).(func() EList[EObject]); ok {
		r = rf()
	} else {
		if ret.Get(0) != nil {
			r = ret.Get(0).(EList[EObject])
		}
	}

	return r
}

// EGet provides mock implementation
func (eObject *MockEObject) EGet(feature EStructuralFeature) any {
	ret := eObject.Called(feature)

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

// EGetResolve provides mock implementation
func (eObject *MockEObject) EGetResolve(feature EStructuralFeature, resolve bool) any {
	ret := eObject.Called(feature, resolve)

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

// EInvoke provides mock implementation
func (eObject *MockEObject) EInvoke(operation EOperation, arguments EList[any]) any {
	ret := eObject.Called(operation, arguments)

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

// EIsProxy provides mock implementation
func (eObject *MockEObject) EIsProxy() bool {
	ret := eObject.Called()

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

// EIsSet provides mock implementation
func (eObject *MockEObject) EIsSet(feature EStructuralFeature) bool {
	ret := eObject.Called(feature)

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

// EResource provides mock implementation
func (eObject *MockEObject) EResource() EResource {
	ret := eObject.Called()

	var r EResource
	if rf, ok := ret.Get(0).(func() EResource); ok {
		r = rf()
	} else {
		if ret.Get(0) != nil {
			r = ret.Get(0).(EResource)
		}
	}

	return r
}

// ESet provides mock implementation
func (eObject *MockEObject) ESet(feature EStructuralFeature, newValue any) {
	eObject.Called(feature, newValue)
}

// EUnset provides mock implementation
func (eObject *MockEObject) EUnset(feature EStructuralFeature) {
	eObject.Called(feature)
}
