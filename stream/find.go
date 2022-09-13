package stream

import "4d63.com/optional"

type findSink struct {
	*terminalSink
	hasValue bool
	value    any
}

func newFindSink() *findSink {
	s := &findSink{}
	s.terminalSink = newTerminalSink(
		func(a any) {
			if !s.hasValue {
				s.hasValue = true
				s.value = a
			}
		},
		cancelationRequested(func() bool {
			return s.hasValue
		}))
	return s
}

func (s *findSink) get() optional.Optional[any] {
	if s.hasValue {
		return optional.Of(s.value)
	}
	return optional.Empty[any]()
}

type findOperation struct {
	mustFindFirst bool
}

func newFindOperation(mustFindFirst bool) *findOperation {
	return &findOperation{mustFindFirst: mustFindFirst}
}

func (op *findOperation) evaluateSequential(stream *stream, iterator Iterator) optional.Optional[any] {
	return wrapAndCopyInto(stream, newFindSink(), iterator).get()
}

func (op *findOperation) evaluateParallel(stream *stream, iterator Iterator) optional.Optional[any] {
	return optional.Empty[any]()
}
