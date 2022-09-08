package stream

type Iterator interface {
	// TryAdvance
	// If a remaining element exists, performs the given action on it,
	// returning true else returns false. If this iterator
	// is ordered the action is performed on the next element in
	// encounter order.
	TryAdvance(action func(any)) bool

	// TrySplit
	// If this iterator can be partitioned, returns an iterator
	// covering elements, that will, upon return from this method, not
	// be covered by this iterator.
	//
	// Unless this iterator covers an infinite number of elements,
	// repeated calls to TrySplit() must eventually return nill.
	//
	// An ideal TrySplit method efficiently (without traversal)
	// divides its elements exactly in half, allowing
	// balanced parallel computation.
	TrySplit() Iterator

	// ForEachRemaining performs the given action for each remaining element,
	// sequentially in the current goroutine, until all elements have been processed
	ForEachRemaining(action func(any))

	// EstimateSize returns an estimate of the number of elements that would be encountered
	// by a ForEachRemaining traversal, or returns -1 if infinite, unknown, or too expensive to compute.
	EstimateSize() int
}
