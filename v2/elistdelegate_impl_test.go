package ecore

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func toAnyList[T any](l EList[T]) EList[any] {
	return ToList(l, ToAny[T], FromAny[T])
}

func TestListDelegate_Add(t *testing.T) {
	l := &MockEList[string]{}
	d := toAnyList[string](l)
	l.On("Add", "s").Once().Return(true)
	assert.True(t, d.Add("s"))
}

func TestListDelegate_Insert(t *testing.T) {
	l := &MockEList[string]{}
	d := toAnyList[string](l)
	l.On("Insert", 1, "s").Once().Return(true)
	assert.True(t, d.Insert(1, "s"))
}

func TestListDelegate_InsertAll(t *testing.T) {
	l := &MockEList[string]{}
	c := &MockECollection[any]{}
	d := toAnyList[string](l)
	l.On("InsertAll", 1, mock.Anything).Once().Return(true)
	assert.True(t, d.InsertAll(1, c))
}

func TestListDelegate_RemoveAt(t *testing.T) {
	l := &MockEList[string]{}
	d := toAnyList[string](l)
	l.On("RemoveAt", 1).Once().Return("s")
	assert.Equal(t, "s", d.RemoveAt(1))
}

func TestListDelegate_MoveObject(t *testing.T) {
	l := &MockEList[string]{}
	d := toAnyList[string](l)
	l.On("MoveObject", 1, "s").Once().Return("s")
	d.MoveObject(1, "s")
}

func TestListDelegate_MoveIndex(t *testing.T) {
	l := &MockEList[string]{}
	d := toAnyList[string](l)
	l.On("MoveIndex", 1, 2).Once().Return("s")
	assert.Equal(t, "s", d.MoveIndex(1, 2))
}

func TestListDelegate_IndexOf(t *testing.T) {
	l := &MockEList[string]{}
	d := toAnyList[string](l)
	l.On("IndexOf", "s").Once().Return(1)
	assert.Equal(t, 1, d.IndexOf("s"))
}

func TestListDelegate_Get(t *testing.T) {
	l := &MockEList[string]{}
	d := toAnyList[string](l)
	l.On("Get", 0).Once().Return("s")
	assert.Equal(t, "s", d.Get(0))
}

func TestListDelegate_Set(t *testing.T) {
	l := &MockEList[string]{}
	d := toAnyList[string](l)
	l.On("Set", 0, "s").Once().Return("")
	assert.Equal(t, "", d.Set(0, "s"))
}

func TestToListWithDelegate(t *testing.T) {
	l := &MockEList[string]{}
	l2 := toAnyList[string](l)
	l3 := FromAnyList[string](l2)
	assert.Equal(t, l, l3)
}