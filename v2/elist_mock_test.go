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

func TestMockEList_Add(t *testing.T) {
	l := &MockEList[int]{}
	l.On("Add", 1).Once().Return(true)
	l.On("Add", 1).Once().Return(func(int) bool {
		return true
	})
	assert.True(t, l.Add(1))
	assert.True(t, l.Add(1))
	mock.AssertExpectationsForObjects(t, l)
}

func TestMockEList_Remove(t *testing.T) {
	l := &MockEList[int]{}
	l.On("Remove", 1).Once().Return(true)
	l.On("Remove", 1).Once().Return(func(int) bool {
		return true
	})
	assert.True(t, l.Remove(1))
	assert.True(t, l.Remove(1))
	mock.AssertExpectationsForObjects(t, l)
}

func TestMockEList_RemoveAt(t *testing.T) {
	l := &MockEList[int]{}
	l.On("RemoveAt", 1).Once().Return(1)
	l.On("RemoveAt", 1).Once().Return(func(int) int {
		return 2
	})
	assert.Equal(t, 1, l.RemoveAt(1))
	assert.Equal(t, 2, l.RemoveAt(1))
	mock.AssertExpectationsForObjects(t, l)
}

func TestMockEList_Insert(t *testing.T) {
	l := &MockEList[int]{}
	l.On("Insert", 1, 2).Once().Return(true)
	l.On("Insert", 1, 2).Once().Return(func(int, int) bool {
		return true
	})
	assert.True(t, l.Insert(1, 2))
	assert.True(t, l.Insert(1, 2))
	mock.AssertExpectationsForObjects(t, l)
}

func TestMockEList_MoveObject(t *testing.T) {
	l := &MockEList[int]{}
	l.On("MoveObject", 1, 2).Once()
	l.MoveObject(1, 2)
	mock.AssertExpectationsForObjects(t, l)
}

func TestMockEList_MoveIndex(t *testing.T) {
	l := &MockEList[int]{}
	l.On("MoveIndex", 1, 2).Once().Return(3)
	l.On("MoveIndex", 1, 2).Once().Return(func(int, int) int {
		return 3
	})
	assert.Equal(t, 3, l.MoveIndex(1, 2))
	assert.Equal(t, 3, l.MoveIndex(1, 2))
	mock.AssertExpectationsForObjects(t, l)
}

func TestMockEList_Get(t *testing.T) {
	l := &MockEList[int]{}
	l.On("Get", 1).Once().Return(1)
	l.On("Get", 1).Once().Return(func(int) int {
		return 0
	})
	assert.Equal(t, 1, l.Get(1))
	assert.Equal(t, 0, l.Get(1))
	mock.AssertExpectationsForObjects(t, l)
}

func TestMockEList_Set(t *testing.T) {
	l := &MockEList[int]{}
	l.On("Set", 1, 2).Once().Return(3)
	l.On("Set", 1, 2).Once().Return(func(int, int) int {
		return 4
	})
	assert.Equal(t, 3, l.Set(1, 2))
	assert.Equal(t, 4, l.Set(1, 2))
	mock.AssertExpectationsForObjects(t, l)
}

func TestMockEList_AddAll(t *testing.T) {
	l := &MockEList[int]{}
	c := &MockEList[int]{}
	l.On("AddAll", c).Once().Return(true)
	l.On("AddAll", c).Once().Return(func(ECollection[int]) bool {
		return true
	})
	assert.True(t, l.AddAll(c))
	assert.True(t, l.AddAll(c))
	mock.AssertExpectationsForObjects(t, l)
}

func TestMockEList_InsertAll(t *testing.T) {
	l := &MockEList[int]{}
	c := &MockEList[int]{}
	l.On("InsertAll", 0, c).Once().Return(true)
	l.On("InsertAll", 0, c).Once().Return(func(int, ECollection[int]) bool {
		return true
	})
	assert.True(t, l.InsertAll(0, c))
	assert.True(t, l.InsertAll(0, c))
	mock.AssertExpectationsForObjects(t, l, c)
}

func TestMockEList_RemoveAll(t *testing.T) {
	l := &MockEList[int]{}
	c := &MockEList[int]{}
	l.On("RemoveAll", c).Once().Return(true)
	l.On("RemoveAll", c).Once().Return(func(ECollection[int]) bool {
		return true
	})
	assert.True(t, l.RemoveAll(c))
	assert.True(t, l.RemoveAll(c))
	mock.AssertExpectationsForObjects(t, l, c)
}

func TestMockEList_RetainAll(t *testing.T) {
	l := &MockEList[int]{}
	c := &MockEList[int]{}
	l.On("RetainAll", c).Once().Return(true)
	l.On("RetainAll", c).Once().Return(func(ECollection[int]) bool {
		return true
	})
	assert.True(t, l.RetainAll(c))
	assert.True(t, l.RetainAll(c))
	mock.AssertExpectationsForObjects(t, l, c)
}

func TestMockEList_Size(t *testing.T) {
	l := &MockEList[int]{}
	l.On("Size").Once().Return(0)
	l.On("Size").Once().Return(func() int {
		return 1
	})
	assert.Equal(t, 0, l.Size())
	assert.Equal(t, 1, l.Size())
	mock.AssertExpectationsForObjects(t, l)
}

func TestMockEList_Clear(t *testing.T) {
	l := &MockEList[int]{}
	l.On("Clear").Once()
	l.Clear()
	mock.AssertExpectationsForObjects(t, l)
}

func TestMockEList_Empty(t *testing.T) {
	l := &MockEList[int]{}
	l.On("Empty").Once().Return(true)
	l.On("Empty").Once().Return(func() bool {
		return false
	})
	assert.True(t, l.Empty())
	assert.False(t, l.Empty())
	mock.AssertExpectationsForObjects(t, l)
}

func TestMockEList_Iterator(t *testing.T) {
	l := &MockEList[int]{}
	it := &MockIterator[int]{}
	l.On("Iterator").Once().Return(it)
	l.On("Iterator").Once().Return(func() EIterator[int] {
		return it
	})
	assert.Equal(t, it, l.Iterator())
	assert.Equal(t, it, l.Iterator())
	mock.AssertExpectationsForObjects(t, l, it)
}

func TestMockEList_ToArray(t *testing.T) {
	l := &MockEList[int]{}
	r := []int{}
	l.On("ToArray").Once().Return(r)
	l.On("ToArray").Once().Return(func() []int {
		return r
	})
	assert.Equal(t, r, l.ToArray())
	assert.Equal(t, r, l.ToArray())
	mock.AssertExpectationsForObjects(t, l)
}

func TestMockEList_Contains(t *testing.T) {
	l := &MockEList[int]{}
	l.On("Contains", 1).Once().Return(false)
	l.On("Contains", 2).Once().Return(func(int) bool {
		return true
	})
	assert.False(t, l.Contains(1))
	assert.True(t, l.Contains(2))
	mock.AssertExpectationsForObjects(t, l)
}

func TestMockEList_IndexOf(t *testing.T) {
	l := &MockEList[int]{}
	l.On("IndexOf", 1).Once().Return(0)
	l.On("IndexOf", 2).Once().Return(func(int) int {
		return 1
	})
	assert.Equal(t, 0, l.IndexOf(1))
	assert.Equal(t, 1, l.IndexOf(2))
	mock.AssertExpectationsForObjects(t, l)
}
