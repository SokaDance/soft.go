package ecore

type EList[T any] interface {
	ECollection[T]

	Insert(int, T) bool

	InsertAll(int, ECollection[T]) bool

	RemoveAt(int) T

	MoveObject(newIndex int, t T)

	MoveIndex(oldIndex int, newIndex int) T

	IndexOf(T) int

	Get(int) T

	Set(int, T) T
}

func ToList[T,U any]( l EList[T]) EList[U] {
	return newListDelegate[T,U](l)
}
