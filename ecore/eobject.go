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

// EObject is the representation of the model object 'EObject'
type EObject interface {
	ENotifier

	EClass() EClass
	EIsProxy() bool
	EResource() EResource
	EContainer() EObject
	EContainingFeature() EStructuralFeature
	EContainmentFeature() EReference
	EContents() EList[EObject]
	EAllContents() EIterator[EObject]
	ECrossReferences() EList[EObject]
	EGet(EStructuralFeature) any
	EGetResolve(EStructuralFeature, bool) any
	ESet(EStructuralFeature, any)
	EIsSet(EStructuralFeature) bool
	EUnset(EStructuralFeature)
	EInvoke(EOperation, EList[any]) any

	// Start of user code EObject
	// End of user code
}
