// *****************************************************************************
// Copyright(c) 2021 MASA Group
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// *****************************************************************************

package ecore

type EMap[K, V any] interface {
	EList[any]

	GetValue(key K) (V, bool)

	Put(key K, value V)

	RemoveKey(key K) V

	ContainsValue(value V) bool

	ContainsKey(key K) bool

	ToMap() map[any]any
}
