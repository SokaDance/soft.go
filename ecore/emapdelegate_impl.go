// *****************************************************************************
// Copyright(c) 2021 MASA Group
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// *****************************************************************************

package ecore

func ToMap[KO comparable, KT comparable, VO any, VT any](
	delegate EMap[KO, VO],
	convertKeyTo func(KO) KT,
	convertKeyFrom func(KT) KO,
	convertValueTo func(VO) VT,
	convertValueFrom func(VT) VO) EMap[KT, VT] {
	if dp, isDelegateProvider := delegate.(EDelegateProvider); isDelegateProvider {
		if m, isMap := dp.GetDelegate().(EMap[KT, VT]); isMap {
			return m
		}
	}
	return &eMapDelegateImpl[KO, KT, VO, VT]{
		delegate:         delegate,
		convertKeyTo:     convertKeyTo,
		convertKeyFrom:   convertKeyFrom,
		convertValueTo:   convertValueTo,
		convertValueFrom: convertValueFrom,
	}
}

type eMapDelegateImpl[KO comparable, KT comparable, VO any, VT any] struct {
	delegate         EMap[KO, VO]
	convertKeyTo     func(KO) KT
	convertKeyFrom   func(KT) KO
	convertValueTo   func(VO) VT
	convertValueFrom func(VT) VO
}

func (m *eMapDelegateImpl[KO, KT, VO, VT]) GetDelegate() any {
	return m.delegate
}

func (m *eMapDelegateImpl[KO, KT, VO, VT]) Add(a any) bool {
	return m.delegate.Add(a)
}

func (m *eMapDelegateImpl[KO, KT, VO, VT]) AddAll(c ECollection[any]) bool {
	return m.delegate.AddAll(c)
}

func (m *eMapDelegateImpl[KO, KT, VO, VT]) Contains(a any) bool {
	return m.delegate.Contains(a)
}

func (m *eMapDelegateImpl[KO, KT, VO, VT]) Clear() {
	m.delegate.Clear()
}

func (m *eMapDelegateImpl[KO, KT, VO, VT]) Empty() bool {
	return m.delegate.Empty()
}

func (m *eMapDelegateImpl[KO, KT, VO, VT]) Size() int {
	return m.delegate.Size()
}

func (m *eMapDelegateImpl[KO, KT, VO, VT]) Get(i int) any {
	return m.delegate.Get(i)
}

func (m *eMapDelegateImpl[KO, KT, VO, VT]) Set(i int, a any) any {
	return m.delegate.Set(i, a)
}

func (m *eMapDelegateImpl[KO, KT, VO, VT]) IndexOf(a any) int {
	return m.delegate.IndexOf(a)
}

func (m *eMapDelegateImpl[KO, KT, VO, VT]) Insert(index int, a any) bool {
	return m.delegate.Insert(index, a)
}

func (m *eMapDelegateImpl[KO, KT, VO, VT]) InsertAll(index int, c ECollection[any]) bool {
	return m.delegate.InsertAll(index, c)
}

func (m *eMapDelegateImpl[KO, KT, VO, VT]) Iterator() EIterator[any] {
	return m.delegate.Iterator()
}

func (m *eMapDelegateImpl[KO, KT, VO, VT]) MoveIndex(oldIndex, newIndex int) any {
	return m.delegate.MoveIndex(oldIndex, newIndex)
}

func (m *eMapDelegateImpl[KO, KT, VO, VT]) MoveObject(index int, a any) {
	m.delegate.MoveObject(index, a)
}

func (m *eMapDelegateImpl[KO, KT, VO, VT]) Remove(a any) bool {
	return m.delegate.Remove(a)
}

func (m *eMapDelegateImpl[KO, KT, VO, VT]) RemoveAll(c ECollection[any]) bool {
	return m.delegate.RemoveAll(c)
}

func (m *eMapDelegateImpl[KO, KT, VO, VT]) RetainAll(c ECollection[any]) bool {
	return m.delegate.RetainAll(c)
}

func (m *eMapDelegateImpl[KO, KT, VO, VT]) RemoveAt(index int) any {
	return m.delegate.RemoveAt(index)
}

func (m *eMapDelegateImpl[KO, KT, VO, VT]) ToArray() []any {
	return m.delegate.ToArray()
}

func (m *eMapDelegateImpl[KO, KT, VO, VT]) GetValue(key KT) (VT, bool) {
	if v, ok := m.delegate.GetValue(m.convertKeyFrom(key)); ok {
		return m.convertValueTo(v), ok
	} else {
		var zero VT
		return zero, false
	}
}

func (m *eMapDelegateImpl[KO, KT, VO, VT]) Put(key KT, value VT) {
	m.delegate.Put(m.convertKeyFrom(key), m.convertValueFrom(value))
}

func (m *eMapDelegateImpl[KO, KT, VO, VT]) RemoveKey(key KT) VT {
	return m.convertValueTo(m.delegate.RemoveKey(m.convertKeyFrom(key)))
}

func (m *eMapDelegateImpl[KO, KT, VO, VT]) ContainsValue(value VT) bool {
	return m.delegate.ContainsValue(m.convertValueFrom(value))
}

func (m *eMapDelegateImpl[KO, KT, VO, VT]) ContainsKey(key KT) bool {
	return m.delegate.ContainsKey(m.convertKeyFrom(key))
}

func (m *eMapDelegateImpl[KO, KT, VO, VT]) ToMap() map[KT]VT {
	d := m.delegate.ToMap()
	r := map[KT]VT{}
	for k, v := range d {
		r[m.convertKeyTo(k)] = m.convertValueTo(v)
	}
	return r
}
