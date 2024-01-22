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

// MockEParameter is an mock type for the EParameter type
type MockEParameter struct {
	mock.Mock
	MockEParameter_Prototype
}

// MockEParameter_Prototype is the mock implementation of all EParameter methods ( inherited and declared )
type MockEParameter_Prototype struct {
	mock *mock.Mock
	MockETypedElement_Prototype
	MockEParameter_Prototype_Methods
}

func (_mp *MockEParameter_Prototype) SetMock(mock *mock.Mock) {
	_mp.mock = mock
	_mp.MockETypedElement_Prototype.SetMock(mock)
	_mp.MockEParameter_Prototype_Methods.SetMock(mock)
}

// MockEParameter_Expecter is the expecter implementation for all EParameter methods ( inherited and declared )
type MockEParameter_Expecter struct {
	MockETypedElement_Expecter
	MockEParameter_Expecter_Methods
}

func (_me *MockEParameter_Expecter) SetMock(mock *mock.Mock) {
	_me.MockETypedElement_Expecter.SetMock(mock)
	_me.MockEParameter_Expecter_Methods.SetMock(mock)
}

func (e *MockEParameter_Prototype) EXPECT() *MockEParameter_Expecter {
	expecter := &MockEParameter_Expecter{}
	expecter.SetMock(e.mock)
	return expecter
}

// MockEParameter_Prototype_Methods is the mock implementation of EParameter declared methods
type MockEParameter_Prototype_Methods struct {
	mock *mock.Mock
}

func (_mdp *MockEParameter_Prototype_Methods) SetMock(mock *mock.Mock) {
	_mdp.mock = mock
}

// MockEParameter_Expecter_Methods is the expecter implementation of EParameter declared methods
type MockEParameter_Expecter_Methods struct {
	mock *mock.Mock
}

func (_mde *MockEParameter_Expecter_Methods) SetMock(mock *mock.Mock) {
	_mde.mock = mock
}

// GetEOperation get the value of eOperation
func (e *MockEParameter_Prototype_Methods) GetEOperation() EOperation {
	ret := e.mock.Called()

	var res EOperation
	if rf, ok := ret.Get(0).(func() EOperation); ok {
		res = rf()
	} else {
		if ret.Get(0) != nil {
			res = ret.Get(0).(EOperation)
		}
	}

	return res
}

type MockEParameter_GetEOperation_Call struct {
	*mock.Call
}

func (e *MockEParameter_Expecter_Methods) GetEOperation() *MockEParameter_GetEOperation_Call {
	return &MockEParameter_GetEOperation_Call{Call: e.mock.On("GetEOperation")}
}

func (c *MockEParameter_GetEOperation_Call) Run(run func()) *MockEParameter_GetEOperation_Call {
	c.Call.Run(func(mock.Arguments) {
		run()
	})
	return c
}

func (c *MockEParameter_GetEOperation_Call) Return(eOperation EOperation) *MockEParameter_GetEOperation_Call {
	c.Call.Return(eOperation)
	return c
}

type mockConstructorTestingTNewMockEParameter interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockEParameter creates a new instance of MockEParameter. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockEParameter(t mockConstructorTestingTNewMockEParameter) *MockEParameter {
	mock := &MockEParameter{}
	mock.SetMock(&mock.Mock)
	mock.Mock.Test(t)
	t.Cleanup(func() { mock.AssertExpectations(t) })
	return mock
}
