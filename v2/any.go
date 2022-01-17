package ecore

func ToAny[T any](t T) any { return t }

func FromAny[T any](a any) T { return a.(T) }

func ToAnyArray[T any](c ECollection[T]) []any {
	return ToArray(c, ToAny[T])
}

func ToAnyCollection[T any](c ECollection[T]) ECollection[any] {
	return ToCollection(c, ToAny[T], FromAny[T])
}

func ToAnyList[T any](l EList[T]) EList[any] {
	return ToList(l, ToAny[T], FromAny[T])
}

func ToAnyObjectList[T any](l EObjectList[T]) EObjectList[any] {
	return ToObjectList(l, ToAny[T], FromAny[T])
}
