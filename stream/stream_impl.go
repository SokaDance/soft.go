package stream

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

func (s *stream) Map(mapper func(any) any) Stream {
	return newStream(s, func(down sink) sink {
		return newChainedSink(down,
			accept(func(a any) {
				down.Accept(mapper(a))
			}),
		)
	})
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

func wrapAndCopyInto[S sink](stream *stream, s S, iterator Iterator) S {
	copyInto(wrapSink[S](stream, s), iterator)
	return s
}

func copyInto[S sink](s S, iterator Iterator) bool {
	canceled := false
	s.Begin(iterator.EstimateSize())
	for finished := false; !finished; {
		canceled = s.CancelationRequested()
		finished = canceled || !iterator.TryAdvance(s.Accept)
	}
	s.End()
	return canceled
}

func wrapSink[S sink](stream *stream, s S) S {
	var result sink = s
	for c := stream; c.previous != nil; c = c.previous {
		result = c.wrap(result)
	}
	return result.(S)
}
