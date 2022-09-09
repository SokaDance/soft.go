package stream

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAnyMatch(t *testing.T) {
	assert.True(t, OfSlice([]any{1, 2, 3}).AnyMatch(func(a any) bool { return a == 2 }))
	assert.False(t, OfSlice([]any{1, 2, 3}).AnyMatch(func(a any) bool { return a == 0 }))
}

func TestAllMatch(t *testing.T) {
	assert.True(t, OfSlice([]any{2, 2, 2}).AllMatch(func(a any) bool { return a == 2 }))
	assert.False(t, OfSlice([]any{2, 2, 2}).AllMatch(func(a any) bool { return a == 3 }))
}

func TestNoneMatch(t *testing.T) {
	assert.True(t, OfSlice([]any{2, 2, 2}).NoneMatch(func(a any) bool { return a == 3 }))
	assert.False(t, OfSlice([]any{2, 2, 2}).NoneMatch(func(a any) bool { return a == 2 }))
}
