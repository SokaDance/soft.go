package ecore

func ToObjectList[T, U any](l EObjectList[T], convertTo func(T) U, convertFrom func(U) T) EObjectList[U] {
	r := &eObjectListDelegate[T, U, EObjectList[T]]{}
	r.delegate = l
	r.convertTo = convertTo
	r.convertFrom = convertFrom
	return r
}

type eObjectListDelegate[T any, U any, D EObjectList[T]] struct {
	eListDelegate[T, U, D]
}

func newObjectListDelegate[T, U any](delegate EObjectList[T]) *eObjectListDelegate[T, U, EObjectList[T]] {
	l := &eObjectListDelegate[T, U, EObjectList[T]]{}
	l.delegate = delegate
	return l
}

func (l *eObjectListDelegate[T, U, D]) GetUnResolvedList() EObjectList[U] {
	return ToObjectList(l.delegate.GetUnResolvedList(), l.convertTo, l.convertFrom)
}
