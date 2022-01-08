package ecore

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestToArray(t *testing.T) {
	c := &MockECollection[int]{}
	it := &MockEIterator[int]{}
	c.On("Size").Once().Return(2)
	c.On("Iterator").Once().Return(it)
	it.On("HasNext").Twice().Return(true)
	it.On("HasNext").Once().Return(false)
	it.On("Next").Once().Return(1)
	it.On("Next").Once().Return(2)
	nc := ToArray[int,any](c)
	assert.Equal(t, []any{1,2}, nc)
	mock.AssertExpectationsForObjects(t, c, it)
}