package stream

import (
	"testing"

	"github.com/stretchr/testify/assert"
	mock "github.com/stretchr/testify/mock"
)

func TestIteratorDelegate_TryAdvance(t *testing.T) {
	mockIterator := NewMockIterator(t)
	it := NewDelegatingIterator(func() Iterator { return mockIterator })
	mockIterator.On("TryAdvance", mock.AnythingOfType("func(interface {})")).Once().Return(true)
	assert.True(t, it.TryAdvance(func(any) {}))
}

func TestIteratorDelegate_TrySplit(t *testing.T) {
	mockIterator := NewMockIterator(t)
	mockResult := NewMockIterator(t)
	it := NewDelegatingIterator(func() Iterator { return mockIterator })
	mockIterator.On("TrySplit").Once().Return(mockResult)
	assert.Equal(t, mockResult, it.TrySplit())
}

func TestIteratorDelegate_EstimateSize(t *testing.T) {
	mockIterator := NewMockIterator(t)
	it := NewDelegatingIterator(func() Iterator { return mockIterator })
	mockIterator.On("EstimateSize").Once().Return(1)
	assert.Equal(t, 1, it.EstimateSize())
}

func TestIteratorDelegate_ForEachremaining(t *testing.T) {
	mockIterator := NewMockIterator(t)
	it := NewDelegatingIterator(func() Iterator { return mockIterator })
	mockIterator.On("ForEachRemaining", mock.AnythingOfType("func(interface {})")).Once()
	it.ForEachRemaining(func(any) {})
}
