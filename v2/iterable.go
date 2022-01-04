package ecore

type Iterable[T any] interface {
	Iterator() Iterator[T]
}