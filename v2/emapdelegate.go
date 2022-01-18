package ecore

type eMapDelegate[KO comparable, KT comparable, VO any, VT any] interface {
	EMap[KT, VT]
	GetDelegate() EMap[KO, VO]
}
