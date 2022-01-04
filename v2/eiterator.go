package ecore

// EIterator is an interator
type EIterator[T any] interface {
	HasNext() bool
	Next() T
}
