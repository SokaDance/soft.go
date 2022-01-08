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

func TestMockECollection_Add(t *testing.T) {
	l := &MockECollection[int]{}
	l.On("Add", 1).Once().Return(true)
	l.On("Add", 1).Once().Return(func(int) bool {
		return true
	})
	assert.True(t, l.Add(1))
	assert.True(t, l.Add(1))
	mock.AssertExpectationsForObjects(t, l)
}

func TestMockECollection_Remove(t *testing.T) {
	l := &MockECollection[int]{}
	l.On("Remove", 1).Once().Return(true)
	l.On("Remove", 1).Once().Return(func(int) bool {
		return true
	})
	assert.True(t, l.Remove(1))
	assert.True(t, l.Remove(1))
	mock.AssertExpectationsForObjects(t, l)
}

func TestMockECollection_AddAll(t *testing.T) {
	l := &MockECollection[int]{}
	c := &MockECollection[int]{}
	l.On("AddAll", c).Once().Return(true)
	l.On("AddAll", c).Once().Return(func(ECollection[int]) bool {
		return true
	})
	assert.True(t, l.AddAll(c))
	assert.True(t, l.AddAll(c))
	mock.AssertExpectationsForObjects(t, l)
}

func TestMockECollection_RemoveAll(t *testing.T) {
	l := &MockECollection[int]{}
	c := &MockECollection[int]{}
	l.On("RemoveAll", c).Once().Return(true)
	l.On("RemoveAll", c).Once().Return(func(ECollection[int]) bool {
		return true
	})
	assert.True(t, l.RemoveAll(c))
	assert.True(t, l.RemoveAll(c))
	mock.AssertExpectationsForObjects(t, l, c)
}

func TestMockECollection_RetainAll(t *testing.T) {
	l := &MockECollection[int]{}
	c := &MockECollection[int]{}
	l.On("RetainAll", c).Once().Return(true)
	l.On("RetainAll", c).Once().Return(func(ECollection[int]) bool {
		return true
	})
	assert.True(t, l.RetainAll(c))
	assert.True(t, l.RetainAll(c))
	mock.AssertExpectationsForObjects(t, l, c)
}

func TestMockECollection_Size(t *testing.T) {
	l := &MockECollection[int]{}
	l.On("Size").Once().Return(0)
	l.On("Size").Once().Return(func() int {
		return 1
	})
	assert.Equal(t, 0, l.Size())
	assert.Equal(t, 1, l.Size())
	mock.AssertExpectationsForObjects(t, l)
}

func TestMockECollection_Clear(t *testing.T) {
	l := &MockECollection[int]{}
	l.On("Clear").Once()
	l.Clear()
	mock.AssertExpectationsForObjects(t, l)
}

func TestMockECollection_Empty(t *testing.T) {
	l := &MockECollection[int]{}
	l.On("Empty").Once().Return(true)
	l.On("Empty").Once().Return(func() bool {
		return false
	})
	assert.True(t, l.Empty())
	assert.False(t, l.Empty())
	mock.AssertExpectationsForObjects(t, l)
}

func TestMockECollection_Iterator(t *testing.T) {
	l := &MockECollection[int]{}
	it := &MockEIterator[int]{}
	l.On("Iterator").Once().Return(it)
	l.On("Iterator").Once().Return(func() EIterator[int] {
		return it
	})
	assert.Equal(t, it, l.Iterator())
	assert.Equal(t, it, l.Iterator())
	mock.AssertExpectationsForObjects(t, l, it)
}

func TestMockECollection_ToArray(t *testing.T) {
	l := &MockECollection[int]{}
	r := []int{}
	l.On("ToArray").Once().Return(r)
	l.On("ToArray").Once().Return(func() []int {
		return r
	})
	assert.Equal(t, r, l.ToArray())
	assert.Equal(t, r, l.ToArray())
	mock.AssertExpectationsForObjects(t, l)
}

func TestMockECollection_Contains(t *testing.T) {
	l := &MockECollection[int]{}
	l.On("Contains", 1).Once().Return(false)
	l.On("Contains", 2).Once().Return(func(int) bool {
		return true
	})
	assert.False(t, l.Contains(1))
	assert.True(t, l.Contains(2))
	mock.AssertExpectationsForObjects(t, l)
}

