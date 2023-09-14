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

// MockEPackage is an mock type for the EPackage type
type MockEPackage struct {
	mock.Mock
	MockEPackage_Prototype
}

// MockEPackage_Prototype is the mock implementation of all EPackage methods ( inherited and declared )
type MockEPackage_Prototype struct {
	mock *mock.Mock
	MockENamedElement_Prototype
	MockEPackage_Prototype_Methods
}

func (_mp *MockEPackage_Prototype) SetMock(mock *mock.Mock) {
	_mp.mock = mock
	_mp.MockENamedElement_Prototype.SetMock(mock)
	_mp.MockEPackage_Prototype_Methods.SetMock(mock)
}

// MockEPackage_Expecter is the expecter implementation for all EPackage methods ( inherited and declared )
type MockEPackage_Expecter struct {
	MockENamedElement_Expecter
	MockEPackage_Expecter_Methods
}

func (_me *MockEPackage_Expecter) SetMock(mock *mock.Mock) {
	_me.MockENamedElement_Expecter.SetMock(mock)
	_me.MockEPackage_Expecter_Methods.SetMock(mock)
}

func (ePackage *MockEPackage_Prototype) EXPECT() *MockEPackage_Expecter {
	expecter := &MockEPackage_Expecter{}
	expecter.SetMock(ePackage.mock)
	return expecter
}

// MockEPackage_Prototype_Methods is the mock implementation of EPackage declared methods
type MockEPackage_Prototype_Methods struct {
	mock *mock.Mock
}

func (_mdp *MockEPackage_Prototype_Methods) SetMock(mock *mock.Mock) {
	_mdp.mock = mock
}

// MockEPackage_Expecter_Methods is the expecter implementation of EPackage declared methods
type MockEPackage_Expecter_Methods struct {
	mock *mock.Mock
}

func (_mde *MockEPackage_Expecter_Methods) SetMock(mock *mock.Mock) {
	_mde.mock = mock
}

// GetEClassifiers get the value of eClassifiers
func (ePackage *MockEPackage_Prototype_Methods) GetEClassifiers() EList {
	ret := ePackage.mock.Called()

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

type MockEPackage_GetEClassifiers_Call struct {
	*mock.Call
}

func (e *MockEPackage_Expecter_Methods) GetEClassifiers() *MockEPackage_GetEClassifiers_Call {
	return &MockEPackage_GetEClassifiers_Call{Call: e.mock.On("GetEClassifiers")}
}

func (c *MockEPackage_GetEClassifiers_Call) Run(run func()) *MockEPackage_GetEClassifiers_Call {
	c.Call.Run(func(mock.Arguments) {
		run()
	})
	return c
}

func (c *MockEPackage_GetEClassifiers_Call) Return(eClassifiers EList) *MockEPackage_GetEClassifiers_Call {
	c.Call.Return(eClassifiers)
	return c
}

// GetEFactoryInstance get the value of eFactoryInstance
func (ePackage *MockEPackage_Prototype_Methods) GetEFactoryInstance() EFactory {
	ret := ePackage.mock.Called()

	var r EFactory
	if rf, ok := ret.Get(0).(func() EFactory); ok {
		r = rf()
	} else {
		if ret.Get(0) != nil {
			r = ret.Get(0).(EFactory)
		}
	}

	return r
}

type MockEPackage_GetEFactoryInstance_Call struct {
	*mock.Call
}

func (e *MockEPackage_Expecter_Methods) GetEFactoryInstance() *MockEPackage_GetEFactoryInstance_Call {
	return &MockEPackage_GetEFactoryInstance_Call{Call: e.mock.On("GetEFactoryInstance")}
}

func (c *MockEPackage_GetEFactoryInstance_Call) Run(run func()) *MockEPackage_GetEFactoryInstance_Call {
	c.Call.Run(func(mock.Arguments) {
		run()
	})
	return c
}

func (c *MockEPackage_GetEFactoryInstance_Call) Return(eFactoryInstance EFactory) *MockEPackage_GetEFactoryInstance_Call {
	c.Call.Return(eFactoryInstance)
	return c
}

// SetEFactoryInstance provides mock implementation for setting the value of eFactoryInstance
func (ePackage *MockEPackage_Prototype_Methods) SetEFactoryInstance(eFactoryInstance EFactory) {
	ePackage.mock.Called(eFactoryInstance)
}

type MockEPackage_SetEFactoryInstance_Call struct {
	*mock.Call
}

// SetEFactoryInstance is a helper method to define mock.On call
// - eFactoryInstance EFactory
func (e *MockEPackage_Expecter_Methods) SetEFactoryInstance(eFactoryInstance any) *MockEPackage_SetEFactoryInstance_Call {
	return &MockEPackage_SetEFactoryInstance_Call{Call: e.mock.On("SetEFactoryInstance", eFactoryInstance)}
}

func (c *MockEPackage_SetEFactoryInstance_Call) Run(run func(eFactoryInstance EFactory)) *MockEPackage_SetEFactoryInstance_Call {
	c.Call.Run(func(args mock.Arguments) {
		run(args[0].(EFactory))
	})
	return c
}

func (c *MockEPackage_SetEFactoryInstance_Call) Return() *MockEPackage_SetEFactoryInstance_Call {
	c.Call.Return()
	return c
}

// GetESubPackages get the value of eSubPackages
func (ePackage *MockEPackage_Prototype_Methods) GetESubPackages() EList {
	ret := ePackage.mock.Called()

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

type MockEPackage_GetESubPackages_Call struct {
	*mock.Call
}

func (e *MockEPackage_Expecter_Methods) GetESubPackages() *MockEPackage_GetESubPackages_Call {
	return &MockEPackage_GetESubPackages_Call{Call: e.mock.On("GetESubPackages")}
}

func (c *MockEPackage_GetESubPackages_Call) Run(run func()) *MockEPackage_GetESubPackages_Call {
	c.Call.Run(func(mock.Arguments) {
		run()
	})
	return c
}

func (c *MockEPackage_GetESubPackages_Call) Return(eSubPackages EList) *MockEPackage_GetESubPackages_Call {
	c.Call.Return(eSubPackages)
	return c
}

// GetESuperPackage get the value of eSuperPackage
func (ePackage *MockEPackage_Prototype_Methods) GetESuperPackage() EPackage {
	ret := ePackage.mock.Called()

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

type MockEPackage_GetESuperPackage_Call struct {
	*mock.Call
}

func (e *MockEPackage_Expecter_Methods) GetESuperPackage() *MockEPackage_GetESuperPackage_Call {
	return &MockEPackage_GetESuperPackage_Call{Call: e.mock.On("GetESuperPackage")}
}

func (c *MockEPackage_GetESuperPackage_Call) Run(run func()) *MockEPackage_GetESuperPackage_Call {
	c.Call.Run(func(mock.Arguments) {
		run()
	})
	return c
}

func (c *MockEPackage_GetESuperPackage_Call) Return(eSuperPackage EPackage) *MockEPackage_GetESuperPackage_Call {
	c.Call.Return(eSuperPackage)
	return c
}

// GetNsPrefix get the value of nsPrefix
func (ePackage *MockEPackage_Prototype_Methods) GetNsPrefix() string {
	ret := ePackage.mock.Called()

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

type MockEPackage_GetNsPrefix_Call struct {
	*mock.Call
}

func (e *MockEPackage_Expecter_Methods) GetNsPrefix() *MockEPackage_GetNsPrefix_Call {
	return &MockEPackage_GetNsPrefix_Call{Call: e.mock.On("GetNsPrefix")}
}

func (c *MockEPackage_GetNsPrefix_Call) Run(run func()) *MockEPackage_GetNsPrefix_Call {
	c.Call.Run(func(mock.Arguments) {
		run()
	})
	return c
}

func (c *MockEPackage_GetNsPrefix_Call) Return(nsPrefix string) *MockEPackage_GetNsPrefix_Call {
	c.Call.Return(nsPrefix)
	return c
}

// SetNsPrefix provides mock implementation for setting the value of nsPrefix
func (ePackage *MockEPackage_Prototype_Methods) SetNsPrefix(nsPrefix string) {
	ePackage.mock.Called(nsPrefix)
}

type MockEPackage_SetNsPrefix_Call struct {
	*mock.Call
}

// SetNsPrefix is a helper method to define mock.On call
// - nsPrefix string
func (e *MockEPackage_Expecter_Methods) SetNsPrefix(nsPrefix any) *MockEPackage_SetNsPrefix_Call {
	return &MockEPackage_SetNsPrefix_Call{Call: e.mock.On("SetNsPrefix", nsPrefix)}
}

func (c *MockEPackage_SetNsPrefix_Call) Run(run func(nsPrefix string)) *MockEPackage_SetNsPrefix_Call {
	c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return c
}

func (c *MockEPackage_SetNsPrefix_Call) Return() *MockEPackage_SetNsPrefix_Call {
	c.Call.Return()
	return c
}

// GetNsURI get the value of nsURI
func (ePackage *MockEPackage_Prototype_Methods) GetNsURI() string {
	ret := ePackage.mock.Called()

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

type MockEPackage_GetNsURI_Call struct {
	*mock.Call
}

func (e *MockEPackage_Expecter_Methods) GetNsURI() *MockEPackage_GetNsURI_Call {
	return &MockEPackage_GetNsURI_Call{Call: e.mock.On("GetNsURI")}
}

func (c *MockEPackage_GetNsURI_Call) Run(run func()) *MockEPackage_GetNsURI_Call {
	c.Call.Run(func(mock.Arguments) {
		run()
	})
	return c
}

func (c *MockEPackage_GetNsURI_Call) Return(nsURI string) *MockEPackage_GetNsURI_Call {
	c.Call.Return(nsURI)
	return c
}

// SetNsURI provides mock implementation for setting the value of nsURI
func (ePackage *MockEPackage_Prototype_Methods) SetNsURI(nsURI string) {
	ePackage.mock.Called(nsURI)
}

type MockEPackage_SetNsURI_Call struct {
	*mock.Call
}

// SetNsURI is a helper method to define mock.On call
// - nsURI string
func (e *MockEPackage_Expecter_Methods) SetNsURI(nsURI any) *MockEPackage_SetNsURI_Call {
	return &MockEPackage_SetNsURI_Call{Call: e.mock.On("SetNsURI", nsURI)}
}

func (c *MockEPackage_SetNsURI_Call) Run(run func(nsURI string)) *MockEPackage_SetNsURI_Call {
	c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return c
}

func (c *MockEPackage_SetNsURI_Call) Return() *MockEPackage_SetNsURI_Call {
	c.Call.Return()
	return c
}

// GetEClassifier provides mock implementation
func (ePackage *MockEPackage_Prototype_Methods) GetEClassifier(name string) EClassifier {
	ret := ePackage.mock.Called(name)

	var r EClassifier
	if rf, ok := ret.Get(0).(func() EClassifier); ok {
		r = rf()
	} else {
		if ret.Get(0) != nil {
			r = ret.Get(0).(EClassifier)
		}
	}

	return r
}

type MockEPackage_GetEClassifier_Call struct {
	*mock.Call
}

// GetEClassifier is a helper method to define mock.On call
// - name string
func (e *MockEPackage_Expecter_Methods) GetEClassifier(name any) *MockEPackage_GetEClassifier_Call {
	return &MockEPackage_GetEClassifier_Call{Call: e.mock.On("GetEClassifier", name)}
}

func (c *MockEPackage_GetEClassifier_Call) Run(run func(string)) *MockEPackage_GetEClassifier_Call {
	c.Call.Run(func(_args mock.Arguments) {
		run(_args[0].(string))
	})
	return c
}

func (c *MockEPackage_GetEClassifier_Call) Return(_a0 EClassifier) *MockEPackage_GetEClassifier_Call {
	c.Call.Return(_a0)
	return c
}

type mockConstructorTestingTNewMockEPackage interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockEPackage creates a new instance of MockEPackage. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockEPackage(t mockConstructorTestingTNewMockEPackage) *MockEPackage {
	mock := &MockEPackage{}
	mock.SetMock(&mock.Mock)
	mock.Mock.Test(t)
	t.Cleanup(func() { mock.AssertExpectations(t) })
	return mock
}
