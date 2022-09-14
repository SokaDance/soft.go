package stream

import "math"

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

type concatIterator struct {
	it1         Iterator
	it2         Iterator
	beforeSplit bool
}

func NewConcatIterator(it1, it2 Iterator) *concatIterator {
	return &concatIterator{
		it1:         it1,
		it2:         it2,
		beforeSplit: true,
	}
}

func (it *concatIterator) TrySplit() Iterator {
	defer func() { it.beforeSplit = false }()
	if it.beforeSplit {
		return it.it1
	}
	return it.it2.TrySplit()
}

func (it *concatIterator) TryAdvance(action func(any)) bool {
	var hasNext bool
	if it.beforeSplit {
		hasNext = it.it1.TryAdvance(action)
		if !hasNext {
			it.beforeSplit = false
			hasNext = it.it2.TryAdvance(action)
		}
	} else {
		hasNext = it.it2.TryAdvance(action)
	}
	return hasNext
}

func (it *concatIterator) ForEachRemaining(action func(any)) {
	if it.beforeSplit {
		it.it1.ForEachRemaining(action)
	}
	it.it2.ForEachRemaining(action)
}

func (it *concatIterator) EstimateSize() int {
	if it.beforeSplit {
		// If one or both estimates are Long.MAX_VALUE then the sum
		// will either be Long.MAX_VALUE or overflow to a negative value
		size := it.it1.EstimateSize() + it.it2.EstimateSize()
		if size >= 0 {
			return size
		}
		return math.MaxInt
	}
	return it.it2.EstimateSize()
}
