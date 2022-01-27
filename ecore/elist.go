// *****************************************************************************
// Copyright(c) 2021 MASA Group
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// *****************************************************************************

package ecore

type EList[T any] interface {
	ECollection[T]

	Insert(int, T) bool

	InsertAll(int, ECollection[T]) bool

	RemoveAt(int) T

	MoveObject(newIndex int, t T)

	MoveIndex(oldIndex int, newIndex int) T

	IndexOf(T) int

	Get(int) T

	Set(int, T) T
}
