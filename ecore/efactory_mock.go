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

import "github.com/stretchr/testify/mock"

// MockEFactory is an mock type for the EFactory type
type MockEFactory struct {
	mock.Mock
	MockEFactory_Prototype
}

// MockEFactory_Prototype is the mock implementation of all EFactory methods ( inherited and declared )
type MockEFactory_Prototype struct {
	mock *mock.Mock
	MockEModelElement_Prototype
	MockEFactory_Prototype_Methods
}

func (_mp *MockEFactory_Prototype) SetMock(mock *mock.Mock) {
	_mp.mock = mock
	_mp.MockEModelElement_Prototype.SetMock(mock)
	_mp.MockEFactory_Prototype_Methods.SetMock(mock)
}

// MockEFactory_Expecter is the expecter implementation for all EFactory methods ( inherited and declared )
type MockEFactory_Expecter struct {
	MockEModelElement_Expecter
	MockEFactory_Expecter_Methods
}

func (_me *MockEFactory_Expecter) SetMock(mock *mock.Mock) {
	_me.MockEModelElement_Expecter.SetMock(mock)
	_me.MockEFactory_Expecter_Methods.SetMock(mock)
}

func (eFactory *MockEFactory_Prototype) EXPECT() *MockEFactory_Expecter {
	expecter := &MockEFactory_Expecter{}
	expecter.SetMock(eFactory.mock)
	return expecter
}

// MockEFactory_Prototype_Methods is the mock implementation of EFactory declared methods
type MockEFactory_Prototype_Methods struct {
	mock *mock.Mock
}

func (_mdp *MockEFactory_Prototype_Methods) SetMock(mock *mock.Mock) {
	_mdp.mock = mock
}

// MockEFactory_Expecter_Methods is the expecter implementation of EFactory declared methods
type MockEFactory_Expecter_Methods struct {
	mock *mock.Mock
}

func (_mde *MockEFactory_Expecter_Methods) SetMock(mock *mock.Mock) {
	_mde.mock = mock
}

// GetEPackage get the value of ePackage
func (eFactory *MockEFactory_Prototype_Methods) GetEPackage() EPackage {
	ret := eFactory.mock.Called()

	var r EPackage
	if rf, ok := ret.Get(0).(func() EPackage); ok {
		r = rf()
	} else {
		if ret.Get(0) != nil {
			r = ret.Get(0).(EPackage)
		}
	}

	return r
}

type MockEFactory_GetEPackage_Call struct {
	*mock.Call
}

func (e *MockEFactory_Expecter_Methods) GetEPackage() *MockEFactory_GetEPackage_Call {
	return &MockEFactory_GetEPackage_Call{Call: e.mock.On("GetEPackage")}
}

func (c *MockEFactory_GetEPackage_Call) Run(run func()) *MockEFactory_GetEPackage_Call {
	c.Call.Run(func(mock.Arguments) {
		run()
	})
	return c
}

func (c *MockEFactory_GetEPackage_Call) Return(ePackage EPackage) *MockEFactory_GetEPackage_Call {
	c.Call.Return(ePackage)
	return c
}

// SetEPackage provides mock implementation for setting the value of ePackage
func (eFactory *MockEFactory_Prototype_Methods) SetEPackage(ePackage EPackage) {
	eFactory.mock.Called(ePackage)
}

type MockEFactory_SetEPackage_Call struct {
	*mock.Call
}

// SetEPackage is a helper method to define mock.On call
// - ePackage EPackage
func (e *MockEFactory_Expecter_Methods) SetEPackage(ePackage any) *MockEFactory_SetEPackage_Call {
	return &MockEFactory_SetEPackage_Call{Call: e.mock.On("SetEPackage", ePackage)}
}

func (c *MockEFactory_SetEPackage_Call) Run(run func(ePackage EPackage)) *MockEFactory_SetEPackage_Call {
	c.Call.Run(func(args mock.Arguments) {
		run(args[0].(EPackage))
	})
	return c
}

func (c *MockEFactory_SetEPackage_Call) Return() *MockEFactory_SetEPackage_Call {
	c.Call.Return()
	return c
}

// ConvertToString provides mock implementation
func (eFactory *MockEFactory_Prototype_Methods) ConvertToString(eDataType EDataType, instanceValue any) string {
	ret := eFactory.mock.Called(eDataType, instanceValue)

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

type MockEFactory_ConvertToString_Call struct {
	*mock.Call
}

// ConvertToString is a helper method to define mock.On call
// - eDataType EDataType
// - instanceValue any
func (e *MockEFactory_Expecter_Methods) ConvertToString(eDataType any, instanceValue any) *MockEFactory_ConvertToString_Call {
	return &MockEFactory_ConvertToString_Call{Call: e.mock.On("ConvertToString", eDataType, instanceValue)}
}

func (c *MockEFactory_ConvertToString_Call) Run(run func(EDataType, any)) *MockEFactory_ConvertToString_Call {
	c.Call.Run(func(_args mock.Arguments) {
		run(_args[0].(EDataType), _args[1])
	})
	return c
}

func (c *MockEFactory_ConvertToString_Call) Return(_a0 string) *MockEFactory_ConvertToString_Call {
	c.Call.Return(_a0)
	return c
}

// Create provides mock implementation
func (eFactory *MockEFactory_Prototype_Methods) Create(eClass EClass) EObject {
	ret := eFactory.mock.Called(eClass)

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

type MockEFactory_Create_Call struct {
	*mock.Call
}

// Create is a helper method to define mock.On call
// - eClass EClass
func (e *MockEFactory_Expecter_Methods) Create(eClass any) *MockEFactory_Create_Call {
	return &MockEFactory_Create_Call{Call: e.mock.On("Create", eClass)}
}

func (c *MockEFactory_Create_Call) Run(run func(EClass)) *MockEFactory_Create_Call {
	c.Call.Run(func(_args mock.Arguments) {
		run(_args[0].(EClass))
	})
	return c
}

func (c *MockEFactory_Create_Call) Return(_a0 EObject) *MockEFactory_Create_Call {
	c.Call.Return(_a0)
	return c
}

// CreateFromString provides mock implementation
func (eFactory *MockEFactory_Prototype_Methods) CreateFromString(eDataType EDataType, literalValue string) any {
	ret := eFactory.mock.Called(eDataType, literalValue)

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

type MockEFactory_CreateFromString_Call struct {
	*mock.Call
}

// CreateFromString is a helper method to define mock.On call
// - eDataType EDataType
// - literalValue string
func (e *MockEFactory_Expecter_Methods) CreateFromString(eDataType any, literalValue any) *MockEFactory_CreateFromString_Call {
	return &MockEFactory_CreateFromString_Call{Call: e.mock.On("CreateFromString", eDataType, literalValue)}
}

func (c *MockEFactory_CreateFromString_Call) Run(run func(EDataType, string)) *MockEFactory_CreateFromString_Call {
	c.Call.Run(func(_args mock.Arguments) {
		run(_args[0].(EDataType), _args[1].(string))
	})
	return c
}

func (c *MockEFactory_CreateFromString_Call) Return(_a0 any) *MockEFactory_CreateFromString_Call {
	c.Call.Return(_a0)
	return c
}

type mockConstructorTestingTNewMockEFactory interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockEFactory creates a new instance of MockEFactory. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockEFactory(t mockConstructorTestingTNewMockEFactory) *MockEFactory {
	mock := &MockEFactory{}
	mock.SetMock(&mock.Mock)
	mock.Mock.Test(t)
	t.Cleanup(func() { mock.AssertExpectations(t) })
	return mock
}
