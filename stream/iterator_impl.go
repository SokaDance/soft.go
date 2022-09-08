package stream

type DelegatingIterator struct {
	supplier func() Iterator
	it       Iterator
}

func NewDelegatingIterator(supplier func() Iterator) *DelegatingIterator {
	return &DelegatingIterator{supplier: supplier}
}

func (di *DelegatingIterator) get() Iterator {
	if di.it == nil {
		di.it = di.supplier()
	}
	return di.it
}

func (di *DelegatingIterator) TryAdvance(action func(any)) bool {
	return di.get().TryAdvance(action)
}

func (di *DelegatingIterator) TrySplit() Iterator {
	return di.get().TrySplit()
}

func (di *DelegatingIterator) ForEachRemaining(action func(any)) {
	di.get().ForEachRemaining(action)
}

func (di *DelegatingIterator) EstimateSize() int {
	return di.get().EstimateSize()
}
