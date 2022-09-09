package stream

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestStream_IsParallel(t *testing.T) {
	s := OfSlice([]any{})
	assert.False(t, s.IsParallel())
	s = s.Parallel()
	require.NotNil(t, s)
	assert.True(t, s.IsParallel())
	s = s.Sequential()
	require.NotNil(t, s)
	assert.False(t, s.IsParallel())
}

func TestStream_AnyMatch_Sequential(t *testing.T) {
	assert.True(t, OfSlice([]any{1, 2, 3}).AnyMatch(func(a any) bool { return a == 2 }))
	assert.False(t, OfSlice([]any{1, 2, 3}).AnyMatch(func(a any) bool { return a == 0 }))
}

// func TestStream_AnyMatch_Parallel(t *testing.T) {
// 	assert.True(t, OfSlice([]any{1, 2, 3}).Parallel().AnyMatch(func(a any) bool { return a == 2 }))
// 	assert.False(t, OfSlice([]any{1, 2, 3}).Parallel().AnyMatch(func(a any) bool { return a == 0 }))
// }

func TestStream_AllMatch_Sequential(t *testing.T) {
	assert.True(t, OfSlice([]any{2, 2, 2}).AllMatch(func(a any) bool { return a == 2 }))
	assert.False(t, OfSlice([]any{2, 2, 2}).AllMatch(func(a any) bool { return a == 3 }))
}

func TestStream_NoneMatch_Sequential(t *testing.T) {
	assert.True(t, OfSlice([]any{2, 2, 2}).NoneMatch(func(a any) bool { return a == 3 }))
	assert.False(t, OfSlice([]any{2, 2, 2}).NoneMatch(func(a any) bool { return a == 2 }))
}
