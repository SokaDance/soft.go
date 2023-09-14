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

// MockEOperation is an mock type for the EOperation type
type MockEOperation struct {
	mock.Mock
	MockEOperation_Prototype
}

// MockEOperation_Prototype is the mock implementation of all EOperation methods ( inherited and declared )
type MockEOperation_Prototype struct {
	mock *mock.Mock
	MockETypedElement_Prototype
	MockEOperation_Prototype_Methods
}

func (_mp *MockEOperation_Prototype) SetMock(mock *mock.Mock) {
	_mp.mock = mock
	_mp.MockETypedElement_Prototype.SetMock(mock)
	_mp.MockEOperation_Prototype_Methods.SetMock(mock)
}

// MockEOperation_Expecter is the expecter implementation for all EOperation methods ( inherited and declared )
type MockEOperation_Expecter struct {
	MockETypedElement_Expecter
	MockEOperation_Expecter_Methods
}

func (_me *MockEOperation_Expecter) SetMock(mock *mock.Mock) {
	_me.MockETypedElement_Expecter.SetMock(mock)
	_me.MockEOperation_Expecter_Methods.SetMock(mock)
}

func (eOperation *MockEOperation_Prototype) EXPECT() *MockEOperation_Expecter {
	expecter := &MockEOperation_Expecter{}
	expecter.SetMock(eOperation.mock)
	return expecter
}

// MockEOperation_Prototype_Methods is the mock implementation of EOperation declared methods
type MockEOperation_Prototype_Methods struct {
	mock *mock.Mock
}

func (_mdp *MockEOperation_Prototype_Methods) SetMock(mock *mock.Mock) {
	_mdp.mock = mock
}

// MockEOperation_Expecter_Methods is the expecter implementation of EOperation declared methods
type MockEOperation_Expecter_Methods struct {
	mock *mock.Mock
}

func (_mde *MockEOperation_Expecter_Methods) SetMock(mock *mock.Mock) {
	_mde.mock = mock
}

// GetEContainingClass get the value of eContainingClass
func (eOperation *MockEOperation_Prototype_Methods) GetEContainingClass() EClass {
	ret := eOperation.mock.Called()

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

type MockEOperation_GetEContainingClass_Call struct {
	*mock.Call
}

func (e *MockEOperation_Expecter_Methods) GetEContainingClass() *MockEOperation_GetEContainingClass_Call {
	return &MockEOperation_GetEContainingClass_Call{Call: e.mock.On("GetEContainingClass")}
}

func (c *MockEOperation_GetEContainingClass_Call) Run(run func()) *MockEOperation_GetEContainingClass_Call {
	c.Call.Run(func(mock.Arguments) {
		run()
	})
	return c
}

func (c *MockEOperation_GetEContainingClass_Call) Return(eContainingClass EClass) *MockEOperation_GetEContainingClass_Call {
	c.Call.Return(eContainingClass)
	return c
}

// GetEExceptions get the value of eExceptions
func (eOperation *MockEOperation_Prototype_Methods) GetEExceptions() EList {
	ret := eOperation.mock.Called()

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

type MockEOperation_GetEExceptions_Call struct {
	*mock.Call
}

func (e *MockEOperation_Expecter_Methods) GetEExceptions() *MockEOperation_GetEExceptions_Call {
	return &MockEOperation_GetEExceptions_Call{Call: e.mock.On("GetEExceptions")}
}

func (c *MockEOperation_GetEExceptions_Call) Run(run func()) *MockEOperation_GetEExceptions_Call {
	c.Call.Run(func(mock.Arguments) {
		run()
	})
	return c
}

func (c *MockEOperation_GetEExceptions_Call) Return(eExceptions EList) *MockEOperation_GetEExceptions_Call {
	c.Call.Return(eExceptions)
	return c
}

// UnsetEExceptions provides mock implementation for unset the value of eExceptions
func (eOperation *MockEOperation_Prototype_Methods) UnsetEExceptions() {
	eOperation.mock.Called()
}

type MockEOperation_UnsetEExceptions_Call struct {
	*mock.Call
}

func (e *MockEOperation_Expecter_Methods) UnsetEExceptions() *MockEOperation_UnsetEExceptions_Call {
	return &MockEOperation_UnsetEExceptions_Call{Call: e.mock.On("UnsetEExceptions")}
}

func (c *MockEOperation_UnsetEExceptions_Call) Run(run func()) *MockEOperation_UnsetEExceptions_Call {
	c.Call.Run(func(mock.Arguments) {
		run()
	})
	return c
}

func (c *MockEOperation_UnsetEExceptions_Call) Return() *MockEOperation_UnsetEExceptions_Call {
	c.Call.Return()
	return c
}

// GetEParameters get the value of eParameters
func (eOperation *MockEOperation_Prototype_Methods) GetEParameters() EList {
	ret := eOperation.mock.Called()

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

type MockEOperation_GetEParameters_Call struct {
	*mock.Call
}

func (e *MockEOperation_Expecter_Methods) GetEParameters() *MockEOperation_GetEParameters_Call {
	return &MockEOperation_GetEParameters_Call{Call: e.mock.On("GetEParameters")}
}

func (c *MockEOperation_GetEParameters_Call) Run(run func()) *MockEOperation_GetEParameters_Call {
	c.Call.Run(func(mock.Arguments) {
		run()
	})
	return c
}

func (c *MockEOperation_GetEParameters_Call) Return(eParameters EList) *MockEOperation_GetEParameters_Call {
	c.Call.Return(eParameters)
	return c
}

// GetOperationID get the value of operationID
func (eOperation *MockEOperation_Prototype_Methods) GetOperationID() int {
	ret := eOperation.mock.Called()

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

type MockEOperation_GetOperationID_Call struct {
	*mock.Call
}

func (e *MockEOperation_Expecter_Methods) GetOperationID() *MockEOperation_GetOperationID_Call {
	return &MockEOperation_GetOperationID_Call{Call: e.mock.On("GetOperationID")}
}

func (c *MockEOperation_GetOperationID_Call) Run(run func()) *MockEOperation_GetOperationID_Call {
	c.Call.Run(func(mock.Arguments) {
		run()
	})
	return c
}

func (c *MockEOperation_GetOperationID_Call) Return(operationID int) *MockEOperation_GetOperationID_Call {
	c.Call.Return(operationID)
	return c
}

// SetOperationID provides mock implementation for setting the value of operationID
func (eOperation *MockEOperation_Prototype_Methods) SetOperationID(operationID int) {
	eOperation.mock.Called(operationID)
}

type MockEOperation_SetOperationID_Call struct {
	*mock.Call
}

// SetOperationID is a helper method to define mock.On call
// - operationID int
func (e *MockEOperation_Expecter_Methods) SetOperationID(operationID any) *MockEOperation_SetOperationID_Call {
	return &MockEOperation_SetOperationID_Call{Call: e.mock.On("SetOperationID", operationID)}
}

func (c *MockEOperation_SetOperationID_Call) Run(run func(operationID int)) *MockEOperation_SetOperationID_Call {
	c.Call.Run(func(args mock.Arguments) {
		run(args[0].(int))
	})
	return c
}

func (c *MockEOperation_SetOperationID_Call) Return() *MockEOperation_SetOperationID_Call {
	c.Call.Return()
	return c
}

// IsOverrideOf provides mock implementation
func (eOperation *MockEOperation_Prototype_Methods) IsOverrideOf(someOperation EOperation) bool {
	ret := eOperation.mock.Called(someOperation)

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

type MockEOperation_IsOverrideOf_Call struct {
	*mock.Call
}

// IsOverrideOf is a helper method to define mock.On call
// - someOperation EOperation
func (e *MockEOperation_Expecter_Methods) IsOverrideOf(someOperation any) *MockEOperation_IsOverrideOf_Call {
	return &MockEOperation_IsOverrideOf_Call{Call: e.mock.On("IsOverrideOf", someOperation)}
}

func (c *MockEOperation_IsOverrideOf_Call) Run(run func(EOperation)) *MockEOperation_IsOverrideOf_Call {
	c.Call.Run(func(_args mock.Arguments) {
		run(_args[0].(EOperation))
	})
	return c
}

func (c *MockEOperation_IsOverrideOf_Call) Return(_a0 bool) *MockEOperation_IsOverrideOf_Call {
	c.Call.Return(_a0)
	return c
}

type mockConstructorTestingTNewMockEOperation interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockEOperation creates a new instance of MockEOperation. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockEOperation(t mockConstructorTestingTNewMockEOperation) *MockEOperation {
	mock := &MockEOperation{}
	mock.SetMock(&mock.Mock)
	mock.Mock.Test(t)
	t.Cleanup(func() { mock.AssertExpectations(t) })
	return mock
}
