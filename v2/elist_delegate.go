package ecore

type eListDelegate[T any, U any, D EList[T]] struct {
	eCollectionDelegate[T,U,D]
}

func newListDelegate[T,U any](delegate EList[T]) *eListDelegate[T,U,EList[T]] {
	l := &eListDelegate[T,U,EList[T]]{}
	l.delegate = delegate
	return l	
}

func (l *eListDelegate[T,U,D]) Insert(index int , u U) bool {
	return l.delegate.Insert(index,any(u).(T))
}

func (l *eListDelegate[T,U,D]) InsertAll(index int, c ECollection[U]) bool {
	return l.delegate.InsertAll(index,ToCollection[U,T](c))
}

func (l *eListDelegate[T,U,D]) RemoveAt(index int) U {
	return any(l.delegate.RemoveAt(index)).(U)
}

func (l *eListDelegate[T,U,D]) MoveObject(newIndex int, u U) {
	l.delegate.MoveObject(newIndex, any(u).(T))
}

func (l *eListDelegate[T,U,D]) MoveIndex(oldIndex int, newIndex int) U {
	return any(l.delegate.MoveIndex(oldIndex, newIndex)).(U)
}

func (l *eListDelegate[T,U,D]) IndexOf(u U) int {
	return l.delegate.IndexOf(any(u).(T))
}

func (l *eListDelegate[T,U,D]) Get(index int) U{
	return any(l.delegate.Get(index)).(U)
}

func (l *eListDelegate[T,U,D]) Set(index int, u U) U {
	return any(l.delegate.Set(index, any(u).(T))).(U)
}
