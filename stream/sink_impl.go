package stream

type sinkImpl struct {
	begin                func(int)
	accept               func(any)
	end                  func()
	cancelationRequested func() bool
}

func (s *sinkImpl) Begin(size int) {
	s.begin(size)
}

func (s *sinkImpl) Accept(a any) {
	s.accept(a)
}

func (s *sinkImpl) End() {
	s.end()
}

func (s *sinkImpl) CancelationRequested() bool {
	return s.cancelationRequested()
}

type sinkOption func(s *sinkImpl)

func begin(onBegin func(int)) sinkOption {
	return func(s *sinkImpl) {
		s.begin = onBegin
	}
}

func cancelationRequested(onCancelationrequested func() bool) sinkOption {
	return func(s *sinkImpl) {
		s.cancelationRequested = onCancelationrequested
	}
}

func accept(onAccept func(any)) sinkOption {
	return func(s *sinkImpl) {
		s.accept = onAccept
	}
}
func end(onEnd func()) sinkOption {
	return func(s *sinkImpl) {
		s.end = onEnd
	}
}

type chainedSink struct {
	*sinkImpl
}

func newChainedSink(down sink, opt ...sinkOption) *chainedSink {
	s := &chainedSink{
		sinkImpl: &sinkImpl{
			begin:                down.Begin,
			accept:               down.Accept,
			end:                  down.End,
			cancelationRequested: down.CancelationRequested,
		},
	}
	for _, o := range opt {
		o(s.sinkImpl)
	}
	return s
}

type terminalSink struct {
	*sinkImpl
}

func newTerminalSink(accept func(any), opt ...sinkOption) *terminalSink {
	s := &terminalSink{
		sinkImpl: &sinkImpl{
			begin:                func(i int) {},
			accept:               accept,
			end:                  func() {},
			cancelationRequested: func() bool { return false },
		},
	}
	for _, o := range opt {
		o(s.sinkImpl)
	}
	return s
}
