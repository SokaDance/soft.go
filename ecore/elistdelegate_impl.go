// *****************************************************************************
// Copyright(c) 2021 MASA Group
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// *****************************************************************************

package ecore

func ToList[T, U any](l EList[T], convertTo func(T) U, convertFrom func(U) T) EList[U] {
	if ld, isDelegate := l.(eListDelegate[U, T]); isDelegate {
		return ld.GetDelegate()
	} else {
		r := &eListDelegateImpl[T, U, EList[T]]{}
		r.delegate = l
		r.convertTo = convertTo
		r.convertFrom = convertFrom
		return r
	}
}

type eListDelegateImpl[T any, U any, C EList[T]] struct {
	eCollectionDelegateImpl[T, U, C]
}

func (l *eListDelegateImpl[T, U, C]) GetDelegate() EList[T] {
	return l.delegate
}

func (l *eListDelegateImpl[T, U, C]) Insert(index int, u U) bool {
	return l.delegate.Insert(index, l.convertFrom(u))
}

func (l *eListDelegateImpl[T, U, C]) InsertAll(index int, c ECollection[U]) bool {
	return l.delegate.InsertAll(index, ToCollection(c, l.convertFrom, l.convertTo))
}

func (l *eListDelegateImpl[T, U, C]) RemoveAt(index int) U {
	return l.convertTo(l.delegate.RemoveAt(index))
}

func (l *eListDelegateImpl[T, U, C]) MoveObject(newIndex int, u U) {
	l.delegate.MoveObject(newIndex, l.convertFrom(u))
}

func (l *eListDelegateImpl[T, U, C]) MoveIndex(oldIndex int, newIndex int) U {
	return l.convertTo(l.delegate.MoveIndex(oldIndex, newIndex))
}

func (l *eListDelegateImpl[T, U, C]) IndexOf(u U) int {
	return l.delegate.IndexOf(l.convertFrom(u))
}

func (l *eListDelegateImpl[T, U, C]) Get(index int) U {
	return l.convertTo(l.delegate.Get(index))
}

func (l *eListDelegateImpl[T, U, C]) Set(index int, u U) U {
	return l.convertTo(l.delegate.Set(index, l.convertFrom(u)))
}
