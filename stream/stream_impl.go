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
	return false
}

func (s *stream) evaluate(op operation) {
	if s.parallel {
		op.evaluateParallel(s.source.iterator)
	} else {
		op.evaluateSequential(s.source.iterator)
	}
}
