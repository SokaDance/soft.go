package ecore

func ToAny[T any](t T) any { return t }

func FromAny[T any](a any) T { return a.(T) }

func ToAnyArray[T any](c ECollection[T]) []any {
	return ToArray(c, ToAny[T])
}

func ToAnyCollection[T any](c ECollection[T]) ECollection[any] {
	return ToCollection(c, ToAny[T], FromAny[T])
}

func FromAnyCollection[T any](c ECollection[any]) ECollection[T] {
	return ToCollection(c, FromAny[T], ToAny[T])
}

func ToAnyList[T any](l EList[T]) EList[any] {
	return ToList(l, ToAny[T], FromAny[T])
}

func FromAnyList[T any](l EList[any]) EList[T] {
	return ToList(l, FromAny[T], ToAny[T])
}

func ToAnyNotifyingList[T any](l ENotifyingList[T]) ENotifyingList[any] {
	return ToNotifyingList(l, ToAny[T], FromAny[T])
}

func FromAnyNotifyingList[T any](l ENotifyingList[any]) ENotifyingList[T] {
	return ToNotifyingList(l, FromAny[T], ToAny[T])
}

func ToAnyObjectList[T any](l EObjectList[T]) EObjectList[any] {
	return ToObjectList(l, ToAny[T], FromAny[T])
}

func FromAnyObjectList[T any](l EObjectList[any]) EObjectList[T] {
	return ToObjectList(l, FromAny[T], ToAny[T])
}

func ToAnyMap[K comparable, V any](m EMap[K, V]) EMap[any, any] {
	return ToMap[K, any, V, any](m, ToAny[K], FromAny[K], ToAny[V], FromAny[V])
}

func FromAnyMap[K comparable, V any](m EMap[any, any]) EMap[K, V] {
	return ToMap[any, K, any, V](m, FromAny[K], ToAny[K], FromAny[V], ToAny[V])
}
