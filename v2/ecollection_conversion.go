package ecore

func ToArray[T, U any](c ECollection[T], convertTo func(T) U) []U {
	result := make([]U, c.Size())
	i := 0
	for it := c.Iterator(); it.HasNext(); i++ {
		result[i] = convertTo(it.Next())
	}
	return result
}

func ToCollection[T, U any](c ECollection[T], convertTo func(T) U, convertFrom func(U) T) ECollection[U] {
	return &eCollectionDelegate[T, U, ECollection[T]]{delegate: c, convertTo: convertTo, convertFrom: convertFrom}
}

type eCollectionDelegate[T any, U any, D ECollection[T]] struct {
	delegate    D
	convertTo   func(T) U
	convertFrom func(U) T
}

func (l *eCollectionDelegate[T, U, D]) Add(u U) bool {
	return l.delegate.Add(l.convertFrom(u))
}

func (l *eCollectionDelegate[T, U, D]) AddAll(c ECollection[U]) bool {
	return l.delegate.AddAll(ToCollection(c, l.convertFrom, l.convertTo))
}

func (l *eCollectionDelegate[T, U, D]) Remove(u U) bool {
	return l.delegate.Remove(l.convertFrom(u))
}

func (l *eCollectionDelegate[T, U, D]) RemoveAll(c ECollection[U]) bool {
	return l.delegate.RemoveAll(ToCollection(c, l.convertFrom, l.convertTo))
}

func (l *eCollectionDelegate[T, U, D]) RetainAll(c ECollection[U]) bool {
	return l.delegate.RetainAll(ToCollection(c, l.convertFrom, l.convertTo))
}

func (l *eCollectionDelegate[T, U, D]) Size() int {
	return l.delegate.Size()
}

func (l *eCollectionDelegate[T, U, D]) Clear() {
	l.delegate.Clear()
}

func (l *eCollectionDelegate[T, U, D]) Empty() bool {
	return l.delegate.Empty()
}

func (l *eCollectionDelegate[T, U, D]) Contains(u U) bool {
	return l.delegate.Contains(l.convertFrom(u))
}

func (l *eCollectionDelegate[T, U, D]) ToArray() []U {
	return ToArray[T](l.delegate, l.convertTo)
}

func (l *eCollectionDelegate[T, U, D]) Iterator() EIterator[U] {
	return &eCollectionDelegateIterator[T, U]{delegate: l.delegate.Iterator(), convertTo: l.convertTo}
}

type eCollectionDelegateIterator[T any, U any] struct {
	delegate  EIterator[T]
	convertTo func(T) U
}

// Next return the current value of the iterator
func (it *eCollectionDelegateIterator[T, U]) Next() U {
	return it.convertTo(it.delegate.Next())
}

// HasNext make the iterator go further in the array
func (it *eCollectionDelegateIterator[T, U]) HasNext() bool {
	return it.delegate.HasNext()
}
