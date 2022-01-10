// *****************************************************************************
// Copyright(c) 2021 MASA Group
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// *****************************************************************************

package ecore

type BasicEMap[K comparable, V comparable] struct {
	*basicEMapList[K, V]
	mapData map[K]V
}

type basicEMapList[K comparable, V comparable] struct {
	basicEList[EMapEntry[K, V]]
	m *BasicEMap[K, V]
}

func newBasicEMapList[K comparable, V comparable](m *BasicEMap[K, V]) *basicEMapList[K, V] {
	l := new(basicEMapList[K, V])
	l.interfaces = l
	l.data = []EMapEntry[K, V]{}
	l.isUnique = true
	l.m = m
	return l
}

func (ml *basicEMapList[K, V]) didAdd(index int, entry EMapEntry[K, V]) {
	ml.m.mapData[entry.GetKey()] = entry.GetValue()
}

func (ml *basicEMapList[K, V]) didSet(index int, newEntry EMapEntry[K, V], oldEntry EMapEntry[K, V]) {
	delete(ml.m.mapData, oldEntry.GetKey())
	ml.m.mapData[newEntry.GetKey()] = newEntry.GetValue()
}

func (ml *basicEMapList[K, V]) didRemove(index int, oldEntry EMapEntry[K, V]) {
	delete(ml.m.mapData, oldEntry.GetKey())
}

func (ml *basicEMapList[K, V]) didClear(oldObjects []EMapEntry[K, V]) {
	ml.m.mapData = make(map[K]V)
}

func NewBasicEMap[K comparable, V comparable]() *BasicEMap[K, V] {
	basicEMap := &BasicEMap[K, V]{}
	basicEMap.Initialize()
	return basicEMap
}

func (m *BasicEMap[K, V]) Initialize() {
	m.basicEMapList = newBasicEMapList(m)
	m.mapData = make(map[K]V)
}

func (m *BasicEMap[K, V]) getEntryForKey(key K) EMapEntry[K, V] {
	for it := m.Iterator(); it.HasNext(); {
		e := it.Next()
		if e.GetKey() == key {
			return e
		}
	}
	return nil
}

func (m *BasicEMap[K, V]) GetValue(key K) V {
	return m.mapData[key]
}

func (m *BasicEMap[K, V]) Put(key K, value V) {
	m.mapData[key] = value
	if e := m.getEntryForKey(key); e != nil {
		e.SetValue(value)
	} else {
		m.Add(m.newEntry(key, value))
	}
}

type eMapEntryImpl[K comparable, V any] struct {
	key   K
	value V
}

func (e *eMapEntryImpl[K, V]) GetKey() K {
	return e.key
}

func (e *eMapEntryImpl[K, V]) SetKey(key K) {
	e.key = key
}

func (e *eMapEntryImpl[K, V]) GetValue() V {
	return e.value
}

func (e *eMapEntryImpl[K, V]) SetValue(value V) {
	e.value = value
}

func (m *BasicEMap[K, V]) newEntry(key K, value V) EMapEntry[K, V] {
	return &eMapEntryImpl[K, V]{key: key, value: value}
}

func (m *BasicEMap[K, V]) RemoveKey(key K) V {
	// remove from map data
	delete(m.mapData, key)

	// remove from list
	if e := m.getEntryForKey(key); e != nil {
		m.Remove(e)
		return e.GetValue()
	}
	var zero V
	return zero
}

func (m *BasicEMap[K, V]) ContainsKey(key K) bool {
	_, ok := m.mapData[key]
	return ok
}

func (m *BasicEMap[K, V]) ContainsValue(value V) bool {
	for _, v := range m.mapData {
		if v == value {
			return true
		}
	}
	return false
}

func (m *BasicEMap[K, V]) ToMap() map[K]V {
	return m.mapData
}
