package diff

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/zyedidia/generic"
	"github.com/zyedidia/generic/set"
)

func TestCompute_Empty(t *testing.T) {
	oldArray := []int{1, 2, 3, 4}
	newArray := []int{1, 2, 3, 4}
	result := Compute(oldArray, newArray, generic.Equals[int], generic.HashInt)
	assert.Equal(t, 0, result.GetDeletes().Size())
	assert.Equal(t, 0, result.GetInserts().Size())
	assert.Equal(t, 0, result.GetUpdates().Size())
	assert.Equal(t, 0, result.GetMoves().Size())
}

func TestCompute_Delete(t *testing.T) {
	oldArray := []int{1, 2, 3, 4}
	newArray := []int{1, 4}
	result := Compute(oldArray, newArray, generic.Equals[int], generic.HashInt)
	assert.True(t, result.GetDeletes().Equal(set.NewMapset(1, 2)))
}

func TestCompute_Insert(t *testing.T) {
	oldArray := []int{1, 2, 3, 4}
	newArray := []int{0, 1, 2, 5, 3, 4, 6}
	result := Compute(oldArray, newArray, generic.Equals[int], generic.HashInt)
	assert.True(t, result.GetInserts().Equal(set.NewMapset(0, 3, 6)))
}

func TestCompute_Move(t *testing.T) {
	oldArray := []int{1, 2, 3, 4}
	newArray := []int{4, 3, 2, 1}
	result := Compute(oldArray, newArray, generic.Equals[int], generic.HashInt)
	assert.True(t, result.GetMoves().Equal(newMoveIndexSet(makeMoveIndex(0, 3), makeMoveIndex(1, 2))))
}

func TestCompute_MoveBis(t *testing.T) {
	oldArray := []int{1, 2, 3, 4}
	newArray := []int{2, 1, 3, 4}
	result := Compute(oldArray, newArray, generic.Equals[int], generic.HashInt)
	assert.True(t, result.GetMoves().Equal(newMoveIndexSet(makeMoveIndex(0, 1))))
}

func TestCompute_DeleteInsert(t *testing.T) {
	oldArray := []int{1, 2, 3, 4}
	newArray := []int{1, 6, 5, 4}
	result := Compute(oldArray, newArray, generic.Equals[int], generic.HashInt)
	assert.True(t, result.GetDeletes().Equal(set.NewMapset(1, 2)))
	assert.True(t, result.GetInserts().Equal(set.NewMapset(1, 2)))
	assert.Equal(t, 0, result.GetMoves().Size())
}

func TestCompute_DeleteMove(t *testing.T) {
	oldArray := []int{1, 2, 3, 4}
	newArray := []int{1, 4, 3}
	result := Compute(oldArray, newArray, generic.Equals[int], generic.HashInt)
	assert.True(t, result.GetDeletes().Equal(set.NewMapset(1)))
	assert.True(t, result.GetMoves().Equal(newMoveIndexSet(makeMoveIndex(1, 3))))
}

func TestCompute_MoveInsert(t *testing.T) {
	oldArray := []int{1, 2, 3, 4}
	newArray := []int{1, 5, 2, 4, 3}
	result := Compute(oldArray, newArray, generic.Equals[int], generic.HashInt)
	assert.True(t, result.GetInserts().Equal(set.NewMapset(1)))
	assert.True(t, result.GetMoves().Equal(newMoveIndexSet(makeMoveIndex(2, 4))))
}

func TestCompute_MoveInsertBis(t *testing.T) {
	oldArray := []int{1, 2, 3, 4}
	newArray := []int{2, 1, 5, 3, 4}
	result := Compute(oldArray, newArray, generic.Equals[int], generic.HashInt)
	assert.True(t, result.GetInserts().Equal(set.NewMapset(2)))
	assert.True(t, result.GetMoves().Equal(newMoveIndexSet(makeMoveIndex(0, 1))))
}

func TestCompute_Update(t *testing.T) {
	mockObject1 := &mock.Mock{}
	mockObject2 := &mock.Mock{}
	mockObject3 := &mock.Mock{}
	oldArray := []*mock.Mock{mockObject1, mockObject3}
	newArray := []*mock.Mock{mockObject2, mockObject3}
	result := Compute(oldArray, newArray, generic.Equals[*mock.Mock], func(m *mock.Mock) uint64 {
		switch m {
		case mockObject1:
			return 1
		case mockObject2:
			return 1
		case mockObject3:
			return 2
		default:
			return 0
		}
	})
	assert.Equal(t, 0, result.GetInserts().Size())
	assert.Equal(t, 0, result.GetMoves().Size())
	assert.Equal(t, 0, result.GetDeletes().Size())
	assert.True(t, result.GetUpdates().Equal(set.NewMapset(0)))
}
