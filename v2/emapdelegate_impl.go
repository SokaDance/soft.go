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
	delegate       EIterator[EMapEntry[KO, VO]]
	convertKeyTo   func(KO) KT
	convertValueTo func(VO) VT
}

func (it *eMapDelegateImplIterator[KO, KT, VO, VT]) HasNext() bool {
	return it.delegate.HasNext()
}

func (it *eMapDelegateImplIterator[KO, KT, VO, VT]) Next() EMapEntry[KT, VT] {
	e := it.delegate.Next()
	return &eMapEntryImpl[KT, VT]{
		key:   it.convertKeyTo(e.GetKey()),
		value: it.convertValueTo(e.GetValue()),
	}
}

func (m *eMapDelegateImpl[KO, KT, VO, VT]) GetDelegate() EMap[KO, VO] {
	return m.delegate
}

func (m *eMapDelegateImpl[KO, KT, VO, VT]) Add(e EMapEntry[KT, VT]) bool {
	entry := &eMapEntryImpl[KO, VO]{
		key:   m.convertKeyFrom(e.GetKey()),
		value: m.convertValueFrom(e.GetValue()),
	}
	if !m.delegate.ContainsKey(entry.GetKey()) {
		m.delegate.Add(entry)
		return true
	}
	return false
}

func (m *eMapDelegateImpl[KO, KT, VO, VT]) fromCollection(c ECollection[EMapEntry[KT, VT]]) ECollection[EMapEntry[KO, VO]] {
	entries := []EMapEntry[KO, VO]{}
	for it := c.Iterator(); it.HasNext(); {
		e := it.Next()
		entries = append(entries, &eMapEntryImpl[KO, VO]{key: m.convertKeyFrom(e.GetKey()), value: m.convertValueFrom(e.GetValue())})
	}
	return NewImmutableEList(entries)
}

func (m *eMapDelegateImpl[KO, KT, VO, VT]) AddAll(c ECollection[EMapEntry[KT, VT]]) bool {
	return m.delegate.AddAll(m.fromCollection(c))
}

func (m *eMapDelegateImpl[KO, KT, VO, VT]) Contains(e EMapEntry[KT, VT]) bool {
	return m.delegate.ContainsKey(m.convertKeyFrom(e.GetKey()))
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

func (m *eMapDelegateImpl[KO, KT, VO, VT]) Get(i int) EMapEntry[KT, VT] {
	e := m.delegate.Get(i)
	return &eMapEntryImpl[KT, VT]{key: m.convertKeyTo(e.GetKey()), value: m.convertValueTo(e.GetValue())}
}

func (m *eMapDelegateImpl[KO, KT, VO, VT]) Set(i int, e EMapEntry[KT, VT]) EMapEntry[KT, VT] {
	n := &eMapEntryImpl[KO, VO]{key: m.convertKeyFrom(e.GetKey()), value: m.convertValueFrom(e.GetValue())}
	o := m.delegate.Set(i, n)
	return &eMapEntryImpl[KT, VT]{key: m.convertKeyTo(o.GetKey()), value: m.convertValueTo(o.GetValue())}
}

func (m *eMapDelegateImpl[KO, KT, VO, VT]) IndexOf(e EMapEntry[KT, VT]) int {
	k := m.convertKeyFrom(e.GetKey())
	i := 0
	for it := m.delegate.Iterator(); it.HasNext(); {
		if it.Next().GetKey() == k {
			return i
		}
		i++
	}
	return -1
}

func (m *eMapDelegateImpl[KO, KT, VO, VT]) Insert(index int, e EMapEntry[KT, VT]) bool {
	n := &eMapEntryImpl[KO, VO]{key: m.convertKeyFrom(e.GetKey()), value: m.convertValueFrom(e.GetValue())}
	b := m.delegate.Insert(index, n)
	return b
}

func (m *eMapDelegateImpl[KO, KT, VO, VT]) InsertAll(index int, c ECollection[EMapEntry[KT, VT]]) bool {
	return m.delegate.InsertAll(index, m.fromCollection(c))
}

func (m *eMapDelegateImpl[KO, KT, VO, VT]) Iterator() EIterator[EMapEntry[KT, VT]] {
	return &eMapDelegateImplIterator[KO, KT, VO, VT]{delegate: m.delegate.Iterator(), convertKeyTo: m.convertKeyTo, convertValueTo: m.convertValueTo}
}

func (m *eMapDelegateImpl[KO, KT, VO, VT]) MoveIndex(oldIndex, newIndex int) EMapEntry[KT, VT] {
	e := m.delegate.MoveIndex(oldIndex, newIndex)
	return &eMapEntryImpl[KT, VT]{key: m.convertKeyTo(e.GetKey()), value: m.convertValueTo(e.GetValue())}
}

func (m *eMapDelegateImpl[KO, KT, VO, VT]) MoveObject(index int, e EMapEntry[KT, VT]) {
	if oldIndex := m.IndexOf(e); oldIndex != -1 {
		m.MoveIndex(oldIndex, index)
	}
}

func (m *eMapDelegateImpl[KO, KT, VO, VT]) Remove(e EMapEntry[KT, VT]) bool {
	if index := m.IndexOf(e); index != -1 {
		m.RemoveAt(index)
		return true
	}
	return false
}

func (m *eMapDelegateImpl[KO, KT, VO, VT]) RemoveAll(c ECollection[EMapEntry[KT, VT]]) bool {
	removed := false
	for it := c.Iterator(); it.HasNext(); {
		removed = removed || m.Remove(it.Next())
	}
	return removed
}

func (m *eMapDelegateImpl[KO, KT, VO, VT]) RetainAll(c ECollection[EMapEntry[KT, VT]]) bool {
	return true
}

func (m *eMapDelegateImpl[KO, KT, VO, VT]) RemoveAt(index int) EMapEntry[KT, VT] {
	e := m.delegate.RemoveAt(index)
	return &eMapEntryImpl[KT, VT]{key: m.convertKeyTo(e.GetKey()), value: m.convertValueTo(e.GetValue())}
}

func (m *eMapDelegateImpl[KO, KT, VO, VT]) ToArray() []EMapEntry[KT, VT] {
	d := m.delegate.ToArray()
	a := make([]EMapEntry[KT, VT], len(d))
	for i, e := range d {
		a[i] = &eMapEntryImpl[KT, VT]{key: m.convertKeyTo(e.GetKey()), value: m.convertValueTo(e.GetValue())}
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
