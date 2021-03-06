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

// EAnnotation is the representation of the model object 'EAnnotation'
type EAnnotation interface {
	EModelElement

	GetSource() string
	SetSource(string)

	GetDetails() EMap

	GetEModelElement() EModelElement
	SetEModelElement(EModelElement)

	GetContents() EList

	GetReferences() EList

	// Start of user code EAnnotation
	// End of user code
}
