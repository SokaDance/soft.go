// *****************************************************************************
// Copyright(c) 2021 MASA Group
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// *****************************************************************************

package ecore

type eMapEntryImpl struct {
	key   any
	value any
}

func (e *eMapEntryImpl) GetKey() any {
	return e.key
}

func (e *eMapEntryImpl) SetKey(key any) {
	e.key = key
}

func (e *eMapEntryImpl) GetValue() any {
	return e.value
}

func (e *eMapEntryImpl) SetValue(value any) {
	e.value = value
}

func newMapEntry(key any, value any) EMapEntry {
	return &eMapEntryImpl{key: key, value: value}
}
