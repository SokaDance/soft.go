// Code generated by soft.generator.go. DO NOT EDIT.

// *****************************************************************************
// Copyright(c) 2021 MASA Group
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// *****************************************************************************

package ecore

import (
	"github.com/stretchr/testify/mock"
)

type MockEObject struct {
	MockEObject_Prototype
	mock.Mock
}

type MockEObject_Prototype struct {
	MockENotifier_Prototype
}

type MockEObject_Expecter struct {
	MockENotifier_Expecter
}

func (eObject *MockEObject_Prototype) EXPECT() *MockEObject_Expecter {
	e := &MockEObject_Expecter{}
	e.Mock = eObject.Mock
	return e
}

// EAllContents provides mock implementation
func (eObject *MockEObject_Prototype) EAllContents() EIterator {
	ret := eObject.Called()

	var r EIterator
	if rf, ok := ret.Get(0).(func() EIterator); ok {
		r = rf()
	} else {
		if ret.Get(0) != nil {
			r = ret.Get(0).(EIterator)
		}
	}

	return r
}

type MockEObject_EAllContents_Call struct {
	*mock.Call
}

// EAllContentsis a helper method to define mock.On call
func (e *MockEObject_Expecter) EAllContents() *MockEObject_EAllContents_Call {
	return &MockEObject_EAllContents_Call{Call: e.Mock.On("EAllContents")}
}

func (c *MockEObject_EAllContents_Call) Run(run func()) *MockEObject_EAllContents_Call {
	c.Call.Run(func(_args mock.Arguments) {
		run()
	})
	return c
}

func (c *MockEObject_EAllContents_Call) Return(_a0 EIterator) *MockEObject_EAllContents_Call {
	c.Call.Return(_a0)
	return c
}

// EClass provides mock implementation
func (eObject *MockEObject_Prototype) EClass() EClass {
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

type MockEObject_EClass_Call struct {
	*mock.Call
}

// EClassis a helper method to define mock.On call
func (e *MockEObject_Expecter) EClass() *MockEObject_EClass_Call {
	return &MockEObject_EClass_Call{Call: e.Mock.On("EClass")}
}

func (c *MockEObject_EClass_Call) Run(run func()) *MockEObject_EClass_Call {
	c.Call.Run(func(_args mock.Arguments) {
		run()
	})
	return c
}

func (c *MockEObject_EClass_Call) Return(_a0 EClass) *MockEObject_EClass_Call {
	c.Call.Return(_a0)
	return c
}

// EContainer provides mock implementation
func (eObject *MockEObject_Prototype) EContainer() EObject {
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

type MockEObject_EContainer_Call struct {
	*mock.Call
}

// EContaineris a helper method to define mock.On call
func (e *MockEObject_Expecter) EContainer() *MockEObject_EContainer_Call {
	return &MockEObject_EContainer_Call{Call: e.Mock.On("EContainer")}
}

func (c *MockEObject_EContainer_Call) Run(run func()) *MockEObject_EContainer_Call {
	c.Call.Run(func(_args mock.Arguments) {
		run()
	})
	return c
}

func (c *MockEObject_EContainer_Call) Return(_a0 EObject) *MockEObject_EContainer_Call {
	c.Call.Return(_a0)
	return c
}

// EContainingFeature provides mock implementation
func (eObject *MockEObject_Prototype) EContainingFeature() EStructuralFeature {
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

type MockEObject_EContainingFeature_Call struct {
	*mock.Call
}

// EContainingFeatureis a helper method to define mock.On call
func (e *MockEObject_Expecter) EContainingFeature() *MockEObject_EContainingFeature_Call {
	return &MockEObject_EContainingFeature_Call{Call: e.Mock.On("EContainingFeature")}
}

func (c *MockEObject_EContainingFeature_Call) Run(run func()) *MockEObject_EContainingFeature_Call {
	c.Call.Run(func(_args mock.Arguments) {
		run()
	})
	return c
}

func (c *MockEObject_EContainingFeature_Call) Return(_a0 EStructuralFeature) *MockEObject_EContainingFeature_Call {
	c.Call.Return(_a0)
	return c
}

// EContainmentFeature provides mock implementation
func (eObject *MockEObject_Prototype) EContainmentFeature() EReference {
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

type MockEObject_EContainmentFeature_Call struct {
	*mock.Call
}

// EContainmentFeatureis a helper method to define mock.On call
func (e *MockEObject_Expecter) EContainmentFeature() *MockEObject_EContainmentFeature_Call {
	return &MockEObject_EContainmentFeature_Call{Call: e.Mock.On("EContainmentFeature")}
}

func (c *MockEObject_EContainmentFeature_Call) Run(run func()) *MockEObject_EContainmentFeature_Call {
	c.Call.Run(func(_args mock.Arguments) {
		run()
	})
	return c
}

func (c *MockEObject_EContainmentFeature_Call) Return(_a0 EReference) *MockEObject_EContainmentFeature_Call {
	c.Call.Return(_a0)
	return c
}

// EContents provides mock implementation
func (eObject *MockEObject_Prototype) EContents() EList {
	ret := eObject.Called()

	var r EList
	if rf, ok := ret.Get(0).(func() EList); ok {
		r = rf()
	} else {
		if ret.Get(0) != nil {
			r = ret.Get(0).(EList)
		}
	}

	return r
}

type MockEObject_EContents_Call struct {
	*mock.Call
}

// EContentsis a helper method to define mock.On call
func (e *MockEObject_Expecter) EContents() *MockEObject_EContents_Call {
	return &MockEObject_EContents_Call{Call: e.Mock.On("EContents")}
}

func (c *MockEObject_EContents_Call) Run(run func()) *MockEObject_EContents_Call {
	c.Call.Run(func(_args mock.Arguments) {
		run()
	})
	return c
}

func (c *MockEObject_EContents_Call) Return(_a0 EList) *MockEObject_EContents_Call {
	c.Call.Return(_a0)
	return c
}

// ECrossReferences provides mock implementation
func (eObject *MockEObject_Prototype) ECrossReferences() EList {
	ret := eObject.Called()

	var r EList
	if rf, ok := ret.Get(0).(func() EList); ok {
		r = rf()
	} else {
		if ret.Get(0) != nil {
			r = ret.Get(0).(EList)
		}
	}

	return r
}

type MockEObject_ECrossReferences_Call struct {
	*mock.Call
}

// ECrossReferencesis a helper method to define mock.On call
func (e *MockEObject_Expecter) ECrossReferences() *MockEObject_ECrossReferences_Call {
	return &MockEObject_ECrossReferences_Call{Call: e.Mock.On("ECrossReferences")}
}

func (c *MockEObject_ECrossReferences_Call) Run(run func()) *MockEObject_ECrossReferences_Call {
	c.Call.Run(func(_args mock.Arguments) {
		run()
	})
	return c
}

func (c *MockEObject_ECrossReferences_Call) Return(_a0 EList) *MockEObject_ECrossReferences_Call {
	c.Call.Return(_a0)
	return c
}

// EGet provides mock implementation
func (eObject *MockEObject_Prototype) EGet(feature EStructuralFeature) any {
	ret := eObject.Called(feature)

	var r any
	if rf, ok := ret.Get(0).(func() any); ok {
		r = rf()
	} else {
		if ret.Get(0) != nil {
			r = ret.Get(0)
		}
	}

	return r
}

type MockEObject_EGet_Call struct {
	*mock.Call
}

// EGetis a helper method to define mock.On call
// - feature EStructuralFeature
func (e *MockEObject_Expecter) EGet(feature any) *MockEObject_EGet_Call {
	return &MockEObject_EGet_Call{Call: e.Mock.On("EGet", feature)}
}

func (c *MockEObject_EGet_Call) Run(run func(EStructuralFeature)) *MockEObject_EGet_Call {
	c.Call.Run(func(_args mock.Arguments) {
		run(_args[0].(EStructuralFeature))
	})
	return c
}

func (c *MockEObject_EGet_Call) Return(_a0 any) *MockEObject_EGet_Call {
	c.Call.Return(_a0)
	return c
}

// EGetResolve provides mock implementation
func (eObject *MockEObject_Prototype) EGetResolve(feature EStructuralFeature, resolve bool) any {
	ret := eObject.Called(feature, resolve)

	var r any
	if rf, ok := ret.Get(0).(func() any); ok {
		r = rf()
	} else {
		if ret.Get(0) != nil {
			r = ret.Get(0)
		}
	}

	return r
}

type MockEObject_EGetResolve_Call struct {
	*mock.Call
}

// EGetResolveis a helper method to define mock.On call
// - feature EStructuralFeature
// - resolve bool
func (e *MockEObject_Expecter) EGetResolve(feature any, resolve any) *MockEObject_EGetResolve_Call {
	return &MockEObject_EGetResolve_Call{Call: e.Mock.On("EGetResolve", feature, resolve)}
}

func (c *MockEObject_EGetResolve_Call) Run(run func(EStructuralFeature, bool)) *MockEObject_EGetResolve_Call {
	c.Call.Run(func(_args mock.Arguments) {
		run(_args[0].(EStructuralFeature), _args[1].(bool))
	})
	return c
}

func (c *MockEObject_EGetResolve_Call) Return(_a0 any) *MockEObject_EGetResolve_Call {
	c.Call.Return(_a0)
	return c
}

// EInvoke provides mock implementation
func (eObject *MockEObject_Prototype) EInvoke(operation EOperation, arguments EList) any {
	ret := eObject.Called(operation, arguments)

	var r any
	if rf, ok := ret.Get(0).(func() any); ok {
		r = rf()
	} else {
		if ret.Get(0) != nil {
			r = ret.Get(0)
		}
	}

	return r
}

type MockEObject_EInvoke_Call struct {
	*mock.Call
}

// EInvokeis a helper method to define mock.On call
// - operation EOperation
// - arguments EList
func (e *MockEObject_Expecter) EInvoke(operation any, arguments any) *MockEObject_EInvoke_Call {
	return &MockEObject_EInvoke_Call{Call: e.Mock.On("EInvoke", operation, arguments)}
}

func (c *MockEObject_EInvoke_Call) Run(run func(EOperation, EList)) *MockEObject_EInvoke_Call {
	c.Call.Run(func(_args mock.Arguments) {
		run(_args[0].(EOperation), _args[1].(EList))
	})
	return c
}

func (c *MockEObject_EInvoke_Call) Return(_a0 any) *MockEObject_EInvoke_Call {
	c.Call.Return(_a0)
	return c
}

// EIsProxy provides mock implementation
func (eObject *MockEObject_Prototype) EIsProxy() bool {
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

type MockEObject_EIsProxy_Call struct {
	*mock.Call
}

// EIsProxyis a helper method to define mock.On call
func (e *MockEObject_Expecter) EIsProxy() *MockEObject_EIsProxy_Call {
	return &MockEObject_EIsProxy_Call{Call: e.Mock.On("EIsProxy")}
}

func (c *MockEObject_EIsProxy_Call) Run(run func()) *MockEObject_EIsProxy_Call {
	c.Call.Run(func(_args mock.Arguments) {
		run()
	})
	return c
}

func (c *MockEObject_EIsProxy_Call) Return(_a0 bool) *MockEObject_EIsProxy_Call {
	c.Call.Return(_a0)
	return c
}

// EIsSet provides mock implementation
func (eObject *MockEObject_Prototype) EIsSet(feature EStructuralFeature) bool {
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

type MockEObject_EIsSet_Call struct {
	*mock.Call
}

// EIsSetis a helper method to define mock.On call
// - feature EStructuralFeature
func (e *MockEObject_Expecter) EIsSet(feature any) *MockEObject_EIsSet_Call {
	return &MockEObject_EIsSet_Call{Call: e.Mock.On("EIsSet", feature)}
}

func (c *MockEObject_EIsSet_Call) Run(run func(EStructuralFeature)) *MockEObject_EIsSet_Call {
	c.Call.Run(func(_args mock.Arguments) {
		run(_args[0].(EStructuralFeature))
	})
	return c
}

func (c *MockEObject_EIsSet_Call) Return(_a0 bool) *MockEObject_EIsSet_Call {
	c.Call.Return(_a0)
	return c
}

// EResource provides mock implementation
func (eObject *MockEObject_Prototype) EResource() EResource {
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

type MockEObject_EResource_Call struct {
	*mock.Call
}

// EResourceis a helper method to define mock.On call
func (e *MockEObject_Expecter) EResource() *MockEObject_EResource_Call {
	return &MockEObject_EResource_Call{Call: e.Mock.On("EResource")}
}

func (c *MockEObject_EResource_Call) Run(run func()) *MockEObject_EResource_Call {
	c.Call.Run(func(_args mock.Arguments) {
		run()
	})
	return c
}

func (c *MockEObject_EResource_Call) Return(_a0 EResource) *MockEObject_EResource_Call {
	c.Call.Return(_a0)
	return c
}

// ESet provides mock implementation
func (eObject *MockEObject_Prototype) ESet(feature EStructuralFeature, newValue any) {
	eObject.Called(feature, newValue)
}

type MockEObject_ESet_Call struct {
	*mock.Call
}

func (e *MockEObject_Expecter) ESet(feature EStructuralFeature, newValue any) *MockEObject_ESet_Call {
	return &MockEObject_ESet_Call{Call: e.Mock.On("ESet", feature, newValue)}
}

func (c *MockEObject_ESet_Call) Run(run func(EStructuralFeature, any)) *MockEObject_ESet_Call {
	c.Call.Run(func(_args mock.Arguments) {
		run(_args[0].(EStructuralFeature), _args[1])
	})
	return c
}

func (c *MockEObject_ESet_Call) Return() *MockEObject_ESet_Call {
	c.Call.Return()
	return c
} // EUnset provides mock implementation
func (eObject *MockEObject_Prototype) EUnset(feature EStructuralFeature) {
	eObject.Called(feature)
}

type MockEObject_EUnset_Call struct {
	*mock.Call
}

func (e *MockEObject_Expecter) EUnset(feature EStructuralFeature) *MockEObject_EUnset_Call {
	return &MockEObject_EUnset_Call{Call: e.Mock.On("EUnset", feature)}
}

func (c *MockEObject_EUnset_Call) Run(run func(EStructuralFeature)) *MockEObject_EUnset_Call {
	c.Call.Run(func(_args mock.Arguments) {
		run(_args[0].(EStructuralFeature))
	})
	return c
}

func (c *MockEObject_EUnset_Call) Return() *MockEObject_EUnset_Call {
	c.Call.Return()
	return c
}

type mockConstructorTestingTNewMockEObject interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockEObject creates a new instance of MockEObject. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockEObject(t mockConstructorTestingTNewMockEObject) *MockEObject {
	mock := &MockEObject{}
	mock.MockEObject_Prototype.Mock = &mock.Mock
	mock.Mock.Test(t)
	t.Cleanup(func() { mock.AssertExpectations(t) })
	return mock
}
