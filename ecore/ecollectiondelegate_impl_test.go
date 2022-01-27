package ecore

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func toAnyCollection[T any](c ECollection[T]) ECollection[any] {
	return ToCollection(c, ToAny[T], FromAny[T])
}

func fromAnyCollection[T any](c ECollection[any]) ECollection[T] {
	return ToCollection(c, FromAny[T], ToAny[T])
}

func TestCollectionDelegate_Add(t *testing.T) {
	l := &MockECollection[string]{}
	d := toAnyCollection[string](l)
	l.On("Add", "s").Once().Return(true)
	assert.True(t, d.Add("s"))
}

func TestCollectionDelegate_AddInvalid(t *testing.T) {
	l := &MockECollection[string]{}
	d := toAnyCollection[string](l)
	assert.Panics(t, func() { d.Add(1) })
}

func TestCollectionDelegate_AddAll(t *testing.T) {
	l := &MockECollection[string]{}
	o := &MockECollection[any]{}
	d := toAnyCollection[string](l)
	l.On("AddAll", mock.Anything).Once().Return(true)
	assert.True(t, d.AddAll(o))
	mock.AssertExpectationsForObjects(t, l, o)
}

func TestCollectionDelegate_Remove(t *testing.T) {
	l := &MockECollection[string]{}
	d := toAnyCollection[string](l)
	l.On("Remove", "s").Once().Return(true)
	assert.True(t, d.Remove("s"))
}

func TestCollectionDelegate_RemoveAll(t *testing.T) {
	l := &MockECollection[string]{}
	o := &MockECollection[any]{}
	d := toAnyCollection[string](l)
	l.On("RemoveAll", mock.Anything).Once().Return(true)
	assert.True(t, d.RemoveAll(o))
}

func TestCollectionDelegate_RetainAll(t *testing.T) {
	l := &MockECollection[string]{}
	o := &MockECollection[any]{}
	d := toAnyCollection[string](l)
	l.On("RetainAll", mock.Anything).Once().Return(true)
	assert.True(t, d.RetainAll(o))
}

func TestCollectionDelegate_Size(t *testing.T) {
	l := &MockECollection[string]{}
	d := toAnyCollection[string](l)
	l.On("Size").Once().Return(1)
	assert.Equal(t, 1, d.Size())
}

func TestCollectionDelegate_Clear(t *testing.T) {
	l := &MockECollection[string]{}
	d := toAnyCollection[string](l)
	l.On("Clear").Once().Return(true)
	d.Clear()
}

func TestCollectionDelegate_Empty(t *testing.T) {
	l := &MockECollection[string]{}
	d := toAnyCollection[string](l)
	l.On("Empty").Once().Return(true)
	assert.True(t, d.Empty())
}

func TestCollectionDelegate_Contains(t *testing.T) {
	l := &MockECollection[string]{}
	v := "s"
	d := toAnyCollection[string](l)
	l.On("Contains", v).Once().Return(true)
	assert.True(t, d.Contains(v))
}

func TestToCollectionWithDelegate(t *testing.T) {
	l := &MockECollection[string]{}
	l2 := toAnyCollection[string](l)
	l3 := fromAnyCollection[string](l2)
	assert.Equal(t, l, l3)
}
