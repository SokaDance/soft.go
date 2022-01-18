package ecore

type eCollectionDelegate[T, U any] interface {
	ECollection[U]
	GetDelegate() ECollection[T]
}
