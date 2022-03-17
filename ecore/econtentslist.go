// *****************************************************************************
// Copyright(c) 2021 MASA Group
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// *****************************************************************************

package ecore

type eContentsList[T EObject] struct {
	emptyImmutableEList[T]
	o        EObject
	features EList[EStructuralFeature]
	resolve  bool
}

type eContentsListIterator[T EObject] struct {
	l             *eContentsList[T]
	prepared      int
	next          T
	featureCursor int
	values        EIterator[T]
}

func (it *eContentsListIterator[T]) Next() T {
	if it.prepared > 1 || it.HasNext() {
		it.prepared = 0
		return it.next
	}
	panic("not such element")
}

func (it *eContentsListIterator[T]) HasNext() bool {
	switch it.prepared {
	case 2:
		return true
	case 1:
		return false
	default:
		features := it.l.features
		resolve := it.l.resolve
		o := it.l.o
		if it.values == nil || !it.values.HasNext() {
			// no feature list or feature list iterator is finished
			for it.featureCursor < features.Size() {
				feature := features.Get(it.featureCursor).(EStructuralFeature)
				it.featureCursor++
				if o.EIsSet(feature) {
					value := o.EGetResolve(feature, resolve)
					if feature.IsMany() {
						// list of values
						anyValues := value.(EList[any])
						// get unresolved list if object list and not resolved iterator
						if objectList, _ := value.(EObjectList[any]); objectList != nil && !resolve {
							anyValues = objectList.GetUnResolvedList()
						}
						values := FromAnyList[T](anyValues)
						if itValues := values.Iterator(); itValues.HasNext() {
							// we have a value
							it.values = itValues
							it.next = itValues.Next()
							it.prepared = 2
							return true
						}
					} else if value != nil {
						// we have a value
						it.values = nil
						it.next = value.(T)
						it.prepared = 2
						return true
					}
				}
			}
			it.values = nil
			it.prepared = 1
			return false
		} else {
			it.next = it.values.Next()
			it.prepared = 2
			return true
		}
	}
}

func newEContentsList[T EObject](o EObject, features EList[EStructuralFeature], resolve bool) *eContentsList[T] {
	return &eContentsList[T]{
		o:        o,
		features: features,
		resolve:  resolve,
	}
}

// Get an element of the array
func (l *eContentsList[T]) Get(index int) T {
	it := l.Iterator()
	for i := 0; i < index; i++ {
		it.Next()
	}
	return it.Next()
}

// Size count the number of element in the array
func (l *eContentsList[T]) Size() int {
	size := 0
	for it := l.features.Iterator(); it.HasNext(); {
		eFeature := it.Next().(EStructuralFeature)
		if l.o.EIsSet(eFeature) {
			value := l.o.EGetResolve(eFeature, false)
			if eFeature.IsMany() {
				list := value.(EList[T])
				size += list.Size()
			} else if value != nil {
				size++
			}
		}
	}
	return size
}

// Empty return true if the array contains 0 element
func (l *eContentsList[T]) Empty() bool {
	for it := l.features.Iterator(); it.HasNext(); {
		eFeature := it.Next().(EStructuralFeature)
		if l.o.EIsSet(eFeature) {
			value := l.o.EGetResolve(eFeature, false)
			if eFeature.IsMany() {
				list := value.(EList[T])
				if !list.Empty() {
					return false
				}
			} else if value != nil {
				return false
			}
		}
	}
	return true
}

// Contains return if an array contains or not an element
func (l *eContentsList[T]) Contains(elem T) bool {
	for it := l.Iterator(); it.HasNext(); {
		if e := it.Next(); equal(e, elem) {
			return true
		}
	}
	return false
}

// IndexOf return the index on an element in an array, else return -1
func (l *eContentsList[T]) IndexOf(elem T) int {
	index := 0
	for it := l.Iterator(); it.HasNext(); {
		if e := it.Next(); equal(e, elem) {
			return index
		}
		index++
	}
	return -1
}

// Iterator through the array
func (l *eContentsList[T]) Iterator() EIterator[T] {
	return &eContentsListIterator[T]{l: l}
}

// ToArray convert to array
func (l *eContentsList[T]) ToArray() []T {
	arr := []T{}
	for it := l.Iterator(); it.HasNext(); {
		arr = append(arr, it.Next())
	}
	return arr
}

func (l *eContentsList[T]) GetUnResolvedList() EObjectList[T] {
	if l.resolve {
		return newEContentsList[T](l.o, l.features, false)
	}
	return l
}
