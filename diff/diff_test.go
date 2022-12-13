package diff

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func equals[T comparable](a, b T) bool {
	return a == b
}

func TestDiff(t *testing.T) {
	for _, tt := range []struct {
		oldArray        []int
		newArray        []int
		entries         []Entry[int]
		mockVisitorInit func(mock *MockDiffVisitor[int])
	}{
		// Add
		{[]int{1, 2, 3, 4}, []int{5, 1, 2, 3, 4}, []Entry[int]{{Index: 0, Element: 5, IsAddition: true}}, func(m *MockDiffVisitor[int]) { m.On("HandleAdd", 0, 5).Once() }},
		{[]int{1, 2, 3, 4}, []int{1, 2, 5, 3, 4}, []Entry[int]{{Index: 2, Element: 5, IsAddition: true}}, func(m *MockDiffVisitor[int]) { m.On("HandleAdd", 2, 5).Once() }},
		{[]int{1, 2, 3, 4}, []int{1, 2, 3, 4, 5}, []Entry[int]{{Index: 4, Element: 5, IsAddition: true}}, func(m *MockDiffVisitor[int]) { m.On("HandleAdd", 4, 5).Once() }},

		// Remove
		{[]int{1, 2, 3, 4}, []int{2, 3, 4}, []Entry[int]{{Index: 0, Element: 1, IsAddition: false}}, func(m *MockDiffVisitor[int]) { m.On("HandleRemove", 0, 1).Once() }},
		{[]int{1, 2, 3, 4}, []int{1, 2, 4}, []Entry[int]{{Index: 2, Element: 3, IsAddition: false}}, func(m *MockDiffVisitor[int]) { m.On("HandleRemove", 2, 3).Once() }},
		{[]int{1, 2, 3, 4}, []int{1, 2, 3}, []Entry[int]{{Index: 3, Element: 4, IsAddition: false}}, func(m *MockDiffVisitor[int]) { m.On("HandleRemove", 3, 4).Once() }},

		// Move
		{[]int{1, 2, 3, 4}, []int{2, 3, 4, 1}, []Entry[int]{{Index: 0, Element: 1, IsAddition: false}, {Index: 3, Element: 1, IsAddition: true}}, func(m *MockDiffVisitor[int]) { m.On("HandleMove", 0, 3, 1).Once() }},
		{[]int{1, 2, 3, 4}, []int{4, 1, 2, 3}, []Entry[int]{{Index: 3, Element: 4, IsAddition: false}, {Index: 0, Element: 4, IsAddition: true}}, func(m *MockDiffVisitor[int]) { m.On("HandleMove", 3, 0, 4).Once() }},
		{[]int{1, 2, 3, 4}, []int{1, 3, 2, 4}, []Entry[int]{{Index: 2, Element: 3, IsAddition: false}, {Index: 1, Element: 3, IsAddition: true}}, func(m *MockDiffVisitor[int]) { m.On("HandleMove", 2, 1, 3).Once() }},
		{[]int{1, 2, 3, 4}, []int{1, 3, 4, 5, 6, 2}, []Entry[int]{{Index: 1, Element: 2, IsAddition: false}, {Index: 3, Element: 2, IsAddition: true}, {Index: 3, Element: 5, IsAddition: true}, {Index: 4, Element: 6, IsAddition: true}}, func(m *MockDiffVisitor[int]) {
			c1 := m.On("HandleMove", 1, 3, 2).Once()
			c2 := m.On("HandleAdd", 3, 5).Once().NotBefore(c1)
			_ = m.On("HandleAdd", 4, 6).Once().NotBefore(c2)
		}},

		// Replace
		{[]int{1, 2, 3, 4}, []int{1, 2, 3, 5}, []Entry[int]{{Index: 3, Element: 5, IsAddition: true}, {Index: 4, Element: 4, IsAddition: false}}, func(m *MockDiffVisitor[int]) { m.On("HandleReplace", 3, 4, 5).Once() }},

		// Add + Remove
		{[]int{1, 2, 3, 4}, []int{1, 2, 5, 3}, []Entry[int]{{Index: 2, Element: 5, IsAddition: true}, {Index: 4, Element: 4, IsAddition: false}}, func(m *MockDiffVisitor[int]) {
			c1 := m.On("HandleAdd", 2, 5).Once()
			_ = m.On("HandleRemove", 4, 4).Once().NotBefore(c1)
		}},
	} {
		diff := Compute(tt.oldArray, tt.newArray, equals[int])
		require.NotNil(t, diff)
		require.Equal(t, tt.entries, diff.Entries)
		mockVisitor := NewMockDiffVisitor[int](t)
		tt.mockVisitorInit(mockVisitor)
		diff.Accept(mockVisitor)
		mockVisitor.AssertExpectations(t)
	}
}

func TestDiff_Accept_Add(t *testing.T) {
	diff := Diff[int]{
		Entries: []Entry[int]{{Index: 2, Element: 1, IsAddition: true}, {Index: 2, Element: 2, IsAddition: true}},
		equals:  equals[int],
	}
	mockVisitor := NewMockDiffVisitor[int](t)
	mockVisitor.On("HandleAdd", 2, 1).Once()
	mockVisitor.On("HandleAdd", 2, 2).Once()
	diff.Accept(mockVisitor)
	mockVisitor.AssertExpectations(t)
}

func TestDiff_Accept_Remove(t *testing.T) {
	diff := Diff[int]{
		Entries: []Entry[int]{{Index: 2, Element: 1, IsAddition: false}, {Index: 2, Element: 2, IsAddition: false}},
		equals:  equals[int],
	}
	mockVisitor := NewMockDiffVisitor[int](t)
	mockVisitor.On("HandleRemove", 2, 1).Once()
	mockVisitor.On("HandleRemove", 2, 2).Once()
	diff.Accept(mockVisitor)
	mockVisitor.AssertExpectations(t)
}

func TestDiff_Accept_AddRemove_Same(t *testing.T) {
	diff := Diff[int]{
		Entries: []Entry[int]{{Index: 2, Element: 5, IsAddition: true}, {Index: 2, Element: 5, IsAddition: false}},
		equals:  equals[int],
	}
	mockVisitor := NewMockDiffVisitor[int](t)
	mockVisitor.On("HandleAdd", 2, 5).Once()
	mockVisitor.On("HandleRemove", 2, 5).Once()
	diff.Accept(mockVisitor)
	mockVisitor.AssertExpectations(t)
}

func TestDiff_Accept_AddRemove_After(t *testing.T) {
	diff := Diff[int]{
		Entries: []Entry[int]{{Index: 4, Element: 5, IsAddition: true}, {Index: 2, Element: 5, IsAddition: false}},
		equals:  equals[int],
	}
	mockVisitor := NewMockDiffVisitor[int](t)
	mockVisitor.On("HandleMove", 2, 3, 5).Once()
	diff.Accept(mockVisitor)
	mockVisitor.AssertExpectations(t)
}
