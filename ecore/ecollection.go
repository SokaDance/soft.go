// *****************************************************************************
// Copyright(c) 2021 MASA Group
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// *****************************************************************************

package ecore

type ECollection[T any] interface {
	EIterable[T]

	Add(t T) bool

	AddAll(c ECollection[T]) bool

	Remove(t T) bool

	RemoveAll(c ECollection[T]) bool

	RetainAll(c ECollection[T]) bool

	Size() int

	Clear()

	Empty() bool

	Contains(t T) bool

	ToArray() []T
}
