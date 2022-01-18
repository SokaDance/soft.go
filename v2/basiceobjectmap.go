// *****************************************************************************
// Copyright(c) 2021 MASA Group
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// *****************************************************************************

package ecore

type BasicEObjectMap[K comparable, V comparable] struct {
	BasicEMap[K, V]
	entryClass EClass
}

func NewBasicEObjectMap[K comparable, V comparable](entryClass EClass) *BasicEObjectMap[K, V] {
	basicEObjectMap := &BasicEObjectMap[K, V]{}
	basicEObjectMap.Initialize()
	basicEObjectMap.entryClass = entryClass
	return basicEObjectMap
}

func (m *BasicEObjectMap[K, V]) Put(key K, value V) {
	m.mapData[key] = value
	m.Add(m.newEntry(key, value))
}

func (m *BasicEObjectMap[K, V]) newEntry(key K, value V) EMapEntry[K, V] {
	eFactory := m.entryClass.GetEPackage().GetEFactoryInstance()
	eEntry := eFactory.Create(m.entryClass).(EMapEntry[K, V])
	eEntry.SetKey(key)
	eEntry.SetValue(value)
	return eEntry
}
