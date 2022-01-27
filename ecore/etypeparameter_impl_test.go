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
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func discardETypeParameter() {
	_ = assert.Equal
	_ = mock.Anything
	_ = testing.Coverage
}

func TestETypeParameterAsETypeParameter(t *testing.T) {
	o := newETypeParameterImpl()
	assert.Equal(t, o, o.asETypeParameter())
}

func TestETypeParameterStaticClass(t *testing.T) {
	o := newETypeParameterImpl()
	assert.Equal(t, GetPackage().GetETypeParameter(), o.EStaticClass())
}

func TestETypeParameterFeatureCount(t *testing.T) {
	o := newETypeParameterImpl()
	assert.Equal(t, ETYPE_PARAMETER_FEATURE_COUNT, o.EStaticFeatureCount())
}

func TestETypeParameterEBoundsGet(t *testing.T) {
	o := newETypeParameterImpl()
	assert.NotNil(t, o.GetEBounds())
}

func TestETypeParameterEGetFromID(t *testing.T) {
	o := newETypeParameterImpl()
	assert.Panics(t, func() { o.EGetFromID(-1, true) })
	assert.Equal(t, o.GetEBounds(), FromAnyList[EGenericType](o.EGetFromID(ETYPE_PARAMETER__EBOUNDS, true)))
}

func TestETypeParameterESetFromID(t *testing.T) {
	o := newETypeParameterImpl()
	assert.Panics(t, func() { o.ESetFromID(-1, nil) })
	{
		// list with a value
		mockList := &MockEList[EGenericType]{}
		mockValue := &MockEGenericType{}
		mockIterator := &MockEIterator[EGenericType]{}
		mockList.On("Iterator").Return(mockIterator).Once()
		mockIterator.On("HasNext").Return(true).Once()
		mockIterator.On("Next").Return(mockValue).Once()
		mockIterator.On("HasNext").Return(false).Once()
		mockValue.On("EInverseAdd", o, EOPPOSITE_FEATURE_BASE-ETYPE_PARAMETER__EBOUNDS, mock.Anything).Return(nil).Once()

		// set list with new contents
		o.ESetFromID(ETYPE_PARAMETER__EBOUNDS, ToAnyList[EGenericType](mockList))
		// checks
		assert.Equal(t, 1, o.GetEBounds().Size())
		assert.Equal(t, mockValue, o.GetEBounds().Get(0))
		mock.AssertExpectationsForObjects(t, mockList, mockIterator, mockValue)
	}

}

func TestETypeParameterEIsSetFromID(t *testing.T) {
	o := newETypeParameterImpl()
	assert.Panics(t, func() { o.EIsSetFromID(-1) })
	assert.False(t, o.EIsSetFromID(ETYPE_PARAMETER__EBOUNDS))
}

func TestETypeParameterEUnsetFromID(t *testing.T) {
	o := newETypeParameterImpl()
	assert.Panics(t, func() { o.EUnsetFromID(-1) })
	{
		o.EUnsetFromID(ETYPE_PARAMETER__EBOUNDS)
		v := o.EGetFromID(ETYPE_PARAMETER__EBOUNDS, false)
		assert.NotNil(t, v)
		l := v.(EList[any])
		assert.True(t, l.Empty())
	}
}

func TestETypeParameterEBasicInverseRemove(t *testing.T) {
	o := newETypeParameterImpl()
	{
		mockObject := new(MockEObject)
		mockNotifications := new(MockENotificationChain)
		assert.Equal(t, mockNotifications, o.EBasicInverseRemove(mockObject, -1, mockNotifications))
	}
	{
		// initialize list with a mock object
		mockObject := &MockEGenericType{}
		mockObject.On("EInverseAdd", o, EOPPOSITE_FEATURE_BASE-ETYPE_PARAMETER__EBOUNDS, mock.Anything).Return(nil).Once()

		l := o.GetEBounds()
		l.Add(mockObject)

		// basic inverse remove
		o.EBasicInverseRemove(mockObject, ETYPE_PARAMETER__EBOUNDS, nil)

		// check it was removed
		assert.False(t, l.Contains(mockObject))
		mock.AssertExpectationsForObjects(t, mockObject)
	}

}
