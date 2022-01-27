package ecore

type eListDelegate[T, U any] interface {
	EList[U]
	GetDelegate() EList[T]
}
