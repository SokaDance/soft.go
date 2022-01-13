package ecore


type eCollectionDelegate[T any ,U any, D ECollection[T]] struct {
	delegate D
}

func newCollectionDelegate[T,U any](delegate ECollection[T]) *eCollectionDelegate[T,U,ECollection[T]] {
	return &eCollectionDelegate[T,U,ECollection[T]]{delegate:delegate}
}

func (l *eCollectionDelegate[T,U,D]) Add(u U) bool {
	return l.delegate.Add(any(u).(T))
}

func (l *eCollectionDelegate[T,U,D]) AddAll(c ECollection[U]) bool {
	return l.delegate.AddAll(ToCollection[U,T](c))
}

func (l *eCollectionDelegate[T,U,D]) Remove(u U) bool {
	return l.delegate.Remove(any(u).(T))
}

func (l *eCollectionDelegate[T,U,D]) RemoveAll(c ECollection[U]) bool {
	return l.delegate.RemoveAll( ToCollection[U,T](c) )
}

func (l *eCollectionDelegate[T,U,D]) RetainAll(c ECollection[U]) bool {
	return l.delegate.RetainAll( ToCollection[U,T](c) )
}

func (l *eCollectionDelegate[T,U,D]) Size() int {
	return l.delegate.Size()
}

func (l *eCollectionDelegate[T,U,D]) Clear() {
	l.delegate.Clear()
}

func (l *eCollectionDelegate[T,U,D]) Empty() bool {
	return l.delegate.Empty()
}

func (l *eCollectionDelegate[T,U,D]) Contains(u U) bool {
	return l.delegate.Contains(any(u).(T))
}

func (l *eCollectionDelegate[T,U,D]) ToArray() []U {
	return ToArray[T,U](l.delegate)
}

type eCollectionDelegateIterator[T any ,U any] struct {
	delegate EIterator[T]
}

// Next return the current value of the iterator
func (it *eCollectionDelegateIterator[T,U]) Next() U {
	return any(it.delegate.Next()).(U)
}

// HasNext make the iterator go further in the array
func (it *eCollectionDelegateIterator[T,U]) HasNext() bool {
	return it.delegate.HasNext()
}

func (l *eCollectionDelegate[T,U,D]) Iterator() EIterator[U] {
	return &eCollectionDelegateIterator[T,U]{delegate : l.delegate.Iterator()}
}