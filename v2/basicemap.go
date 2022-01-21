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
	EList[any]
	mapList *basicEMapList[K, V]
	mapData map[K]V
}

type basicEMapList[K comparable, V comparable] struct {
	basicEList[EMapEntry]
	m *BasicEMap[K, V]
}

func newBasicEMapList[K comparable, V comparable](m *BasicEMap[K, V]) *basicEMapList[K, V] {
	l := new(basicEMapList[K, V])
	l.interfaces = l
	l.data = []EMapEntry{}
	l.isUnique = true
	l.m = m
	return l
}

func (ml *basicEMapList[K, V]) didAdd(index int, entry EMapEntry) {
	ml.m.mapData[entry.GetKey().(K)] = entry.GetValue().(V)
}

func (ml *basicEMapList[K, V]) didSet(index int, newEntry EMapEntry, oldEntry EMapEntry) {
	delete(ml.m.mapData, oldEntry.GetKey().(K))
	ml.m.mapData[newEntry.GetKey().(K)] = newEntry.GetValue().(V)
}

func (ml *basicEMapList[K, V]) didRemove(index int, oldEntry EMapEntry) {
	delete(ml.m.mapData, oldEntry.GetKey().(K))
}

func (ml *basicEMapList[K, V]) didClear(oldObjects []EMapEntry) {
	ml.m.mapData = make(map[K]V)
}

func NewBasicEMap[K comparable, V comparable]() *BasicEMap[K, V] {
	basicEMap := &BasicEMap[K, V]{}
	basicEMap.Initialize()
	return basicEMap
}

func (m *BasicEMap[K, V]) Initialize() {
	m.mapList = newBasicEMapList(m)
	m.mapData = make(map[K]V)
	m.EList = ToAnyList[EMapEntry](m.mapList)
}

func (m *BasicEMap[K, V]) getEntryForKey(key K) EMapEntry {
	for it := m.mapList.basicEList.Iterator(); it.HasNext(); {
		e := it.Next()
		if e.GetKey().(K) == key {
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
		m.Add(newMapEntry(key, value))
	}
}

func (m *BasicEMap[K, V]) RemoveKey(key K) V {
	// remove from map data
	delete(m.mapData, key)

	// remove from list
	if e := m.getEntryForKey(key); e != nil {
		m.Remove(e)
		return e.GetValue().(V)
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
