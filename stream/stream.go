package stream

type Stream interface {
	// stateless operation
	Filter(predicate func(any) bool) Stream
	FlatMap(mapper func(any) Stream) Stream
	Map(mapper func(any) any) Stream
	Peek(action func(any)) Stream

	// statefull operation
	Sorted(comparator func(any, any) int) Stream
	Limit(int) Stream
	Skip(int) Stream

	// terminal operation
	ForEach(action func(any))
}
