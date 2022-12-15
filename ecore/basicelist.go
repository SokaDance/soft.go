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

type abstractEList interface {
	doGet(index int) any

	doSet(index int, elem any) any

	doAdd(elem any)

	doAddAll(list EList) bool

	doInsert(index int, elem any)

	doInsertAll(index int, list EList) bool

	doClear() []any

	doMove(oldIndex, newIndew int) any

	doRemove(index int) any

	doRemoveRange(fromIndex, toIndex int)

	DidAdd(index int, elem any)

	DidSet(index int, newElem any, oldElem any)

	DidRemove(index int, old any)

	DidClear(oldObjects []any)

	DidMove(newIndex int, movedObject any, oldIndex int)

	DidChange()
}

// basicEList is an array of a dynamic size
type BasicEList struct {
	interfaces any
	data       []any
	isUnique   bool
}

// NewEmptyBasicEList return a new ArrayEList
func NewEmptyBasicEList() *BasicEList {
	a := new(BasicEList)
	a.interfaces = a
	a.data = []any{}
	a.isUnique = false
	return a
}

// NewBasicEList return a new ArrayEList
func NewBasicEList(data []any) *BasicEList {
	a := new(BasicEList)
	a.interfaces = a
	a.data = data
	a.isUnique = false
	return a
}

// NewUniqueBasicEList return a new ArrayEList with isUnique set as true
func NewUniqueBasicEList(data []any) *BasicEList {
	a := new(BasicEList)
	a.interfaces = a
	a.data = data
	a.isUnique = true
	return a
}

// Remove all elements from list that are not in ref list
func getNonDuplicates(list EList, ref EList) *BasicEList {
	newList := NewBasicEList([]any{})
	for it := list.Iterator(); it.HasNext(); {
		value := it.Next()
		if !newList.Contains(value) && !ref.Contains(value) {
			newList.Add(value)
		}
	}
	return newList
}

func (list *BasicEList) SetInterfaces(interfaces any) {
	list.interfaces = interfaces
}

func (list *BasicEList) asEList() EList {
	return list.interfaces.(EList)
}

func (list *BasicEList) asAbstractEList() abstractEList {
	return list.interfaces.(abstractEList)
}

func (list *BasicEList) Add(elem any) bool {
	if list.isUnique && list.Contains(elem) {
		return false
	}
	list.asAbstractEList().doAdd(elem)
	return true
}

// Add a new element to the array
func (list *BasicEList) doAdd(e any) {
	size := len(list.data)
	list.data = append(list.data, e)
	// events
	interfaces := list.asAbstractEList()
	interfaces.DidAdd(size, e)
	interfaces.DidChange()
}

// AddAll elements of an array in the current one
func (list *BasicEList) AddAll(collection EList) bool {
	if list.isUnique {
		collection = getNonDuplicates(collection, list)
		if collection.Size() == 0 {
			return false
		}
	}
	list.asAbstractEList().doAddAll(collection)
	return true
}

func (list *BasicEList) doAddAll(collection EList) bool {
	data := collection.ToArray()
	list.data = append(list.data, data...)
	interfaces := list.asAbstractEList()
	// events
	for i, element := range data {
		interfaces.DidAdd(i, element)
		interfaces.DidChange()
	}
	return len(data) != 0
}

// Insert an element in the array
func (list *BasicEList) Insert(index int, elem any) bool {
	if index < 0 || index > list.Size() {
		panic("Index out of bounds: index=" + strconv.Itoa(index) + " size=" + strconv.Itoa(list.Size()))
	}
	if list.isUnique && list.Contains(elem) {
		return false
	}
	list.asAbstractEList().doInsert(index, elem)
	return true
}

func (list *BasicEList) doInsert(index int, e any) {
	list.data = append(list.data, nil)
	copy(list.data[index+1:], list.data[index:])
	list.data[index] = e
	// events
	interfaces := list.asAbstractEList()
	interfaces.DidAdd(index, e)
	interfaces.DidChange()
}

// InsertAll element of an array at a given position
func (list *BasicEList) InsertAll(index int, collection EList) bool {
	if index < 0 || index > list.Size() {
		panic("Index out of bounds: index=" + strconv.Itoa(index) + " size=" + strconv.Itoa(list.Size()))
	}
	if list.isUnique {
		collection = getNonDuplicates(collection, list)
		if collection.Size() == 0 {
			return false
		}
	}
	list.asAbstractEList().doInsertAll(index, collection)
	return true
}

func (list *BasicEList) doInsertAll(index int, collection EList) bool {
	data := collection.ToArray()
	list.data = append(list.data[:index], append(data, list.data[index:]...)...)
	// events
	interfaces := list.asAbstractEList()
	for i, element := range data {
		interfaces.DidAdd(i+index, element)
		interfaces.DidChange()
	}
	return len(data) != 0
}

// Move an element to the given index
func (list *BasicEList) MoveObject(newIndex int, elem any) {
	oldIndex := list.asEList().IndexOf(elem)
	if oldIndex == -1 {
		panic("Object not found")
	}
	list.asAbstractEList().doMove(oldIndex, newIndex)
}

// Swap move an element from oldIndex to newIndex
func (list *BasicEList) Move(oldIndex, newIndex int) any {
	return list.asAbstractEList().doMove(oldIndex, newIndex)
}

func (list *BasicEList) doMove(oldIndex, newIndex int) any {
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
		interfaces := list.asAbstractEList()
		interfaces.DidMove(newIndex, object, oldIndex)
		interfaces.DidChange()
	}
	return object
}

// RemoveAt remove an element at a given position
func (list *BasicEList) RemoveAt(index int) any {
	return list.asAbstractEList().doRemove(index)
}

// Remove an element in an array
func (list *BasicEList) Remove(elem any) bool {
	index := list.asEList().IndexOf(elem)
	if index == -1 {
		return false
	}
	list.asAbstractEList().doRemove(index)
	return true
}

func (list *BasicEList) doRemove(index int) any {
	if index < 0 || index >= list.Size() {
		panic("Index out of bounds: index=" + strconv.Itoa(index) + " size=" + strconv.Itoa(list.Size()))
	}
	// retrieve removed object
	object := list.data[index]

	// remove index
	copy(list.data[index:], list.data[index+1:])
	list.data[len(list.data)-1] = nil
	list.data = list.data[:len(list.data)-1]

	// events
	interfaces := list.asAbstractEList()
	interfaces.DidRemove(index, object)
	interfaces.DidChange()
	return object
}

func (list *BasicEList) RemoveAll(collection EList) bool {
	modified := false
	for i := list.Size() - 1; i >= 0; i-- {
		if collection.Contains(list.asEList().Get(i)) {
			list.RemoveAt(i)
			modified = true
		}
	}
	return modified
}

// Get an element of the array
func (list *BasicEList) Get(index int) any {
	if index < 0 || index >= list.Size() {
		panic("Index out of bounds: index=" + strconv.Itoa(index) + " size=" + strconv.Itoa(list.Size()))
	}
	return list.asAbstractEList().doGet(index)
}

func (list *BasicEList) doGet(index int) any {
	return list.data[index]
}

func (list *BasicEList) doSet(index int, elem any) any {
	old := list.data[index]
	list.data[index] = elem
	// events
	interfaces := list.asAbstractEList()
	interfaces.DidSet(index, elem, old)
	interfaces.DidChange()
	return old
}

// Set an element of the array
func (list *BasicEList) Set(index int, elem any) any {
	if index < 0 || index >= list.Size() {
		panic("Index out of bounds: index=" + strconv.Itoa(index) + " size=" + strconv.Itoa(list.Size()))
	}
	if list.isUnique {
		currIndex := list.IndexOf(elem)
		if currIndex >= 0 && currIndex != index {
			panic("element already in list")
		}
	}
	return list.asAbstractEList().doSet(index, elem)
}

// Size count the number of element in the array
func (list *BasicEList) Size() int {
	return len(list.data)
}

// Clear remove all elements of the array
func (list *BasicEList) Clear() {
	list.asAbstractEList().doClear()
}

func (list *BasicEList) doClear() []any {
	oldData := list.data
	list.data = make([]any, 0)

	// events
	list.asAbstractEList().DidClear(oldData)
	return oldData
}

// Empty return true if the array contains 0 element
func (list *BasicEList) Empty() bool {
	return list.Size() == 0
}

// Contains return if an array contains or not an element
func (list *BasicEList) Contains(elem any) bool {
	return list.asEList().IndexOf(elem) != -1
}

// IndexOf return the index on an element in an array, else return -1
func (list *BasicEList) IndexOf(elem any) int {
	for i, value := range list.data {
		if value == elem {
			return i
		}
	}
	return -1
}

func (list *BasicEList) SubList(from int, to int) EList {
	return newBasicSubEList(list, from, to)
}

// Iterator through the array
func (list *BasicEList) Iterator() EIterator {
	return &listIterator{list: list}
}

func (list *BasicEList) ToArray() []any {
	return list.data
}

func (list *BasicEList) DidAdd(index int, elem any) {

}

func (list *BasicEList) DidSet(index int, newElem any, oldElem any) {

}

func (list *BasicEList) DidRemove(index int, old any) {

}

func (list *BasicEList) DidClear(oldObjects []any) {

}

func (list *BasicEList) DidMove(newIndex int, movedObject any, oldIndex int) {

}

func (list *BasicEList) DidChange() {

}

type basicSubEList struct {
	interfaces any
	offset     int
	size       int
	root       *BasicEList
	parent     *basicSubEList
}

func newBasicSubEList(root *BasicEList, fromIndex int, toIndex int) *basicSubEList {
	l := &basicSubEList{
		root:   root,
		offset: fromIndex,
		size:   toIndex - fromIndex,
	}
	l.interfaces = l
	return l
}

func (list *basicSubEList) asEList() EList {
	return list.interfaces.(EList)
}

func (list *basicSubEList) updateSize(sizeChange int) {
	for slist := list; slist != nil; slist = slist.parent {
		slist.size += sizeChange
	}
}

func (list *basicSubEList) Get(index int) any {
	return list.root.asEList().Get(list.offset + index)
}

func (list *basicSubEList) Set(index int, elem any) any {
	return list.root.asEList().Set(list.offset+index, elem)
}

func (list *basicSubEList) Add(elem any) bool {
	return list.asEList().Insert(list.offset+list.size, elem)
}

func (list *basicSubEList) AddAll(other EList) bool {
	return list.asEList().InsertAll(list.offset+list.size, other)
}

func (list *basicSubEList) Insert(index int, elem any) bool {
	b := list.root.asEList().Insert(list.offset+index, elem)
	if b {
		list.updateSize(1)
	}
	return b
}

func (list *basicSubEList) InsertAll(index int, collection EList) bool {
	if index < 0 || index > list.size {
		panic("Index out of bounds: index=" + strconv.Itoa(index) + " size=" + strconv.Itoa(list.Size()))
	}
	if list.root.isUnique {
		collection = getNonDuplicates(collection, list)
		if collection.Size() == 0 {
			return false
		}
	}
	b := list.root.asEList().InsertAll(list.offset+index, collection)
	if b {
		list.updateSize(collection.Size())
	}
	return b
}

func (list *basicSubEList) MoveObject(index int, elem any) {
	list.root.asEList().MoveObject(list.offset+index, elem)
}

func (list *basicSubEList) Move(oldIndex int, newIndex int) any {
	return list.root.asEList().Move(list.offset+oldIndex, list.offset+newIndex)
}

func (list *basicSubEList) RemoveAt(index int) any {
	e := list.root.asEList().RemoveAt(list.offset + index)
	list.updateSize(-1)
	return e
}

func (list *basicSubEList) Remove(elem any) bool {
	index := list.asEList().IndexOf(elem)
	if index != -1 {
		list.asEList().RemoveAt(index)
		return true
	}
	return false
}

func (list *basicSubEList) RemoveAll(collection EList) bool {
	modified := false
	for i := list.Size() - 1; i >= 0; i-- {
		if collection.Contains(list.asEList().Get(i)) {
			list.asEList().RemoveAt(i)
			modified = true
		}
	}
	return modified
}

func (list *basicSubEList) Size() int {
	return list.size
}

func (list *basicSubEList) Clear() {

}

func (list *basicSubEList) Empty() bool {
	return list.size == 0
}

func (list *basicSubEList) Contains(elem any) bool {
	return list.interfaces.(EList).IndexOf(elem) != -1
}

func (list *basicSubEList) IndexOf(elem any) int {
	for i := list.offset; i < list.size; i++ {
		if value := list.root.asEList(); value == elem {
			return i - list.offset
		}
	}
	return -1
}

func (list *basicSubEList) SubList(fromIndex int, toIndex int) EList {
	l := &basicSubEList{
		root:   list.root,
		parent: list,
		offset: list.offset + fromIndex,
		size:   toIndex - fromIndex,
	}
	l.interfaces = l
	return l
}

func (list *basicSubEList) Iterator() EIterator {

}

func (list *basicSubEList) ToArray() []any {

}
