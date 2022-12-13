package diff

type EqualsFn[T any] func(a, b T) bool

type Entry[T any] struct {
	Index      int
	Element    T
	IsAddition bool
}

type Diff[T any] []Entry[T]

type DiffVisitor[T any] interface {
	HandleAdd(index int, element T)
	HandleMove(oldIndex int, newIndex int, element T)
	HandleRemove(index int, element T)
	HandleReplace(index int, oldElement T, newElement T)
}

func (d Diff[T]) Accept(v DiffVisitor[T]) {

}

func insert[T any](arr []T, index int, t T) []T {
	var empty T
	arr = append(arr, empty)
	copy(arr[index+1:], arr[index:])
	arr[index] = t
	return arr
}

func remove[T any](arr []T, index int) []T {
	var empty T
	copy(arr[index:], arr[index+1:])
	arr[len(arr)-1] = empty
	return arr[:len(arr)-1]
}

func indexOf[T any](array []T, v T, index int, equals EqualsFn[T]) int {
	for i := index; i < len(array); i++ {
		if equals(array[i], v) {
			return i
		}
	}
	return -1
}

func Compute[T any](immutableOldArray []T, newArray []T, equals EqualsFn[T]) Diff[T] {
	index := 0
	entries := []Entry[T]{}
	oldArray := make([]T, len(immutableOldArray))
	copy(oldArray, immutableOldArray)
	for _, newValue := range newArray {
		if len(oldArray) <= index {
			// append newValue to newArray
			entries = append(entries, Entry[T]{Index: index, Element: newValue, IsAddition: true})
		} else {
			done := false
			for !done {
				done = true
				oldValue := oldArray[index]
				if !equals(oldValue, newValue) {
					if oldIndexOfNewValue := indexOf(oldArray, newValue, index, equals); oldIndexOfNewValue != -1 {
						if newIndexOfOldValue := indexOf(newArray, oldValue, index, equals); newIndexOfOldValue == -1 {
							// removing oldValue from list[index]
							entries = append(entries, Entry[T]{Index: index, Element: oldValue, IsAddition: false})
							oldArray = remove(oldArray, index)
							done = false
						} else if newIndexOfOldValue > oldIndexOfNewValue {
							// moving oldValue from list[index] to
							// [newIndexOfOldValue]
							if len(oldArray) <= newIndexOfOldValue {
								// The element cannot be moved to the
								// correct index
								// now, however later iterations will insert
								// elements
								// in front of it, eventually moving it into
								// the
								// correct spot.
								newIndexOfOldValue = len(oldArray) - 1
							}
							entries = append(entries, Entry[T]{Index: index, Element: oldValue, IsAddition: false})
							oldArray = remove(oldArray, index)
							entries = append(entries, Entry[T]{Index: newIndexOfOldValue, Element: oldValue, IsAddition: true})
							oldArray = insert(oldArray, newIndexOfOldValue, oldValue)
							done = false
						} else {
							// move newValue from list[oldIndexOfNewValue]
							// to [index]
							entries = append(entries, Entry[T]{Index: oldIndexOfNewValue, Element: newValue, IsAddition: false})
							oldArray = remove(oldArray, oldIndexOfNewValue)
							entries = append(entries, Entry[T]{Index: index, Element: newValue, IsAddition: true})
							oldArray = insert(oldArray, index, newValue)
						}

					} else {
						// add newValue at list[index]
						oldArray = insert(oldArray, index, newValue)
						entries = append(entries, Entry[T]{Index: index, Element: newValue, IsAddition: true})
					}
				}
			}
		}
		index++
	}
	for i := len(oldArray) - 1; i >= index; {
		// remove excess trailing elements not present in newArray
		entries = append(entries, Entry[T]{Index: i, Element: oldArray[i], IsAddition: false})
	}
	return Diff[T](entries)
}
