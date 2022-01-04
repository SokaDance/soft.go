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

type abstractEList[T comparable] interface {
	doGet(index int) T

	doSet(index int, t T) T

	doAdd(t T)

	doAddAll(c Collection[T]) bool

	doInsert(index int, t T)

	doInsertAll(index int, c Collection[T]) bool

	doClear() []T

	doMove(oldIndex, newIndew int) T

	doRemove(index int) T

	didAdd(index int, t T)

	didSet(index int, newElem T, oldElem T)

	didRemove(index int, old T)

	didClear(oldObjects []T)

	didMove(newIndex int, movedObject T, oldIndex int)

	didChange()
}

// basicEList is an array of a dynamic size
type basicEList[T comparable] struct {
	interfaces interface{}
	data       []T
	isUnique   bool
}

// NewEmptyBasicEList return a new ArrayEList
func NewEmptyBasicEList[T comparable]() *basicEList[T] {
	a := new(basicEList[T])
	a.interfaces = a
	a.data = []T{}
	a.isUnique = false
	return a
}

// NewBasicEList return a new ArrayEList
func NewBasicEList[T comparable](data []T) *basicEList[T] {
	a := new(basicEList[T])
	a.interfaces = a
	a.data = data
	a.isUnique = false
	return a
}

// NewUniqueBasicEList return a new ArrayEList with isUnique set as true
func NewUniqueBasicEList[T comparable](data []T) *basicEList[T] {
	a := new(basicEList[T])
	a.interfaces = a
	a.data = data
	a.isUnique = true
	return a
}

func (list *basicEList[T]) SetInterfaces(interfaces interface{}) {
	list.interfaces = interfaces
}

func (list *basicEList[T]) asAbstractEList() abstractEList[T] {
	return 	list.interfaces.(abstractEList[T])
}

func (list *basicEList[T]) asEList() EList[T] {
	return 	list.interfaces.(EList[T])
}

// Remove all elements in collection that already are in list
func  (list *basicEList[T]) getNonDuplicates(c Collection[T]) Collection[T] {
	newList := NewBasicEList([]T{})
	refList := list.asEList()
	for it := c.Iterator(); it.HasNext(); {
		value := it.Next()
		if !newList.Contains(value) && !refList.Contains(value) {
			newList.Add(value)
		}
	}
	return newList
}

func (list *basicEList[T]) Add(t T) bool {
	if list.isUnique && list.Contains(t) {
		return false
	}
	list.asAbstractEList().doAdd(t)
	return true
}

// Add a new element to the array
func (list *basicEList[T]) doAdd(e T) {
	size := len(list.data)
	list.data = append(list.data, e)
	// events
	abstractEList := list.asAbstractEList()
	abstractEList.didAdd(size, e)
	abstractEList.didChange()
}

// AddAll elements of an array in the current one
func (list *basicEList[T]) AddAll(collection Collection[T]) bool {
	if list.isUnique {
		collection = list.getNonDuplicates(collection)
		if collection.Size() == 0 {
			return false
		}
	}
	list.asAbstractEList().doAddAll(collection)
	return true
}

func (list *basicEList[T]) doAddAll(collection Collection[T]) bool {
	data := collection.ToArray()
	list.data = append(list.data, data...)
	abstractEList := list.asAbstractEList()
	// events
	for i, element := range data {
		abstractEList.didAdd(i, element)
		abstractEList.didChange()
	}
	return len(data) != 0
}

// Insert an element in the array
func (list *basicEList[T]) Insert(index int, t T) bool {
	if index < 0 || index > list.Size() {
		panic("Index out of bounds: index=" + strconv.Itoa(index) + " size=" + strconv.Itoa(list.Size()))
	}
	if list.isUnique && list.Contains(t) {
		return false
	}
	list.asAbstractEList().doInsert(index, t)
	return true
}

func (list *basicEList[T]) doInsert(index int, t T) {
	list.data = append(list.data, t)
	copy(list.data[index+1:], list.data[index:])
	list.data[index] = t
	// events
	abstractEList := list.asAbstractEList()
	abstractEList.didAdd(index, t)
	abstractEList.didChange()
}

// InsertAll element of an array at a given position
func (list *basicEList[T]) InsertAll(index int, collection Collection[T]) bool {
	if index < 0 || index > list.Size() {
		panic("Index out of bounds: index=" + strconv.Itoa(index) + " size=" + strconv.Itoa(list.Size()))
	}
	if list.isUnique {
		collection = list.getNonDuplicates(collection)
		if collection.Size() == 0 {
			return false
		}
	}
	list.asAbstractEList().doInsertAll(index, collection)
	return true
}

func (list *basicEList[T]) doInsertAll(index int, collection Collection[T]) bool {
	data := collection.ToArray()
	list.data = append(list.data[:index], append(data, list.data[index:]...)...)
	// events
	abstractEList := list.asAbstractEList()
	for i, element := range data {
		abstractEList.didAdd(i+index, element)
		abstractEList.didChange()
	}
	return len(data) != 0
}

// Move an element to the given index
func (list *basicEList[T]) MoveObject(newIndex int, t T) {
	oldIndex := list.asEList().IndexOf(t)
	if oldIndex == -1 {
		panic("Object not found")
	}
	list.asAbstractEList().doMove(oldIndex, newIndex)
}

// Swap move an element from oldIndex to newIndex
func (list *basicEList[T]) MoveIndex(oldIndex, newIndex int) T {
	return list.asAbstractEList().doMove(oldIndex, newIndex)
}

func (list *basicEList[T]) doMove(oldIndex, newIndex int) T {
	if oldIndex < 0 || oldIndex >= list.Size() ||
		newIndex < 0 || newIndex > list.Size() {
		panic("Index out of bounds: oldIndex=" + strconv.Itoa(oldIndex) + " newIndex=" + strconv.Itoa(newIndex) + " size=" + strconv.Itoa(list.Size()))
	}

	object := list.data[oldIndex]
	if oldIndex != newIndex {
		if newIndex < oldIndex {
			copy(list.data[newIndex+1:], list.data[newIndex:oldIndex])
		} else {
			copy(list.data[oldIndex:], list.data[oldIndex+1:newIndex+1])
		}
		list.data[newIndex] = object

		// events
		abstractEList := list.asAbstractEList()
		abstractEList.didMove(newIndex, object, oldIndex)
		abstractEList.didChange()
	}
	return object
}

// RemoveAt remove an element at a given position
func (list *basicEList[T]) RemoveAt(index int) T {
	return list.asAbstractEList().doRemove(index)
}

// Remove an element in an array
func (list *basicEList[T]) Remove(t T) bool {
	index := list.asEList().IndexOf(t)
	if index == -1 {
		return false
	}
	list.asAbstractEList().doRemove(index)
	return true
}

func (list *basicEList[T]) doRemove(index int) T {
	if index < 0 || index >= list.Size() {
		panic("Index out of bounds: index=" + strconv.Itoa(index) + " size=" + strconv.Itoa(list.Size()))
	}
	object := list.data[index]
	list.data = append(list.data[:index], list.data[index+1:]...)
	// events
	abstractEList := list.asAbstractEList()
	abstractEList.didRemove(index, object)
	abstractEList.didChange()
	return object
}

func (list *basicEList[T]) RemoveAll(collection Collection[T]) bool {
	modified := false
	for i := list.Size() - 1; i >= 0; {
		if collection.Contains(list.Get(i)) {
			list.RemoveAt(i)
			modified = true
		}
		i--
	}
	return modified
}

func (list *basicEList[T]) RetainAll(collection Collection[T]) bool {
	modified := false
	for i := list.Size() - 1; i >= 0; {
		if (!collection.Contains(list.Get(i))) {
			list.RemoveAt(i)
			modified = true
		}
		i--
	}
	return modified
}


// Get an element of the array
func (list *basicEList[T]) Get(index int) T {
	if index < 0 || index >= list.Size() {
		panic("Index out of bounds: index=" + strconv.Itoa(index) + " size=" + strconv.Itoa(list.Size()))
	}
	return list.asAbstractEList().doGet(index)
}

func (list *basicEList[T]) doGet(index int) T {
	return list.data[index]
}

func (list *basicEList[T]) doSet(index int, t T) T {
	old := list.data[index]
	list.data[index] = t
	// events
	abstractEList := list.asAbstractEList()
	abstractEList.didSet(index, t, old)
	abstractEList.didChange()
	return old
}

// Set an element of the array
func (list *basicEList[T]) Set(index int, t T) T {
	if index < 0 || index >= list.Size() {
		panic("Index out of bounds: index=" + strconv.Itoa(index) + " size=" + strconv.Itoa(list.Size()))
	}
	if list.isUnique {
		currIndex := list.IndexOf(t)
		if currIndex >= 0 && currIndex != index {
			panic("element already in list")
		}
	}
	return list.asAbstractEList().doSet(index, t)
}

// Size count the number of element in the array
func (list *basicEList[T]) Size() int {
	return len(list.data)
}

// Clear remove all elements of the array
func (list *basicEList[T]) Clear() {
	list.asAbstractEList().doClear()
}

func (list *basicEList[T]) doClear() []T {
	oldData := list.data
	list.data = make([]T, 0)

	// events
	list.asAbstractEList().didClear(oldData)
	return oldData
}

// Empty return true if the array contains 0 element
func (list *basicEList[T]) Empty() bool {
	return list.Size() == 0
}

// Contains return if an array contains or not an element
func (list *basicEList[T]) Contains(t T) bool {
	return list.interfaces.(EList[T]).IndexOf(t) != -1
}

// IndexOf return the index on an element in an array, else return -1
func (list *basicEList[T]) IndexOf(t T) int {
	for i, value := range list.data {
		if value == t {
			return i
		}
	}
	return -1
}

// Iterator through the array
func (list *basicEList[T]) Iterator() Iterator[T] {
	return &eListIterator[T]{list: list}
}

func (list *basicEList[T]) ToArray() []T {
	return list.data
}

func (list *basicEList[T]) didAdd(index int, t T) {

}

func (list *basicEList[T]) didSet(index int, newElem T, oldElem T) {

}

func (list *basicEList[T]) didRemove(index int, old T) {

}

func (list *basicEList[T]) didClear(oldObjects []T) {

}

func (list *basicEList[T]) didMove(newIndex int, movedObject T, oldIndex int) {

}

func (list *basicEList[T]) didChange() {

}
