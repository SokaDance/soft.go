// *****************************************************************************
// Copyright(c) 2021 MASA Group
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// *****************************************************************************

package ecore

func ToAny[T any](t T) any { return t }

func FromAny[T any](a any) T {
	if a == nil {
		var zero T
		return zero
	}
	return a.(T)
}

func ToAnyCollection[T any](c ECollection[T]) any {
	switch c := any(c).(type) {
	case ECollection[any]:
		return c
	case ECollection[T]:
		return ToCollection(c, ToAny[T], FromAny[T])
	default:
		return nil
	}
}

func FromAnyCollection[T any](v any) ECollection[T] {
	switch c := v.(type) {
	case ECollection[T]:
		return c
	case ECollection[any]:
		return ToCollection(c, FromAny[T], ToAny[T])
	default:
		return nil
	}
}

func ToAnyList[T any](l EList[T]) any {
	switch c := any(l).(type) {
	case EList[any]:
		return c
	case EObjectList[T]:
		return ToObjectList(c, ToAny[T], FromAny[T])
	case ENotifyingList[T]:
		return ToNotifyingList(c, ToAny[T], FromAny[T])
	case EList[T]:
		return ToList(c, ToAny[T], FromAny[T])
	default:
		return nil
	}
}

func FromAnyList[T any](v any) EList[T] {
	switch c := v.(type) {
	case EList[T]:
		return c
	case EObjectList[any]:
		return ToObjectList(c, FromAny[T], ToAny[T])
	case ENotifyingList[any]:
		return ToNotifyingList(c, FromAny[T], ToAny[T])
	case EList[any]:
		return ToList(c, FromAny[T], ToAny[T])
	default:
		return nil
	}
}

func ToAnyMap[K comparable, V any](m EMap[K, V]) any {
	switch c := any(m).(type) {
	case EMap[any, any]:
		return c
	case EMap[K, V]:
		return ToMap(m, ToAny[K], FromAny[K], ToAny[V], FromAny[V])
	default:
		return nil
	}
}

func FromAnyMap[K comparable, V any](v any) EMap[K, V] {
	switch m := v.(type) {
	case EMap[K, V]:
		return m
	case EMap[any, any]:
		return ToMap(m, FromAny[K], ToAny[K], FromAny[V], ToAny[V])
	default:
		return nil
	}
}
