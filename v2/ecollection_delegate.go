package ecore


type eCollectionDelegate[T,U any] struct {
	delegate ECollection[T]
}

func (l *eCollectionDelegate[T,U]) Add(u U) bool {
	return l.delegate.Add(any(u).(T))
}

func (l *eCollectionDelegate[T,U]) AddAll(c ECollection[U]) bool {
	return l.delegate.AddAll(ToCollection[U,T](c))
}

func (l *eCollectionDelegate[T,U]) Remove(u U) bool {
	return l.delegate.Remove(any(u).(T))
}

func (l *eCollectionDelegate[T,U]) RemoveAll(c ECollection[U]) bool {
	return l.delegate.RemoveAll( ToCollection[U,T](c) )
}

func (l *eCollectionDelegate[T,U]) RetainAll(c ECollection[U]) bool {
	return l.delegate.RetainAll( ToCollection[U,T](c) )
}

func (l *eCollectionDelegate[T,U]) Size() int {
	return l.delegate.Size()
}

func (l *eCollectionDelegate[T,U]) Clear() {
	l.delegate.Clear()
}

func (l *eCollectionDelegate[T,U]) Empty() bool {
	return l.delegate.Empty()
}

func (l *eCollectionDelegate[T,U]) Contains(u U) bool {
	return l.delegate.Contains(any(u).(T))
}

func (l *eCollectionDelegate[T,U]) ToArray() []U {
	return ToArray[T,U](l.delegate)
}

type eCollectionDelegateIterator[T,U any] struct {
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

func (l *eCollectionDelegate[T,U]) Iterator() EIterator[U] {
	return &eCollectionDelegateIterator[T,U]{delegate : l.delegate.Iterator()}
}