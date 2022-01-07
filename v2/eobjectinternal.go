// *****************************************************************************
// Copyright(c) 2021 MASA Group
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// *****************************************************************************

package ecore

// EObjectInternal ...
type EObjectInternal interface {
	EObject

	EDynamicProperties() EDynamicProperties

	EStaticClass() EClass
	EStaticFeatureCount() int

	EInternalContainer() EObject
	EInternalContainerFeatureID() int
	EInternalResource() EResource
	ESetInternalContainer(container EObject, containerFeatureID int)
	ESetInternalResource(resource EResource)

	ESetResource(resource EResource, notifications ENotificationChain) ENotificationChain

	EInverseAdd(otherEnd EObject, featureID int, notifications ENotificationChain) ENotificationChain
	EInverseRemove(otherEnd EObject, featureID int, notifications ENotificationChain) ENotificationChain

	EFeatureID(feature EStructuralFeature) int
	EDerivedFeatureID(container EObject, featureID int) int
	EOperationID(operation EOperation) int
	EDerivedOperationID(container EObject, operationID int) int
	EGetFromID(featureID int, resolve bool) any
	ESetFromID(featureID int, newValue any)
	EUnsetFromID(featureID int)
	EIsSetFromID(featureID int) bool
	EInvokeFromID(operationID int, arguments EList[any]) any

	EBasicInverseAdd(otherEnd EObject, featureID int, notifications ENotificationChain) ENotificationChain
	EBasicInverseRemove(otherEnd EObject, featureID int, notifications ENotificationChain) ENotificationChain

	EObjectForFragmentSegment(string) EObject
	EURIFragmentSegment(EStructuralFeature, EObject) string

	EProxyURI() *URI
	ESetProxyURI(uri *URI)
	EResolveProxy(proxy EObject) EObject
}