package ecore

type eObjectListDelegate[T any,U any,D EObjectList[T]] struct {
	eListDelegate[T,U,D]
}

func newObjectListDelegate[T,U any](delegate EObjectList[T]) *eObjectListDelegate[T,U,EObjectList[T]] {
	l := &eObjectListDelegate[T,U,EObjectList[T]]{}
	l.delegate = delegate
	return l
}

func (l *eObjectListDelegate[T,U,D]) GetUnResolvedList() EList[U] {
	return ToList[T,U](l.delegate.GetUnResolvedList())
}