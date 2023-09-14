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

import "github.com/stretchr/testify/assert"
import "github.com/stretchr/testify/mock"
import "testing"

func discardETypedElement() {
	_ = assert.Equal
	_ = mock.Anything
	_ = testing.Coverage
}

func TestETypedElementAsETypedElement(t *testing.T) {
	o := newETypedElementImpl()
	assert.Equal(t, o, o.asETypedElement())
}

func TestETypedElementStaticClass(t *testing.T) {
	o := newETypedElementImpl()
	assert.Equal(t, GetPackage().GetETypedElement(), o.EStaticClass())
}

func TestETypedElementFeatureCount(t *testing.T) {
	o := newETypedElementImpl()
	assert.Equal(t, ETYPED_ELEMENT_FEATURE_COUNT, o.EStaticFeatureCount())
}

func TestETypedElementETypeGet(t *testing.T) {
	o := newETypedElementImpl()

	// get default value
	assert.Nil(t, o.GetEType())

	// initialize object with a mock value
	mockValue := NewMockEClassifier(t)
	o.eType = mockValue

	// events
	mockAdapter := NewMockEAdapter(t)
	mockAdapter.EXPECT().SetTarget(o).Once()
	o.EAdapters().Add(mockAdapter)
	mock.AssertExpectationsForObjects(t, mockAdapter)

	// set object resource
	mockResourceSet := NewMockEResourceSet(t)
	mockResource := NewMockEResource(t)
	o.ESetInternalResource(mockResource)

	// get non resolved value
	mockValue.EXPECT().EIsProxy().Return(false).Once()
	assert.Equal(t, mockValue, o.GetEType())
	mock.AssertExpectationsForObjects(t, mockValue, mockAdapter, mockResource, mockResourceSet)

	// get a resolved value
	mockURI := NewURI("test:///file.t")
	mockResolved := NewMockEClassifier(t)
	mockResolved.EXPECT().EProxyURI().Return(nil).Once()
	mockResource.EXPECT().GetResourceSet().Return(mockResourceSet).Once()
	mockResourceSet.EXPECT().GetEObject(mockURI, true).Return(mockResolved).Once()
	mockValue.EXPECT().EIsProxy().Return(true).Once()
	mockValue.EXPECT().EProxyURI().Return(mockURI).Twice()
	mockAdapter.EXPECT().NotifyChanged(mock.MatchedBy(func(notification ENotification) bool {
		return notification.GetEventType() == RESOLVE && notification.GetFeatureID() == ETYPED_ELEMENT__ETYPE && notification.GetOldValue() == mockValue && notification.GetNewValue() == mockResolved
	})).Once()
	assert.Equal(t, mockResolved, o.GetEType())
	mock.AssertExpectationsForObjects(t, mockAdapter, mockValue, mockResolved, mockAdapter, mockResource, mockResourceSet)
}

func TestETypedElementETypeSet(t *testing.T) {
	o := newETypedElementImpl()
	v := NewMockEClassifier(t)
	mockAdapter := NewMockEAdapter(t)
	mockAdapter.EXPECT().SetTarget(o).Once()
	mockAdapter.EXPECT().NotifyChanged(mock.Anything).Once()
	o.EAdapters().Add(mockAdapter)
	o.SetEType(v)
	mockAdapter.AssertExpectations(t)
}

func TestETypedElementETypeUnSet(t *testing.T) {
	o := newETypedElementImpl()
	mockAdapter := NewMockEAdapter(t)
	mockAdapter.EXPECT().SetTarget(o).Once()
	o.EAdapters().Add(mockAdapter)

	mockAdapter.EXPECT().NotifyChanged(mock.MatchedBy(func(notification ENotification) bool {
		return notification.GetEventType() == UNSET && notification.GetFeatureID() == ETYPED_ELEMENT__ETYPE
	})).Once()
	o.UnsetEType()
	assert.Nil(t, o.GetEType())
	mock.AssertExpectationsForObjects(t, mockAdapter)
}

func TestETypedElementManyGet(t *testing.T) {
	o := newETypedElementImpl()
	assert.Panics(t, func() { o.IsMany() })
}

func TestETypedElementOrderedGet(t *testing.T) {
	o := newETypedElementImpl()
	// get default value
	assert.Equal(t, bool(true), o.IsOrdered())
	// get initialized value
	v := bool(true)
	o.isOrdered = v
	assert.Equal(t, v, o.IsOrdered())
}

func TestETypedElementOrderedSet(t *testing.T) {
	o := newETypedElementImpl()
	v := bool(true)
	mockAdapter := NewMockEAdapter(t)
	mockAdapter.EXPECT().SetTarget(o).Once()
	mockAdapter.EXPECT().NotifyChanged(mock.Anything).Once()
	o.EAdapters().Add(mockAdapter)
	o.SetOrdered(v)
	mockAdapter.AssertExpectations(t)
}

func TestETypedElementRequiredGet(t *testing.T) {
	o := newETypedElementImpl()
	assert.Panics(t, func() { o.IsRequired() })
}

func TestETypedElementUniqueGet(t *testing.T) {
	o := newETypedElementImpl()
	// get default value
	assert.Equal(t, bool(true), o.IsUnique())
	// get initialized value
	v := bool(true)
	o.isUnique = v
	assert.Equal(t, v, o.IsUnique())
}

func TestETypedElementUniqueSet(t *testing.T) {
	o := newETypedElementImpl()
	v := bool(true)
	mockAdapter := NewMockEAdapter(t)
	mockAdapter.EXPECT().SetTarget(o).Once()
	mockAdapter.EXPECT().NotifyChanged(mock.Anything).Once()
	o.EAdapters().Add(mockAdapter)
	o.SetUnique(v)
	mockAdapter.AssertExpectations(t)
}

func TestETypedElementLowerBoundGet(t *testing.T) {
	o := newETypedElementImpl()
	// get default value
	assert.Equal(t, int(0), o.GetLowerBound())
	// get initialized value
	v := int(45)
	o.lowerBound = v
	assert.Equal(t, v, o.GetLowerBound())
}

func TestETypedElementLowerBoundSet(t *testing.T) {
	o := newETypedElementImpl()
	v := int(45)
	mockAdapter := NewMockEAdapter(t)
	mockAdapter.EXPECT().SetTarget(o).Once()
	mockAdapter.EXPECT().NotifyChanged(mock.Anything).Once()
	o.EAdapters().Add(mockAdapter)
	o.SetLowerBound(v)
	mockAdapter.AssertExpectations(t)
}

func TestETypedElementUpperBoundGet(t *testing.T) {
	o := newETypedElementImpl()
	// get default value
	assert.Equal(t, int(1), o.GetUpperBound())
	// get initialized value
	v := int(45)
	o.upperBound = v
	assert.Equal(t, v, o.GetUpperBound())
}

func TestETypedElementUpperBoundSet(t *testing.T) {
	o := newETypedElementImpl()
	v := int(45)
	mockAdapter := NewMockEAdapter(t)
	mockAdapter.EXPECT().SetTarget(o).Once()
	mockAdapter.EXPECT().NotifyChanged(mock.Anything).Once()
	o.EAdapters().Add(mockAdapter)
	o.SetUpperBound(v)
	mockAdapter.AssertExpectations(t)
}

func TestETypedElementEGetFromID(t *testing.T) {
	o := newETypedElementImpl()
	assert.Panics(t, func() { o.EGetFromID(-1, true) })
	assert.Equal(t, o.GetEType(), o.EGetFromID(ETYPED_ELEMENT__ETYPE, true))
	assert.Equal(t, o.GetLowerBound(), o.EGetFromID(ETYPED_ELEMENT__LOWER_BOUND, true))
	assert.Panics(t, func() { o.EGetFromID(ETYPED_ELEMENT__MANY, true) })
	assert.Panics(t, func() { o.EGetFromID(ETYPED_ELEMENT__MANY, false) })
	assert.Equal(t, o.IsOrdered(), o.EGetFromID(ETYPED_ELEMENT__ORDERED, true))
	assert.Panics(t, func() { o.EGetFromID(ETYPED_ELEMENT__REQUIRED, true) })
	assert.Panics(t, func() { o.EGetFromID(ETYPED_ELEMENT__REQUIRED, false) })
	assert.Equal(t, o.IsUnique(), o.EGetFromID(ETYPED_ELEMENT__UNIQUE, true))
	assert.Equal(t, o.GetUpperBound(), o.EGetFromID(ETYPED_ELEMENT__UPPER_BOUND, true))
}

func TestETypedElementESetFromID(t *testing.T) {
	o := newETypedElementImpl()
	assert.Panics(t, func() { o.ESetFromID(-1, nil) })
	{
		v := NewMockEClassifier(t)
		o.ESetFromID(ETYPED_ELEMENT__ETYPE, v)
		assert.Equal(t, v, o.EGetFromID(ETYPED_ELEMENT__ETYPE, false))

		o.ESetFromID(ETYPED_ELEMENT__ETYPE, nil)
	}
	{
		v := int(45)
		o.ESetFromID(ETYPED_ELEMENT__LOWER_BOUND, v)
		assert.Equal(t, v, o.EGetFromID(ETYPED_ELEMENT__LOWER_BOUND, false))
	}
	{
		v := bool(true)
		o.ESetFromID(ETYPED_ELEMENT__ORDERED, v)
		assert.Equal(t, v, o.EGetFromID(ETYPED_ELEMENT__ORDERED, false))
	}
	{
		v := bool(true)
		o.ESetFromID(ETYPED_ELEMENT__UNIQUE, v)
		assert.Equal(t, v, o.EGetFromID(ETYPED_ELEMENT__UNIQUE, false))
	}
	{
		v := int(45)
		o.ESetFromID(ETYPED_ELEMENT__UPPER_BOUND, v)
		assert.Equal(t, v, o.EGetFromID(ETYPED_ELEMENT__UPPER_BOUND, false))
	}

}

func TestETypedElementEIsSetFromID(t *testing.T) {
	o := newETypedElementImpl()
	assert.Panics(t, func() { o.EIsSetFromID(-1) })
	assert.False(t, o.EIsSetFromID(ETYPED_ELEMENT__ETYPE))
	assert.False(t, o.EIsSetFromID(ETYPED_ELEMENT__LOWER_BOUND))
	assert.Panics(t, func() { o.EIsSetFromID(ETYPED_ELEMENT__MANY) })
	assert.False(t, o.EIsSetFromID(ETYPED_ELEMENT__ORDERED))
	assert.Panics(t, func() { o.EIsSetFromID(ETYPED_ELEMENT__REQUIRED) })
	assert.False(t, o.EIsSetFromID(ETYPED_ELEMENT__UNIQUE))
	assert.False(t, o.EIsSetFromID(ETYPED_ELEMENT__UPPER_BOUND))
}

func TestETypedElementEUnsetFromID(t *testing.T) {
	o := newETypedElementImpl()
	assert.Panics(t, func() { o.EUnsetFromID(-1) })
	{
		o.EUnsetFromID(ETYPED_ELEMENT__ETYPE)
		assert.Nil(t, o.EGetFromID(ETYPED_ELEMENT__ETYPE, false))
	}
	{
		o.EUnsetFromID(ETYPED_ELEMENT__LOWER_BOUND)
		v := o.EGetFromID(ETYPED_ELEMENT__LOWER_BOUND, false)
		assert.Equal(t, int(0), v)
	}
	{
		o.EUnsetFromID(ETYPED_ELEMENT__ORDERED)
		v := o.EGetFromID(ETYPED_ELEMENT__ORDERED, false)
		assert.Equal(t, bool(true), v)
	}
	{
		o.EUnsetFromID(ETYPED_ELEMENT__UNIQUE)
		v := o.EGetFromID(ETYPED_ELEMENT__UNIQUE, false)
		assert.Equal(t, bool(true), v)
	}
	{
		o.EUnsetFromID(ETYPED_ELEMENT__UPPER_BOUND)
		v := o.EGetFromID(ETYPED_ELEMENT__UPPER_BOUND, false)
		assert.Equal(t, int(1), v)
	}
}
