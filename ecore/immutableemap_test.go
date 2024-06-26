package ecore

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestImmutableEMap_Constructor(t *testing.T) {
	m := NewImmutableEMap(nil)
	require.NotNil(t, m)
}

func TestImmutableEMap_GetValue(t *testing.T) {
	mockEntry := NewMockEMapEntry(t)
	mockEntry.EXPECT().GetKey().Return(1)
	mockEntry.EXPECT().GetValue().Return(2)
	im := NewImmutableEMap([]any{mockEntry})
	require.NotNil(t, m)
	require.Equal(t, 2, im.GetValue(1))
	require.Equal(t, nil, im.GetValue(3))
}

func TestImmutableEMap_Put(t *testing.T) {
	m := NewImmutableEMap(nil)
	require.NotNil(t, m)
	assert.Panics(t, func() {
		m.Put(1, 2)
	})
}

func TestImmutableEMap_RemoveKey(t *testing.T) {
	m := NewImmutableEMap(nil)
	require.NotNil(t, m)
	assert.Panics(t, func() {
		m.RemoveKey(1)
	})
}

func TestImmutableEMap_ContainsValue(t *testing.T) {
	mockEntry := NewMockEMapEntry(t)
	mockEntry.EXPECT().GetKey().Return(1)
	mockEntry.EXPECT().GetValue().Return(2)
	im := NewImmutableEMap([]any{mockEntry})
	require.NotNil(t, m)
	require.NotNil(t, m)
	require.True(t, im.ContainsValue(2))
	require.False(t, im.ContainsValue(3))
}

func TestImmutableEMap_ContainsKey(t *testing.T) {
	mockEntry := NewMockEMapEntry(t)
	mockEntry.EXPECT().GetKey().Return(1)
	mockEntry.EXPECT().GetValue().Return(2)
	im := NewImmutableEMap([]any{mockEntry})
	require.NotNil(t, m)
	require.True(t, im.ContainsKey(1))
	require.False(t, im.ContainsKey(2))
}

func TestImmutableEMap_ToMap(t *testing.T) {
	mockEntry := NewMockEMapEntry(t)
	mockEntry.EXPECT().GetKey().Return(1)
	mockEntry.EXPECT().GetValue().Return(2)
	im := NewImmutableEMap([]any{mockEntry})
	require.NotNil(t, m)
	require.Equal(t, map[any]any{1: 2}, im.ToMap())
}
