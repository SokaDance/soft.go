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

func discardEPackage() {
	_ = assert.Equal
	_ = mock.Anything
	_ = testing.Coverage
}

func TestEPackageAsEPackage(t *testing.T) {
	o := newEPackageImpl()
	assert.Equal(t, o, o.asEPackage())
}

func TestEPackageStaticClass(t *testing.T) {
	o := newEPackageImpl()
	assert.Equal(t, GetPackage().GetEPackage(), o.EStaticClass())
}

func TestEPackageFeatureCount(t *testing.T) {
	o := newEPackageImpl()
	assert.Equal(t, EPACKAGE_FEATURE_COUNT, o.EStaticFeatureCount())
}

func TestEPackageEClassifiersGet(t *testing.T) {
	o := newEPackageImpl()
	assert.NotNil(t, o.GetEClassifiers())
	assert.Panics(t, func() { _ = o.GetEClassifiers().Get(0).(EClassifier) })
}

func TestEPackageEFactoryInstanceGet(t *testing.T) {
	o := newEPackageImpl()
	// get default value
	assert.Nil(t, o.GetEFactoryInstance())

	// get initialized value
	v := NewMockEFactory(t)
	o.eFactoryInstance = v
	assert.Equal(t, v, o.GetEFactoryInstance())
}

func TestEPackageEFactoryInstanceSet(t *testing.T) {
	o := newEPackageImpl()

	// add listener
	mockAdapter := NewMockEAdapter(t)
	mockAdapter.EXPECT().SetTarget(o).Once()
	o.EAdapters().Add(mockAdapter)
	mock.AssertExpectationsForObjects(t, mockAdapter)

	// first value
	mockValue1 := NewMockEFactory(t)
	mockValue1.EXPECT().EInverseAdd(o, EFACTORY__EPACKAGE, nil).Return(nil).Once()
	mockAdapter.EXPECT().NotifyChanged(mock.Anything).Once()
	o.SetEFactoryInstance(mockValue1)
	mock.AssertExpectationsForObjects(t, mockAdapter, mockValue1)

	// second value
	mockValue2 := NewMockEFactory(t)
	mockValue1.EXPECT().EInverseRemove(o, EFACTORY__EPACKAGE, nil).Return(nil).Once()
	mockValue2.EXPECT().EInverseAdd(o, EFACTORY__EPACKAGE, nil).Return(nil).Once()
	mockAdapter.EXPECT().NotifyChanged(mock.Anything).Once()
	o.SetEFactoryInstance(mockValue2)
	mock.AssertExpectationsForObjects(t, mockAdapter, mockValue1, mockValue2)
}

func TestEPackageEFactoryInstanceBasicSet(t *testing.T) {
	o := newEPackageImpl()

	// add listener
	mockAdapter := NewMockEAdapter(t)
	mockAdapter.EXPECT().SetTarget(o).Once()
	o.EAdapters().Add(mockAdapter)
	mock.AssertExpectationsForObjects(t, mockAdapter)

	mockValue := NewMockEFactory(t)
	mockNotifications := NewMockENotificationChain(t)
	mockNotifications.EXPECT().Add(mock.MatchedBy(func(notification ENotification) bool {
		return notification.GetEventType() == SET && notification.GetFeatureID() == EPACKAGE__EFACTORY_INSTANCE
	})).Return(true).Once()
	o.basicSetEFactoryInstance(mockValue, mockNotifications)
	mock.AssertExpectationsForObjects(t, mockAdapter, mockValue, mockNotifications)
}

func TestEPackageESubPackagesGet(t *testing.T) {
	o := newEPackageImpl()
	assert.NotNil(t, o.GetESubPackages())
	assert.Panics(t, func() { _ = o.GetESubPackages().Get(0).(EPackage) })
}

func TestEPackageESuperPackageGet(t *testing.T) {
	// default
	o := newEPackageImpl()
	assert.Nil(t, o.GetESuperPackage())

	// set a mock container
	v := NewMockEPackage(t)
	o.ESetInternalContainer(v, EPACKAGE__ESUPER_PACKAGE)

	// no proxy
	v.EXPECT().EIsProxy().Return(false).Once()
	assert.Equal(t, v, o.GetESuperPackage())
}

func TestEPackageNsPrefixGet(t *testing.T) {
	o := newEPackageImpl()
	// get default value
	assert.Equal(t, string(""), o.GetNsPrefix())
	// get initialized value
	v := string("Test String")
	o.nsPrefix = v
	assert.Equal(t, v, o.GetNsPrefix())
}

func TestEPackageNsPrefixSet(t *testing.T) {
	o := newEPackageImpl()
	v := string("Test String")
	mockAdapter := NewMockEAdapter(t)
	mockAdapter.EXPECT().SetTarget(o).Once()
	mockAdapter.EXPECT().NotifyChanged(mock.Anything).Once()
	o.EAdapters().Add(mockAdapter)
	o.SetNsPrefix(v)
	mockAdapter.AssertExpectations(t)
}

func TestEPackageNsURIGet(t *testing.T) {
	o := newEPackageImpl()
	// get default value
	assert.Equal(t, string(""), o.GetNsURI())
	// get initialized value
	v := string("Test String")
	o.nsURI = v
	assert.Equal(t, v, o.GetNsURI())
}

func TestEPackageNsURISet(t *testing.T) {
	o := newEPackageImpl()
	v := string("Test String")
	mockAdapter := NewMockEAdapter(t)
	mockAdapter.EXPECT().SetTarget(o).Once()
	mockAdapter.EXPECT().NotifyChanged(mock.Anything).Once()
	o.EAdapters().Add(mockAdapter)
	o.SetNsURI(v)
	mockAdapter.AssertExpectations(t)
}

func TestEPackageGetEClassifierOperation(t *testing.T) {
	o := newEPackageImpl()
	assert.Panics(t, func() { o.GetEClassifier("") })
}

func TestEPackageEGetFromID(t *testing.T) {
	o := newEPackageImpl()
	assert.Panics(t, func() { o.EGetFromID(-1, true) })
	assert.Equal(t, o.GetEClassifiers(), o.EGetFromID(EPACKAGE__ECLASSIFIERS, true))
	assert.Equal(t, o.GetEClassifiers().(EObjectList).GetUnResolvedList(), o.EGetFromID(EPACKAGE__ECLASSIFIERS, false))
	assert.Equal(t, o.GetEFactoryInstance(), o.EGetFromID(EPACKAGE__EFACTORY_INSTANCE, true))
	assert.Equal(t, o.GetESubPackages(), o.EGetFromID(EPACKAGE__ESUB_PACKAGES, true))
	assert.Equal(t, o.GetESubPackages().(EObjectList).GetUnResolvedList(), o.EGetFromID(EPACKAGE__ESUB_PACKAGES, false))
	assert.Equal(t, o.GetESuperPackage(), o.EGetFromID(EPACKAGE__ESUPER_PACKAGE, true))
	assert.Equal(t, o.GetNsPrefix(), o.EGetFromID(EPACKAGE__NS_PREFIX, true))
	assert.Equal(t, o.GetNsURI(), o.EGetFromID(EPACKAGE__NS_URI, true))
}

func TestEPackageESetFromID(t *testing.T) {
	o := newEPackageImpl()
	assert.Panics(t, func() { o.ESetFromID(-1, nil) })
	{
		// list with a value
		mockValue := NewMockEClassifier(t)
		l := NewImmutableEList([]any{mockValue})
		mockValue.EXPECT().EInverseAdd(o, ECLASSIFIER__EPACKAGE, mock.Anything).Return(nil).Once()

		// set list with new contents
		o.ESetFromID(EPACKAGE__ECLASSIFIERS, l)
		// checks
		assert.Equal(t, 1, o.GetEClassifiers().Size())
		assert.Equal(t, mockValue, o.GetEClassifiers().Get(0))
		mock.AssertExpectationsForObjects(t, mockValue)
	}
	{
		mockValue := NewMockEFactory(t)
		mockValue.EXPECT().EInverseAdd(o, EFACTORY__EPACKAGE, nil).Return(nil).Once()
		o.ESetFromID(EPACKAGE__EFACTORY_INSTANCE, mockValue)
		assert.Equal(t, mockValue, o.EGetFromID(EPACKAGE__EFACTORY_INSTANCE, false))
		mock.AssertExpectationsForObjects(t, mockValue)
	}
	{
		// list with a value
		mockValue := NewMockEPackage(t)
		l := NewImmutableEList([]any{mockValue})
		mockValue.EXPECT().EInverseAdd(o, EPACKAGE__ESUPER_PACKAGE, mock.Anything).Return(nil).Once()

		// set list with new contents
		o.ESetFromID(EPACKAGE__ESUB_PACKAGES, l)
		// checks
		assert.Equal(t, 1, o.GetESubPackages().Size())
		assert.Equal(t, mockValue, o.GetESubPackages().Get(0))
		mock.AssertExpectationsForObjects(t, mockValue)
	}
	{
		v := string("Test String")
		o.ESetFromID(EPACKAGE__NS_PREFIX, v)
		assert.Equal(t, v, o.EGetFromID(EPACKAGE__NS_PREFIX, false))
	}
	{
		v := string("Test String")
		o.ESetFromID(EPACKAGE__NS_URI, v)
		assert.Equal(t, v, o.EGetFromID(EPACKAGE__NS_URI, false))
	}

}

func TestEPackageEIsSetFromID(t *testing.T) {
	o := newEPackageImpl()
	assert.Panics(t, func() { o.EIsSetFromID(-1) })
	assert.False(t, o.EIsSetFromID(EPACKAGE__ECLASSIFIERS))
	assert.False(t, o.EIsSetFromID(EPACKAGE__EFACTORY_INSTANCE))
	assert.False(t, o.EIsSetFromID(EPACKAGE__ESUB_PACKAGES))
	assert.False(t, o.EIsSetFromID(EPACKAGE__ESUPER_PACKAGE))
	assert.False(t, o.EIsSetFromID(EPACKAGE__NS_PREFIX))
	assert.False(t, o.EIsSetFromID(EPACKAGE__NS_URI))
}

func TestEPackageEUnsetFromID(t *testing.T) {
	o := newEPackageImpl()
	assert.Panics(t, func() { o.EUnsetFromID(-1) })
	{
		o.EUnsetFromID(EPACKAGE__ECLASSIFIERS)
		v := o.EGetFromID(EPACKAGE__ECLASSIFIERS, false)
		assert.NotNil(t, v)
		l := v.(EList)
		assert.True(t, l.Empty())
	}
	{
		o.EUnsetFromID(EPACKAGE__EFACTORY_INSTANCE)
		assert.Nil(t, o.EGetFromID(EPACKAGE__EFACTORY_INSTANCE, false))
	}
	{
		o.EUnsetFromID(EPACKAGE__ESUB_PACKAGES)
		v := o.EGetFromID(EPACKAGE__ESUB_PACKAGES, false)
		assert.NotNil(t, v)
		l := v.(EList)
		assert.True(t, l.Empty())
	}
	{
		o.EUnsetFromID(EPACKAGE__NS_PREFIX)
		v := o.EGetFromID(EPACKAGE__NS_PREFIX, false)
		assert.Equal(t, string(""), v)
	}
	{
		o.EUnsetFromID(EPACKAGE__NS_URI)
		v := o.EGetFromID(EPACKAGE__NS_URI, false)
		assert.Equal(t, string(""), v)
	}
}

func TestEPackageEInvokeFromID(t *testing.T) {
	o := newEPackageImpl()
	assert.Panics(t, func() { o.EInvokeFromID(-1, nil) })
	assert.Panics(t, func() { o.EInvokeFromID(EPACKAGE__GET_ECLASSIFIER_ESTRING, nil) })
}

func TestEPackageEBasicInverseAdd(t *testing.T) {
	o := newEPackageImpl()
	{
		mockObject := NewMockEObject(t)
		mockNotifications := NewMockENotificationChain(t)
		assert.Equal(t, mockNotifications, o.EBasicInverseAdd(mockObject, -1, mockNotifications))
	}
	{
		mockObject := NewMockEClassifier(t)
		o.EBasicInverseAdd(mockObject, EPACKAGE__ECLASSIFIERS, nil)
		l := o.GetEClassifiers()
		assert.True(t, l.Contains(mockObject))
		mock.AssertExpectationsForObjects(t, mockObject)
	}
	{
		mockObject := NewMockEFactory(t)
		o.EBasicInverseAdd(mockObject, EPACKAGE__EFACTORY_INSTANCE, nil)
		assert.Equal(t, mockObject, o.GetEFactoryInstance())
		mock.AssertExpectationsForObjects(t, mockObject)

		mockObject2 := NewMockEFactory(t)
		mockObject.EXPECT().EInverseRemove(o, EOPPOSITE_FEATURE_BASE-EPACKAGE__EFACTORY_INSTANCE, nil).Return(nil).Once()
		o.EBasicInverseAdd(mockObject2, EPACKAGE__EFACTORY_INSTANCE, nil)
		assert.Equal(t, mockObject2, o.GetEFactoryInstance())
		mock.AssertExpectationsForObjects(t, mockObject2)
	}
	{
		mockObject := NewMockEPackage(t)
		o.EBasicInverseAdd(mockObject, EPACKAGE__ESUB_PACKAGES, nil)
		l := o.GetESubPackages()
		assert.True(t, l.Contains(mockObject))
		mock.AssertExpectationsForObjects(t, mockObject)
	}
	{
		mockObject := NewMockEPackage(t)
		mockObject.EXPECT().EResource().Return(nil).Once()
		mockObject.EXPECT().EIsProxy().Return(false).Once()
		o.EBasicInverseAdd(mockObject, EPACKAGE__ESUPER_PACKAGE, nil)
		assert.Equal(t, mockObject, o.GetESuperPackage())
		mock.AssertExpectationsForObjects(t, mockObject)

		mockOther := NewMockEPackage(t)
		mockOther.EXPECT().EResource().Return(nil).Once()
		mockOther.EXPECT().EIsProxy().Return(false).Once()
		mockObject.EXPECT().EResource().Return(nil).Once()
		mockObject.EXPECT().EInverseRemove(o, EPACKAGE__ESUB_PACKAGES, nil).Return(nil).Once()
		o.EBasicInverseAdd(mockOther, EPACKAGE__ESUPER_PACKAGE, nil)
		assert.Equal(t, mockOther, o.GetESuperPackage())
		mock.AssertExpectationsForObjects(t, mockObject, mockOther)
	}

}

func TestEPackageEBasicInverseRemove(t *testing.T) {
	o := newEPackageImpl()
	{
		mockObject := NewMockEObject(t)
		mockNotifications := NewMockENotificationChain(t)
		assert.Equal(t, mockNotifications, o.EBasicInverseRemove(mockObject, -1, mockNotifications))
	}
	{
		// initialize list with a mock object
		mockObject := NewMockEClassifier(t)
		mockObject.EXPECT().EInverseAdd(o, ECLASSIFIER__EPACKAGE, mock.Anything).Return(nil).Once()

		l := o.GetEClassifiers()
		l.Add(mockObject)

		// basic inverse remove
		o.EBasicInverseRemove(mockObject, EPACKAGE__ECLASSIFIERS, nil)

		// check it was removed
		assert.False(t, l.Contains(mockObject))
		mock.AssertExpectationsForObjects(t, mockObject)
	}
	{
		mockObject := NewMockEFactory(t)
		o.EBasicInverseRemove(mockObject, EPACKAGE__EFACTORY_INSTANCE, nil)
		mock.AssertExpectationsForObjects(t, mockObject)
	}
	{
		// initialize list with a mock object
		mockObject := NewMockEPackage(t)
		mockObject.EXPECT().EInverseAdd(o, EPACKAGE__ESUPER_PACKAGE, mock.Anything).Return(nil).Once()

		l := o.GetESubPackages()
		l.Add(mockObject)

		// basic inverse remove
		o.EBasicInverseRemove(mockObject, EPACKAGE__ESUB_PACKAGES, nil)

		// check it was removed
		assert.False(t, l.Contains(mockObject))
		mock.AssertExpectationsForObjects(t, mockObject)
	}
	{
		mockObject := NewMockEPackage(t)
		o.EBasicInverseRemove(mockObject, EPACKAGE__ESUPER_PACKAGE, nil)
		mock.AssertExpectationsForObjects(t, mockObject)
	}

}
