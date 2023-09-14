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

func discardEStructuralFeature() {
	_ = assert.Equal
	_ = mock.Anything
	_ = testing.Coverage
}

func TestEStructuralFeatureAsEStructuralFeature(t *testing.T) {
	o := newEStructuralFeatureImpl()
	assert.Equal(t, o, o.asEStructuralFeature())
}

func TestEStructuralFeatureStaticClass(t *testing.T) {
	o := newEStructuralFeatureImpl()
	assert.Equal(t, GetPackage().GetEStructuralFeature(), o.EStaticClass())
}

func TestEStructuralFeatureFeatureCount(t *testing.T) {
	o := newEStructuralFeatureImpl()
	assert.Equal(t, ESTRUCTURAL_FEATURE_FEATURE_COUNT, o.EStaticFeatureCount())
}

func TestEStructuralFeatureDefaultValueGet(t *testing.T) {
	o := newEStructuralFeatureImpl()
	assert.Panics(t, func() { o.GetDefaultValue() })
}

func TestEStructuralFeatureDefaultValueSet(t *testing.T) {
	o := newEStructuralFeatureImpl()
	v := any(nil)
	assert.Panics(t, func() { o.SetDefaultValue(v) })
}

func TestEStructuralFeatureDefaultValueLiteralGet(t *testing.T) {
	o := newEStructuralFeatureImpl()
	// get default value
	assert.Equal(t, string(""), o.GetDefaultValueLiteral())
	// get initialized value
	v := string("Test String")
	o.defaultValueLiteral = v
	assert.Equal(t, v, o.GetDefaultValueLiteral())
}

func TestEStructuralFeatureDefaultValueLiteralSet(t *testing.T) {
	o := newEStructuralFeatureImpl()
	v := string("Test String")
	mockAdapter := NewMockEAdapter(t)
	mockAdapter.EXPECT().SetTarget(o).Once()
	mockAdapter.EXPECT().NotifyChanged(mock.Anything).Once()
	o.EAdapters().Add(mockAdapter)
	o.SetDefaultValueLiteral(v)
	mockAdapter.AssertExpectations(t)
}

func TestEStructuralFeatureEContainingClassGet(t *testing.T) {
	// default
	o := newEStructuralFeatureImpl()
	assert.Nil(t, o.GetEContainingClass())

	// set a mock container
	v := NewMockEClass(t)
	o.ESetInternalContainer(v, ESTRUCTURAL_FEATURE__ECONTAINING_CLASS)

	// no proxy
	v.EXPECT().EIsProxy().Return(false).Once()
	assert.Equal(t, v, o.GetEContainingClass())
}

func TestEStructuralFeatureFeatureIDGet(t *testing.T) {
	o := newEStructuralFeatureImpl()
	// get default value
	assert.Equal(t, int(-1), o.GetFeatureID())
	// get initialized value
	v := int(45)
	o.featureID = v
	assert.Equal(t, v, o.GetFeatureID())
}

func TestEStructuralFeatureFeatureIDSet(t *testing.T) {
	o := newEStructuralFeatureImpl()
	v := int(45)
	mockAdapter := NewMockEAdapter(t)
	mockAdapter.EXPECT().SetTarget(o).Once()
	mockAdapter.EXPECT().NotifyChanged(mock.Anything).Once()
	o.EAdapters().Add(mockAdapter)
	o.SetFeatureID(v)
	mockAdapter.AssertExpectations(t)
}

func TestEStructuralFeatureChangeableGet(t *testing.T) {
	o := newEStructuralFeatureImpl()
	// get default value
	assert.Equal(t, bool(true), o.IsChangeable())
	// get initialized value
	v := bool(true)
	o.isChangeable = v
	assert.Equal(t, v, o.IsChangeable())
}

func TestEStructuralFeatureChangeableSet(t *testing.T) {
	o := newEStructuralFeatureImpl()
	v := bool(true)
	mockAdapter := NewMockEAdapter(t)
	mockAdapter.EXPECT().SetTarget(o).Once()
	mockAdapter.EXPECT().NotifyChanged(mock.Anything).Once()
	o.EAdapters().Add(mockAdapter)
	o.SetChangeable(v)
	mockAdapter.AssertExpectations(t)
}

func TestEStructuralFeatureDerivedGet(t *testing.T) {
	o := newEStructuralFeatureImpl()
	// get default value
	assert.Equal(t, bool(false), o.IsDerived())
	// get initialized value
	v := bool(true)
	o.isDerived = v
	assert.Equal(t, v, o.IsDerived())
}

func TestEStructuralFeatureDerivedSet(t *testing.T) {
	o := newEStructuralFeatureImpl()
	v := bool(true)
	mockAdapter := NewMockEAdapter(t)
	mockAdapter.EXPECT().SetTarget(o).Once()
	mockAdapter.EXPECT().NotifyChanged(mock.Anything).Once()
	o.EAdapters().Add(mockAdapter)
	o.SetDerived(v)
	mockAdapter.AssertExpectations(t)
}

func TestEStructuralFeatureTransientGet(t *testing.T) {
	o := newEStructuralFeatureImpl()
	// get default value
	assert.Equal(t, bool(false), o.IsTransient())
	// get initialized value
	v := bool(true)
	o.isTransient = v
	assert.Equal(t, v, o.IsTransient())
}

func TestEStructuralFeatureTransientSet(t *testing.T) {
	o := newEStructuralFeatureImpl()
	v := bool(true)
	mockAdapter := NewMockEAdapter(t)
	mockAdapter.EXPECT().SetTarget(o).Once()
	mockAdapter.EXPECT().NotifyChanged(mock.Anything).Once()
	o.EAdapters().Add(mockAdapter)
	o.SetTransient(v)
	mockAdapter.AssertExpectations(t)
}

func TestEStructuralFeatureUnsettableGet(t *testing.T) {
	o := newEStructuralFeatureImpl()
	// get default value
	assert.Equal(t, bool(false), o.IsUnsettable())
	// get initialized value
	v := bool(true)
	o.isUnsettable = v
	assert.Equal(t, v, o.IsUnsettable())
}

func TestEStructuralFeatureUnsettableSet(t *testing.T) {
	o := newEStructuralFeatureImpl()
	v := bool(true)
	mockAdapter := NewMockEAdapter(t)
	mockAdapter.EXPECT().SetTarget(o).Once()
	mockAdapter.EXPECT().NotifyChanged(mock.Anything).Once()
	o.EAdapters().Add(mockAdapter)
	o.SetUnsettable(v)
	mockAdapter.AssertExpectations(t)
}

func TestEStructuralFeatureVolatileGet(t *testing.T) {
	o := newEStructuralFeatureImpl()
	// get default value
	assert.Equal(t, bool(false), o.IsVolatile())
	// get initialized value
	v := bool(true)
	o.isVolatile = v
	assert.Equal(t, v, o.IsVolatile())
}

func TestEStructuralFeatureVolatileSet(t *testing.T) {
	o := newEStructuralFeatureImpl()
	v := bool(true)
	mockAdapter := NewMockEAdapter(t)
	mockAdapter.EXPECT().SetTarget(o).Once()
	mockAdapter.EXPECT().NotifyChanged(mock.Anything).Once()
	o.EAdapters().Add(mockAdapter)
	o.SetVolatile(v)
	mockAdapter.AssertExpectations(t)
}

func TestEStructuralFeatureGetContainerClassOperation(t *testing.T) {
	o := newEStructuralFeatureImpl()
	assert.Panics(t, func() { o.GetContainerClass() })
}

func TestEStructuralFeatureEGetFromID(t *testing.T) {
	o := newEStructuralFeatureImpl()
	assert.Panics(t, func() { o.EGetFromID(-1, true) })
	assert.Equal(t, o.IsChangeable(), o.EGetFromID(ESTRUCTURAL_FEATURE__CHANGEABLE, true))
	assert.Panics(t, func() { o.EGetFromID(ESTRUCTURAL_FEATURE__DEFAULT_VALUE, true) })
	assert.Panics(t, func() { o.EGetFromID(ESTRUCTURAL_FEATURE__DEFAULT_VALUE, false) })
	assert.Equal(t, o.GetDefaultValueLiteral(), o.EGetFromID(ESTRUCTURAL_FEATURE__DEFAULT_VALUE_LITERAL, true))
	assert.Equal(t, o.IsDerived(), o.EGetFromID(ESTRUCTURAL_FEATURE__DERIVED, true))
	assert.Equal(t, o.GetEContainingClass(), o.EGetFromID(ESTRUCTURAL_FEATURE__ECONTAINING_CLASS, true))
	assert.Equal(t, o.GetFeatureID(), o.EGetFromID(ESTRUCTURAL_FEATURE__FEATURE_ID, true))
	assert.Equal(t, o.IsTransient(), o.EGetFromID(ESTRUCTURAL_FEATURE__TRANSIENT, true))
	assert.Equal(t, o.IsUnsettable(), o.EGetFromID(ESTRUCTURAL_FEATURE__UNSETTABLE, true))
	assert.Equal(t, o.IsVolatile(), o.EGetFromID(ESTRUCTURAL_FEATURE__VOLATILE, true))
}

func TestEStructuralFeatureESetFromID(t *testing.T) {
	o := newEStructuralFeatureImpl()
	assert.Panics(t, func() { o.ESetFromID(-1, nil) })
	{
		v := bool(true)
		o.ESetFromID(ESTRUCTURAL_FEATURE__CHANGEABLE, v)
		assert.Equal(t, v, o.EGetFromID(ESTRUCTURAL_FEATURE__CHANGEABLE, false))
	}
	assert.Panics(t, func() { o.ESetFromID(ESTRUCTURAL_FEATURE__DEFAULT_VALUE, nil) })
	{
		v := string("Test String")
		o.ESetFromID(ESTRUCTURAL_FEATURE__DEFAULT_VALUE_LITERAL, v)
		assert.Equal(t, v, o.EGetFromID(ESTRUCTURAL_FEATURE__DEFAULT_VALUE_LITERAL, false))
	}
	{
		v := bool(true)
		o.ESetFromID(ESTRUCTURAL_FEATURE__DERIVED, v)
		assert.Equal(t, v, o.EGetFromID(ESTRUCTURAL_FEATURE__DERIVED, false))
	}
	{
		v := int(45)
		o.ESetFromID(ESTRUCTURAL_FEATURE__FEATURE_ID, v)
		assert.Equal(t, v, o.EGetFromID(ESTRUCTURAL_FEATURE__FEATURE_ID, false))
	}
	{
		v := bool(true)
		o.ESetFromID(ESTRUCTURAL_FEATURE__TRANSIENT, v)
		assert.Equal(t, v, o.EGetFromID(ESTRUCTURAL_FEATURE__TRANSIENT, false))
	}
	{
		v := bool(true)
		o.ESetFromID(ESTRUCTURAL_FEATURE__UNSETTABLE, v)
		assert.Equal(t, v, o.EGetFromID(ESTRUCTURAL_FEATURE__UNSETTABLE, false))
	}
	{
		v := bool(true)
		o.ESetFromID(ESTRUCTURAL_FEATURE__VOLATILE, v)
		assert.Equal(t, v, o.EGetFromID(ESTRUCTURAL_FEATURE__VOLATILE, false))
	}

}

func TestEStructuralFeatureEIsSetFromID(t *testing.T) {
	o := newEStructuralFeatureImpl()
	assert.Panics(t, func() { o.EIsSetFromID(-1) })
	assert.False(t, o.EIsSetFromID(ESTRUCTURAL_FEATURE__CHANGEABLE))
	assert.Panics(t, func() { o.EIsSetFromID(ESTRUCTURAL_FEATURE__DEFAULT_VALUE) })
	assert.False(t, o.EIsSetFromID(ESTRUCTURAL_FEATURE__DEFAULT_VALUE_LITERAL))
	assert.False(t, o.EIsSetFromID(ESTRUCTURAL_FEATURE__DERIVED))
	assert.False(t, o.EIsSetFromID(ESTRUCTURAL_FEATURE__ECONTAINING_CLASS))
	assert.False(t, o.EIsSetFromID(ESTRUCTURAL_FEATURE__FEATURE_ID))
	assert.False(t, o.EIsSetFromID(ESTRUCTURAL_FEATURE__TRANSIENT))
	assert.False(t, o.EIsSetFromID(ESTRUCTURAL_FEATURE__UNSETTABLE))
	assert.False(t, o.EIsSetFromID(ESTRUCTURAL_FEATURE__VOLATILE))
}

func TestEStructuralFeatureEUnsetFromID(t *testing.T) {
	o := newEStructuralFeatureImpl()
	assert.Panics(t, func() { o.EUnsetFromID(-1) })
	{
		o.EUnsetFromID(ESTRUCTURAL_FEATURE__CHANGEABLE)
		v := o.EGetFromID(ESTRUCTURAL_FEATURE__CHANGEABLE, false)
		assert.Equal(t, bool(true), v)
	}
	{
		assert.Panics(t, func() { o.EUnsetFromID(ESTRUCTURAL_FEATURE__DEFAULT_VALUE) })
	}
	{
		o.EUnsetFromID(ESTRUCTURAL_FEATURE__DEFAULT_VALUE_LITERAL)
		v := o.EGetFromID(ESTRUCTURAL_FEATURE__DEFAULT_VALUE_LITERAL, false)
		assert.Equal(t, string(""), v)
	}
	{
		o.EUnsetFromID(ESTRUCTURAL_FEATURE__DERIVED)
		v := o.EGetFromID(ESTRUCTURAL_FEATURE__DERIVED, false)
		assert.Equal(t, bool(false), v)
	}
	{
		o.EUnsetFromID(ESTRUCTURAL_FEATURE__FEATURE_ID)
		v := o.EGetFromID(ESTRUCTURAL_FEATURE__FEATURE_ID, false)
		assert.Equal(t, int(-1), v)
	}
	{
		o.EUnsetFromID(ESTRUCTURAL_FEATURE__TRANSIENT)
		v := o.EGetFromID(ESTRUCTURAL_FEATURE__TRANSIENT, false)
		assert.Equal(t, bool(false), v)
	}
	{
		o.EUnsetFromID(ESTRUCTURAL_FEATURE__UNSETTABLE)
		v := o.EGetFromID(ESTRUCTURAL_FEATURE__UNSETTABLE, false)
		assert.Equal(t, bool(false), v)
	}
	{
		o.EUnsetFromID(ESTRUCTURAL_FEATURE__VOLATILE)
		v := o.EGetFromID(ESTRUCTURAL_FEATURE__VOLATILE, false)
		assert.Equal(t, bool(false), v)
	}
}

func TestEStructuralFeatureEInvokeFromID(t *testing.T) {
	o := newEStructuralFeatureImpl()
	assert.Panics(t, func() { o.EInvokeFromID(-1, nil) })
	assert.Panics(t, func() { o.EInvokeFromID(ESTRUCTURAL_FEATURE__GET_CONTAINER_CLASS, nil) })
}

func TestEStructuralFeatureEBasicInverseAdd(t *testing.T) {
	o := newEStructuralFeatureImpl()
	{
		mockObject := NewMockEObject(t)
		mockNotifications := NewMockENotificationChain(t)
		assert.Equal(t, mockNotifications, o.EBasicInverseAdd(mockObject, -1, mockNotifications))
	}
	{
		mockObject := NewMockEClass(t)
		mockObject.EXPECT().EResource().Return(nil).Once()
		mockObject.EXPECT().EIsProxy().Return(false).Once()
		o.EBasicInverseAdd(mockObject, ESTRUCTURAL_FEATURE__ECONTAINING_CLASS, nil)
		assert.Equal(t, mockObject, o.GetEContainingClass())
		mock.AssertExpectationsForObjects(t, mockObject)

		mockOther := NewMockEClass(t)
		mockOther.EXPECT().EResource().Return(nil).Once()
		mockOther.EXPECT().EIsProxy().Return(false).Once()
		mockObject.EXPECT().EResource().Return(nil).Once()
		mockObject.EXPECT().EInverseRemove(o, ECLASS__ESTRUCTURAL_FEATURES, nil).Return(nil).Once()
		o.EBasicInverseAdd(mockOther, ESTRUCTURAL_FEATURE__ECONTAINING_CLASS, nil)
		assert.Equal(t, mockOther, o.GetEContainingClass())
		mock.AssertExpectationsForObjects(t, mockObject, mockOther)
	}

}

func TestEStructuralFeatureEBasicInverseRemove(t *testing.T) {
	o := newEStructuralFeatureImpl()
	{
		mockObject := NewMockEObject(t)
		mockNotifications := NewMockENotificationChain(t)
		assert.Equal(t, mockNotifications, o.EBasicInverseRemove(mockObject, -1, mockNotifications))
	}
	{
		mockObject := NewMockEClass(t)
		o.EBasicInverseRemove(mockObject, ESTRUCTURAL_FEATURE__ECONTAINING_CLASS, nil)
		mock.AssertExpectationsForObjects(t, mockObject)
	}

}
