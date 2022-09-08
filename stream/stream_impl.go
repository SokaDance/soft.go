package stream

type streamImpl struct {
	interfaces any
	source     *streamImpl
	iterator   Iterator
	previous   *streamImpl
	next       *streamImpl
	depth      int
	parallel   bool
	stateFull  bool
}

func newSourceStreamImpl(iterator Iterator, parallel bool) *streamImpl {
	s := &streamImpl{}
	s.interfaces = s
	s.iterator = iterator
	s.parallel = parallel
	s.source = s
	return s
}

func newNextStreamImpl(previous *streamImpl, stateFull bool) *streamImpl {
	s := &streamImpl{}
	s.interfaces = s
	s.stateFull = stateFull
	s.depth = previous.depth + 1
	s.previous = previous
	s.source = previous.source
	previous.next = s
	return s
}

func (s *streamImpl) asStream() Stream {
	return s.interfaces.(Stream)
}

func (s *streamImpl) IsParallel() bool {
	return s.parallel
}

func (s *streamImpl) Parallel() Stream {
	s.source.parallel = true
	return s.asStream()
}

func (s *streamImpl) Sequential() Stream {
	s.source.parallel = false
	return s.asStream()
}

type head struct {
	*streamImpl
}

func newHead(iterator Iterator, parallel bool) *head {
	h := &head{
		streamImpl: newSourceStreamImpl(iterator, parallel),
	}
	h.interfaces = h
	return h
}

type statelessOp struct {
	*streamImpl
}

type statefullOp struct {
	*streamImpl
}
