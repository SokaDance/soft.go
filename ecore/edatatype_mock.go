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

type MockEDataType struct {
	MockEDataType_Prototype
	mock.Mock
}

type MockEDataType_Prototype struct {
	MockEClassifier_Prototype
}

type MockEDataType_Expecter struct {
	MockEClassifier_Expecter
}

func (eDataType *MockEDataType_Prototype) EXPECT() *MockEDataType_Expecter {
	e := &MockEDataType_Expecter{}
	e.Mock = eDataType.Mock
	return e
}

// IsSerializable get the value of isSerializable
func (eDataType *MockEDataType_Prototype) IsSerializable() bool {
	ret := eDataType.Called()

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

type MockEDataType_IsSerializable_Call struct {
	*mock.Call
}

func (e *MockEDataType_Expecter) IsSerializable() *MockEDataType_IsSerializable_Call {
	return &MockEDataType_IsSerializable_Call{Call: e.Mock.On("IsSerializable")}
}

func (c *MockEDataType_IsSerializable_Call) Run(run func()) *MockEDataType_IsSerializable_Call {
	c.Call.Run(func(mock.Arguments) {
		run()
	})
	return c
}

func (c *MockEDataType_IsSerializable_Call) Return(isSerializable bool) *MockEDataType_IsSerializable_Call {
	c.Call.Return(isSerializable)
	return c
}

// SetSerializable provides mock implementation for setting the value of isSerializable
func (eDataType *MockEDataType_Prototype) SetSerializable(isSerializable bool) {
	eDataType.Called(isSerializable)
}

type MockEDataType_SetSerializable_Call struct {
	*mock.Call
}

// SetSerializableis a helper method to define mock.On call
// - isSerializable bool
func (e *MockEDataType_Expecter) SetSerializable(isSerializable any) *MockEDataType_SetSerializable_Call {
	return &MockEDataType_SetSerializable_Call{Call: e.Mock.On("SetSerializable", isSerializable)}
}

func (c *MockEDataType_SetSerializable_Call) Run(run func(isSerializable bool)) *MockEDataType_SetSerializable_Call {
	c.Call.Run(func(args mock.Arguments) {
		run(args[0].(bool))
	})
	return c
}

func (c *MockEDataType_SetSerializable_Call) Return() *MockEDataType_SetSerializable_Call {
	c.Call.Return()
	return c
}

type mockConstructorTestingTNewMockEDataType interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockEDataType creates a new instance of MockEDataType. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockEDataType(t mockConstructorTestingTNewMockEDataType) *MockEDataType {
	mock := &MockEDataType{}
	mock.MockEDataType_Prototype.Mock = &mock.Mock
	mock.Mock.Test(t)
	t.Cleanup(func() { mock.AssertExpectations(t) })
	return mock
}
