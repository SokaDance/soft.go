// *****************************************************************************
// Copyright(c) 2021 MASA Group
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// *****************************************************************************

package ecore

type BasicEMap struct {
	EList
	interfaces any
	mapData    map[any]any
}

type eMapEntryFactory interface {
	newEntry(key any, value any) EMapEntry
}

type basicEMapList struct {
	BasicEList
	m *BasicEMap
}

func newBasicEMapList(m *BasicEMap) *basicEMapList {
	l := new(basicEMapList)
	l.interfaces = l
	l.data = []any{}
	l.isUnique = true
	l.m = m
	return l
}

func (ml *basicEMapList) DidAdd(index int, elem any) {
	entry := elem.(EMapEntry)
	ml.m.doAdd(entry)
}

func (ml *basicEMapList) DidSet(index int, newElem any, oldElem any) {
	newEntry := newElem.(EMapEntry)
	oldEntry := oldElem.(EMapEntry)
	ml.m.doRemove(oldEntry)
	ml.m.doAdd(newEntry)
}

func (ml *basicEMapList) DidRemove(index int, oldElem any) {
	oldEntry := oldElem.(EMapEntry)
	ml.m.doRemove(oldEntry)
}

func (ml *basicEMapList) DidClear(oldObjects []any) {
	ml.m.doClear()
}

func NewBasicEMap() *BasicEMap {
	basicEMap := &BasicEMap{}
	basicEMap.EList = newBasicEMapList(basicEMap)
	basicEMap.interfaces = basicEMap
	return basicEMap
}

func (m *BasicEMap) asEMapEntryFactory() eMapEntryFactory {
	return m.interfaces.(eMapEntryFactory)
}

func (m *BasicEMap) getEntryForKey(key any) EMapEntry {
	for it := m.Iterator(); it.HasNext(); {
		e := it.Next().(EMapEntry)
		if e.GetKey() == key {
			return e
		}
	}
	return nil
}

func (m *BasicEMap) GetValue(key any) any {
	m.initDataMap()
	return m.mapData[key]
}

func (m *BasicEMap) Put(key any, value any) {
	if e := m.getEntryForKey(key); e != nil {
		e.SetValue(value)

		if m.mapData != nil {
			m.mapData[key] = value
		}
	} else {
		m.Add(m.asEMapEntryFactory().newEntry(key, value))
	}
}

type eMapEntryImpl struct {
	key   any
	value any
}

func (e *eMapEntryImpl) GetKey() any {
	return e.key
}

func (e *eMapEntryImpl) SetKey(key any) {
	e.key = key
}

func (e *eMapEntryImpl) GetValue() any {
	return e.value
}

func (e *eMapEntryImpl) SetValue(value any) {
	e.value = value
}

func (m *BasicEMap) newEntry(key any, value any) EMapEntry {
	return &eMapEntryImpl{key: key, value: value}
}

func (m *BasicEMap) RemoveKey(key any) any {
	if e := m.getEntryForKey(key); e != nil {
		m.Remove(e)
		return e.GetValue()
	}
	return nil
}

func (m *BasicEMap) ContainsValue(value any) bool {
	for it := m.Iterator(); it.HasNext(); {
		e := it.Next().(EMapEntry)
		if e.GetValue() == value {
			return true
		}
	}
	return false
}

func (m *BasicEMap) ContainsKey(key any) bool {
	m.initDataMap()
	_, ok := m.mapData[key]
	return ok
}

func (m *BasicEMap) ToMap() map[any]any {
	m.initDataMap()
	return m.mapData
}

func (m *BasicEMap) initDataMap() {
	if m.mapData == nil {
		m.mapData = map[any]any{}
		for itEntry := m.Iterator(); itEntry.HasNext(); {
			entry := itEntry.Next().(EMapEntry)
			m.mapData[entry.GetKey()] = entry.GetValue()
		}
	}
}

func (m *BasicEMap) doAdd(e EMapEntry) {
	if m.mapData != nil {
		m.mapData[e.GetKey()] = e.GetValue()
	}
}

func (m *BasicEMap) doRemove(e EMapEntry) {
	delete(m.mapData, e.GetKey())
}

func (m *BasicEMap) doClear() {
	m.mapData = nil
}
