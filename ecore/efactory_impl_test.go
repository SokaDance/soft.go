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

func discardEFactory() {
	_ = assert.Equal
	_ = mock.Anything
	_ = testing.Coverage
	_ = NewMockEFactory
}

func TestEFactoryAsEFactory(t *testing.T) {
	o := newEFactoryImpl()
	assert.Equal(t, o, o.asEFactory())
}

func TestEFactoryStaticClass(t *testing.T) {
	o := newEFactoryImpl()
	assert.Equal(t, GetPackage().GetEFactory(), o.EStaticClass())
}

func TestEFactoryFeatureCount(t *testing.T) {
	o := newEFactoryImpl()
	assert.Equal(t, EFACTORY_FEATURE_COUNT, o.EStaticFeatureCount())
}

func TestEFactoryEPackageGet(t *testing.T) {
	// default
	o := newEFactoryImpl()
	assert.Nil(t, o.GetEPackage())

	// set a mock container
	v := NewMockEPackage(t)
	o.ESetInternalContainer(v, EFACTORY__EPACKAGE)

	// no proxy
	v.EXPECT().EIsProxy().Return(false).Once()
	assert.Equal(t, v, o.GetEPackage())
}

func TestEFactoryEPackageSet(t *testing.T) {
	// object
	o := newEFactoryImpl()

	// add listener
	mockAdapter := NewMockEAdapter(t)
	mockAdapter.EXPECT().SetTarget(o).Once()
	o.EAdapters().Add(mockAdapter)
	mock.AssertExpectationsForObjects(t, mockAdapter)

	// set with the mock value
	mockValue := NewMockEPackage(t)
	mockResource := NewMockEResource(t)
	mockValue.EXPECT().EInverseAdd(o, EPACKAGE__EFACTORY_INSTANCE, nil).Return(nil).Once()
	mockValue.EXPECT().EResource().Return(mockResource).Once()
	mockResource.EXPECT().Attached(o).Once()
	mockAdapter.EXPECT().NotifyChanged(mock.Anything).Once()
	o.SetEPackage(mockValue)
	mock.AssertExpectationsForObjects(t, mockAdapter, mockValue, mockResource)

	// set with the same mock value
	mockAdapter.EXPECT().NotifyChanged(mock.Anything).Once()
	o.SetEPackage(mockValue)
	mock.AssertExpectationsForObjects(t, mockAdapter, mockValue, mockResource)

	// another value - in a different resource
	mockValue2 := NewMockEPackage(t)
	mockResource2 := NewMockEResource(t)
	mockValue.EXPECT().EInverseRemove(o, EPACKAGE__EFACTORY_INSTANCE, nil).Return(nil).Once()
	mockValue.EXPECT().EResource().Return(mockResource).Once()
	mockValue2.EXPECT().EInverseAdd(o, EPACKAGE__EFACTORY_INSTANCE, nil).Return(nil).Once()
	mockValue2.EXPECT().EResource().Return(mockResource2).Once()
	mockResource.EXPECT().Detached(o).Once()
	mockResource2.EXPECT().Attached(o).Once()
	mockAdapter.EXPECT().NotifyChanged(mock.Anything).Once()
	o.SetEPackage(mockValue2)
	mock.AssertExpectationsForObjects(t, mockAdapter, mockValue, mockResource, mockValue2, mockResource2)
}

func TestEFactoryEPackageBasicSet(t *testing.T) {
	o := newEFactoryImpl()

	// add listener
	mockAdapter := NewMockEAdapter(t)
	mockAdapter.EXPECT().SetTarget(o).Once()
	o.EAdapters().Add(mockAdapter)
	mock.AssertExpectationsForObjects(t, mockAdapter)

	mockValue := NewMockEPackage(t)
	mockNotifications := NewMockENotificationChain(t)
	mockValue.EXPECT().EResource().Return(nil).Once()
	mockNotifications.EXPECT().Add(mock.MatchedBy(func(notification ENotification) bool {
		return notification.GetEventType() == SET && notification.GetFeatureID() == EFACTORY__EPACKAGE
	})).Return(true).Once()
	o.basicSetEPackage(mockValue, mockNotifications)
	mock.AssertExpectationsForObjects(t, mockAdapter, mockValue, mockNotifications)
}

func TestEFactoryConvertToStringOperation(t *testing.T) {
	o := newEFactoryImpl()
	assert.Panics(t, func() { o.ConvertToString(nil, nil) })
}
func TestEFactoryCreateOperation(t *testing.T) {
	o := newEFactoryImpl()
	assert.Panics(t, func() { o.Create(nil) })
}
func TestEFactoryCreateFromStringOperation(t *testing.T) {
	o := newEFactoryImpl()
	assert.Panics(t, func() { o.CreateFromString(nil, "") })
}

func TestEFactoryEGetFromID(t *testing.T) {
	o := newEFactoryImpl()
	assert.Panics(t, func() { o.EGetFromID(-1, true) })
	assert.Equal(t, o.GetEPackage(), o.EGetFromID(EFACTORY__EPACKAGE, true))
}

func TestEFactoryESetFromID(t *testing.T) {
	o := newEFactoryImpl()
	assert.Panics(t, func() { o.ESetFromID(-1, nil) })
	{
		mockValue := NewMockEPackage(t)
		mockValue.EXPECT().EIsProxy().Return(false).Once()
		mockValue.EXPECT().EResource().Return(nil).Once()
		mockValue.EXPECT().EInverseAdd(o, EPACKAGE__EFACTORY_INSTANCE, nil).Return(nil).Once()
		o.ESetFromID(EFACTORY__EPACKAGE, mockValue)
		assert.Equal(t, mockValue, o.EGetFromID(EFACTORY__EPACKAGE, false))
		mock.AssertExpectationsForObjects(t, mockValue)
	}

}

func TestEFactoryEIsSetFromID(t *testing.T) {
	o := newEFactoryImpl()
	assert.Panics(t, func() { o.EIsSetFromID(-1) })
	assert.False(t, o.EIsSetFromID(EFACTORY__EPACKAGE))
}

func TestEFactoryEUnsetFromID(t *testing.T) {
	o := newEFactoryImpl()
	assert.Panics(t, func() { o.EUnsetFromID(-1) })
	{
		o.EUnsetFromID(EFACTORY__EPACKAGE)
		assert.Nil(t, o.EGetFromID(EFACTORY__EPACKAGE, false))
	}
}

func TestEFactoryEInvokeFromID(t *testing.T) {
	o := newEFactoryImpl()
	assert.Panics(t, func() { o.EInvokeFromID(-1, nil) })
	assert.Panics(t, func() { o.EInvokeFromID(EFACTORY__CONVERT_TO_STRING_EDATATYPE_EJAVAOBJECT, nil) })
	assert.Panics(t, func() { o.EInvokeFromID(EFACTORY__CREATE_ECLASS, nil) })
	assert.Panics(t, func() { o.EInvokeFromID(EFACTORY__CREATE_FROM_STRING_EDATATYPE_ESTRING, nil) })
}

func TestEFactoryEBasicInverseAdd(t *testing.T) {
	o := newEFactoryImpl()
	{
		mockObject := NewMockEObject(t)
		mockNotifications := NewMockENotificationChain(t)
		assert.Equal(t, mockNotifications, o.EBasicInverseAdd(mockObject, -1, mockNotifications))
	}
	{
		mockObject := NewMockEPackage(t)
		mockObject.EXPECT().EResource().Return(nil).Once()
		mockObject.EXPECT().EIsProxy().Return(false).Once()
		o.EBasicInverseAdd(mockObject, EFACTORY__EPACKAGE, nil)
		assert.Equal(t, mockObject, o.GetEPackage())
		mock.AssertExpectationsForObjects(t, mockObject)

		mockOther := NewMockEPackage(t)
		mockOther.EXPECT().EResource().Return(nil).Once()
		mockOther.EXPECT().EIsProxy().Return(false).Once()
		mockObject.EXPECT().EResource().Return(nil).Once()
		mockObject.EXPECT().EInverseRemove(o, EPACKAGE__EFACTORY_INSTANCE, nil).Return(nil).Once()
		o.EBasicInverseAdd(mockOther, EFACTORY__EPACKAGE, nil)
		assert.Equal(t, mockOther, o.GetEPackage())
		mock.AssertExpectationsForObjects(t, mockObject, mockOther)
	}

}

func TestEFactoryEBasicInverseRemove(t *testing.T) {
	o := newEFactoryImpl()
	{
		mockObject := NewMockEObject(t)
		mockNotifications := NewMockENotificationChain(t)
		assert.Equal(t, mockNotifications, o.EBasicInverseRemove(mockObject, -1, mockNotifications))
	}
	{
		mockObject := NewMockEPackage(t)
		o.EBasicInverseRemove(mockObject, EFACTORY__EPACKAGE, nil)
		mock.AssertExpectationsForObjects(t, mockObject)
	}

}
