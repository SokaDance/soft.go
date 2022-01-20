package ecore

func ToMap[KO comparable, KT comparable, VO any, VT any](
	delegate EMap[KO, VO],
	convertKeyTo func(KO) KT,
	convertKeyFrom func(KT) KO,
	convertValueTo func(VO) VT,
	convertValueFrom func(VT) VO) EMap[KT, VT] {
	if md, isDelegate := delegate.(eMapDelegate[KT, KO, VT, VO]); isDelegate {
		return md.GetDelegate()
	}
	return &eMapDelegateImpl[KO, KT, VO, VT]{
		delegate:         delegate,
		convertKeyTo:     convertKeyTo,
		convertKeyFrom:   convertKeyFrom,
		convertValueTo:   convertValueTo,
		convertValueFrom: convertValueFrom,
	}
}

type eMapDelegateImpl[KO comparable, KT comparable, VO any, VT any] struct {
	delegate         EMap[KO, VO]
	convertKeyTo     func(KO) KT
	convertKeyFrom   func(KT) KO
	convertValueTo   func(VO) VT
	convertValueFrom func(VT) VO
}

type eMapDelegateImplIterator[KO comparable, KT comparable, VO any, VT any] struct {
	delegate       EIterator[any]
	convertKeyTo   func(KO) KT
	convertValueTo func(VO) VT
}

func (it *eMapDelegateImplIterator[KO, KT, VO, VT]) HasNext() bool {
	return it.delegate.HasNext()
}

func (it *eMapDelegateImplIterator[KO, KT, VO, VT]) Next() any {
	e := it.delegate.Next().(EMapEntry[KO, VO])
	return newMapEntry(
		it.convertKeyTo(e.GetKey()),
		it.convertValueTo(e.GetValue()),
	)
}

func (m *eMapDelegateImpl[KO, KT, VO, VT]) GetDelegate() EMap[KO, VO] {
	return m.delegate
}

func (m *eMapDelegateImpl[KO, KT, VO, VT]) toDelegateMapEntry(a any) EMapEntry[KO, VO] {
	switch e := a.(type) {
	case EMapEntry[KT, VT]:
		return newMapEntry(m.convertKeyFrom(e.GetKey()), m.convertValueFrom(e.GetValue()))
	case EMapEntry[KO, VO]:
		return e
	}
	return nil
}

func (m *eMapDelegateImpl[KO, KT, VO, VT]) toTargetMapEntry(a any) EMapEntry[KT, VT] {
	switch e := a.(type) {
	case EMapEntry[KT, VT]:
		return e
	case EMapEntry[KO, VO]:
		return newMapEntry(m.convertKeyTo(e.GetKey()), m.convertValueTo(e.GetValue()))
	}
	return nil
}

func (m *eMapDelegateImpl[KO, KT, VO, VT]) toDelegateCollection(c ECollection[any]) ECollection[any] {
	entries := []any{}
	for it := c.Iterator(); it.HasNext(); {
		e := it.Next()
		entries = append(entries, m.toDelegateMapEntry(e))
	}
	return NewImmutableEList(entries)
}

func (m *eMapDelegateImpl[KO, KT, VO, VT]) Add(a any) bool {
	entry := m.toDelegateMapEntry(a)
	if !m.delegate.ContainsKey(entry.GetKey()) {
		m.delegate.Add(entry)
		return true
	}
	return false
}

func (m *eMapDelegateImpl[KO, KT, VO, VT]) AddAll(c ECollection[any]) bool {
	return m.delegate.AddAll(m.toDelegateCollection(c))
}

func (m *eMapDelegateImpl[KO, KT, VO, VT]) Contains(a any) bool {
	entry := m.toDelegateMapEntry(a)
	return m.delegate.ContainsKey(entry.GetKey())
}

func (m *eMapDelegateImpl[KO, KT, VO, VT]) Clear() {
	m.delegate.Clear()
}

func (m *eMapDelegateImpl[KO, KT, VO, VT]) Empty() bool {
	return m.delegate.Empty()
}

func (m *eMapDelegateImpl[KO, KT, VO, VT]) Size() int {
	return m.delegate.Size()
}

func (m *eMapDelegateImpl[KO, KT, VO, VT]) Get(i int) any {
	return m.toTargetMapEntry(m.delegate.Get(i))
}

func (m *eMapDelegateImpl[KO, KT, VO, VT]) Set(i int, a any) any {
	return m.toTargetMapEntry(m.delegate.Set(i, m.toDelegateMapEntry(a)))
}

func (m *eMapDelegateImpl[KO, KT, VO, VT]) IndexOf(a any) int {
	i := 0
	e := m.toDelegateMapEntry(a)
	for it := m.delegate.Iterator(); it.HasNext(); {
		o := it.Next().(EMapEntry[KO, VO])
		if o.GetKey() == e.GetKey() {
			return i
		}
		i++
	}
	return -1
}

func (m *eMapDelegateImpl[KO, KT, VO, VT]) Insert(index int, a any) bool {
	return m.delegate.Insert(index, m.toDelegateMapEntry(a))
}

func (m *eMapDelegateImpl[KO, KT, VO, VT]) InsertAll(index int, c ECollection[any]) bool {
	return m.delegate.InsertAll(index, m.toDelegateCollection(c))
}

func (m *eMapDelegateImpl[KO, KT, VO, VT]) Iterator() EIterator[any] {
	return &eMapDelegateImplIterator[KO, KT, VO, VT]{delegate: m.delegate.Iterator(), convertKeyTo: m.convertKeyTo, convertValueTo: m.convertValueTo}
}

func (m *eMapDelegateImpl[KO, KT, VO, VT]) MoveIndex(oldIndex, newIndex int) any {
	return m.toTargetMapEntry(m.delegate.MoveIndex(oldIndex, newIndex))
}

func (m *eMapDelegateImpl[KO, KT, VO, VT]) MoveObject(index int, a any) {
	if oldIndex := m.IndexOf(a); oldIndex != -1 {
		m.MoveIndex(oldIndex, index)
	}
}

func (m *eMapDelegateImpl[KO, KT, VO, VT]) Remove(a any) bool {
	if index := m.IndexOf(a); index != -1 {
		m.RemoveAt(index)
		return true
	}
	return false
}

func (m *eMapDelegateImpl[KO, KT, VO, VT]) RemoveAll(c ECollection[any]) bool {
	removed := false
	for it := c.Iterator(); it.HasNext(); {
		removed = removed || m.Remove(it.Next())
	}
	return removed
}

func (m *eMapDelegateImpl[KO, KT, VO, VT]) RetainAll(c ECollection[any]) bool {
	return true
}

func (m *eMapDelegateImpl[KO, KT, VO, VT]) RemoveAt(index int) any {
	return m.toTargetMapEntry(m.delegate.RemoveAt(index))
}

func (m *eMapDelegateImpl[KO, KT, VO, VT]) ToArray() []any {
	d := m.delegate.ToArray()
	a := make([]any, len(d))
	for i, e := range d {
		a[i] = m.toTargetMapEntry(e)
	}
	return a
}

func (m *eMapDelegateImpl[KO, KT, VO, VT]) GetValue(key KT) VT {
	return m.convertValueTo(m.delegate.GetValue(m.convertKeyFrom(key)))
}

func (m *eMapDelegateImpl[KO, KT, VO, VT]) Put(key KT, value VT) {
	m.delegate.Put(m.convertKeyFrom(key), m.convertValueFrom(value))
}

func (m *eMapDelegateImpl[KO, KT, VO, VT]) RemoveKey(key KT) VT {
	return m.convertValueTo(m.delegate.RemoveKey(m.convertKeyFrom(key)))
}

func (m *eMapDelegateImpl[KO, KT, VO, VT]) ContainsValue(value VT) bool {
	return m.delegate.ContainsValue(m.convertValueFrom(value))
}

func (m *eMapDelegateImpl[KO, KT, VO, VT]) ContainsKey(key KT) bool {
	return m.delegate.ContainsKey(m.convertKeyFrom(key))
}

func (m *eMapDelegateImpl[KO, KT, VO, VT]) ToMap() map[KT]VT {
	d := m.delegate.ToMap()
	r := map[KT]VT{}
	for k, v := range d {
		r[m.convertKeyTo(k)] = m.convertValueTo(v)
	}
	return r
}
