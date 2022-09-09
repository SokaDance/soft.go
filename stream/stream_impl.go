package stream

type stream struct {
	source   *stream
	previous *stream
	iterator Iterator
	parallel bool
	wrap     func(sink) sink
}

func newHead(iterator Iterator, wrap func(sink) sink) *stream {
	s := &stream{}
	s.iterator = iterator
	s.source = s
	return s
}

func newStream(previous *stream, wrap func(sink) sink) *stream {
	s := &stream{}
	s.previous = previous
	s.source = previous.source
	return s
}

func (s *stream) IsParallel() bool {
	return s.parallel
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

func (s *stream) AnyMatch(predicate func(any) bool) bool {
	return evaluate[bool](s, newMatchOperation(matchAny, predicate))
}

func evaluate[R any](s *stream, op operation[R]) R {
	if s.parallel {
		return op.evaluateParallel(s, s.source.iterator)
	} else {
		return op.evaluateSequential(s, s.source.iterator)
	}
}

func wrapAndCopyInto[S sink](stream *stream, s S, iterator Iterator) S {
	copyInto(wrapSink(stream, s), iterator)
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
	for c := stream; c != nil; c = c.previous {
		result = c.wrap(result)
	}
	return result.(S)
}
