package ecore

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToObjectListWithDelegate(t *testing.T) {
	l := &MockEObjectList[EObject]{}
	l2 := ToAnyList[EObject](l)
	l3 := FromAnyList[EObject](l2)
	assert.Equal(t, l, l3)
}
