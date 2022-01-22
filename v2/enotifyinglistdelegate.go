package ecore

type eNotifyingListDelegate[T,U any] interface {
	ENotifyingList[U]
	GetDelegate() ENotifyingList[T]
}