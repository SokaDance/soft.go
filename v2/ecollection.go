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

func ToArray[T,U any](c ECollection[T]) []U {
	result := make([]U,c.Size())
	i := 0
	for it := c.Iterator() ; it.HasNext(); i++{
		result[i] = any(it.Next()).(U)
	}
	return result
}
