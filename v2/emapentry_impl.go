package ecore

type eMapEntryImpl[K comparable, V any] struct {
	key   K
	value V
}

func (e *eMapEntryImpl[K, V]) GetKey() K {
	return e.key
}

func (e *eMapEntryImpl[K, V]) SetKey(key K) {
	e.key = key
}

func (e *eMapEntryImpl[K, V]) GetValue() V {
	return e.value
}

func (e *eMapEntryImpl[K, V]) SetValue(value V) {
	e.value = value
}

func (e *eMapEntryImpl[K, V]) SetAnyKey(key any) {
	e.key = key.(K)
}

func (e *eMapEntryImpl[K, V]) SetAnyValue(value any) {
	e.value = value.(V)
}

func newMapEntry[K comparable, V any](key K, value V) EMapEntry[K, V] {
	return &eMapEntryImpl[K, V]{key: key, value: value}
}
