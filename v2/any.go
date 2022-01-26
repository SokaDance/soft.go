package ecore

func ToAny[T any](t T) any { return t }

func FromAny[T any](a any) T {
	if a == nil {
		var zero T
		return zero
	}
	return a.(T)
}

func ToAnyList[T any](l EList[T]) any {
	switch c := any(l).(type) {
	case EList[any]:
		return c
	case EObjectList[T]:
		return ToObjectList(c, ToAny[T], FromAny[T])
	case ENotifyingList[T]:
		return ToNotifyingList(c, ToAny[T], FromAny[T])
	case EList[T]:
		return ToList(c, ToAny[T], FromAny[T])
	default:
		return nil
	}
}

func FromAnyList[T any](v any) EList[T] {
	switch c := v.(type) {
	case EList[T]:
		return c
	case EObjectList[any]:
		return ToObjectList(c, FromAny[T], ToAny[T])
	case ENotifyingList[any]:
		return ToNotifyingList(c, FromAny[T], ToAny[T])
	case EList[any]:
		return ToList(c, FromAny[T], ToAny[T])
	default:
		return nil
	}
}

func ToAnyMap[K comparable, V any](m EMap[K, V]) any {
	switch c := any(m).(type) {
	case EMap[any, any]:
		return c
	default:
		return ToMap(m, ToAny[K], FromAny[K], ToAny[V], FromAny[V])
	}
}

func FromAnyMap[K comparable, V any](v any) EMap[K, V] {
	switch m := v.(type) {
	case EMap[any, any]:
		return ToMap(m, FromAny[K], ToAny[K], FromAny[V], ToAny[V])
	case EMap[K, V]:
		return m
	default:
		return nil
	}
}