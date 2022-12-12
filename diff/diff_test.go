package diff

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
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

const nbIteration = 1
const collectionSize = 5

func performRandomActions(original []*mock.Mock, minimumCountAfterActions, maximumCountAfterActions int, objectToID map[*mock.Mock]uint64) []*mock.Mock {
	destination := make([]*mock.Mock, len(original))
	copy(destination, original)
	destination = performInsertActions(destination, rand.Intn(maximumCountAfterActions)/2, objectToID)
	destination = performDeleteActions(destination, rand.Intn(maximumCountAfterActions)/2)
	destination = performMoveActions(destination, rand.Intn(maximumCountAfterActions)/3)
	destination = performUpdateActions(destination, rand.Intn(maximumCountAfterActions)/3, objectToID)
	if len(destination) < minimumCountAfterActions {
		destination = performInsertActions(destination, minimumCountAfterActions, objectToID)
	}
	if len(destination) > maximumCountAfterActions {
		destination = performInsertActions(destination, minimumCountAfterActions, objectToID)
	}
	if len(destination) > maximumCountAfterActions {
		destination = performDeleteActions(destination, len(destination)-maximumCountAfterActions)
	}
	return destination
}

func scliceInsert[T any](array []T, i int, v T) []T {
	var empty T
	array = append(array, empty)
	copy(array[i+1:], array[i:])
	array[i] = v
	return array
}

func scliceDelete[T any](array []T, i int) []T {
	var empty T
	copy(array[i:], array[i+1:])
	array[len(array)-1] = empty
	array = array[:len(array)-1]
	return array
}

func scliceMove[T any](array []T, from, to int) []T {
	if from != to {
		object := array[from]
		if to < from {
			copy(array[to+1:], array[to:from])
		} else {
			copy(array[from:], array[from+1:to+1])
		}
		array[to] = object
	}
	return array
}

func performInsertActions(array []*mock.Mock, count int, objectToID map[*mock.Mock]uint64) []*mock.Mock {
	for i := 0; i < count; i++ {
		// create object
		mockObject := &mock.Mock{}
		objectToID[mockObject] = uint64(len(array) + i)
		// compute random index
		index := rand.Intn(len(array))
		// apply to array
		array = scliceInsert(array, index, mockObject)
	}
	return array
}

func performDeleteActions(array []*mock.Mock, count int) []*mock.Mock {
	for i := 0; i < count; i++ {
		// compute random index
		index := rand.Intn(len(array))
		// apply to array
		array = scliceDelete(array, index)
	}
	return array
}

func performUpdateActions(array []*mock.Mock, count int, objectToID map[*mock.Mock]uint64) []*mock.Mock {
	for i := 0; i < count; i++ {
		// create new object
		mockObject := &mock.Mock{}
		// compute random index
		index := rand.Intn(len(array))
		// apply to array
		oldObject := array[index]
		array[index] = mockObject
		objectToID[mockObject] = objectToID[oldObject]
	}
	return array
}

func performMoveActions(array []*mock.Mock, count int) []*mock.Mock {
	for i := 0; i < count; i++ {
		// compute random index
		from := rand.Intn(len(array))
		to := rand.Intn(len(array))
		// apply to array
		array = scliceMove(array, from, to)
	}
	return array
}

func applyResult(currentObjects []*mock.Mock, expectedObjects []*mock.Mock, result Result) []*mock.Mock {
	resultObjects := make([]*mock.Mock, len(currentObjects))
	copy(resultObjects, currentObjects)
	result.GetDeletes().Each(func(i int) {
		resultObjects = scliceDelete(resultObjects, i)
	})
	result.GetInserts().Each(func(i int) {
		resultObjects = scliceInsert(resultObjects, i, expectedObjects[i])
	})
	result.GetMoves().Each(func(mi MoveIndex) {
		resultObjects = scliceMove(resultObjects, mi.From, mi.To)
	})
	return resultObjects
}

func TestCompute_Stress(t *testing.T) {

	objectToID := map[*mock.Mock]uint64{}

	// build objects
	currentObjects := make([]*mock.Mock, collectionSize)
	for i := 0; i < collectionSize; i++ {
		mockObject := &mock.Mock{}
		objectToID[mockObject] = uint64(i)
		currentObjects[i] = mockObject
	}

	// perform random actions & compute diff
	for i := 0; i < nbIteration; i++ {
		expectedObjects := performRandomActions(currentObjects, collectionSize/10, collectionSize*10, objectToID)
		result := Compute(currentObjects, expectedObjects, generic.Equals[*mock.Mock], func(m *mock.Mock) uint64 { return objectToID[m] })
		resultObjects := applyResult(currentObjects, expectedObjects, result)
		require.Equal(t, expectedObjects, resultObjects)
		currentObjects = resultObjects
	}

}
