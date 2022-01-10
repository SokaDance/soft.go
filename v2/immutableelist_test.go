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
)

func TestEmptyImmutableAdd(t *testing.T) {
	a := &emptyImmutableEList[int]{}
	assert.Panics(t, func() { a.Add(1) })
}

func TestEmptyImmutableAddAll(t *testing.T) {
	a := &emptyImmutableEList[int]{}
	b := NewImmutableEList([]int{1, 2})
	assert.Panics(t, func() { a.AddAll(b) })
}

func TestEmptyImmutableInsert(t *testing.T) {
	a := &emptyImmutableEList[int]{}
	assert.Panics(t, func() { a.Insert(0, 1) })
}

func TestEmptyImmutableInsertAll(t *testing.T) {
	a := &emptyImmutableEList[int]{}
	b := NewImmutableEList([]int{1, 2})
	assert.Panics(t, func() { a.InsertAll(0, b) })
}

func TestEmptyImmutableMoveObject(t *testing.T) {
	a := &emptyImmutableEList[int]{}
	assert.Panics(t, func() { a.MoveObject(0, 1) })
}

func TestEmptyImmutableMove(t *testing.T) {
	a := NewImmutableEList[int](nil)
	assert.Panics(t, func() { a.MoveIndex(0, 1) })
}

func TestEmptyImmutableGet(t *testing.T) {
	a := &emptyImmutableEList[int]{}
	assert.Panics(t, func() { a.Get(0) })
}

func TestEmptyImmutableSet(t *testing.T) {
	a := &emptyImmutableEList[int]{}
	assert.Panics(t, func() { a.Set(0, 1) })
}

func TestEmptyImmutableRemoveAt(t *testing.T) {
	a := &emptyImmutableEList[int]{}
	assert.Panics(t, func() { a.RemoveAt(0) })
}

func TestEmptyImmutableRemove(t *testing.T) {
	a := &emptyImmutableEList[int]{}
	assert.Panics(t, func() { a.Remove(0) })
}

func TestEmptyImmutableSize(t *testing.T) {
	a := &emptyImmutableEList[int]{}
	assert.Equal(t, 0, a.Size())
}

func TestEmptyImmutableEmpty(t *testing.T) {
	a := &emptyImmutableEList[int]{}
	assert.True(t, a.Empty())
}

func TestEmptyImmutableClear(t *testing.T) {
	a := &emptyImmutableEList[int]{}
	assert.Panics(t, func() { a.Clear() })
}

func TestEmptyImmutableRemoveAll(t *testing.T) {
	a := &emptyImmutableEList[int]{}
	assert.Panics(t, func() { a.RemoveAll(nil) })
}

func TestEmptyImmutableContains(t *testing.T) {
	a := &emptyImmutableEList[int]{}
	assert.False(t, a.Contains(3))
}

func TestEmptyImmutableIndexOf(t *testing.T) {
	a := &emptyImmutableEList[int]{}
	assert.Equal(t, -1, a.IndexOf(3))
}

func TestEmptyImmutableToArray(t *testing.T) {
	a := &emptyImmutableEList[int]{}
	assert.Equal(t, []int{}, a.ToArray())
}

func TestEmptyImmutableGetUnResolved(t *testing.T) {
	l := &emptyImmutableEList[int]{}
	assert.Equal(t, l, l.GetUnResolvedList())
}

func TestImmutableEListAdd(t *testing.T) {
	a := NewImmutableEList[int](nil)
	assert.Panics(t, func() { a.Add(1) })
}

func TestImmutableEListAddAll(t *testing.T) {
	a := NewImmutableEList[int](nil)
	b := NewImmutableEList([]int{1, 2})
	assert.Panics(t, func() { a.AddAll(b) })
}

func TestImmutableEListInsert(t *testing.T) {
	a := NewImmutableEList[int](nil)
	assert.Panics(t, func() { a.Insert(0, 1) })
}

func TestImmutableEListInsertAll(t *testing.T) {
	a := NewImmutableEList[int](nil)
	b := NewImmutableEList([]int{1, 2})
	assert.Panics(t, func() { a.InsertAll(0, b) })
}

func TestImmutableEListMoveObject(t *testing.T) {
	a := NewImmutableEList[int](nil)
	assert.Panics(t, func() { a.MoveObject(0, 1) })
}

func TestImmutableEListMove(t *testing.T) {
	a := NewImmutableEList[int](nil)
	assert.Panics(t, func() { a.MoveIndex(0, 1) })
}

func TestImmutableEListGet(t *testing.T) {
	a := NewImmutableEList([]int{1, 2})
	assert.Equal(t, 1, a.Get(0))
	assert.Equal(t, 2, a.Get(1))
	assert.Panics(t, func() { a.Get(2) })
}

func TestImmutableEListSet(t *testing.T) {
	a := NewImmutableEList[int](nil)
	assert.Panics(t, func() { a.Set(0, 1) })
}

func TestImmutableEListRemoveAt(t *testing.T) {
	a := NewImmutableEList[int](nil)
	assert.Panics(t, func() { a.RemoveAt(0) })
}

func TestImmutableEListRemove(t *testing.T) {
	a := NewImmutableEList[int](nil)
	assert.Panics(t, func() { a.Remove(0) })
}

func TestImmutableEListSize(t *testing.T) {
	a := NewImmutableEList([]int{1, 2})
	assert.Equal(t, 2, a.Size())
}

func TestImmutableEListEmpty(t *testing.T) {
	a := NewImmutableEList[int](nil)
	assert.True(t, a.Empty())
}

func TestImmutableEListClear(t *testing.T) {
	a := NewImmutableEList[int](nil)
	assert.Panics(t, func() { a.Clear() })
}

func TestImmutableEListRemoveAll(t *testing.T) {
	a := NewImmutableEList[int](nil)
	assert.Panics(t, func() { a.RemoveAll(nil) })
}

func TestImmutableEListContains(t *testing.T) {
	a := NewImmutableEList([]int{1, 2})
	assert.True(t, a.Contains(2))
	assert.False(t, a.Contains(3))
	b := NewImmutableEList[int](nil)
	assert.False(t, b.Contains(3))
}

func TestImmutableEListIterate(t *testing.T) {
	var iDatas []int
	iDatas = append(iDatas, 0, 2, 4, 6)
	arr := NewImmutableEList(iDatas)
	i := 0
	for it := arr.Iterator(); it.HasNext(); {
		assert.Equal(t, it.Next(), i)
		i += 2
	}
	assert.Equal(t, 8, i)
}

func TestImmutableEListGetUnResolved(t *testing.T) {
	l := NewImmutableEList([]int{})
	assert.Equal(t, l, l.GetUnResolvedList())
}
