package stream

import "4d63.com/optional"

type countSink struct {
	*terminalSink
	count int
}

func newCountSink() *countSink {
	s := &countSink{}
	s.terminalSink = newTerminalSink(
		func(a any) {
			s.count++
		},
		begin(func(i int) {
			s.count = 0
		},
		))
	return s
}

type countOperation struct {
}

func newCountOperation() *countOperation {
	return &countOperation{}
}

func (op *countOperation) evaluateSequential(stream *stream, iterator Iterator) int {
	return wrapAndCopyInto(stream, newCountSink(), iterator).count
}

func (op *countOperation) evaluateParallel(stream *stream, iterator Iterator) int {
	return 0
}

type reducingSink struct {
	*terminalSink
	state any
	empty bool
}

func newReducingSink(reducer func(any, any) any) *reducingSink {
	s := &reducingSink{}
	s.terminalSink = newTerminalSink(
		func(a any) {
			if s.empty {
				s.empty = false
				s.state = a
			} else {
				s.state = reducer(s.state, a)
			}
		},
		begin(func(i int) {
			s.empty = true
			s.state = nil
		},
		))
	return s
}

func (s *reducingSink) get() optional.Optional[any] {
	if s.empty {
		return optional.Empty[any]()
	}
	return optional.Of(s.state)
}

type reducingOperation struct {
	reducer func(any, any) any
}

func newReducingOperation(reducer func(any, any) any) *reducingOperation {
	return &reducingOperation{reducer: reducer}
}

func (op *reducingOperation) evaluateSequential(stream *stream, iterator Iterator) optional.Optional[any] {
	return wrapAndCopyInto(stream, newReducingSink(op.reducer), iterator).get()
}

func (op *reducingOperation) evaluateParallel(stream *stream, iterator Iterator) optional.Optional[any] {
	return nil
}

type reducingSinkWithIdentity struct {
	*terminalSink
	state any
}

func newReducingSinkWithIdentity(identity any, reducer func(any, any) any) *reducingSinkWithIdentity {
	s := &reducingSinkWithIdentity{}
	s.terminalSink = newTerminalSink(
		func(a any) {
			s.state = reducer(s.state, a)
		},
		begin(func(i int) {
			s.state = identity
		},
		))
	return s
}

type reducingOperationWithIdentity struct {
	identity any
	reducer  func(any, any) any
}

func newReducingOperationWithIdentity(identity any, reducer func(any, any) any) *reducingOperationWithIdentity {
	return &reducingOperationWithIdentity{identity: identity, reducer: reducer}
}

func (op *reducingOperationWithIdentity) evaluateSequential(stream *stream, iterator Iterator) any {
	return wrapAndCopyInto(stream, newReducingSinkWithIdentity(op.identity, op.reducer), iterator).state
}

func (op *reducingOperationWithIdentity) evaluateParallel(stream *stream, iterator Iterator) any {
	return nil
}
