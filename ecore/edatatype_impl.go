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

// eDataTypeImpl is the implementation of the model object 'EDataType'
type eDataTypeImpl struct {
	eClassifierExt
	isSerializable bool
}

// newEDataTypeImpl is the constructor of a eDataTypeImpl
func newEDataTypeImpl() *eDataTypeImpl {
	eDataType := new(eDataTypeImpl)
	eDataType.SetInterfaces(eDataType)
	eDataType.Initialize()
	return eDataType
}

func (eDataType *eDataTypeImpl) Initialize() {
	eDataType.eClassifierExt.Initialize()
	eDataType.isSerializable = true

}

func (eDataType *eDataTypeImpl) asEDataType() EDataType {
	return eDataType.GetInterfaces().(EDataType)
}

func (eDataType *eDataTypeImpl) EStaticClass() EClass {
	return GetPackage().GetEDataType()
}

func (eDataType *eDataTypeImpl) EStaticFeatureCount() int {
	return EDATA_TYPE_FEATURE_COUNT
}

// IsSerializable get the value of isSerializable
func (eDataType *eDataTypeImpl) IsSerializable() bool {
	return eDataType.isSerializable
}

// SetSerializable set the value of isSerializable
func (eDataType *eDataTypeImpl) SetSerializable(newIsSerializable bool) {
	oldIsSerializable := eDataType.isSerializable
	eDataType.isSerializable = newIsSerializable
	if eDataType.ENotificationRequired() {
		eDataType.ENotify(NewNotificationByFeatureID(eDataType.AsEObject(), SET, EDATA_TYPE__SERIALIZABLE, oldIsSerializable, newIsSerializable, NO_INDEX))
	}
}

func (eDataType *eDataTypeImpl) EGetFromID(featureID int, resolve bool) interface{} {
	switch featureID {
	case EDATA_TYPE__SERIALIZABLE:
		return eDataType.asEDataType().IsSerializable()
	default:
		return eDataType.eClassifierExt.EGetFromID(featureID, resolve)
	}
}

func (eDataType *eDataTypeImpl) ESetFromID(featureID int, newValue interface{}) {
	switch featureID {
	case EDATA_TYPE__SERIALIZABLE:
		eDataType.asEDataType().SetSerializable(newValue.(bool))
	default:
		eDataType.eClassifierExt.ESetFromID(featureID, newValue)
	}
}

func (eDataType *eDataTypeImpl) EUnsetFromID(featureID int) {
	switch featureID {
	case EDATA_TYPE__SERIALIZABLE:
		eDataType.asEDataType().SetSerializable(true)
	default:
		eDataType.eClassifierExt.EUnsetFromID(featureID)
	}
}

func (eDataType *eDataTypeImpl) EIsSetFromID(featureID int) bool {
	switch featureID {
	case EDATA_TYPE__SERIALIZABLE:
		return eDataType.isSerializable != true
	default:
		return eDataType.eClassifierExt.EIsSetFromID(featureID)
	}
}
