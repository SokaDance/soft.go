package ecore

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMapDelegate_Add(t *testing.T) {
	m := &MockEMap[string, string]{}
	d := ToAnyMap[string, string](m)
	e := &MockEMapEntry{}
	m.On("Add", e).Once().Return(true)
	assert.True(t, d.Add(e))
}

func TestMapDelegate_AddAll(t *testing.T) {
	m := &MockEMap[string, string]{}
	d := ToAnyMap[string, string](m)
	c := &MockECollection[any]{}
	m.On("AddAll", c).Once().Return(true)
	assert.True(t, d.AddAll(c))
}

func TestMapDelegate_Contains(t *testing.T) {
	m := &MockEMap[string, string]{}
	d := ToAnyMap[string, string](m)
	e := &MockEMapEntry{}
	m.On("Contains", e).Once().Return(true)
	assert.True(t, d.Contains(e))
}

func TestMapDelegate_Clear(t *testing.T) {
	m := &MockEMap[string, string]{}
	d := ToAnyMap[string, string](m)
	m.On("Clear").Once()
	d.Clear()
}

func TestMapDelegate_Empty(t *testing.T) {
	m := &MockEMap[string, string]{}
	d := ToAnyMap[string, string](m)
	m.On("Empty").Once().Return(true)
	assert.True(t, d.Empty())
}

func TestMapDelegate_Size(t *testing.T) {
	m := &MockEMap[string, string]{}
	d := ToAnyMap[string, string](m)
	m.On("Size").Once().Return(1)
	assert.Equal(t, 1, d.Size())
}

func TestMapDelegate_Get(t *testing.T) {
	m := &MockEMap[string, string]{}
	d := ToAnyMap[string, string](m)
	e := &MockEMapEntry{}
	m.On("Get", 1).Once().Return(e)
	assert.Equal(t, e, d.Get(1))
}

func TestMapDelegate_Set(t *testing.T) {
	m := &MockEMap[string, string]{}
	d := ToAnyMap[string, string](m)
	e := &MockEMapEntry{}
	o := &MockEMapEntry{}
	m.On("Set", 1, e).Once().Return(o)
	assert.Equal(t, o, d.Set(1, e))
}

func TestMapDelegate_IndexOf(t *testing.T) {
	m := &MockEMap[string, string]{}
	d := ToAnyMap[string, string](m)
	e := &MockEMapEntry{}
	m.On("IndexOf", e).Once().Return(1)
	assert.Equal(t, 1, d.IndexOf(e))
}

func TestMapDelegate_Insert(t *testing.T) {
	m := &MockEMap[string, string]{}
	d := ToAnyMap[string, string](m)
	e := &MockEMapEntry{}
	m.On("Insert", 1, e).Once().Return(true)
	assert.True(t, d.Insert(1, e))
}

func TestMapDelegate_InsertAll(t *testing.T) {
	m := &MockEMap[string, string]{}
	d := ToAnyMap[string, string](m)
	c := &MockECollection[any]{}
	m.On("InsertAll", 1, c).Once().Return(true)
	assert.True(t, d.InsertAll(1, c))
}

func TestMapDelegate_Iterator(t *testing.T) {
	m := &MockEMap[string, string]{}
	d := ToAnyMap[string, string](m)
	c := &MockEIterator[any]{}
	m.On("Iterator").Once().Return(c)
	assert.Equal(t, c, d.Iterator())
}

func TestMapDelegate_MoveIndex(t *testing.T) {
	m := &MockEMap[string, string]{}
	d := ToAnyMap[string, string](m)
	e := &MockEMapEntry{}
	m.On("MoveIndex", 1, 2).Once().Return(e)
	assert.Equal(t, e, d.MoveIndex(1, 2))
}

func TestMapDelegate_MoveObject(t *testing.T) {
	m := &MockEMap[string, string]{}
	d := ToAnyMap[string, string](m)
	e := &MockEMapEntry{}
	m.On("MoveObject", 1, e).Once()
	d.MoveObject(1, e)
}

func TestMapDelegate_Remove(t *testing.T) {
	m := &MockEMap[string, string]{}
	d := ToAnyMap[string, string](m)
	e := &MockEMapEntry{}
	m.On("Remove", e).Once().Return(true)
	assert.True(t, d.Remove(e))
}

func TestMapDelegate_RemoveAll(t *testing.T) {
	m := &MockEMap[string, string]{}
	d := ToAnyMap[string, string](m)
	c := &MockECollection[any]{}
	m.On("RemoveAll", c).Once().Return(true)
	assert.True(t, d.RemoveAll(c))
}

func TestMapDelegate_RetainAll(t *testing.T) {
	m := &MockEMap[string, string]{}
	d := ToAnyMap[string, string](m)
	c := &MockECollection[any]{}
	m.On("RetainAll", c).Once().Return(true)
	assert.True(t, d.RetainAll(c))
}

func TestMapDelegate_RemoveAt(t *testing.T) {
	m := &MockEMap[string, string]{}
	d := ToAnyMap[string, string](m)
	e := &MockEMapEntry{}
	m.On("RemoveAt", 1).Once().Return(e)
	assert.Equal(t, e, d.RemoveAt(1))
}

func TestMapDelegate_ToArray(t *testing.T) {
	m := &MockEMap[string, string]{}
	d := ToAnyMap[string, string](m)
	a := []any{&MockEMapEntry{}}
	m.On("ToArray").Once().Return(a)
	assert.Equal(t, a, d.ToArray())
}

func TestMapDelegate_GetValue(t *testing.T) {
	m := &MockEMap[string, string]{}
	d := ToAnyMap[string, string](m)
	{
		m.On("GetValue", "key").Once().Return("value", true)
		v, ok := d.GetValue("key")
		assert.Equal(t, "value", v)
		assert.True(t, ok)
	}
	{
		m.On("GetValue", "key").Once().Return("", false)
		v, ok := d.GetValue("key")
		assert.Equal(t, nil, v)
		assert.False(t, ok)
	}
}

func TestMapDelegate_Put(t *testing.T) {
	m := &MockEMap[string, string]{}
	d := ToAnyMap[string, string](m)
	m.On("Put", "key", "value").Once()
	d.Put("key", "value")
}

func TestMapDelegate_RemoveKey(t *testing.T) {
	m := &MockEMap[string, string]{}
	d := ToAnyMap[string, string](m)
	m.On("RemoveKey", "key").Once().Return("value")
	assert.Equal(t, "value", d.RemoveKey("key"))
}

func TestMapDelegate_ContainsValue(t *testing.T) {
	m := &MockEMap[string, string]{}
	d := ToAnyMap[string, string](m)
	m.On("ContainsValue", "value").Once().Return(true)
	assert.True(t, d.ContainsValue("value"))
}

func TestMapDelegate_ContainsKey(t *testing.T) {
	m := &MockEMap[string, string]{}
	d := ToAnyMap[string, string](m)
	m.On("ContainsKey", "key").Once().Return(true)
	assert.True(t, d.ContainsKey("key"))
}

func TestMapDelegate_ToMap(t *testing.T) {
	m := &MockEMap[string, string]{}
	d := ToAnyMap[string, string](m)
	m.On("ToMap").Once().Return(map[any]any{"key": "value"})
	assert.Equal(t, map[any]any{"key": "value"}, d.ToMap())
}

func TestMapDelegate_WithDelegate(t *testing.T) {
	m1 := &MockEMap[string, string]{}
	m2 := ToAnyMap[string, string](m1)
	m3 := FromAnyMap[string, string](m2)
	assert.Equal(t, m1, m3)
}
