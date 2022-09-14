package stream

import "4d63.com/optional"

type Stream interface {
	// Iterator Returns an iterator for the elements of this stream.
	Iterator() Iterator

	// Parallel returns an equivalent stream that is parallel.  May return
	// itself, either because the stream was already parallel, or because
	// the underlying stream state was modified to be parallel.
	Parallel() Stream

	// Sequential returns an equivalent stream that is sequential.  May return
	// itself, either because the stream was already sequential, or because
	// the underlying stream state was modified to be sequential.
	Sequential() Stream

	// IsParallel returns whether this stream, if a terminal operation were to be executed,
	// would execute in parallel.  Calling this method after invoking an
	// terminal stream operation method may yield unpredictable results.
	IsParallel() bool

	// stateless operation
	Filter(predicate func(any) bool) Stream
	FlatMap(mapper func(any) Stream) Stream
	Map(mapper func(any) any) Stream
	Peek(action func(any)) Stream

	// statefull operation
	Distinct() Stream
	// Sorted(comparator func(any, any) int) Stream
	Limit(int) Stream
	Skip(int) Stream

	// terminal operation
	ToSlice() []any
	ForEach(action func(any))
	AnyMatch(predicate func(any) bool) bool
	AllMatch(predicate func(any) bool) bool
	NoneMatch(predicate func(any) bool) bool
	FindFirst() optional.Optional[any]
	FindAny() optional.Optional[any]
	Min(comparator func(any, any) int) optional.Optional[any]
	Max(comparator func(any, any) int) optional.Optional[any]
	Reduce(accumulator func(any, any) any) optional.Optional[any]
	ReduceWith(initValue any, accumulator func(any, any) any) any
	Count() int
}

func OfIterator(it Iterator) Stream {
	return newHead(it)
}

func OfSlice(slice []any) Stream {
	return OfIterator(NewSliceIterator(slice))
}
