package ecore

// Iterator is an interator
type Iterator[T any] interface {
	HasNext() bool
	Next() T
}
