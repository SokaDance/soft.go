package ecore

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFromAny(t *testing.T) {
	v := 1
	a := any(v)
	assert.Equal(t, 1, FromAny[int](a))
	assert.Equal(t, a, FromAny[any](a))
}
