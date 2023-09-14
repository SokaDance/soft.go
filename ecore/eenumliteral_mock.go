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

// MockEEnumLiteral is an mock type for the EEnumLiteral type
type MockEEnumLiteral struct {
	mock.Mock
	MockEEnumLiteral_Prototype
}

// MockEEnumLiteral_Prototype is the mock implementation of all EEnumLiteral methods ( inherited and declared )
type MockEEnumLiteral_Prototype struct {
	mock *mock.Mock
	MockENamedElement_Prototype
	MockEEnumLiteral_Prototype_Methods
}

func (_mp *MockEEnumLiteral_Prototype) SetMock(mock *mock.Mock) {
	_mp.mock = mock
	_mp.MockENamedElement_Prototype.SetMock(mock)
	_mp.MockEEnumLiteral_Prototype_Methods.SetMock(mock)
}

// MockEEnumLiteral_Expecter is the expecter implementation for all EEnumLiteral methods ( inherited and declared )
type MockEEnumLiteral_Expecter struct {
	MockENamedElement_Expecter
	MockEEnumLiteral_Expecter_Methods
}

func (_me *MockEEnumLiteral_Expecter) SetMock(mock *mock.Mock) {
	_me.MockENamedElement_Expecter.SetMock(mock)
	_me.MockEEnumLiteral_Expecter_Methods.SetMock(mock)
}

func (eEnumLiteral *MockEEnumLiteral_Prototype) EXPECT() *MockEEnumLiteral_Expecter {
	expecter := &MockEEnumLiteral_Expecter{}
	expecter.SetMock(eEnumLiteral.mock)
	return expecter
}

// MockEEnumLiteral_Prototype_Methods is the mock implementation of EEnumLiteral declared methods
type MockEEnumLiteral_Prototype_Methods struct {
	mock *mock.Mock
}

func (_mdp *MockEEnumLiteral_Prototype_Methods) SetMock(mock *mock.Mock) {
	_mdp.mock = mock
}

// MockEEnumLiteral_Expecter_Methods is the expecter implementation of EEnumLiteral declared methods
type MockEEnumLiteral_Expecter_Methods struct {
	mock *mock.Mock
}

func (_mde *MockEEnumLiteral_Expecter_Methods) SetMock(mock *mock.Mock) {
	_mde.mock = mock
}

// GetEEnum get the value of eEnum
func (eEnumLiteral *MockEEnumLiteral_Prototype_Methods) GetEEnum() EEnum {
	ret := eEnumLiteral.mock.Called()

	var r EEnum
	if rf, ok := ret.Get(0).(func() EEnum); ok {
		r = rf()
	} else {
		if ret.Get(0) != nil {
			r = ret.Get(0).(EEnum)
		}
	}

	return r
}

type MockEEnumLiteral_GetEEnum_Call struct {
	*mock.Call
}

func (e *MockEEnumLiteral_Expecter_Methods) GetEEnum() *MockEEnumLiteral_GetEEnum_Call {
	return &MockEEnumLiteral_GetEEnum_Call{Call: e.mock.On("GetEEnum")}
}

func (c *MockEEnumLiteral_GetEEnum_Call) Run(run func()) *MockEEnumLiteral_GetEEnum_Call {
	c.Call.Run(func(mock.Arguments) {
		run()
	})
	return c
}

func (c *MockEEnumLiteral_GetEEnum_Call) Return(eEnum EEnum) *MockEEnumLiteral_GetEEnum_Call {
	c.Call.Return(eEnum)
	return c
}

// GetInstance get the value of instance
func (eEnumLiteral *MockEEnumLiteral_Prototype_Methods) GetInstance() any {
	ret := eEnumLiteral.mock.Called()

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

type MockEEnumLiteral_GetInstance_Call struct {
	*mock.Call
}

func (e *MockEEnumLiteral_Expecter_Methods) GetInstance() *MockEEnumLiteral_GetInstance_Call {
	return &MockEEnumLiteral_GetInstance_Call{Call: e.mock.On("GetInstance")}
}

func (c *MockEEnumLiteral_GetInstance_Call) Run(run func()) *MockEEnumLiteral_GetInstance_Call {
	c.Call.Run(func(mock.Arguments) {
		run()
	})
	return c
}

func (c *MockEEnumLiteral_GetInstance_Call) Return(instance any) *MockEEnumLiteral_GetInstance_Call {
	c.Call.Return(instance)
	return c
}

// SetInstance provides mock implementation for setting the value of instance
func (eEnumLiteral *MockEEnumLiteral_Prototype_Methods) SetInstance(instance any) {
	eEnumLiteral.mock.Called(instance)
}

type MockEEnumLiteral_SetInstance_Call struct {
	*mock.Call
}

// SetInstance is a helper method to define mock.On call
// - instance any
func (e *MockEEnumLiteral_Expecter_Methods) SetInstance(instance any) *MockEEnumLiteral_SetInstance_Call {
	return &MockEEnumLiteral_SetInstance_Call{Call: e.mock.On("SetInstance", instance)}
}

func (c *MockEEnumLiteral_SetInstance_Call) Run(run func(instance any)) *MockEEnumLiteral_SetInstance_Call {
	c.Call.Run(func(args mock.Arguments) {
		run(args[0])
	})
	return c
}

func (c *MockEEnumLiteral_SetInstance_Call) Return() *MockEEnumLiteral_SetInstance_Call {
	c.Call.Return()
	return c
}

// GetLiteral get the value of literal
func (eEnumLiteral *MockEEnumLiteral_Prototype_Methods) GetLiteral() string {
	ret := eEnumLiteral.mock.Called()

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

type MockEEnumLiteral_GetLiteral_Call struct {
	*mock.Call
}

func (e *MockEEnumLiteral_Expecter_Methods) GetLiteral() *MockEEnumLiteral_GetLiteral_Call {
	return &MockEEnumLiteral_GetLiteral_Call{Call: e.mock.On("GetLiteral")}
}

func (c *MockEEnumLiteral_GetLiteral_Call) Run(run func()) *MockEEnumLiteral_GetLiteral_Call {
	c.Call.Run(func(mock.Arguments) {
		run()
	})
	return c
}

func (c *MockEEnumLiteral_GetLiteral_Call) Return(literal string) *MockEEnumLiteral_GetLiteral_Call {
	c.Call.Return(literal)
	return c
}

// SetLiteral provides mock implementation for setting the value of literal
func (eEnumLiteral *MockEEnumLiteral_Prototype_Methods) SetLiteral(literal string) {
	eEnumLiteral.mock.Called(literal)
}

type MockEEnumLiteral_SetLiteral_Call struct {
	*mock.Call
}

// SetLiteral is a helper method to define mock.On call
// - literal string
func (e *MockEEnumLiteral_Expecter_Methods) SetLiteral(literal any) *MockEEnumLiteral_SetLiteral_Call {
	return &MockEEnumLiteral_SetLiteral_Call{Call: e.mock.On("SetLiteral", literal)}
}

func (c *MockEEnumLiteral_SetLiteral_Call) Run(run func(literal string)) *MockEEnumLiteral_SetLiteral_Call {
	c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return c
}

func (c *MockEEnumLiteral_SetLiteral_Call) Return() *MockEEnumLiteral_SetLiteral_Call {
	c.Call.Return()
	return c
}

// GetValue get the value of value
func (eEnumLiteral *MockEEnumLiteral_Prototype_Methods) GetValue() int {
	ret := eEnumLiteral.mock.Called()

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

type MockEEnumLiteral_GetValue_Call struct {
	*mock.Call
}

func (e *MockEEnumLiteral_Expecter_Methods) GetValue() *MockEEnumLiteral_GetValue_Call {
	return &MockEEnumLiteral_GetValue_Call{Call: e.mock.On("GetValue")}
}

func (c *MockEEnumLiteral_GetValue_Call) Run(run func()) *MockEEnumLiteral_GetValue_Call {
	c.Call.Run(func(mock.Arguments) {
		run()
	})
	return c
}

func (c *MockEEnumLiteral_GetValue_Call) Return(value int) *MockEEnumLiteral_GetValue_Call {
	c.Call.Return(value)
	return c
}

// SetValue provides mock implementation for setting the value of value
func (eEnumLiteral *MockEEnumLiteral_Prototype_Methods) SetValue(value int) {
	eEnumLiteral.mock.Called(value)
}

type MockEEnumLiteral_SetValue_Call struct {
	*mock.Call
}

// SetValue is a helper method to define mock.On call
// - value int
func (e *MockEEnumLiteral_Expecter_Methods) SetValue(value any) *MockEEnumLiteral_SetValue_Call {
	return &MockEEnumLiteral_SetValue_Call{Call: e.mock.On("SetValue", value)}
}

func (c *MockEEnumLiteral_SetValue_Call) Run(run func(value int)) *MockEEnumLiteral_SetValue_Call {
	c.Call.Run(func(args mock.Arguments) {
		run(args[0].(int))
	})
	return c
}

func (c *MockEEnumLiteral_SetValue_Call) Return() *MockEEnumLiteral_SetValue_Call {
	c.Call.Return()
	return c
}

type mockConstructorTestingTNewMockEEnumLiteral interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockEEnumLiteral creates a new instance of MockEEnumLiteral. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockEEnumLiteral(t mockConstructorTestingTNewMockEEnumLiteral) *MockEEnumLiteral {
	mock := &MockEEnumLiteral{}
	mock.SetMock(&mock.Mock)
	mock.Mock.Test(t)
	t.Cleanup(func() { mock.AssertExpectations(t) })
	return mock
}
