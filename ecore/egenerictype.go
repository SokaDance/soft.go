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

// EGenericType is the representation of the model object 'EGenericType'
type EGenericType interface {
	EObject

	IsInstance(any) bool

	GetEUpperBound() EGenericType
	SetEUpperBound(EGenericType)

	GetETypeArguments() EList[EGenericType]

	GetERawType() EClassifier

	GetELowerBound() EGenericType
	SetELowerBound(EGenericType)

	GetETypeParameter() ETypeParameter
	SetETypeParameter(ETypeParameter)

	GetEClassifier() EClassifier
	SetEClassifier(EClassifier)

	// Start of user code EGenericType
	// End of user code
}
