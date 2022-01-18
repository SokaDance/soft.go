package ecore

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestToAnyArray(t *testing.T) {
	c := &MockECollection[int]{}
	it := &MockEIterator[int]{}
	c.On("Size").Once().Return(2)
	c.On("Iterator").Once().Return(it)
	it.On("HasNext").Twice().Return(true)
	it.On("HasNext").Once().Return(false)
	it.On("Next").Once().Return(1)
	it.On("Next").Once().Return(2)
	aa := ToAnyArray[int](c)
	assert.Equal(t, []any{1, 2}, aa)
	mock.AssertExpectationsForObjects(t, c, it)
}

func TestCollectionDelegate_Add(t *testing.T) {
	l := &MockECollection[string]{}
	d := ToAnyCollection[string](l)
	l.On("Add", "s").Once().Return(true)
	assert.True(t, d.Add("s"))
}

func TestCollectionDelegate_AddInvalid(t *testing.T) {
	l := &MockECollection[string]{}
	d := ToAnyCollection[string](l)
	assert.Panics(t, func() { d.Add(1) })
}

func TestCollectionDelegate_AddAll(t *testing.T) {
	l := &MockECollection[string]{}
	o := &MockECollection[any]{}
	d := ToAnyCollection[string](l)
	l.On("AddAll", mock.Anything).Once().Return(true)
	assert.True(t, d.AddAll(o))
	mock.AssertExpectationsForObjects(t, l, o)
}

func TestCollectionDelegate_Remove(t *testing.T) {
	l := &MockECollection[string]{}
	d := ToAnyCollection[string](l)
	l.On("Remove", "s").Once().Return(true)
	assert.True(t, d.Remove("s"))
}

func TestCollectionDelegate_RemoveAll(t *testing.T) {
	l := &MockECollection[string]{}
	o := &MockECollection[any]{}
	d := ToAnyCollection[string](l)
	l.On("RemoveAll", mock.Anything).Once().Return(true)
	assert.True(t, d.RemoveAll(o))
}

func TestCollectionDelegate_RetainAll(t *testing.T) {
	l := &MockECollection[string]{}
	o := &MockECollection[any]{}
	d := ToAnyCollection[string](l)
	l.On("RetainAll", mock.Anything).Once().Return(true)
	assert.True(t, d.RetainAll(o))
}

func TestCollectionDelegate_Size(t *testing.T) {
	l := &MockECollection[string]{}
	d := ToAnyCollection[string](l)
	l.On("Size").Once().Return(1)
	assert.Equal(t, 1, d.Size())
}

func TestCollectionDelegate_Clear(t *testing.T) {
	l := &MockECollection[string]{}
	d := ToAnyCollection[string](l)
	l.On("Clear").Once().Return(true)
	d.Clear()
}

func TestCollectionDelegate_Empty(t *testing.T) {
	l := &MockECollection[string]{}
	d := ToAnyCollection[string](l)
	l.On("Empty").Once().Return(true)
	assert.True(t, d.Empty())
}

func TestCollectionDelegate_Contains(t *testing.T) {
	l := &MockECollection[string]{}
	v := "s"
	d := ToAnyCollection[string](l)
	l.On("Contains", v).Once().Return(true)
	assert.True(t, d.Contains(v))
}

func TestToCollectionWithDelegate(t *testing.T) {
	l := &MockECollection[string]{}
	l2 := ToAnyCollection[string](l)
	l3 := FromAnyCollection[string](l2)
	assert.Equal(t, l, l3)
}
