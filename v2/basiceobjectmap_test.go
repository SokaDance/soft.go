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
	MockEMapEntry[EObject, EObject]
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
	mockEntry.On("SetKey", mockKey).Once()
	mockEntry.On("SetValue", mockValue).Once()
	mockEntry.On("GetKey").Once().Return(mockKey)
	mockEntry.On("GetValue").Once().Return(mockValue)
	m.Put(mockKey, mockValue)
	mock.AssertExpectationsForObjects(t, mockClass, mockPackage, mockFactory, mockEntry, mockKey, mockValue)
}
