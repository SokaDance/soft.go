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
	"reflect"
)

// EStructuralFeatureImpl is the implementation of the model object 'EStructuralFeature'
type EStructuralFeatureImpl struct {
	ETypedElementExt
	defaultValueLiteral string
	featureID           int
	isChangeable        bool
	isDerived           bool
	isTransient         bool
	isUnsettable        bool
	isVolatile          bool
}

// newEStructuralFeatureImpl is the constructor of a EStructuralFeatureImpl
func newEStructuralFeatureImpl() *EStructuralFeatureImpl {
	e := new(EStructuralFeatureImpl)
	e.SetInterfaces(e)
	e.Initialize()
	return e
}

func (e *EStructuralFeatureImpl) Initialize() {
	e.ETypedElementExt.Initialize()
	e.defaultValueLiteral = ""
	e.featureID = -1
	e.isChangeable = true
	e.isDerived = false
	e.isTransient = false
	e.isUnsettable = false
	e.isVolatile = false

}

func (e *EStructuralFeatureImpl) asEStructuralFeature() EStructuralFeature {
	return e.GetInterfaces().(EStructuralFeature)
}

func (e *EStructuralFeatureImpl) EStaticClass() EClass {
	return GetPackage().GetEStructuralFeature()
}

func (e *EStructuralFeatureImpl) EStaticFeatureCount() int {
	return ESTRUCTURAL_FEATURE_FEATURE_COUNT
}

// GetContainerClass default implementation
func (e *EStructuralFeatureImpl) GetContainerClass() reflect.Type {
	panic("GetContainerClass not implemented")
}

// GetDefaultValue get the value of defaultValue
func (e *EStructuralFeatureImpl) GetDefaultValue() any {
	panic("GetDefaultValue not implemented")
}

// SetDefaultValue set the value of defaultValue
func (e *EStructuralFeatureImpl) SetDefaultValue(newDefaultValue any) {
	panic("SetDefaultValue not implemented")
}

// GetDefaultValueLiteral get the value of defaultValueLiteral
func (e *EStructuralFeatureImpl) GetDefaultValueLiteral() string {
	return e.defaultValueLiteral
}

// SetDefaultValueLiteral set the value of defaultValueLiteral
func (e *EStructuralFeatureImpl) SetDefaultValueLiteral(newDefaultValueLiteral string) {
	oldDefaultValueLiteral := e.defaultValueLiteral
	e.defaultValueLiteral = newDefaultValueLiteral
	if e.ENotificationRequired() {
		e.ENotify(NewNotificationByFeatureID(e.AsEObject(), SET, ESTRUCTURAL_FEATURE__DEFAULT_VALUE_LITERAL, oldDefaultValueLiteral, newDefaultValueLiteral, NO_INDEX))
	}
}

// GetEContainingClass get the value of eContainingClass
func (e *EStructuralFeatureImpl) GetEContainingClass() EClass {
	if e.EContainerFeatureID() == ESTRUCTURAL_FEATURE__ECONTAINING_CLASS {
		return e.EContainer().(EClass)
	}
	return nil
}

// GetFeatureID get the value of featureID
func (e *EStructuralFeatureImpl) GetFeatureID() int {
	return e.featureID
}

// SetFeatureID set the value of featureID
func (e *EStructuralFeatureImpl) SetFeatureID(newFeatureID int) {
	oldFeatureID := e.featureID
	e.featureID = newFeatureID
	if e.ENotificationRequired() {
		e.ENotify(NewNotificationByFeatureID(e.AsEObject(), SET, ESTRUCTURAL_FEATURE__FEATURE_ID, oldFeatureID, newFeatureID, NO_INDEX))
	}
}

// IsChangeable get the value of isChangeable
func (e *EStructuralFeatureImpl) IsChangeable() bool {
	return e.isChangeable
}

// SetChangeable set the value of isChangeable
func (e *EStructuralFeatureImpl) SetChangeable(newIsChangeable bool) {
	oldIsChangeable := e.isChangeable
	e.isChangeable = newIsChangeable
	if e.ENotificationRequired() {
		e.ENotify(NewNotificationByFeatureID(e.AsEObject(), SET, ESTRUCTURAL_FEATURE__CHANGEABLE, oldIsChangeable, newIsChangeable, NO_INDEX))
	}
}

// IsDerived get the value of isDerived
func (e *EStructuralFeatureImpl) IsDerived() bool {
	return e.isDerived
}

// SetDerived set the value of isDerived
func (e *EStructuralFeatureImpl) SetDerived(newIsDerived bool) {
	oldIsDerived := e.isDerived
	e.isDerived = newIsDerived
	if e.ENotificationRequired() {
		e.ENotify(NewNotificationByFeatureID(e.AsEObject(), SET, ESTRUCTURAL_FEATURE__DERIVED, oldIsDerived, newIsDerived, NO_INDEX))
	}
}

// IsTransient get the value of isTransient
func (e *EStructuralFeatureImpl) IsTransient() bool {
	return e.isTransient
}

// SetTransient set the value of isTransient
func (e *EStructuralFeatureImpl) SetTransient(newIsTransient bool) {
	oldIsTransient := e.isTransient
	e.isTransient = newIsTransient
	if e.ENotificationRequired() {
		e.ENotify(NewNotificationByFeatureID(e.AsEObject(), SET, ESTRUCTURAL_FEATURE__TRANSIENT, oldIsTransient, newIsTransient, NO_INDEX))
	}
}

// IsUnsettable get the value of isUnsettable
func (e *EStructuralFeatureImpl) IsUnsettable() bool {
	return e.isUnsettable
}

// SetUnsettable set the value of isUnsettable
func (e *EStructuralFeatureImpl) SetUnsettable(newIsUnsettable bool) {
	oldIsUnsettable := e.isUnsettable
	e.isUnsettable = newIsUnsettable
	if e.ENotificationRequired() {
		e.ENotify(NewNotificationByFeatureID(e.AsEObject(), SET, ESTRUCTURAL_FEATURE__UNSETTABLE, oldIsUnsettable, newIsUnsettable, NO_INDEX))
	}
}

// IsVolatile get the value of isVolatile
func (e *EStructuralFeatureImpl) IsVolatile() bool {
	return e.isVolatile
}

// SetVolatile set the value of isVolatile
func (e *EStructuralFeatureImpl) SetVolatile(newIsVolatile bool) {
	oldIsVolatile := e.isVolatile
	e.isVolatile = newIsVolatile
	if e.ENotificationRequired() {
		e.ENotify(NewNotificationByFeatureID(e.AsEObject(), SET, ESTRUCTURAL_FEATURE__VOLATILE, oldIsVolatile, newIsVolatile, NO_INDEX))
	}
}

func (e *EStructuralFeatureImpl) EGetFromID(featureID int, resolve bool) any {
	switch featureID {
	case ESTRUCTURAL_FEATURE__CHANGEABLE:
		return e.asEStructuralFeature().IsChangeable()
	case ESTRUCTURAL_FEATURE__DEFAULT_VALUE:
		return e.asEStructuralFeature().GetDefaultValue()
	case ESTRUCTURAL_FEATURE__DEFAULT_VALUE_LITERAL:
		return e.asEStructuralFeature().GetDefaultValueLiteral()
	case ESTRUCTURAL_FEATURE__DERIVED:
		return e.asEStructuralFeature().IsDerived()
	case ESTRUCTURAL_FEATURE__ECONTAINING_CLASS:
		return e.asEStructuralFeature().GetEContainingClass()
	case ESTRUCTURAL_FEATURE__FEATURE_ID:
		return e.asEStructuralFeature().GetFeatureID()
	case ESTRUCTURAL_FEATURE__TRANSIENT:
		return e.asEStructuralFeature().IsTransient()
	case ESTRUCTURAL_FEATURE__UNSETTABLE:
		return e.asEStructuralFeature().IsUnsettable()
	case ESTRUCTURAL_FEATURE__VOLATILE:
		return e.asEStructuralFeature().IsVolatile()
	default:
		return e.ETypedElementExt.EGetFromID(featureID, resolve)
	}
}

func (e *EStructuralFeatureImpl) ESetFromID(featureID int, newValue any) {
	switch featureID {
	case ESTRUCTURAL_FEATURE__CHANGEABLE:
		e.asEStructuralFeature().SetChangeable(newValue.(bool))
	case ESTRUCTURAL_FEATURE__DEFAULT_VALUE:
		e.asEStructuralFeature().SetDefaultValue(newValue)
	case ESTRUCTURAL_FEATURE__DEFAULT_VALUE_LITERAL:
		e.asEStructuralFeature().SetDefaultValueLiteral(newValue.(string))
	case ESTRUCTURAL_FEATURE__DERIVED:
		e.asEStructuralFeature().SetDerived(newValue.(bool))
	case ESTRUCTURAL_FEATURE__FEATURE_ID:
		e.asEStructuralFeature().SetFeatureID(newValue.(int))
	case ESTRUCTURAL_FEATURE__TRANSIENT:
		e.asEStructuralFeature().SetTransient(newValue.(bool))
	case ESTRUCTURAL_FEATURE__UNSETTABLE:
		e.asEStructuralFeature().SetUnsettable(newValue.(bool))
	case ESTRUCTURAL_FEATURE__VOLATILE:
		e.asEStructuralFeature().SetVolatile(newValue.(bool))
	default:
		e.ETypedElementExt.ESetFromID(featureID, newValue)
	}
}

func (e *EStructuralFeatureImpl) EUnsetFromID(featureID int) {
	switch featureID {
	case ESTRUCTURAL_FEATURE__CHANGEABLE:
		e.asEStructuralFeature().SetChangeable(true)
	case ESTRUCTURAL_FEATURE__DEFAULT_VALUE:
		e.asEStructuralFeature().SetDefaultValue(nil)
	case ESTRUCTURAL_FEATURE__DEFAULT_VALUE_LITERAL:
		e.asEStructuralFeature().SetDefaultValueLiteral("")
	case ESTRUCTURAL_FEATURE__DERIVED:
		e.asEStructuralFeature().SetDerived(false)
	case ESTRUCTURAL_FEATURE__FEATURE_ID:
		e.asEStructuralFeature().SetFeatureID(-1)
	case ESTRUCTURAL_FEATURE__TRANSIENT:
		e.asEStructuralFeature().SetTransient(false)
	case ESTRUCTURAL_FEATURE__UNSETTABLE:
		e.asEStructuralFeature().SetUnsettable(false)
	case ESTRUCTURAL_FEATURE__VOLATILE:
		e.asEStructuralFeature().SetVolatile(false)
	default:
		e.ETypedElementExt.EUnsetFromID(featureID)
	}
}

func (e *EStructuralFeatureImpl) EIsSetFromID(featureID int) bool {
	switch featureID {
	case ESTRUCTURAL_FEATURE__CHANGEABLE:
		return e.isChangeable != true
	case ESTRUCTURAL_FEATURE__DEFAULT_VALUE:
		return e.asEStructuralFeature().GetDefaultValue() != nil
	case ESTRUCTURAL_FEATURE__DEFAULT_VALUE_LITERAL:
		return e.defaultValueLiteral != ""
	case ESTRUCTURAL_FEATURE__DERIVED:
		return e.isDerived != false
	case ESTRUCTURAL_FEATURE__ECONTAINING_CLASS:
		return e.asEStructuralFeature().GetEContainingClass() != nil
	case ESTRUCTURAL_FEATURE__FEATURE_ID:
		return e.featureID != -1
	case ESTRUCTURAL_FEATURE__TRANSIENT:
		return e.isTransient != false
	case ESTRUCTURAL_FEATURE__UNSETTABLE:
		return e.isUnsettable != false
	case ESTRUCTURAL_FEATURE__VOLATILE:
		return e.isVolatile != false
	default:
		return e.ETypedElementExt.EIsSetFromID(featureID)
	}
}

func (e *EStructuralFeatureImpl) EInvokeFromID(operationID int, arguments EList) any {
	switch operationID {
	case ESTRUCTURAL_FEATURE__GET_CONTAINER_CLASS:
		return e.asEStructuralFeature().GetContainerClass()
	default:
		return e.ETypedElementExt.EInvokeFromID(operationID, arguments)
	}
}

func (e *EStructuralFeatureImpl) EBasicInverseAdd(otherEnd EObject, featureID int, notifications ENotificationChain) ENotificationChain {
	switch featureID {
	case ESTRUCTURAL_FEATURE__ECONTAINING_CLASS:
		msgs := notifications
		if e.EInternalContainer() != nil {
			msgs = e.EBasicRemoveFromContainer(msgs)
		}
		return e.EBasicSetContainer(otherEnd, ESTRUCTURAL_FEATURE__ECONTAINING_CLASS, msgs)
	default:
		return e.ETypedElementExt.EBasicInverseAdd(otherEnd, featureID, notifications)
	}
}

func (e *EStructuralFeatureImpl) EBasicInverseRemove(otherEnd EObject, featureID int, notifications ENotificationChain) ENotificationChain {
	switch featureID {
	case ESTRUCTURAL_FEATURE__ECONTAINING_CLASS:
		return e.EBasicSetContainer(nil, ESTRUCTURAL_FEATURE__ECONTAINING_CLASS, notifications)
	default:
		return e.ETypedElementExt.EBasicInverseRemove(otherEnd, featureID, notifications)
	}
}
