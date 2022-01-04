package ecore

type EList[T any] interface {
	Collection[T]

	Insert(int, T) bool

	InsertAll(int, Collection[T]) bool

	RemoveAt(int) T

	MoveObject(newIndex int, t T)

	MoveIndex(oldIndex int, newIndex int) T

	Get(int) T

	Set(int, T) T
}