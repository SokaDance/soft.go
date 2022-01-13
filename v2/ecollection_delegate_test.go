package ecore

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCollectionDelegate_Add(t *testing.T) {
	l := &MockEList[string]{}
	d := ToCollection[string,any](l)
	l.On("Add","s").Once().Return(true)
	assert.True(t, d.Add("s"))
}

func TestCollectionDelegate_AddAll(t *testing.T) {
	l := &MockEList[string]{}
	o := &MockEList[any]{}
	d := ToCollection[string,any](l)
	l.On("AddAll",mock.Anything).Once().Return(true)
	assert.True(t, d.AddAll(o))
	mock.AssertExpectationsForObjects(t, l , o)
}

func TestCollectionDelegate_Remove(t *testing.T) {
	l := &MockEList[string]{}
	d := ToCollection[string,any](l)
	l.On("Remove","s").Once().Return(true)
	assert.True(t, d.Remove("s"))
}

func TestCollectionDelegate_RemoveAll(t *testing.T) {
	l := &MockEList[string]{}
	o := &MockEList[any]{}
	d := ToCollection[string,any](l)
	l.On("RemoveAll",mock.Anything).Once().Return(true)
	assert.True(t, d.RemoveAll(o))
}

func TestCollectionDelegate_RetainAll(t *testing.T) {
	l := &MockEList[string]{}
	o := &MockEList[any]{}
	d := ToCollection[string,any](l)
	l.On("RetainAll",mock.Anything).Once().Return(true)
	assert.True(t, d.RetainAll(o))
}

func TestCollectionDelegate_Size(t *testing.T) {
	l := &MockEList[string]{}
	d := ToCollection[string,any](l)
	l.On("Size").Once().Return(1)
	assert.Equal(t, 1, d.Size())
}

func TestCollectionDelegate_Clear(t *testing.T) {
	l := &MockEList[string]{}
	d := ToCollection[string,any](l)
	l.On("Clear").Once().Return(true)
	d.Clear()
}

func TestCollectionDelegate_Empty(t *testing.T) {
	l := &MockEList[string]{}
	d := ToCollection[string,any](l)
	l.On("Empty").Once().Return(true)
	assert.True(t, d.Empty())
}

func TestCollectionDelegate_Contains(t *testing.T) {
	l := &MockEList[string]{}
	v := "s"
	d := ToCollection[string,any](l)
	l.On("Contains", v).Once().Return(true)
	assert.True(t, d.Contains(v))
}
