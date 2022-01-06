package ecore

type EIterable[T any] interface {
	Iterator() EIterator[T]
}
