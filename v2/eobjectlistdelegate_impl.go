package ecore

func ToObjectList[T, U any](l EObjectList[T], convertTo func(T) U, convertFrom func(U) T) EObjectList[U] {
	if ld, isDelegate := l.(eObjectListDelegate[U, T]); isDelegate {
		return ld.GetDelegate()
	} else {
		r := &eObjectListDelegateImpl[T, U, EObjectList[T]]{}
		r.delegate = l
		r.convertTo = convertTo
		r.convertFrom = convertFrom
		return r
	}
}

type eObjectListDelegateImpl[T any, U any, C EObjectList[T]] struct {
	eListDelegateImpl[T, U, C]
}

func (l *eObjectListDelegateImpl[T, U, C]) GetDelegate() EObjectList[T] {
	return l.delegate
}

func (l *eObjectListDelegateImpl[T, U, C]) GetUnResolvedList() EObjectList[U] {
	return ToObjectList(l.delegate.GetUnResolvedList(), l.convertTo, l.convertFrom)
}
