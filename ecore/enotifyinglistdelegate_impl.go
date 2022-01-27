package ecore

func ToNotifyingList[T, U any](l ENotifyingList[T], convertTo func(T) U, convertFrom func(U) T) ENotifyingList[U] {
	if ld, isDelegate := l.(eNotifyingListDelegate[U, T]); isDelegate {
		return ld.GetDelegate()
	} else {
		r := &eNotifyingListDelegateImpl[T, U, ENotifyingList[T]]{}
		r.delegate = l
		r.convertTo = convertTo
		r.convertFrom = convertFrom
		return r
	}
}

type eNotifyingListDelegateImpl[T any, U any, C ENotifyingList[T]] struct {
	eListDelegateImpl[T, U, C]
}

func (l *eNotifyingListDelegateImpl[T, U, C]) GetDelegate() ENotifyingList[T] {
	return l.delegate
}

func (l *eNotifyingListDelegateImpl[T, U, C]) GetNotifier() ENotifier {
	return l.delegate.GetNotifier()
}

func (l *eNotifyingListDelegateImpl[T, U, C]) GetFeature() EStructuralFeature {
	return l.delegate.GetFeature()
}

func (l *eNotifyingListDelegateImpl[T, U, C]) GetFeatureID() int {
	return l.delegate.GetFeatureID()
}

func (l *eNotifyingListDelegateImpl[T, U, C]) AddWithNotification(object U, notifications ENotificationChain) ENotificationChain {
	return l.delegate.AddWithNotification(l.convertFrom(object), notifications)
}

func (l *eNotifyingListDelegateImpl[T, U, C]) RemoveWithNotification(object U, notifications ENotificationChain) ENotificationChain {
	return l.delegate.RemoveWithNotification(l.convertFrom(object), notifications)
}

func (l *eNotifyingListDelegateImpl[T, U, C]) SetWithNotification(index int, object U, notifications ENotificationChain) ENotificationChain {
	return l.delegate.SetWithNotification(index, l.convertFrom(object), notifications)
}
