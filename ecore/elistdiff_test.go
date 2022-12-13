package ecore

import (
	"testing"

	"github.com/masagroup/soft.go/diff"
)

func equalsAny(a, b any) bool {
	return a == b
}

func TestApplyDiffToList_Add(t *testing.T) {
	diff := diff.MakeDiff[any](
		[]diff.Entry[any]{{Index: 2, Element: 1, IsAddition: true}, {Index: 2, Element: 2, IsAddition: true}},
		equalsAny,
	)
	mockList := &MockEList{}
	mockList.On("Insert", 2, 1).Return(true).Once()
	mockList.On("Insert", 2, 2).Return(true).Once()
	ApplyDiffToList(diff, mockList)
	mockList.AssertExpectations(t)
}

func TestApplyDiffToList_Remove(t *testing.T) {
	diff := diff.MakeDiff[any](
		[]diff.Entry[any]{{Index: 2, Element: 1, IsAddition: false}, {Index: 2, Element: 2, IsAddition: false}},
		equalsAny,
	)
	mockList := &MockEList{}
	mockList.On("Remove", 2).Return(true).Once()
	mockList.On("Remove", 2).Return(true).Once()
	ApplyDiffToList(diff, mockList)
	mockList.AssertExpectations(t)
}

func TestApplyDiffToList_Move(t *testing.T) {
	diff := diff.MakeDiff[any](
		[]diff.Entry[any]{{Index: 4, Element: 5, IsAddition: true}, {Index: 2, Element: 5, IsAddition: false}},
		equalsAny,
	)
	mockList := &MockEList{}
	mockList.On("Move", 2, 3).Return(true).Once()
	ApplyDiffToList(diff, mockList)
	mockList.AssertExpectations(t)
}

func TestApplyDiffToList_Replace(t *testing.T) {
	diff := diff.MakeDiff[any](
		[]diff.Entry[any]{{Index: 4, Element: 4, IsAddition: false}, {Index: 4, Element: 5, IsAddition: true}},
		equalsAny,
	)
	mockList := &MockEList{}
	mockList.On("Set", 4, 5).Return(true).Once()
	ApplyDiffToList(diff, mockList)
	mockList.AssertExpectations(t)
}
