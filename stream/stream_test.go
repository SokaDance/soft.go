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

func TestStream_ForEach_Sequential(t *testing.T) {
	res := 0
	OfSlice([]any{1, 2, 3}).ForEach(func(a any) { res += a.(int) })
	assert.Equal(t, 6, res)
}

func TestStream_Count_Sequential(t *testing.T) {
	assert.Equal(t, 3, OfSlice([]any{1, 2, 3}).Count())
}

func TestStream_Filter_Sequential(t *testing.T) {
	assert.Equal(t, 0, OfSlice([]any{1, 3}).Filter(func(a any) bool { return a == 2 }).Count())
	assert.Equal(t, 1, OfSlice([]any{1, 2, 3}).Filter(func(a any) bool { return a == 2 }).Count())
	assert.Equal(t, 2, OfSlice([]any{1, 2, 3, 2}).Filter(func(a any) bool { return a == 2 }).Count())
}

func TestStream_AnyMatch_Sequential(t *testing.T) {
	assert.True(t, OfSlice([]any{1, 2, 3}).AnyMatch(func(a any) bool { return a == 2 }))
	assert.False(t, OfSlice([]any{1, 2, 3}).AnyMatch(func(a any) bool { return a == 0 }))
}

func TestStream_AllMatch_Sequential(t *testing.T) {
	assert.True(t, OfSlice([]any{2, 2, 2}).AllMatch(func(a any) bool { return a == 2 }))
	assert.False(t, OfSlice([]any{2, 2, 2}).AllMatch(func(a any) bool { return a == 3 }))
}

func TestStream_NoneMatch_Sequential(t *testing.T) {
	assert.True(t, OfSlice([]any{2, 2, 2}).NoneMatch(func(a any) bool { return a == 3 }))
	assert.False(t, OfSlice([]any{2, 2, 2}).NoneMatch(func(a any) bool { return a == 2 }))
}

func TestStream_Min_Sequential(t *testing.T) {
	assert.Equal(t, 1, OfSlice([]any{3, 2, 1}).Min(func(a1, a2 any) int { return a1.(int) - a2.(int) }).ElseZero())
}

func TestStream_Max_Sequential(t *testing.T) {
	assert.Equal(t, 3, OfSlice([]any{3, 2, 1}).Max(func(a1, a2 any) int { return a1.(int) - a2.(int) }).ElseZero())
}

func TestStream_Reduce(t *testing.T) {
	assert.Equal(t, 6, OfSlice([]any{1, 2, 3}).Reduce(func(a1, a2 any) any { return a1.(int) + a2.(int) }).ElseZero())
}

func TestStream_ReduceWith(t *testing.T) {
	assert.Equal(t, 7, OfSlice([]any{1, 2, 3}).ReduceWith(1, func(a1, a2 any) any { return a1.(int) + a2.(int) }))
}
