// *****************************************************************************
// Copyright(c) 2021 MASA Group
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// *****************************************************************************

package ecore

func ToObjectList[T, U any](l EObjectList[T], convertTo func(T) U, convertFrom func(U) T) EObjectList[U] {
	if dp, isDelegateProvider := l.(EDelegateProvider); isDelegateProvider {
		if list, isList := dp.GetDelegate().(EObjectList[U]); isList {
			return list
		}
	}
	r := &eObjectListDelegateImpl[T, U, EObjectListConstraint[T]]{}
	r.delegate = l.(EObjectListConstraint[T])
	r.convertTo = convertTo
	r.convertFrom = convertFrom
	return r
}

type EObjectListConstraint[T any] interface {
	EObjectList[T]
	ENotifyingList[T]
}

type eObjectListDelegateImpl[T any, U any, C EObjectListConstraint[T]] struct {
	eNotifyingListDelegateImpl[T, U, C]
}

func (l *eObjectListDelegateImpl[T, U, C]) GetUnResolvedList() EObjectList[U] {
	return ToObjectList(l.delegate.GetUnResolvedList(), l.convertTo, l.convertFrom)
}
