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

func discardEClass() {
	_ = assert.Equal
	_ = mock.Anything
	_ = testing.Coverage
}

func TestEClassAsEClass(t *testing.T) {
	o := newEClassImpl()
	assert.Equal(t, o, o.asEClass())
}

func TestEClassStaticClass(t *testing.T) {
	o := newEClassImpl()
	assert.Equal(t, GetPackage().GetEClass(), o.EStaticClass())
}

func TestEClassFeatureCount(t *testing.T) {
	o := newEClassImpl()
	assert.Equal(t, ECLASS_FEATURE_COUNT, o.EStaticFeatureCount())
}

func TestEClassEAllAttributesGet(t *testing.T) {
	o := newEClassImpl()
	assert.NotNil(t, o.GetEAllAttributes())
	assert.Panics(t, func() { _ = o.GetEAllAttributes().Get(0).(EAttribute) })
}

func TestEClassEAllContainmentsGet(t *testing.T) {
	o := newEClassImpl()
	assert.NotNil(t, o.GetEAllContainments())
	assert.Panics(t, func() { _ = o.GetEAllContainments().Get(0).(EReference) })
}

func TestEClassEAllCrossReferencesGet(t *testing.T) {
	o := newEClassImpl()
	assert.NotNil(t, o.GetEAllCrossReferences())
	assert.Panics(t, func() { _ = o.GetEAllCrossReferences().Get(0).(EReference) })
}

func TestEClassEAllOperationsGet(t *testing.T) {
	o := newEClassImpl()
	assert.NotNil(t, o.GetEAllOperations())
	assert.Panics(t, func() { _ = o.GetEAllOperations().Get(0).(EOperation) })
}

func TestEClassEAllReferencesGet(t *testing.T) {
	o := newEClassImpl()
	assert.NotNil(t, o.GetEAllReferences())
	assert.Panics(t, func() { _ = o.GetEAllReferences().Get(0).(EReference) })
}

func TestEClassEAllStructuralFeaturesGet(t *testing.T) {
	o := newEClassImpl()
	assert.NotNil(t, o.GetEAllStructuralFeatures())
	assert.Panics(t, func() { _ = o.GetEAllStructuralFeatures().Get(0).(EStructuralFeature) })
}

func TestEClassEAllSuperTypesGet(t *testing.T) {
	o := newEClassImpl()
	assert.NotNil(t, o.GetEAllSuperTypes())
	assert.Panics(t, func() { _ = o.GetEAllSuperTypes().Get(0).(EClass) })
}

func TestEClassEAttributesGet(t *testing.T) {
	o := newEClassImpl()
	assert.NotNil(t, o.GetEAttributes())
	assert.Panics(t, func() { _ = o.GetEAttributes().Get(0).(EAttribute) })
}

func TestEClassEContainmentFeaturesGet(t *testing.T) {
	o := newEClassImpl()
	assert.NotNil(t, o.GetEContainmentFeatures())
	assert.Panics(t, func() { _ = o.GetEContainmentFeatures().Get(0).(EStructuralFeature) })
}

func TestEClassECrossReferenceFeaturesGet(t *testing.T) {
	o := newEClassImpl()
	assert.NotNil(t, o.GetECrossReferenceFeatures())
	assert.Panics(t, func() { _ = o.GetECrossReferenceFeatures().Get(0).(EStructuralFeature) })
}

func TestEClassEIDAttributeGet(t *testing.T) {
	o := newEClassImpl()
	assert.Panics(t, func() { o.GetEIDAttribute() })
}

func TestEClassEOperationsGet(t *testing.T) {
	o := newEClassImpl()
	assert.NotNil(t, o.GetEOperations())
	assert.Panics(t, func() { _ = o.GetEOperations().Get(0).(EOperation) })
}

func TestEClassEReferencesGet(t *testing.T) {
	o := newEClassImpl()
	assert.NotNil(t, o.GetEReferences())
	assert.Panics(t, func() { _ = o.GetEReferences().Get(0).(EReference) })
}

func TestEClassEStructuralFeaturesGet(t *testing.T) {
	o := newEClassImpl()
	assert.NotNil(t, o.GetEStructuralFeatures())
	assert.Panics(t, func() { _ = o.GetEStructuralFeatures().Get(0).(EStructuralFeature) })
}

func TestEClassESuperTypesGet(t *testing.T) {
	o := newEClassImpl()
	assert.NotNil(t, o.GetESuperTypes())
	assert.Panics(t, func() { _ = o.GetESuperTypes().Get(0).(EClass) })
}

func TestEClassAbstractGet(t *testing.T) {
	o := newEClassImpl()
	// get default value
	assert.Equal(t, bool(false), o.IsAbstract())
	// get initialized value
	v := bool(true)
	o.isAbstract = v
	assert.Equal(t, v, o.IsAbstract())
}

func TestEClassAbstractSet(t *testing.T) {
	o := newEClassImpl()
	v := bool(true)
	mockAdapter := NewMockEAdapter(t)
	mockAdapter.EXPECT().SetTarget(o).Once()
	mockAdapter.EXPECT().NotifyChanged(mock.Anything).Once()
	o.EAdapters().Add(mockAdapter)
	o.SetAbstract(v)
	mockAdapter.AssertExpectations(t)
}

func TestEClassInterfaceGet(t *testing.T) {
	o := newEClassImpl()
	// get default value
	assert.Equal(t, bool(false), o.IsInterface())
	// get initialized value
	v := bool(true)
	o.isInterface = v
	assert.Equal(t, v, o.IsInterface())
}

func TestEClassInterfaceSet(t *testing.T) {
	o := newEClassImpl()
	v := bool(true)
	mockAdapter := NewMockEAdapter(t)
	mockAdapter.EXPECT().SetTarget(o).Once()
	mockAdapter.EXPECT().NotifyChanged(mock.Anything).Once()
	o.EAdapters().Add(mockAdapter)
	o.SetInterface(v)
	mockAdapter.AssertExpectations(t)
}

func TestEClassGetEOperationOperation(t *testing.T) {
	o := newEClassImpl()
	assert.Panics(t, func() { o.GetEOperation(0) })
}
func TestEClassGetEStructuralFeatureOperation(t *testing.T) {
	o := newEClassImpl()
	assert.Panics(t, func() { o.GetEStructuralFeature(0) })
}
func TestEClassGetEStructuralFeatureFromNameOperation(t *testing.T) {
	o := newEClassImpl()
	assert.Panics(t, func() { o.GetEStructuralFeatureFromName("") })
}
func TestEClassGetFeatureCountOperation(t *testing.T) {
	o := newEClassImpl()
	assert.Panics(t, func() { o.GetFeatureCount() })
}
func TestEClassGetFeatureIDOperation(t *testing.T) {
	o := newEClassImpl()
	assert.Panics(t, func() { o.GetFeatureID(nil) })
}
func TestEClassGetFeatureTypeOperation(t *testing.T) {
	o := newEClassImpl()
	assert.Panics(t, func() { o.GetFeatureType(nil) })
}
func TestEClassGetOperationCountOperation(t *testing.T) {
	o := newEClassImpl()
	assert.Panics(t, func() { o.GetOperationCount() })
}
func TestEClassGetOperationIDOperation(t *testing.T) {
	o := newEClassImpl()
	assert.Panics(t, func() { o.GetOperationID(nil) })
}
func TestEClassGetOverrideOperation(t *testing.T) {
	o := newEClassImpl()
	assert.Panics(t, func() { o.GetOverride(nil) })
}
func TestEClassIsSuperTypeOfOperation(t *testing.T) {
	o := newEClassImpl()
	assert.Panics(t, func() { o.IsSuperTypeOf(nil) })
}

func TestEClassEGetFromID(t *testing.T) {
	o := newEClassImpl()
	assert.Panics(t, func() { o.EGetFromID(-1, true) })
	assert.Equal(t, o.IsAbstract(), o.EGetFromID(ECLASS__ABSTRACT, true))
	assert.Equal(t, o.GetEAllAttributes(), o.EGetFromID(ECLASS__EALL_ATTRIBUTES, true))
	assert.Equal(t, o.GetEAllAttributes().(EObjectList).GetUnResolvedList(), o.EGetFromID(ECLASS__EALL_ATTRIBUTES, false))
	assert.Equal(t, o.GetEAllContainments(), o.EGetFromID(ECLASS__EALL_CONTAINMENTS, true))
	assert.Equal(t, o.GetEAllContainments().(EObjectList).GetUnResolvedList(), o.EGetFromID(ECLASS__EALL_CONTAINMENTS, false))
	assert.Equal(t, o.GetEAllCrossReferences(), o.EGetFromID(ECLASS__EALL_CROSS_REFERENCES, true))
	assert.Equal(t, o.GetEAllCrossReferences().(EObjectList).GetUnResolvedList(), o.EGetFromID(ECLASS__EALL_CROSS_REFERENCES, false))
	assert.Equal(t, o.GetEAllOperations(), o.EGetFromID(ECLASS__EALL_OPERATIONS, true))
	assert.Equal(t, o.GetEAllOperations().(EObjectList).GetUnResolvedList(), o.EGetFromID(ECLASS__EALL_OPERATIONS, false))
	assert.Equal(t, o.GetEAllReferences(), o.EGetFromID(ECLASS__EALL_REFERENCES, true))
	assert.Equal(t, o.GetEAllReferences().(EObjectList).GetUnResolvedList(), o.EGetFromID(ECLASS__EALL_REFERENCES, false))
	assert.Equal(t, o.GetEAllStructuralFeatures(), o.EGetFromID(ECLASS__EALL_STRUCTURAL_FEATURES, true))
	assert.Equal(t, o.GetEAllStructuralFeatures().(EObjectList).GetUnResolvedList(), o.EGetFromID(ECLASS__EALL_STRUCTURAL_FEATURES, false))
	assert.Equal(t, o.GetEAllSuperTypes(), o.EGetFromID(ECLASS__EALL_SUPER_TYPES, true))
	assert.Equal(t, o.GetEAllSuperTypes().(EObjectList).GetUnResolvedList(), o.EGetFromID(ECLASS__EALL_SUPER_TYPES, false))
	assert.Equal(t, o.GetEAttributes(), o.EGetFromID(ECLASS__EATTRIBUTES, true))
	assert.Equal(t, o.GetEAttributes().(EObjectList).GetUnResolvedList(), o.EGetFromID(ECLASS__EATTRIBUTES, false))
	assert.Equal(t, o.GetEContainmentFeatures(), o.EGetFromID(ECLASS__ECONTAINMENT_FEATURES, true))
	assert.Equal(t, o.GetEContainmentFeatures().(EObjectList).GetUnResolvedList(), o.EGetFromID(ECLASS__ECONTAINMENT_FEATURES, false))
	assert.Equal(t, o.GetECrossReferenceFeatures(), o.EGetFromID(ECLASS__ECROSS_REFERENCE_FEATURES, true))
	assert.Equal(t, o.GetECrossReferenceFeatures().(EObjectList).GetUnResolvedList(), o.EGetFromID(ECLASS__ECROSS_REFERENCE_FEATURES, false))
	assert.Panics(t, func() { o.EGetFromID(ECLASS__EID_ATTRIBUTE, true) })
	assert.Panics(t, func() { o.EGetFromID(ECLASS__EID_ATTRIBUTE, false) })
	assert.Equal(t, o.GetEOperations(), o.EGetFromID(ECLASS__EOPERATIONS, true))
	assert.Equal(t, o.GetEOperations().(EObjectList).GetUnResolvedList(), o.EGetFromID(ECLASS__EOPERATIONS, false))
	assert.Equal(t, o.GetEReferences(), o.EGetFromID(ECLASS__EREFERENCES, true))
	assert.Equal(t, o.GetEReferences().(EObjectList).GetUnResolvedList(), o.EGetFromID(ECLASS__EREFERENCES, false))
	assert.Equal(t, o.GetEStructuralFeatures(), o.EGetFromID(ECLASS__ESTRUCTURAL_FEATURES, true))
	assert.Equal(t, o.GetEStructuralFeatures().(EObjectList).GetUnResolvedList(), o.EGetFromID(ECLASS__ESTRUCTURAL_FEATURES, false))
	assert.Equal(t, o.GetESuperTypes(), o.EGetFromID(ECLASS__ESUPER_TYPES, true))
	assert.Equal(t, o.GetESuperTypes().(EObjectList).GetUnResolvedList(), o.EGetFromID(ECLASS__ESUPER_TYPES, false))
	assert.Equal(t, o.IsInterface(), o.EGetFromID(ECLASS__INTERFACE, true))
}

func TestEClassESetFromID(t *testing.T) {
	o := newEClassImpl()
	assert.Panics(t, func() { o.ESetFromID(-1, nil) })
	{
		v := bool(true)
		o.ESetFromID(ECLASS__ABSTRACT, v)
		assert.Equal(t, v, o.EGetFromID(ECLASS__ABSTRACT, false))
	}
	{
		// list with a value
		mockValue := NewMockEOperation(t)
		l := NewImmutableEList([]any{mockValue})
		mockValue.EXPECT().EInverseAdd(o, EOPERATION__ECONTAINING_CLASS, mock.Anything).Return(nil).Once()

		// set list with new contents
		o.ESetFromID(ECLASS__EOPERATIONS, l)
		// checks
		assert.Equal(t, 1, o.GetEOperations().Size())
		assert.Equal(t, mockValue, o.GetEOperations().Get(0))
		mock.AssertExpectationsForObjects(t, mockValue)
	}
	{
		// list with a value
		mockValue := NewMockEStructuralFeature(t)
		l := NewImmutableEList([]any{mockValue})
		mockValue.EXPECT().EInverseAdd(o, ESTRUCTURAL_FEATURE__ECONTAINING_CLASS, mock.Anything).Return(nil).Once()

		// set list with new contents
		o.ESetFromID(ECLASS__ESTRUCTURAL_FEATURES, l)
		// checks
		assert.Equal(t, 1, o.GetEStructuralFeatures().Size())
		assert.Equal(t, mockValue, o.GetEStructuralFeatures().Get(0))
		mock.AssertExpectationsForObjects(t, mockValue)
	}
	{
		// list with a value
		mockValue := NewMockEClass(t)
		l := NewImmutableEList([]any{mockValue})
		mockValue.EXPECT().EIsProxy().Return(false).Once()

		// set list with new contents
		o.ESetFromID(ECLASS__ESUPER_TYPES, l)
		// checks
		assert.Equal(t, 1, o.GetESuperTypes().Size())
		assert.Equal(t, mockValue, o.GetESuperTypes().Get(0))
		mock.AssertExpectationsForObjects(t, mockValue)
	}
	{
		v := bool(true)
		o.ESetFromID(ECLASS__INTERFACE, v)
		assert.Equal(t, v, o.EGetFromID(ECLASS__INTERFACE, false))
	}

}

func TestEClassEIsSetFromID(t *testing.T) {
	o := newEClassImpl()
	assert.Panics(t, func() { o.EIsSetFromID(-1) })
	assert.False(t, o.EIsSetFromID(ECLASS__ABSTRACT))
	assert.False(t, o.EIsSetFromID(ECLASS__EALL_ATTRIBUTES))
	assert.False(t, o.EIsSetFromID(ECLASS__EALL_CONTAINMENTS))
	assert.False(t, o.EIsSetFromID(ECLASS__EALL_CROSS_REFERENCES))
	assert.False(t, o.EIsSetFromID(ECLASS__EALL_OPERATIONS))
	assert.False(t, o.EIsSetFromID(ECLASS__EALL_REFERENCES))
	assert.False(t, o.EIsSetFromID(ECLASS__EALL_STRUCTURAL_FEATURES))
	assert.False(t, o.EIsSetFromID(ECLASS__EALL_SUPER_TYPES))
	assert.False(t, o.EIsSetFromID(ECLASS__EATTRIBUTES))
	assert.False(t, o.EIsSetFromID(ECLASS__ECONTAINMENT_FEATURES))
	assert.False(t, o.EIsSetFromID(ECLASS__ECROSS_REFERENCE_FEATURES))
	assert.False(t, o.EIsSetFromID(ECLASS__EID_ATTRIBUTE))
	assert.False(t, o.EIsSetFromID(ECLASS__EOPERATIONS))
	assert.False(t, o.EIsSetFromID(ECLASS__EREFERENCES))
	assert.False(t, o.EIsSetFromID(ECLASS__ESTRUCTURAL_FEATURES))
	assert.False(t, o.EIsSetFromID(ECLASS__ESUPER_TYPES))
	assert.False(t, o.EIsSetFromID(ECLASS__INTERFACE))
}

func TestEClassEUnsetFromID(t *testing.T) {
	o := newEClassImpl()
	assert.Panics(t, func() { o.EUnsetFromID(-1) })
	{
		o.EUnsetFromID(ECLASS__ABSTRACT)
		v := o.EGetFromID(ECLASS__ABSTRACT, false)
		assert.Equal(t, bool(false), v)
	}
	{
		o.EUnsetFromID(ECLASS__EOPERATIONS)
		v := o.EGetFromID(ECLASS__EOPERATIONS, false)
		assert.NotNil(t, v)
		l := v.(EList)
		assert.True(t, l.Empty())
	}
	{
		o.EUnsetFromID(ECLASS__ESTRUCTURAL_FEATURES)
		v := o.EGetFromID(ECLASS__ESTRUCTURAL_FEATURES, false)
		assert.NotNil(t, v)
		l := v.(EList)
		assert.True(t, l.Empty())
	}
	{
		o.EUnsetFromID(ECLASS__ESUPER_TYPES)
		v := o.EGetFromID(ECLASS__ESUPER_TYPES, false)
		assert.NotNil(t, v)
		l := v.(EList)
		assert.True(t, l.Empty())
	}
	{
		o.EUnsetFromID(ECLASS__INTERFACE)
		v := o.EGetFromID(ECLASS__INTERFACE, false)
		assert.Equal(t, bool(false), v)
	}
}

func TestEClassEInvokeFromID(t *testing.T) {
	o := newEClassImpl()
	assert.Panics(t, func() { o.EInvokeFromID(-1, nil) })
	assert.Panics(t, func() { o.EInvokeFromID(ECLASS__GET_EOPERATION_EINT, nil) })
	assert.Panics(t, func() { o.EInvokeFromID(ECLASS__GET_ESTRUCTURAL_FEATURE_EINT, nil) })
	assert.Panics(t, func() { o.EInvokeFromID(ECLASS__GET_ESTRUCTURAL_FEATURE_ESTRING, nil) })
	assert.Panics(t, func() { o.EInvokeFromID(ECLASS__GET_FEATURE_COUNT, nil) })
	assert.Panics(t, func() { o.EInvokeFromID(ECLASS__GET_FEATURE_ID_ESTRUCTURALFEATURE, nil) })
	assert.Panics(t, func() { o.EInvokeFromID(ECLASS__GET_FEATURE_TYPE_ESTRUCTURALFEATURE, nil) })
	assert.Panics(t, func() { o.EInvokeFromID(ECLASS__GET_OPERATION_COUNT, nil) })
	assert.Panics(t, func() { o.EInvokeFromID(ECLASS__GET_OPERATION_ID_EOPERATION, nil) })
	assert.Panics(t, func() { o.EInvokeFromID(ECLASS__GET_OVERRIDE_EOPERATION, nil) })
	assert.Panics(t, func() { o.EInvokeFromID(ECLASS__IS_SUPER_TYPE_OF_ECLASS, nil) })
}

func TestEClassEBasicInverseAdd(t *testing.T) {
	o := newEClassImpl()
	{
		mockObject := NewMockEObject(t)
		mockNotifications := NewMockENotificationChain(t)
		assert.Equal(t, mockNotifications, o.EBasicInverseAdd(mockObject, -1, mockNotifications))
	}
	{
		mockObject := NewMockEOperation(t)
		o.EBasicInverseAdd(mockObject, ECLASS__EOPERATIONS, nil)
		l := o.GetEOperations()
		assert.True(t, l.Contains(mockObject))
		mock.AssertExpectationsForObjects(t, mockObject)
	}
	{
		mockObject := NewMockEStructuralFeature(t)
		o.EBasicInverseAdd(mockObject, ECLASS__ESTRUCTURAL_FEATURES, nil)
		l := o.GetEStructuralFeatures()
		assert.True(t, l.Contains(mockObject))
		mock.AssertExpectationsForObjects(t, mockObject)
	}

}

func TestEClassEBasicInverseRemove(t *testing.T) {
	o := newEClassImpl()
	{
		mockObject := NewMockEObject(t)
		mockNotifications := NewMockENotificationChain(t)
		assert.Equal(t, mockNotifications, o.EBasicInverseRemove(mockObject, -1, mockNotifications))
	}
	{
		// initialize list with a mock object
		mockObject := NewMockEOperation(t)
		mockObject.EXPECT().EInverseAdd(o, EOPERATION__ECONTAINING_CLASS, mock.Anything).Return(nil).Once()

		l := o.GetEOperations()
		l.Add(mockObject)

		// basic inverse remove
		o.EBasicInverseRemove(mockObject, ECLASS__EOPERATIONS, nil)

		// check it was removed
		assert.False(t, l.Contains(mockObject))
		mock.AssertExpectationsForObjects(t, mockObject)
	}
	{
		// initialize list with a mock object
		mockObject := NewMockEStructuralFeature(t)
		mockObject.EXPECT().EInverseAdd(o, ESTRUCTURAL_FEATURE__ECONTAINING_CLASS, mock.Anything).Return(nil).Once()

		l := o.GetEStructuralFeatures()
		l.Add(mockObject)

		// basic inverse remove
		o.EBasicInverseRemove(mockObject, ECLASS__ESTRUCTURAL_FEATURES, nil)

		// check it was removed
		assert.False(t, l.Contains(mockObject))
		mock.AssertExpectationsForObjects(t, mockObject)
	}

}
