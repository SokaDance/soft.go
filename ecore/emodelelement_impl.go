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

// eModelElementImpl is the implementation of the model object 'EModelElement'
type eModelElementImpl struct {
	CompactEObjectContainer
	eAnnotations EList[EAnnotation]
}
type eModelElementImplInitializers interface {
	initEAnnotations() EList[EAnnotation]
}

// newEModelElementImpl is the constructor of a eModelElementImpl
func newEModelElementImpl() *eModelElementImpl {
	eModelElement := new(eModelElementImpl)
	eModelElement.SetInterfaces(eModelElement)
	eModelElement.Initialize()
	return eModelElement
}

func (eModelElement *eModelElementImpl) Initialize() {
	eModelElement.CompactEObjectContainer.Initialize()

}

func (eModelElement *eModelElementImpl) asEModelElement() EModelElement {
	return eModelElement.GetInterfaces().(EModelElement)
}

func (eModelElement *eModelElementImpl) asInitializers() eModelElementImplInitializers {
	return eModelElement.AsEObject().(eModelElementImplInitializers)
}

func (eModelElement *eModelElementImpl) EStaticClass() EClass {
	return GetPackage().GetEModelElement()
}

func (eModelElement *eModelElementImpl) EStaticFeatureCount() int {
	return EMODEL_ELEMENT_FEATURE_COUNT
}

// GetEAnnotation default implementation
func (eModelElement *eModelElementImpl) GetEAnnotation(string) EAnnotation {
	panic("GetEAnnotation not implemented")
}

// GetEAnnotations get the value of eAnnotations
func (eModelElement *eModelElementImpl) GetEAnnotations() EList[EAnnotation] {
	if eModelElement.eAnnotations == nil {
		eModelElement.eAnnotations = eModelElement.asInitializers().initEAnnotations()
	}
	return eModelElement.eAnnotations
}

func (eModelElement *eModelElementImpl) initEAnnotations() EList[EAnnotation] {
	return NewBasicEObjectList[EAnnotation](eModelElement.AsEObjectInternal(), EMODEL_ELEMENT__EANNOTATIONS, EANNOTATION__EMODEL_ELEMENT, true, true, true, false, false)
}

func (eModelElement *eModelElementImpl) EGetFromID(featureID int, resolve bool) any {
	switch featureID {
	case EMODEL_ELEMENT__EANNOTATIONS:
		return ToAnyList(eModelElement.asEModelElement().GetEAnnotations())
	default:
		return eModelElement.CompactEObjectContainer.EGetFromID(featureID, resolve)
	}
}

func (eModelElement *eModelElementImpl) ESetFromID(featureID int, value any) {
	switch featureID {
	case EMODEL_ELEMENT__EANNOTATIONS:
		newCollection := FromAnyCollection[EAnnotation](value)
		l := eModelElement.asEModelElement().GetEAnnotations()
		l.Clear()
		l.AddAll(newCollection)
	default:
		eModelElement.CompactEObjectContainer.ESetFromID(featureID, value)
	}
}

func (eModelElement *eModelElementImpl) EUnsetFromID(featureID int) {
	switch featureID {
	case EMODEL_ELEMENT__EANNOTATIONS:
		eModelElement.asEModelElement().GetEAnnotations().Clear()
	default:
		eModelElement.CompactEObjectContainer.EUnsetFromID(featureID)
	}
}

func (eModelElement *eModelElementImpl) EIsSetFromID(featureID int) bool {
	switch featureID {
	case EMODEL_ELEMENT__EANNOTATIONS:
		return eModelElement.eAnnotations != nil && eModelElement.eAnnotations.Size() != 0
	default:
		return eModelElement.CompactEObjectContainer.EIsSetFromID(featureID)
	}
}

func (eModelElement *eModelElementImpl) EInvokeFromID(operationID int, arguments EList[any]) any {
	switch operationID {
	case EMODEL_ELEMENT__GET_EANNOTATION_ESTRING:
		return eModelElement.asEModelElement().GetEAnnotation(arguments.Get(0).(string))
	default:
		return eModelElement.CompactEObjectContainer.EInvokeFromID(operationID, arguments)
	}
}

func (eModelElement *eModelElementImpl) EBasicInverseAdd(otherEnd EObject, featureID int, notifications ENotificationChain) ENotificationChain {
	switch featureID {
	case EMODEL_ELEMENT__EANNOTATIONS:
		list := eModelElement.GetEAnnotations().(ENotifyingList[EAnnotation])
		end := otherEnd.(EAnnotation)
		return list.AddWithNotification(end, notifications)
	default:
		return eModelElement.CompactEObjectContainer.EBasicInverseAdd(otherEnd, featureID, notifications)
	}
}

func (eModelElement *eModelElementImpl) EBasicInverseRemove(otherEnd EObject, featureID int, notifications ENotificationChain) ENotificationChain {
	switch featureID {
	case EMODEL_ELEMENT__EANNOTATIONS:
		list := eModelElement.GetEAnnotations().(ENotifyingList[EAnnotation])
		end := otherEnd.(EAnnotation)
		return list.RemoveWithNotification(end, notifications)
	default:
		return eModelElement.CompactEObjectContainer.EBasicInverseRemove(otherEnd, featureID, notifications)
	}
}
