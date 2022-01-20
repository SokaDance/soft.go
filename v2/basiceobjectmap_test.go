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
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestBasicEObjectMap_Constructor(t *testing.T) {
	mockClass := &MockEClass{}
	m := NewBasicEObjectMap[EObject, EObject](mockClass)
	assert.NotNil(t, m)

	var mp EMap[EObject, EObject] = m
	assert.NotNil(t, mp)

	var ml EList[any] = m
	assert.NotNil(t, ml)
}

type MockEObjectEMapEntry struct {
	MockEObject
}

func (m *MockEObjectEMapEntry) GetKey() EObject {
	ret := m.Called()

	var r0 EObject
	if rf, ok := ret.Get(0).(func() EObject); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(EObject)
		}
	}

	return r0
}

// GetValue provides a mock function with given fields:
func (m *MockEObjectEMapEntry) GetValue() EObject {
	ret := m.Called()

	var r0 EObject
	if rf, ok := ret.Get(0).(func() EObject); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(EObject)
		}
	}

	return r0
}

// SetKey provides a mock function with given fields: _a0
func (m *MockEObjectEMapEntry) SetKey(_a0 EObject) {
	m.Called(_a0)
}

// SetValue provides a mock function with given fields: _a0
func (m *MockEObjectEMapEntry) SetValue(_a0 EObject) {
	m.Called(_a0)
}

func (m *MockEObjectEMapEntry) SetAnyValue(v any) {
	m.Called(v)
}

func (m *MockEObjectEMapEntry) SetAnyKey(k any) {
	m.Called(k)
}

func TestBasicEObjectMap_Put(t *testing.T) {
	mockClass := &MockEClass{}
	mockPackage := &MockEPackage{}
	mockFactory := &MockEFactory{}
	mockEntry := &MockEObjectEMapEntry{}
	mockKey := &MockEObject{}
	mockValue := &MockEObject{}
	m := NewBasicEObjectMap[EObject, EObject](mockClass)

	mockClass.On("GetEPackage").Once().Return(mockPackage)
	mockPackage.On("GetEFactoryInstance").Once().Return(mockFactory)
	mockFactory.On("Create", mockClass).Once().Return(mockEntry)
	mockEntry.On("SetAnyKey", mockKey).Once()
	mockEntry.On("SetAnyValue", mockValue).Once()
	mockEntry.On("GetKey").Once().Return(func () EObject{ return mockKey} )
	mockEntry.On("GetValue").Once().Return(func () EObject{ return mockValue})
	m.Put(mockKey, mockValue)
	mock.AssertExpectationsForObjects(t, mockClass, mockPackage, mockFactory, mockEntry, mockKey, mockValue)
}
