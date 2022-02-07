package ecore

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestEMapEntryImpl_Constructor(t *testing.T) {
	e := newMapEntry(1, 2)
	assert.NotNil(t, e)
}

func TestEMapEntryImpl_Accessors(t *testing.T) {
	e := newMapEntry(1, 2)
	require.NotNil(t, e)
	assert.Equal(t, 1, e.GetKey())
	assert.Equal(t, 2, e.GetValue())

	e.SetKey(2)
	assert.Equal(t, 2, e.GetKey())

	e.SetValue(2)
	assert.Equal(t, 2, e.GetValue())
}
