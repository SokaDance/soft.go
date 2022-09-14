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

func TestStream_ToSlice(t *testing.T) {
	assert.Equal(t, []any{}, OfSlice(nil).ToSlice())
	assert.Equal(t, []any{}, OfSlice([]any{}).ToSlice())
	assert.Equal(t, []any{1, 2, 3}, OfSlice([]any{1, 2, 3}).ToSlice())
}

func TestStream_Map_Sequential(t *testing.T) {
	assert.Equal(t, []any{-1, -2, -3}, OfSlice([]any{1, 2, 3}).Map(func(a any) any { return -a.(int) }).ToSlice())
}

func TestStream_FlatMap_Sequential(t *testing.T) {
	assert.Equal(t, []any{1, 2, 3, 2, 4, 6, 3, 6, 9}, OfSlice([]any{1, 2, 3}).FlatMap(func(a any) Stream {
		i := a.(int)
		return OfSlice([]any{i * 1, i * 2, i * 3})
	}).ToSlice())
}

func TestStream_Filter_Sequential(t *testing.T) {
	assert.Equal(t, 0, OfSlice([]any{1, 3}).Filter(func(a any) bool { return a == 2 }).Count())
	assert.Equal(t, 1, OfSlice([]any{1, 2, 3}).Filter(func(a any) bool { return a == 2 }).Count())
	assert.Equal(t, 2, OfSlice([]any{1, 2, 3, 2}).Filter(func(a any) bool { return a == 2 }).Count())
}

func TestStream_Skip(t *testing.T) {
	assert.Equal(t, []any{1, 2, 3}, OfSlice([]any{1, 2, 3}).Skip(0).ToSlice())
	assert.Equal(t, []any{2, 3}, OfSlice([]any{1, 2, 3}).Skip(1).ToSlice())
	assert.Equal(t, []any{3}, OfSlice([]any{1, 2, 3}).Skip(2).ToSlice())
	assert.Equal(t, []any{}, OfSlice([]any{1, 2, 3}).Skip(3).ToSlice())
	assert.Equal(t, []any{}, OfSlice([]any{1, 2, 3}).Skip(4).ToSlice())
}

func TestStream_Limit(t *testing.T) {
	assert.Equal(t, []any{}, OfSlice([]any{1, 2, 3}).Limit(0).ToSlice())
	assert.Equal(t, []any{1}, OfSlice([]any{1, 2, 3}).Limit(1).ToSlice())
	assert.Equal(t, []any{1, 2}, OfSlice([]any{1, 2, 3}).Limit(2).ToSlice())
	assert.Equal(t, []any{1, 2, 3}, OfSlice([]any{1, 2, 3}).Limit(3).ToSlice())
	assert.Equal(t, []any{1, 2, 3}, OfSlice([]any{1, 2, 3}).Limit(4).ToSlice())
}

func TestStream_Distinct(t *testing.T) {
	assert.Equal(t, []any{1, 2, 3}, OfSlice([]any{1, 2, 3, 2, 3}).Distinct().ToSlice())
}

func TestStream_Peek(t *testing.T) {
	result := 0
	assert.Equal(t, []any{1, 2, 3}, OfSlice([]any{1, 2, 3}).Peek(func(a any) { result += a.(int) }).ToSlice())
	assert.Equal(t, 6, result)
}

func TestStream_Sorted(t *testing.T) {
	assert.Equal(t, []any{1, 2, 3}, OfSlice([]any{3, 1, 2}).Sorted(func(a1, a2 any) bool { return a1.(int) < a2.(int) }).ToSlice())
}

func TestStream_ForEach_Sequential(t *testing.T) {
	res := 0
	OfSlice([]any{1, 2, 3}).ForEach(func(a any) { res += a.(int) })
	assert.Equal(t, 6, res)
}

func TestStream_Count_Sequential(t *testing.T) {
	assert.Equal(t, 3, OfSlice([]any{1, 2, 3}).Count())
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
	assert.Equal(t, 1, OfSlice([]any{2, 3, 1}).Min(func(a1, a2 any) int { return a1.(int) - a2.(int) }).ElseZero())
}

func TestStream_Max_Sequential(t *testing.T) {
	assert.Equal(t, 3, OfSlice([]any{2, 3, 1}).Max(func(a1, a2 any) int { return a1.(int) - a2.(int) }).ElseZero())
}

func TestStream_Reduce(t *testing.T) {
	assert.False(t, OfSlice([]any{}).Reduce(func(a1, a2 any) any { return a1.(int) + a2.(int) }).IsPresent())
	assert.Equal(t, 6, OfSlice([]any{1, 2, 3}).Reduce(func(a1, a2 any) any { return a1.(int) + a2.(int) }).ElseZero())
}

func TestStream_ReduceWith(t *testing.T) {
	assert.Equal(t, 7, OfSlice([]any{1, 2, 3}).ReduceWith(1, func(a1, a2 any) any { return a1.(int) + a2.(int) }))
}

func TestStream_Concat(t *testing.T) {
	assert.Equal(t, []any{1, 2, 3, 4}, Concat(OfSlice([]any{1, 2}), OfSlice([]any{3, 4})).ToSlice())
}

func TestStream_FindFirst(t *testing.T) {
	assert.False(t, OfSlice([]any{}).FindFirst().IsPresent())
	assert.Equal(t, 1, OfSlice([]any{1, 2, 3}).FindFirst().ElseZero())
}

func TestStream_FindAny(t *testing.T) {
	assert.False(t, OfSlice([]any{}).FindAny().IsPresent())
	assert.Equal(t, 1, OfSlice([]any{1, 2, 3}).FindAny().ElseZero())
}

func TestStream_Iterator_Source(t *testing.T) {
	result := []any{}
	it := OfSlice([]any{1, 2, 3}).Iterator()
	for it.TryAdvance(func(a any) { result = append(result, a) }) {
	}
	assert.Equal(t, []any{1, 2, 3}, result)
}

func TestStream_Iterator_Chained(t *testing.T) {
	result := []any{}
	it := OfSlice([]any{1, 2, 3}).Map(func(a any) any { return -a.(int) }).Iterator()
	for it.TryAdvance(func(a any) { result = append(result, a) }) {
	}
	assert.Equal(t, []any{-1, -2, -3}, result)
}
