package diff

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"github.com/zyedidia/generic"
)

func hashInt(i int) uint64 {
	return uint64(i)
}

var equalsInt = generic.Equals[int]

func TestCompute_Empty(t *testing.T) {
	oldArray := []int{1, 2, 3, 4}
	newArray := []int{1, 2, 3, 4}
	result := Compute(oldArray, newArray, equalsInt, hashInt)
	assert.Equal(t, Result{
		Operations:  []Operation{},
		OldIndexFor: map[uint64]int{1: 0, 2: 1, 3: 2, 4: 3},
		NewIndexFor: map[uint64]int{1: 0, 2: 1, 3: 2, 4: 3},
	}, result)
}

func TestCompute_Delete(t *testing.T) {
	oldArray := []int{1, 2, 3, 4}
	newArray := []int{1, 4}
	result := Compute(oldArray, newArray, equalsInt, hashInt)
	assert.Equal(t, Result{
		Operations:  []Operation{Delete(1), Delete(2)},
		OldIndexFor: map[uint64]int{1: 0, 2: 1, 3: 2, 4: 3},
		NewIndexFor: map[uint64]int{1: 0, 4: 1},
	}, result)
}

func TestCompute_Insert(t *testing.T) {
	oldArray := []int{1, 2, 3, 4}
	newArray := []int{0, 1, 2, 5, 3, 4, 6}
	result := Compute(oldArray, newArray, equalsInt, hashInt)
	assert.Equal(t, Result{
		Operations:  []Operation{Insert(0), Insert(3), Insert(6)},
		OldIndexFor: map[uint64]int{1: 0, 2: 1, 3: 2, 4: 3},
		NewIndexFor: map[uint64]int{0: 0, 1: 1, 2: 2, 5: 3, 3: 4, 4: 5, 6: 6},
	}, result)
}

func TestCompute_Move(t *testing.T) {
	oldArray := []int{1, 2, 3, 4}
	newArray := []int{4, 3, 2, 1}
	result := Compute(oldArray, newArray, equalsInt, hashInt)
	assert.Equal(t, Result{
		Operations:  []Operation{Move{From: 3, To: 0}, Move{From: 2, To: 1}, Move{From: 0, To: 3}, Move{From: 1, To: 2}},
		OldIndexFor: map[uint64]int{1: 0, 2: 1, 3: 2, 4: 3},
		NewIndexFor: map[uint64]int{4: 0, 3: 1, 2: 2, 1: 3},
	}, result)
}

func TestCompute_MoveBis(t *testing.T) {
	oldArray := []int{1, 2, 3, 4}
	newArray := []int{2, 1, 3, 4}
	result := Compute(oldArray, newArray, equalsInt, hashInt)
	assert.Equal(t, Result{
		Operations:  []Operation{Move{From: 1, To: 0}, Move{From: 0, To: 1}},
		OldIndexFor: map[uint64]int{1: 0, 2: 1, 3: 2, 4: 3},
		NewIndexFor: map[uint64]int{2: 0, 1: 1, 3: 2, 4: 3},
	}, result)
}

func TestCompute_DeleteInsert(t *testing.T) {
	oldArray := []int{1, 2, 3, 4}
	newArray := []int{1, 6, 5, 4}
	result := Compute(oldArray, newArray, equalsInt, hashInt)
	assert.Equal(t, Result{
		Operations:  []Operation{Delete(1), Delete(2), Insert(1), Insert(2)},
		OldIndexFor: map[uint64]int{1: 0, 2: 1, 3: 2, 4: 3},
		NewIndexFor: map[uint64]int{1: 0, 6: 1, 5: 2, 4: 3},
	}, result)
}

func TestCompute_DeleteMove(t *testing.T) {
	oldArray := []int{1, 2, 3, 4}
	newArray := []int{1, 4, 3}
	result := Compute(oldArray, newArray, equalsInt, hashInt)
	assert.Equal(t, Result{
		Operations:  []Operation{Delete(1), Move{3, 1}},
		OldIndexFor: map[uint64]int{1: 0, 2: 1, 3: 2, 4: 3},
		NewIndexFor: map[uint64]int{1: 0, 4: 1, 3: 2},
	}, result)
}

func TestCompute_MoveInsert(t *testing.T) {
	oldArray := []int{1, 2, 3, 4}
	newArray := []int{1, 5, 2, 4, 3}
	result := Compute(oldArray, newArray, equalsInt, hashInt)
	assert.Equal(t, Result{
		Operations:  []Operation{Insert(1), Move{2, 4}},
		OldIndexFor: map[uint64]int{1: 0, 2: 1, 3: 2, 4: 3},
		NewIndexFor: map[uint64]int{1: 0, 5: 1, 2: 2, 4: 3, 3: 4},
	}, result)
}

func TestCompute_MoveInsertBis(t *testing.T) {
	oldArray := []int{1, 2, 3, 4}
	newArray := []int{2, 1, 5, 3, 4}
	result := Compute(oldArray, newArray, equalsInt, hashInt)
	assert.Equal(t, Result{
		Operations:  []Operation{Move{1, 0}, Move{0, 1}, Insert(2)},
		OldIndexFor: map[uint64]int{1: 0, 2: 1, 3: 2, 4: 3},
		NewIndexFor: map[uint64]int{2: 0, 1: 1, 5: 2, 3: 3, 4: 4},
	}, result)
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
	assert.Equal(t, Result{
		Operations:  []Operation{Update(0)},
		OldIndexFor: map[uint64]int{1: 0, 2: 1},
		NewIndexFor: map[uint64]int{1: 0, 2: 1},
	}, result)
}

const nbIteration = 50
const collectionSize = 100

func performRandomActions(original []int, minimumCountAfterActions, maximumCountAfterActions int) []int {
	destination := make([]int, len(original))
	copy(destination, original)
	destination = performInsertActions(destination, rand.Intn(maximumCountAfterActions)/2)
	destination = performDeleteActions(destination, rand.Intn(maximumCountAfterActions)/2)
	destination = performMoveActions(destination, rand.Intn(maximumCountAfterActions)/3)
	if len(destination) < minimumCountAfterActions {
		destination = performInsertActions(destination, minimumCountAfterActions)
	}
	if len(destination) > maximumCountAfterActions {
		destination = performInsertActions(destination, minimumCountAfterActions)
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

var id = 0

func newObject() int {
	defer func() {
		id++
	}()
	return id
}

func performInsertActions(array []int, count int) []int {
	for i := 0; i < count; i++ {
		// create object
		mockObject := newObject()
		// compute random index
		index := rand.Intn(len(array))
		// apply to array
		array = scliceInsert(array, index, mockObject)
	}
	return array
}

func performDeleteActions(array []int, count int) []int {
	for i := 0; i < count; i++ {
		// compute random index
		index := rand.Intn(len(array))
		// apply to array
		array = scliceDelete(array, index)
	}
	return array
}

func performMoveActions(array []int, count int) []int {
	for i := 0; i < count; i++ {
		// compute random index
		from := rand.Intn(len(array))
		to := rand.Intn(len(array))
		// apply to array
		array = scliceMove(array, from, to)
	}
	return array
}

func applyResult(currentObjects []int, expectedObjects []int, result Result) []int {
	resultObjects := make([]int, len(currentObjects))
	copy(resultObjects, currentObjects)
	for _, operation := range result.Operations {
		switch op := operation.(type) {
		case Delete:
			index := int(op)
			resultObjects = scliceDelete(resultObjects, index)
		case Insert:
			index := int(op)
			resultObjects = scliceInsert(resultObjects, index, expectedObjects[index])
		case Move:
			resultObjects[op.To] = currentObjects[op.From]
		}
	}
	return resultObjects
}

func TestCompute_Stress(t *testing.T) {

	// build objects
	currentObjects := make([]int, collectionSize)
	for i := 0; i < collectionSize; i++ {
		currentObjects[i] = newObject()
	}

	// perform random actions & compute diff
	for i := 0; i < nbIteration; i++ {
		expectedObjects := performRandomActions(currentObjects, collectionSize/10, collectionSize*10)
		result := Compute(currentObjects, expectedObjects, generic.Equals[int], generic.HashInt)
		resultObjects := applyResult(currentObjects, expectedObjects, result)
		require.Equal(t, expectedObjects, resultObjects)
		currentObjects = resultObjects
	}

}
