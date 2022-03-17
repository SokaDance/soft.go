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

func discardEAnnotation() {
	_ = assert.Equal
	_ = mock.Anything
	_ = testing.Coverage
}

func TestEAnnotationAsEAnnotation(t *testing.T) {
	o := newEAnnotationImpl()
	assert.Equal(t, o, o.asEAnnotation())
}

func TestEAnnotationStaticClass(t *testing.T) {
	o := newEAnnotationImpl()
	assert.Equal(t, GetPackage().GetEAnnotationClass(), o.EStaticClass())
}

func TestEAnnotationFeatureCount(t *testing.T) {
	o := newEAnnotationImpl()
	assert.Equal(t, EANNOTATION_FEATURE_COUNT, o.EStaticFeatureCount())
}

func TestEAnnotationContentsGet(t *testing.T) {
	o := newEAnnotationImpl()
	assert.NotNil(t, o.GetContents())
}

func TestEAnnotationDetailsGet(t *testing.T) {
	o := newEAnnotationImpl()
	assert.NotNil(t, o.GetDetails())
}

func TestEAnnotationEModelElementGet(t *testing.T) {
	// default
	o := newEAnnotationImpl()
	assert.Nil(t, o.GetEModelElement())

	// set a mock container
	v := &MockEModelElement{}
	o.ESetInternalContainer(v, EANNOTATION__EMODEL_ELEMENT)

	// no proxy
	v.On("EIsProxy").Return(false)
	assert.Equal(t, v, o.GetEModelElement())
}

func TestEAnnotationEModelElementSet(t *testing.T) {
	// object
	o := newEAnnotationImpl()

	// add listener
	mockAdapter := new(MockEAdapter)
	mockAdapter.On("SetTarget", o).Once()
	o.EAdapters().Add(mockAdapter)
	mock.AssertExpectationsForObjects(t, mockAdapter)

	// set with the mock value
	mockValue := &MockEModelElement{}
	mockResource := new(MockEResource)
	mockValue.On("EInverseAdd", o, EMODEL_ELEMENT__EANNOTATIONS, nil).Return(nil).Once()
	mockValue.On("EResource").Return(mockResource).Once()
	mockResource.On("Attached", o).Once()
	mockAdapter.On("NotifyChanged", mock.Anything).Once()
	o.SetEModelElement(mockValue)
	mock.AssertExpectationsForObjects(t, mockAdapter, mockValue, mockResource)

	// set with the same mock value
	mockAdapter.On("NotifyChanged", mock.Anything).Once()
	o.SetEModelElement(mockValue)
	mock.AssertExpectationsForObjects(t, mockAdapter, mockValue, mockResource)

	// another value - in a different resource
	mockValue2 := &MockEModelElement{}
	mockResource2 := new(MockEResource)
	mockValue.On("EInverseRemove", o, EMODEL_ELEMENT__EANNOTATIONS, nil).Return(nil).Once()
	mockValue.On("EResource").Return(mockResource).Once()
	mockValue2.On("EInverseAdd", o, EMODEL_ELEMENT__EANNOTATIONS, nil).Return(nil).Once()
	mockValue2.On("EResource").Return(mockResource2).Once()
	mockResource.On("Detached", o).Once()
	mockResource2.On("Attached", o).Once()
	mockAdapter.On("NotifyChanged", mock.Anything).Once()
	o.SetEModelElement(mockValue2)
	mock.AssertExpectationsForObjects(t, mockAdapter, mockValue, mockResource, mockValue2, mockResource2)
}

func TestEAnnotationEModelElementBasicSet(t *testing.T) {
	o := newEAnnotationImpl()

	// add listener
	mockAdapter := new(MockEAdapter)
	mockAdapter.On("SetTarget", o).Once()
	o.EAdapters().Add(mockAdapter)
	mock.AssertExpectationsForObjects(t, mockAdapter)

	mockValue := &MockEModelElement{}
	mockNotifications := new(MockENotificationChain)
	mockValue.On("EResource").Return(nil).Once()
	mockNotifications.On("Add", mock.MatchedBy(func(notification ENotification) bool {
		return notification.GetEventType() == SET && notification.GetFeatureID() == EANNOTATION__EMODEL_ELEMENT
	})).Return(true).Once()
	o.basicSetEModelElement(mockValue, mockNotifications)
	mock.AssertExpectationsForObjects(t, mockAdapter, mockValue, mockNotifications)
}

func TestEAnnotationReferencesGet(t *testing.T) {
	o := newEAnnotationImpl()
	assert.NotNil(t, o.GetReferences())
}

func TestEAnnotationSourceGet(t *testing.T) {
	o := newEAnnotationImpl()
	// get default value
	assert.Equal(t, string(""), o.GetSource())
	// get initialized value
	v := "Test String"
	o.source = v
	assert.Equal(t, v, o.GetSource())
}

func TestEAnnotationSourceSet(t *testing.T) {
	o := newEAnnotationImpl()
	v := "Test String"
	mockAdapter := new(MockEAdapter)
	mockAdapter.On("SetTarget", o).Once()
	mockAdapter.On("NotifyChanged", mock.Anything).Once()
	o.EAdapters().Add(mockAdapter)
	o.SetSource(v)
	mockAdapter.AssertExpectations(t)
}

func TestEAnnotationEGetFromID(t *testing.T) {
	o := newEAnnotationImpl()
	assert.Panics(t, func() { o.EGetFromID(-1, true) })
	assert.Equal(t, o.GetContents(), FromAnyList[EObject](o.EGetFromID(EANNOTATION__CONTENTS, true)))
	assert.Equal(t, o.GetDetails(), FromAnyMap[string, string](o.EGetFromID(EANNOTATION__DETAILS, true)))
	{
		// set a mock container
		v := &MockEModelElement{}
		v.On("EIsProxy").Return(false)
		o.ESetInternalContainer(v, EANNOTATION__EMODEL_ELEMENT)
		assert.Equal(t, v, o.EGetFromID(EANNOTATION__EMODEL_ELEMENT, false))
	}
	assert.Equal(t, o.GetReferences(), FromAnyList[EObject](o.EGetFromID(EANNOTATION__REFERENCES, true)))
	assert.Equal(t, o.GetReferences().(EObjectList[EObject]).GetUnResolvedList(), FromAnyList[EObject](o.EGetFromID(EANNOTATION__REFERENCES, false)))
}

func TestEAnnotationESetFromID(t *testing.T) {
	o := newEAnnotationImpl()
	assert.Panics(t, func() { o.ESetFromID(-1, nil) })
	{
		// list with a value
		mockList := &MockEList[EObject]{}
		mockValue := &MockEObjectInternal{}
		mockIterator := &MockEIterator[EObject]{}
		mockList.On("Iterator").Return(mockIterator).Once()
		mockIterator.On("HasNext").Return(true).Once()
		mockIterator.On("Next").Return(mockValue).Once()
		mockIterator.On("HasNext").Return(false).Once()
		mockValue.On("EInverseAdd", o, EOPPOSITE_FEATURE_BASE-EANNOTATION__CONTENTS, mock.Anything).Return(nil).Once()

		// set list with new contents
		o.ESetFromID(EANNOTATION__CONTENTS, ToAnyList[EObject](mockList))
		// checks
		assert.Equal(t, 1, o.GetContents().Size())
		assert.Equal(t, mockValue, o.GetContents().Get(0))
		mock.AssertExpectationsForObjects(t, mockList, mockIterator, mockValue)
	}
	{
		// map with a value
		mockMap := &MockEMap[string, string]{}
		mockEntry := &MockEMapEntry{}
		mockIterator := &MockEIterator[any]{}
		mockKey := "Test String"
		mockValue := "Test String"
		mockMap.On("Iterator").Return(mockIterator).Once()
		mockIterator.On("HasNext").Return(true).Once()
		mockIterator.On("Next").Return(mockEntry).Once()
		mockIterator.On("HasNext").Return(false).Once()
		mockEntry.On("GetKey").Return(mockKey).Once()
		mockEntry.On("GetValue").Return(mockValue).Once()

		// set list with new contents
		o.ESetFromID(EANNOTATION__DETAILS, ToAnyMap[string, string](mockMap))
		assert.Equal(t, map[any]any{mockKey: mockValue}, o.GetDetails().ToMap())
		mock.AssertExpectationsForObjects(t, mockMap, mockEntry)
	}
	{
		mockValue := &MockEModelElement{}
		mockValue.On("EIsProxy").Return(false).Once()
		mockValue.On("EResource").Return(nil).Once()
		mockValue.On("EInverseAdd", o, EMODEL_ELEMENT__EANNOTATIONS, nil).Return(nil).Once()
		o.ESetFromID(EANNOTATION__EMODEL_ELEMENT, mockValue)
		assert.Equal(t, mockValue, o.EGetFromID(EANNOTATION__EMODEL_ELEMENT, false))
		mock.AssertExpectationsForObjects(t, mockValue)
	}
	{
		// list with a value
		mockList := &MockEList[EObject]{}
		mockValue := &MockEObjectInternal{}
		mockIterator := &MockEIterator[EObject]{}
		mockList.On("Iterator").Return(mockIterator).Once()
		mockIterator.On("HasNext").Return(true).Once()
		mockIterator.On("Next").Return(mockValue).Once()
		mockIterator.On("HasNext").Return(false).Once()
		mockValue.On("EIsProxy").Return(false).Once()

		// set list with new contents
		o.ESetFromID(EANNOTATION__REFERENCES, ToAnyList[EObject](mockList))
		// checks
		assert.Equal(t, 1, o.GetReferences().Size())
		assert.Equal(t, mockValue, o.GetReferences().Get(0))
		mock.AssertExpectationsForObjects(t, mockList, mockIterator, mockValue)
	}
	{
		v := "Test String"
		o.ESetFromID(EANNOTATION__SOURCE, v)
		assert.Equal(t, v, o.EGetFromID(EANNOTATION__SOURCE, false))
	}

}

func TestEAnnotationEIsSetFromID(t *testing.T) {
	o := newEAnnotationImpl()
	assert.Panics(t, func() { o.EIsSetFromID(-1) })
	assert.False(t, o.EIsSetFromID(EANNOTATION__CONTENTS))
	assert.False(t, o.EIsSetFromID(EANNOTATION__DETAILS))
	assert.False(t, o.EIsSetFromID(EANNOTATION__EMODEL_ELEMENT))
	assert.False(t, o.EIsSetFromID(EANNOTATION__REFERENCES))
	assert.False(t, o.EIsSetFromID(EANNOTATION__SOURCE))
}

func TestEAnnotationEUnsetFromID(t *testing.T) {
	o := newEAnnotationImpl()
	assert.Panics(t, func() { o.EUnsetFromID(-1) })
	{
		o.EUnsetFromID(EANNOTATION__CONTENTS)
		v := o.EGetFromID(EANNOTATION__CONTENTS, false)
		assert.NotNil(t, v)
		l := v.(EList[any])
		assert.True(t, l.Empty())
	}
	{
		o.EUnsetFromID(EANNOTATION__DETAILS)
		v := o.EGetFromID(EANNOTATION__DETAILS, false)
		assert.NotNil(t, v)
		l := v.(EList[any])
		assert.True(t, l.Empty())
	}
	{
		o.EUnsetFromID(EANNOTATION__EMODEL_ELEMENT)
		assert.Nil(t, o.EGetFromID(EANNOTATION__EMODEL_ELEMENT, false))
	}
	{
		o.EUnsetFromID(EANNOTATION__REFERENCES)
		v := o.EGetFromID(EANNOTATION__REFERENCES, false)
		assert.NotNil(t, v)
		l := v.(EList[any])
		assert.True(t, l.Empty())
	}
	{
		o.EUnsetFromID(EANNOTATION__SOURCE)
		v := o.EGetFromID(EANNOTATION__SOURCE, false)
		assert.Equal(t, string(""), v)
	}
}

func TestEAnnotationEBasicInverseAdd(t *testing.T) {
	o := newEAnnotationImpl()
	{
		mockObject := new(MockEObject)
		mockNotifications := new(MockENotificationChain)
		assert.Equal(t, mockNotifications, o.EBasicInverseAdd(mockObject, -1, mockNotifications))
	}
	{
		mockObject := &MockEModelElement{}
		mockObject.On("EResource").Return(nil).Once()
		mockObject.On("EIsProxy").Return(false).Once()
		o.EBasicInverseAdd(mockObject, EANNOTATION__EMODEL_ELEMENT, nil)
		assert.Equal(t, mockObject, o.GetEModelElement())
		mock.AssertExpectationsForObjects(t, mockObject)

		mockOther := &MockEModelElement{}
		mockOther.On("EResource").Return(nil).Once()
		mockOther.On("EIsProxy").Return(false).Once()
		mockObject.On("EResource").Return(nil).Once()
		mockObject.On("EInverseRemove", o, EMODEL_ELEMENT__EANNOTATIONS, nil).Return(nil).Once()
		o.EBasicInverseAdd(mockOther, EANNOTATION__EMODEL_ELEMENT, nil)
		assert.Equal(t, mockOther, o.GetEModelElement())
		mock.AssertExpectationsForObjects(t, mockObject, mockOther)
	}

}

func TestEAnnotationEBasicInverseRemove(t *testing.T) {
	o := newEAnnotationImpl()
	{
		mockObject := new(MockEObject)
		mockNotifications := new(MockENotificationChain)
		assert.Equal(t, mockNotifications, o.EBasicInverseRemove(mockObject, -1, mockNotifications))
	}
	{
		// initialize list with a mock object
		mockObject := &MockEObjectInternal{}
		mockObject.On("EInverseAdd", o, EOPPOSITE_FEATURE_BASE-EANNOTATION__CONTENTS, mock.Anything).Return(nil).Once()

		l := o.GetContents()
		l.Add(mockObject)

		// basic inverse remove
		o.EBasicInverseRemove(mockObject, EANNOTATION__CONTENTS, nil)

		// check it was removed
		assert.False(t, l.Contains(mockObject))
		mock.AssertExpectationsForObjects(t, mockObject)
	}
	{
		mockObject := &MockEStringToStringMapEntry{}
		o.EBasicInverseRemove(mockObject, EANNOTATION__DETAILS, nil)
		mock.AssertExpectationsForObjects(t, mockObject)
		assert.True(t, o.GetDetails().Empty())
	}
	{
		mockObject := &MockEModelElement{}
		o.EBasicInverseRemove(mockObject, EANNOTATION__EMODEL_ELEMENT, nil)
		mock.AssertExpectationsForObjects(t, mockObject)
	}

}
