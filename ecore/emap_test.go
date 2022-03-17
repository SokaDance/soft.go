package ecore

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func assertEMapValueGet[K any, V any](t *testing.T, m EMap[K, V], k K, v V, exists bool) {
	rv, rexists := m.GetValue(k)
	assert.Equal(t, v, rv)
	assert.Equal(t, exists, rexists)
}
