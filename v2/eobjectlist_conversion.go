package ecore

func ToObjectList[T,U any](l EObjectList[T], convertTo func (T) U, convertFrom func(U) T ) EObjectList[U] {
	r := &eObjectListDelegate[T,U,EObjectList[T]]{}
	r.delegate = l
	r.convertTo = convertTo
	r.convertFrom = convertFrom
	return r	
}

func ToAnyObjectList[T any]( l EObjectList[T] ) EObjectList[any] {
	return ToObjectList[T,any](l,toAny[T],fromAny[T])
}

type eObjectListDelegate[T any,U any,D EObjectList[T]] struct {
	eListDelegate[T,U,D]
}

func newObjectListDelegate[T,U any](delegate EObjectList[T]) *eObjectListDelegate[T,U,EObjectList[T]] {
	l := &eObjectListDelegate[T,U,EObjectList[T]]{}
	l.delegate = delegate
	return l
}

func (l *eObjectListDelegate[T,U,D]) GetUnResolvedList() EList[U] {
	return ToList[T,U](l.delegate.GetUnResolvedList(),l.convertTo,l.convertFrom)
}
