package stream

type delegatingIterator struct {
	supplier func() Iterator
	it       Iterator
}

func NewDelegatingIterator(supplier func() Iterator) *delegatingIterator {
	return &delegatingIterator{supplier: supplier}
}

func (di *delegatingIterator) get() Iterator {
	if di.it == nil {
		di.it = di.supplier()
	}
	return di.it
}

func (di *delegatingIterator) TryAdvance(action func(any)) bool {
	return di.get().TryAdvance(action)
}

func (di *delegatingIterator) TrySplit() Iterator {
	return di.get().TrySplit()
}

func (di *delegatingIterator) ForEachRemaining(action func(any)) {
	di.get().ForEachRemaining(action)
}

func (di *delegatingIterator) EstimateSize() int {
	return di.get().EstimateSize()
}

type sliceIterator struct {
	slice         []any
	index         int // current index, modified on Advance/Split
	estimatedSize int // estimated size, to help to split evenly
}

func NewSliceIterator(slice []any) *sliceIterator {
	return &sliceIterator{slice: slice, index: 0, estimatedSize: -1}
}

func (it *sliceIterator) TryAdvance(action func(any)) bool {
	if it.index >= 0 && it.index < len(it.slice) {
		a := it.slice[it.index]
		it.index++
		action(a)
		return true
	}
	return false
}

func (it *sliceIterator) TrySplit() Iterator {
	lo := it.index
	mid := (lo + len(it.slice)) >> 1
	if lo >= mid {
		return nil
	}
	if it.estimatedSize == -1 {
		it.index = mid
		return &sliceIterator{slice: it.slice[lo:it.index], index: 0, estimatedSize: -1}
	}
	prefixEstimatedSize := it.estimatedSize >> 1
	it.estimatedSize -= prefixEstimatedSize
	it.index = mid
	return &sliceIterator{slice: it.slice[lo:it.index], estimatedSize: prefixEstimatedSize}
}

func (it *sliceIterator) ForEachRemaining(action func(any)) {
	for i := it.index; i < len(it.slice); i++ {
		action(it.slice[i])
	}
}

func (it *sliceIterator) EstimateSize() int {
	if it.estimatedSize < 0 {
		return len(it.slice)
	}
	return it.estimatedSize
}
