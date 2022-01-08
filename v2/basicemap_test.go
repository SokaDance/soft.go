// *****************************************************************************
// Copyright(c) 2021 MASA Group
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// *****************************************************************************

package ecore

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestBasicEMap_Constructor(t *testing.T) {
	m := NewBasicEMap[string,any]()
	assert.NotNil(t, m)

	var mp EMap[string,any] = m
	assert.NotNil(t, mp)

	var ml EList[EMapEntry[string,any]] = m
	assert.NotNil(t, ml)
}

func TestBasicEMap_Put(t *testing.T) {
	m := NewBasicEMap[int,string]()
	m.Put(2, "2")
	assert.Equal(t, map[int]string{2: "2"}, m.ToMap())
}

func TestBasicEMap_GetValue(t *testing.T) {
	m := NewBasicEMap[int,string]()
	assert.Equal(t, "", m.GetValue(2))

	m.Put(2, "2")
	assert.Equal(t, "2", m.GetValue(2))
}

func TestBasicEMap_RemoveKey(t *testing.T) {
	m := NewBasicEMap[int,string]()
	m.Put(2, "2")

	assert.Equal(t, "2", m.RemoveKey(2))
	assert.Equal(t, "", m.GetValue(2))
	assert.Equal(t, "", m.RemoveKey(2))
}

func TestBasicEMap_PutOverwrite(t *testing.T) {
	m := NewBasicEMap[int,string]()
	assert.Equal(t, "", m.GetValue(2))
	m.Put(2, "3")
	assert.Equal(t, "3", m.GetValue(2))
	m.Put(2, "2")
	assert.Equal(t, "2", m.GetValue(2))
	assert.Equalf(t, 1, m.Size(), "Don't store old cell.")
}

func TestBasicEMap_ContainsKey(t *testing.T) {
	m := NewBasicEMap[int,string]()
	assert.False(t, m.ContainsKey(2))

	m.Put(2, "2")
	assert.True(t, m.ContainsKey(2))

	m.RemoveKey(2)
	assert.False(t, m.ContainsKey(2))
}

func TestBasicEMap_ContainsValue(t *testing.T) {
	m := NewBasicEMap[int,string]()
	assert.False(t, m.ContainsValue("2"))

	m.Put(2, "2")
	assert.True(t, m.ContainsValue("2"))

	m.RemoveKey(2)
	assert.False(t, m.ContainsValue("2"))
}

func TestBasicEMap_AddEntry(t *testing.T) {
	m := NewBasicEMap[int,string]()
	mockEntry := &MockEMapEntry[int,string]{}
	mockEntry.On("GetKey").Once().Return(2)
	mockEntry.On("GetValue").Once().Return("2")
	m.Add(mockEntry)
	assert.Equal(t, map[int]string{2: "2"}, m.ToMap())
	mock.AssertExpectationsForObjects(t, mockEntry)
}

func TestBasicEMap_SetEntry(t *testing.T) {
	m := NewBasicEMap[int,string]()
	mockEntry := &MockEMapEntry[int,string]{}
	mockEntry.On("GetKey").Once().Return(2)
	mockEntry.On("GetValue").Once().Return("2")
	m.Add(mockEntry)
	mock.AssertExpectationsForObjects(t, mockEntry)

	mockOtherEntry := &MockEMapEntry[int,string]{}
	mockEntry.On("GetKey").Once().Return(2)
	mockOtherEntry.On("GetKey").Once().Return(3)
	mockOtherEntry.On("GetValue").Once().Return("3")
	m.Set(0, mockOtherEntry)
	assert.Equal(t, map[int]string{3: "3"}, m.ToMap())
	mock.AssertExpectationsForObjects(t, mockEntry, mockOtherEntry)
}

func TestBasicEMap_RemoveEntry(t *testing.T) {
	m := NewBasicEMap[int,string]()
	mockEntry1 := &MockEMapEntry[int,string]{}
	mockEntry1.On("GetKey").Once().Return(2)
	mockEntry1.On("GetValue").Once().Return("2")
	mockEntry2 := &MockEMapEntry[int,string]{}
	mockEntry2.On("GetKey").Once().Return(3)
	mockEntry2.On("GetValue").Once().Return("3")
	m.Add(mockEntry1)
	m.Add(mockEntry2)
	assert.Equal(t, map[int]string{2: "2", 3: "3"}, m.ToMap())
	mock.AssertExpectationsForObjects(t, mockEntry1, mockEntry2)

	mockEntry1.On("GetKey").Once().Return(2)
	m.RemoveAt(0)
	assert.Equal(t, map[int]string{3: "3"}, m.ToMap())
	mock.AssertExpectationsForObjects(t, mockEntry1, mockEntry2)
}

func TestBasicEMap_Clear(t *testing.T) {
	m := NewBasicEMap[int,string]()
	mockEntry1 := &MockEMapEntry[int,string]{}
	mockEntry1.On("GetKey").Once().Return(2)
	mockEntry1.On("GetValue").Once().Return("2")
	mockEntry2 := &MockEMapEntry[int,string]{}
	mockEntry2.On("GetKey").Once().Return(3)
	mockEntry2.On("GetValue").Once().Return("3")
	m.Add(mockEntry1)
	m.Add(mockEntry2)
	mock.AssertExpectationsForObjects(t, mockEntry1, mockEntry2)

	m.Clear()
	assert.Equal(t, map[int]string{}, m.ToMap())
	assert.Equal(t, []EMapEntry[int,string]{}, m.ToArray())
	mock.AssertExpectationsForObjects(t, mockEntry1, mockEntry2)
}

func TestBasicEMap_UpdateEntry(t *testing.T) {
	m := NewBasicEMap[int,string]()
	m.Put(2, "2")
	e := m.Get(0)
	e.SetKey(3)
	e.SetValue("3")
}
