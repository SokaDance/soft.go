// *****************************************************************************
// Copyright(c) 2021 MASA Group
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// *****************************************************************************

package ecore

func ToArray[T, U any](c ECollection[T], convertTo func(T) U) []U {
	result := make([]U, c.Size())
	i := 0
	for it := c.Iterator(); it.HasNext(); i++ {
		result[i] = convertTo(it.Next())
	}
	return result
}

func ToCollection[T, U any](c ECollection[T], convertTo func(T) U, convertFrom func(U) T) ECollection[U] {
	if cd, isDelegate := c.(eCollectionDelegate[U, T]); isDelegate {
		return cd.GetDelegate()
	}
	return &eCollectionDelegateImpl[T, U, ECollection[T]]{delegate: c, convertTo: convertTo, convertFrom: convertFrom}
}

type eCollectionDelegateImpl[T any, U any, C ECollection[T]] struct {
	delegate    C
	convertTo   func(T) U
	convertFrom func(U) T
}

func (l *eCollectionDelegateImpl[T, U, C]) GetDelegate() ECollection[T] {
	return l.delegate
}

func (l *eCollectionDelegateImpl[T, U, C]) Add(u U) bool {
	return l.delegate.Add(l.convertFrom(u))
}

func (l *eCollectionDelegateImpl[T, U, C]) AddAll(c ECollection[U]) bool {
	return l.delegate.AddAll(ToCollection(c, l.convertFrom, l.convertTo))
}

func (l *eCollectionDelegateImpl[T, U, C]) Remove(u U) bool {
	return l.delegate.Remove(l.convertFrom(u))
}

func (l *eCollectionDelegateImpl[T, U, C]) RemoveAll(c ECollection[U]) bool {
	return l.delegate.RemoveAll(ToCollection(c, l.convertFrom, l.convertTo))
}

func (l *eCollectionDelegateImpl[T, U, C]) RetainAll(c ECollection[U]) bool {
	return l.delegate.RetainAll(ToCollection(c, l.convertFrom, l.convertTo))
}

func (l *eCollectionDelegateImpl[T, U, C]) Size() int {
	return l.delegate.Size()
}

func (l *eCollectionDelegateImpl[T, U, C]) Clear() {
	l.delegate.Clear()
}

func (l *eCollectionDelegateImpl[T, U, C]) Empty() bool {
	return l.delegate.Empty()
}

func (l *eCollectionDelegateImpl[T, U, C]) Contains(u U) bool {
	return l.delegate.Contains(l.convertFrom(u))
}

func (l *eCollectionDelegateImpl[T, U, C]) ToArray() []U {
	return ToArray[T](l.delegate, l.convertTo)
}

func (l *eCollectionDelegateImpl[T, U, C]) Iterator() EIterator[U] {
	return &eCollectionDelegateImplIterator[T, U]{delegate: l.delegate.Iterator(), convertTo: l.convertTo}
}

type eCollectionDelegateImplIterator[T any, U any] struct {
	delegate  EIterator[T]
	convertTo func(T) U
}

// Next return the current value of the iterator
func (it *eCollectionDelegateImplIterator[T, U]) Next() U {
	return it.convertTo(it.delegate.Next())
}

// HasNext make the iterator go further in the array
func (it *eCollectionDelegateImplIterator[T, U]) HasNext() bool {
	return it.delegate.HasNext()
}
