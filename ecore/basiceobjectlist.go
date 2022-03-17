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

type basicEObjectList[T any] struct {
	BasicENotifyingList[T]
	owner            EObjectInternal
	featureID        int
	inverseFeatureID int
	containment      bool
	inverse          bool
	opposite         bool
	proxies          bool
	unset            bool
}

func NewBasicEObjectList[T any](owner EObjectInternal, featureID int, inverseFeatureID int, containment, inverse, opposite, proxies, unset bool) *basicEObjectList[T] {
	l := new(basicEObjectList[T])
	l.interfaces = l
	l.data = []T{}
	l.isUnique = true
	l.owner = owner
	l.featureID = featureID
	l.inverseFeatureID = inverseFeatureID
	l.containment = containment
	l.inverse = inverse
	l.opposite = opposite
	l.proxies = proxies
	l.unset = unset
	return l
}

// GetNotifier ...
func (list *basicEObjectList[T]) GetNotifier() ENotifier {
	return list.owner
}

// GetFeature ...
func (list *basicEObjectList[T]) GetFeature() EStructuralFeature {
	if list.owner != nil {
		return list.owner.EClass().GetEStructuralFeature(list.featureID)
	}
	return nil
}

// GetFeatureID ...
func (list *basicEObjectList[T]) GetFeatureID() int {
	return list.featureID
}

// GetUnResolvedList ...
func (list *basicEObjectList[T]) GetUnResolvedList() EObjectList[T] {
	if list.proxies {
		u := new(unResolvedBasicEObjectList[T])
		u.delegate = list
		return u
	}
	return list
}

func (list *basicEObjectList[T]) IndexOf(elem T) int {
	if list.proxies {
		for i, value := range list.data {
			if equal(value, elem) || equal(list.resolve(i, value), elem) {
				return i
			}
		}
		return -1
	}
	return list.basicEList.IndexOf(elem)
}

func (list *basicEObjectList[T]) doGet(index int) T {
	return list.resolve(index, list.basicEList.doGet(index))
}

func (list *basicEObjectList[T]) resolve(index int, object T) T {
	resolved := list.resolveProxy(any(object).(EObject)).(T)
	if !equal(resolved, object) {
		list.basicEList.doSet(index, resolved)
		var notifications ENotificationChain
		if list.containment {
			internal := list.interfaces.(eNotifyingListInternal[T])
			notifications = internal.inverseRemove(object, notifications)
			if resolvedInternal, _ := any(resolved).(EObjectInternal); resolvedInternal != nil && resolvedInternal.EInternalContainer() == nil {
				notifications = internal.inverseAdd(resolved, notifications)
			}
		}
		list.createAndDispatchNotification(notifications, RESOLVE, object, resolved, index)
	}
	return resolved
}

func (list *basicEObjectList[T]) resolveProxy(eObject EObject) EObject {
	if list.proxies && eObject.EIsProxy() {
		eObject = list.owner.EResolveProxy(eObject)
	}
	return eObject
}

func (list *basicEObjectList[T]) inverseAdd(object T, notifications ENotificationChain) ENotificationChain {
	internal, _ := any(object).(EObjectInternal)
	if internal != nil && list.inverse {
		if list.opposite {
			return internal.EInverseAdd(list.owner, list.inverseFeatureID, notifications)
		} else {
			return internal.EInverseAdd(list.owner, EOPPOSITE_FEATURE_BASE-list.featureID, notifications)
		}
	}
	return notifications
}

func (list *basicEObjectList[T]) inverseRemove(object T, notifications ENotificationChain) ENotificationChain {
	internal, _ := any(object).(EObjectInternal)
	if internal != nil && list.inverse {
		if list.opposite {
			return internal.EInverseRemove(list.owner, list.inverseFeatureID, notifications)
		} else {
			return internal.EInverseRemove(list.owner, EOPPOSITE_FEATURE_BASE-list.featureID, notifications)
		}
	}
	return notifications
}

type unResolvedBasicEObjectList[T any] struct {
	delegate *basicEObjectList[T]
}

func (l *unResolvedBasicEObjectList[T]) Add(elem T) bool {
	if l.delegate.isUnique && l.Contains(elem) {
		return false
	}
	l.delegate.interfaces.(abstractEList[T]).doAdd(elem)
	return true
}

// AddAll elements of an list in the current one
func (l *unResolvedBasicEObjectList[T]) AddAll(c ECollection[T]) bool {
	if l.delegate.isUnique {
		c = getNonDuplicates[T](l, c)
		if c.Empty() {
			return false
		}
	}
	l.delegate.interfaces.(abstractEList[T]).doAddAll(c)
	return true
}

// Insert an element in the list
func (l *unResolvedBasicEObjectList[T]) Insert(index int, elem T) bool {
	if index < 0 || index > l.Size() {
		panic("Index out of bounds: index=" + strconv.Itoa(index) + " size=" + strconv.Itoa(l.Size()))
	}
	if l.delegate.isUnique && l.Contains(elem) {
		return false
	}
	l.delegate.interfaces.(abstractEList[T]).doInsert(index, elem)
	return true
}

// InsertAll element of an list at a given position
func (l *unResolvedBasicEObjectList[T]) InsertAll(index int, c ECollection[T]) bool {
	if index < 0 || index > l.Size() {
		panic("Index out of bounds: index=" + strconv.Itoa(index) + " size=" + strconv.Itoa(l.Size()))
	}
	if l.delegate.isUnique {
		c = getNonDuplicates[T](l, c)
		if c.Empty() {
			return false
		}
	}
	l.delegate.interfaces.(abstractEList[T]).doInsertAll(index, c)
	return true
}

// Move an element to the given index
func (l *unResolvedBasicEObjectList[T]) MoveObject(newIndex int, elem T) {
	oldIndex := l.IndexOf(elem)
	if oldIndex == -1 {
		panic("Object not found")
	}
	l.delegate.interfaces.(abstractEList[T]).doMove(oldIndex, newIndex)
}

// Swap move an element from oldIndex to newIndex
func (l *unResolvedBasicEObjectList[T]) MoveIndex(oldIndex, newIndex int) T {
	return l.delegate.interfaces.(abstractEList[T]).doMove(oldIndex, newIndex)
}

// RemoveAt remove an element at a given position
func (l *unResolvedBasicEObjectList[T]) RemoveAt(index int) T {
	return l.delegate.interfaces.(abstractEList[T]).doRemove(index)
}

// Remove an element in an list
func (l *unResolvedBasicEObjectList[T]) Remove(elem T) bool {
	index := l.IndexOf(elem)
	if index == -1 {
		return false
	}
	l.delegate.interfaces.(abstractEList[T]).doRemove(index)
	return true
}

func (l *unResolvedBasicEObjectList[T]) RemoveAll(collection ECollection[T]) bool {
	modified := false
	for i := l.Size() - 1; i >= 0; {
		if collection.Contains(l.Get(i)) {
			l.RemoveAt(i)
			modified = true
		}
		i--
	}
	return modified
}

func (l *unResolvedBasicEObjectList[T]) RetainAll(collection ECollection[T]) bool {
	modified := false
	for i := l.Size() - 1; i >= 0; {
		if !collection.Contains(l.Get(i)) {
			l.RemoveAt(i)
			modified = true
		}
		i--
	}
	return modified
}

// Get an element of the list
func (l *unResolvedBasicEObjectList[T]) Get(index int) T {
	if index < 0 || index >= l.Size() {
		panic("Index out of bounds: index=" + strconv.Itoa(index) + " size=" + strconv.Itoa(l.Size()))
	}
	return l.delegate.data[index]
}

// Set an element of the list
func (l *unResolvedBasicEObjectList[T]) Set(index int, elem T) T {
	if index < 0 || index >= l.Size() {
		panic("Index out of bounds: index=" + strconv.Itoa(index) + " size=" + strconv.Itoa(l.Size()))
	}
	if l.delegate.isUnique {
		currIndex := l.IndexOf(elem)
		if currIndex >= 0 && currIndex != index {
			panic("element already in list")
		}
	}
	return l.delegate.interfaces.(abstractEList[T]).doSet(index, elem)
}

// Size count the number of element in the list
func (l *unResolvedBasicEObjectList[T]) Size() int {
	return l.delegate.Size()
}

// Clear remove all elements of the list
func (l *unResolvedBasicEObjectList[T]) Clear() {
	l.delegate.Clear()
}

// Empty return true if the list contains 0 element
func (l *unResolvedBasicEObjectList[T]) Empty() bool {
	return l.delegate.Empty()
}

// Contains return if an list contains or not an element
func (l *unResolvedBasicEObjectList[T]) Contains(elem T) bool {
	return l.IndexOf(elem) != -1
}

// IndexOf return the index on an element in an list, else return -1
func (l *unResolvedBasicEObjectList[T]) IndexOf(elem T) int {
	return l.delegate.basicEList.IndexOf(elem)
}

// Iterator through the list
func (l *unResolvedBasicEObjectList[T]) Iterator() EIterator[T] {
	return &eListIterator[T]{list: l}
}

func (l *unResolvedBasicEObjectList[T]) ToArray() []T {
	return l.delegate.ToArray()
}

func (l *unResolvedBasicEObjectList[T]) GetNotifier() ENotifier {
	return l.delegate.GetNotifier()
}

func (l *unResolvedBasicEObjectList[T]) GetFeature() EStructuralFeature {
	return l.delegate.GetFeature()
}

func (l *unResolvedBasicEObjectList[T]) GetFeatureID() int {
	return l.delegate.GetFeatureID()
}

func (l *unResolvedBasicEObjectList[T]) AddWithNotification(object T, notifications ENotificationChain) ENotificationChain {
	return l.delegate.AddWithNotification(object, notifications)
}

func (l *unResolvedBasicEObjectList[T]) RemoveWithNotification(object T, notifications ENotificationChain) ENotificationChain {
	index := l.delegate.basicEList.IndexOf(object)
	if index != -1 {
		oldObject := l.delegate.basicEList.doRemove(index)
		return l.delegate.createAndAddNotification(notifications, REMOVE, oldObject, nil, index)
	}
	return notifications
}

func (l *unResolvedBasicEObjectList[T]) SetWithNotification(index int, object T, notifications ENotificationChain) ENotificationChain {
	return l.delegate.SetWithNotification(index, object, notifications)
}

func (l *unResolvedBasicEObjectList[T]) GetUnResolvedList() EObjectList[T] {
	return l
}
