package ecore

type eMapEntryImpl struct {
	key   any
	value any
}

func (e *eMapEntryImpl) GetKey() any {
	return e.key
}

func (e *eMapEntryImpl) SetKey(key any) {
	e.key = key
}

func (e *eMapEntryImpl) GetValue() any {
	return e.value
}

func (e *eMapEntryImpl) SetValue(value any) {
	e.value = value
}

func newMapEntry(key any, value any) EMapEntry {
	return &eMapEntryImpl{key: key, value: value}
}
