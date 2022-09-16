package stream

type matchKind struct {
	stopOnPredicateMatches bool
	shortCircuitResult     bool
}

var matchAny matchKind = matchKind{true, true}
var matchAll matchKind = matchKind{false, false}
var matchNone matchKind = matchKind{true, false}

type matchSink struct {
	*terminalSink
	stop  bool
	value bool
}

func newMatchSink(kind matchKind, predicate func(any) bool) *matchSink {
	s := &matchSink{}
	s.value = !kind.shortCircuitResult
	s.terminalSink = newTerminalSink(
		func(a any) {
			if !s.stop && predicate(a) == kind.stopOnPredicateMatches {
				s.stop = true
				s.value = kind.shortCircuitResult
			}
		},
		cancelationRequested(func() bool {
			return s.stop
		}),
	)
	return s
}

type matchOperation struct {
	kind      matchKind
	predicate func(any) bool
}

func newMatchOperation(kind matchKind, predicate func(any) bool) *matchOperation {
	return &matchOperation{kind: kind, predicate: predicate}
}

func (op *matchOperation) evaluateSequential(stream *stream, iterator Iterator) bool {
	return wrapAndCopyInto(stream, newMatchSink(op.kind, op.predicate), iterator).value
}

func (op *matchOperation) evaluateParallel(stream *stream, iterator Iterator) bool {
	return newRootTask[bool](
		iterator,
		func(t *task[bool]) bool {
			result := wrapAndCopyInto(stream, newMatchSink(op.kind, op.predicate), iterator).value
			if result == op.kind.shortCircuitResult {
				t.setResult(result)
			}
			return result
		},
		func(t *task[bool]) bool {
			return !op.kind.shortCircuitResult
		},
	).Invoke()
}
