package ecore

type eObjectListDelegate[T, U any] interface {
	EObjectList[U]
	GetDelegate() EObjectList[T]
}
