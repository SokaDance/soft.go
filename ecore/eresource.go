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
	"io"
)

const (
	RESOURCE__RESOURCE_SET = 0

	RESOURCE__URI = 1

	RESOURCE__CONTENTS = 2

	RESOURCE__IS_LOADED = 4

	RESOURCE__ERRORS = 5

	RESOURCE__WARNINGS = 6
)

//EResource ...
type EResource interface {
	ENotifier

	GetResourceSet() EResourceSet

	GetURI() *URI
	SetURI(*URI)

	GetContents() EList[EObject]
	GetAllContents() EIterator[EObject]

	GetEObject(string) EObject
	GetURIFragment(EObject) string

	Attached(object EObject)
	Detached(object EObject)

	IsLoading() bool
	IsLoaded() bool

	Load()
	LoadWithOptions(options map[string]any)
	LoadWithReader(r io.Reader, options map[string]any)

	Unload()

	Save()
	SaveWithOptions(options map[string]any)
	SaveWithWriter(w io.Writer, options map[string]any)

	GetErrors() EList[EDiagnostic]
	GetWarnings() EList[EDiagnostic]

	SetObjectIDManager(EObjectIDManager)
	GetObjectIDManager() EObjectIDManager
}
