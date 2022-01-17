package ecore

func ToList[T, U any](l EList[T], convertTo func(T) U, convertFrom func(U) T) EList[U] {
	r := &eListDelegate[T, U, EList[T]]{}
	r.delegate = l
	r.convertTo = convertTo
	r.convertFrom = convertFrom
	return r
}

type eListDelegate[T any, U any, D EList[T]] struct {
	eCollectionDelegate[T, U, D]
}

func (l *eListDelegate[T, U, D]) Insert(index int, u U) bool {
	return l.delegate.Insert(index, l.convertFrom(u))
}

func (l *eListDelegate[T, U, D]) InsertAll(index int, c ECollection[U]) bool {
	return l.delegate.InsertAll(index, ToCollection(c, l.convertFrom, l.convertTo))
}

func (l *eListDelegate[T, U, D]) RemoveAt(index int) U {
	return l.convertTo(l.delegate.RemoveAt(index))
}

func (l *eListDelegate[T, U, D]) MoveObject(newIndex int, u U) {
	l.delegate.MoveObject(newIndex, l.convertFrom(u))
}

func (l *eListDelegate[T, U, D]) MoveIndex(oldIndex int, newIndex int) U {
	return l.convertTo(l.delegate.MoveIndex(oldIndex, newIndex))
}

func (l *eListDelegate[T, U, D]) IndexOf(u U) int {
	return l.delegate.IndexOf(l.convertFrom(u))
}

func (l *eListDelegate[T, U, D]) Get(index int) U {
	return l.convertTo(l.delegate.Get(index))
}

func (l *eListDelegate[T, U, D]) Set(index int, u U) U {
	return l.convertTo(l.delegate.Set(index, l.convertFrom(u)))
}
