package ecore

type ECollection[T any] interface {
	EIterable[T]

	Add(t T) bool

	AddAll(c ECollection[T]) bool

	Remove(t T) bool

	RemoveAll(c ECollection[T]) bool

	RetainAll(c ECollection[T]) bool

	Size() int

	Clear()

	Empty() bool

	Contains(t T) bool

	ToArray() []T
}

