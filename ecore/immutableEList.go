// *****************************************************************************
// Copyright(c) 2021 MASA Group
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// *****************************************************************************

package ecore

import "strconv"

type emptyImmutableEList[T any] struct {
}

func (l *emptyImmutableEList[T]) Add(t T) bool {
	panic("Immutable list can't be modified")
}

func (l *emptyImmutableEList[T]) AddAll(c ECollection[T]) bool {
	panic("Immutable list can't be modified")
}

func (l *emptyImmutableEList[T]) Insert(index int, t T) bool {
	panic("Immutable list can't be modified")
}

func (l *emptyImmutableEList[T]) InsertAll(index int, c ECollection[T]) bool {
	panic("Immutable list can't be modified")
}

func (l *emptyImmutableEList[T]) MoveObject(newIndex int, t T) {
	panic("Immutable list can't be modified")
}

func (l *emptyImmutableEList[T]) MoveIndex(oldIndex, newIndex int) T {
	panic("Immutable list can't be modified")
}

// Get an element of the array
func (l *emptyImmutableEList[T]) Get(index int) T {
	panic("Index out of bounds: index=" + strconv.Itoa(index) + " size=" + strconv.Itoa(l.Size()))
}

func (l *emptyImmutableEList[T]) Set(index int, t T) T {
	panic("Immutable list can't be modified")
}

func (l *emptyImmutableEList[T]) RemoveAt(index int) T {
	panic("Immutable list can't be modified")
}

func (l *emptyImmutableEList[T]) Remove(t T) bool {
	panic("Immutable list can't be modified")
}

func (l *emptyImmutableEList[T]) RemoveAll(c ECollection[T]) bool {
	panic("Immutable list can't be modified")
}

func (l *emptyImmutableEList[T]) RetainAll(c ECollection[T]) bool {
	panic("Immutable list can't be modified")
}

// Size count the number of element in the array
func (l *emptyImmutableEList[T]) Size() int {
	return 0
}

func (l *emptyImmutableEList[T]) Clear() {
	panic("Immutable list can't be modified")
}

// Empty return true if the array contains 0 element
func (l *emptyImmutableEList[T]) Empty() bool {
	return true
}

// Contains return if an array contains or not an element
func (l *emptyImmutableEList[T]) Contains(t T) bool {
	return false
}

// IndexOf return the index on an element in an array, else return -1
func (l *emptyImmutableEList[T]) IndexOf(t T) int {
	return -1
}

// Iterator through the array
func (l *emptyImmutableEList[T]) Iterator() EIterator[T] {
	return &eListIterator[T]{list: l}
}

// ToArray convert to array
func (l *emptyImmutableEList[T]) ToArray() []T {
	return []T{}
}

func NewEmptyImmutableEList[T any]() *emptyImmutableEList[T] {
	return &emptyImmutableEList[T]{}
}

type immutableEList[T comparable] struct {
	emptyImmutableEList[T]
	data []T
}

// NewImmutableEList return a new ImmutableEList
func NewImmutableEList[T comparable](data []T) *immutableEList[T] {
	return &immutableEList[T]{data: data}
}

// Get an element of the array
func (l *immutableEList[T]) Get(index int) T {
	if index < 0 || index >= l.Size() {
		panic("Index out of bounds: index=" + strconv.Itoa(index) + " size=" + strconv.Itoa(l.Size()))
	}
	return l.data[index]
}

// Size count the number of element in the array
func (l *immutableEList[T]) Size() int {
	return len(l.data)
}

// Empty return true if the array contains 0 element
func (l *immutableEList[T]) Empty() bool {
	return l.Size() == 0
}

// Contains return if an array contains or not an element
func (l *immutableEList[T]) Contains(t T) bool {
	return l.IndexOf(t) != -1
}

// IndexOf return the index on an element in an array, else return -1
func (l *immutableEList[T]) IndexOf(t T) int {
	for i, value := range l.data {
		if value == t {
			return i
		}
	}
	return -1
}

// Iterator through the array
func (l *immutableEList[T]) Iterator() EIterator[T] {
	return &eListIterator[T]{list: l}
}

// ToArray convert to array
func (l *immutableEList[T]) ToArray() []T {
	return l.data
}
