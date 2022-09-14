package stream

import "4d63.com/optional"

type stream struct {
	source   *stream
	previous *stream
	iterator Iterator
	parallel bool
	wrap     func(sink) sink
}

func newHead(iterator Iterator) *stream {
	s := &stream{}
	s.iterator = iterator
	s.source = s
	return s
}

func newStream(previous *stream, wrap func(sink) sink) *stream {
	s := &stream{}
	s.previous = previous
	s.source = previous.source
	s.wrap = wrap
	return s
}

func (s *stream) Iterator() Iterator {
	if s == s.source {
		return s.iterator
	} else {
		return nil
	}
}

func (s *stream) IsParallel() bool {
	return s.source.parallel
}

func (s *stream) Parallel() Stream {
	s.source.parallel = true
	return s
}

func (s *stream) Sequential() Stream {
	s.source.parallel = false
	return s
}

func (s *stream) Filter(predicate func(any) bool) Stream {
	return newStream(s, func(down sink) sink {
		return newChainedSink(down,
			begin(func(int) {
				down.Begin(-1)
			}),
			accept(func(a any) {
				if predicate(a) {
					down.Accept(a)
				}
			}),
		)
	})
}

func (s *stream) FlatMap(mapper func(any) Stream) Stream {
	return newStream(s, func(down sink) sink {
		cancellationRequestedCalled := false
		return newChainedSink(down,
			begin(func(int) {
				down.Begin(-1)
			}),
			accept(func(a any) {
				if result := mapper(a); result != nil {
					if !cancellationRequestedCalled {
						result.Sequential().ForEach(down.Accept)
					} else {
						it := result.Sequential().Iterator()
						for !down.CancelationRequested() && it.TryAdvance(down.Accept) {
						}
					}
				}
			}),
			cancelationRequested(func() bool {
				cancellationRequestedCalled = true
				return down.CancelationRequested()
			}),
		)
	})
}

func (s *stream) Map(mapper func(any) any) Stream {
	return newStream(s, func(down sink) sink {
		return newChainedSink(down,
			accept(func(a any) {
				down.Accept(mapper(a))
			}),
		)
	})
}

func (s *stream) Peek(action func(any)) Stream {
	return newStream(s, func(down sink) sink {
		return newChainedSink(down,
			accept(func(a any) {
				action(a)
				down.Accept(a)
			}),
		)
	})
}

func (s *stream) Distinct() Stream {
	return newStream(s, func(down sink) sink {
		var seen map[any]struct{}
		return newChainedSink(down,
			begin(func(i int) {
				seen = map[any]struct{}{}
				down.Begin(-1)
			}),
			end(func() {
				seen = nil
			}),
			accept(func(a any) {
				_, exists := seen[a]
				seen[a] = struct{}{}
				if !exists {
					down.Accept(a)
				}
			}),
		)
	})
}

func (s *stream) ToSlice() []any {
	return evaluate[[]any](s, newToSliceOperation())
}

func (s *stream) ForEach(action func(any)) {
	evaluate[any](s, newForEachOperation(action))
}

func (s *stream) AnyMatch(predicate func(any) bool) bool {
	return evaluate[bool](s, newMatchOperation(matchAny, predicate))
}

func (s *stream) AllMatch(predicate func(any) bool) bool {
	return evaluate[bool](s, newMatchOperation(matchAll, predicate))
}

func (s *stream) NoneMatch(predicate func(any) bool) bool {
	return evaluate[bool](s, newMatchOperation(matchNone, predicate))
}

func (s *stream) FindFirst() optional.Optional[any] {
	return evaluate[optional.Optional[any]](s, newFindOperation(true))
}

func (s *stream) FindAny() optional.Optional[any] {
	return evaluate[optional.Optional[any]](s, newFindOperation(false))
}

func (s *stream) Min(comparator func(any, any) int) optional.Optional[any] {
	return s.Reduce(func(a, b any) any {
		if comparator(a, b) <= 0 {
			return a
		}
		return b
	})
}

func (s *stream) Max(comparator func(any, any) int) optional.Optional[any] {
	return s.Reduce(func(a, b any) any {
		if comparator(a, b) >= 0 {
			return a
		}
		return b
	})
}

func (s *stream) Reduce(accumulator func(any, any) any) optional.Optional[any] {
	return evaluate[optional.Optional[any]](s, newReducingOperation(accumulator))
}

func (s *stream) ReduceWith(initValue any, accumulator func(any, any) any) any {
	return evaluate[any](s, newReducingOperationWithIdentity(initValue, accumulator))
}

func (s *stream) Count() int {
	return evaluate[int](s, newCountOperation())
}

func evaluate[R any](stream *stream, op operation[R]) R {
	if stream.IsParallel() {
		return op.evaluateParallel(stream, stream.source.iterator)
	} else {
		return op.evaluateSequential(stream, stream.source.iterator)
	}
}

func copyInto(s sink, iterator Iterator) bool {
	canceled := false
	s.Begin(iterator.EstimateSize())
	for finished := false; !finished; {
		canceled = s.CancelationRequested()
		finished = canceled || !iterator.TryAdvance(s.Accept)
	}
	s.End()
	return canceled
}

func wrapSink(stream *stream, s sink) sink {
	var result sink = s
	for c := stream; c.previous != nil; c = c.previous {
		result = c.wrap(result)
	}
	return result
}

func wrapAndCopyInto[S sink](stream *stream, s S, iterator Iterator) S {
	copyInto(wrapSink(stream, s), iterator)
	return s
}
