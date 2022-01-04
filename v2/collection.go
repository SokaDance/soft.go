package ecore

type Collection[T any] interface {
	Add(t T) bool

	AddAll(c Collection[T]) bool

	Remove(t T) bool

	RemoveAll(c Collection[T]) bool

	RetainAll(c Collection[T]) bool

	Size() int

	Clear()

	Empty() bool

	Contains(t T) bool

	ToArray() []T
}

