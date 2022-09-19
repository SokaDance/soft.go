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

type matchTask struct {
	*taskImpl[bool]
	op     *matchOperation
	stream *stream
}

func newMatchTask(op *matchOperation, stream *stream, it Iterator) *matchTask {
	t := &matchTask{
		taskImpl: newRootTask[bool](it),
		op:       op,
		stream:   stream,
	}
	t.setInternal(t)
	return t
}

func (t *matchTask) computeResult() bool {
	result := wrapAndCopyInto(t.stream, newMatchSink(t.op.kind, t.op.predicate), t.iterator()).value
	if result == t.op.kind.shortCircuitResult {
		t.setSharedResult(result)
	}
	return result
}

func (t *matchTask) defaultResult() bool {
	return !t.op.kind.shortCircuitResult
}

func (t *matchTask) newChildTask(it Iterator) *taskImpl[bool] {
	child := &matchTask{
		taskImpl: newChildTask[bool](t.taskImpl, it),
		op:       t.op,
		stream:   t.stream,
	}
	child.setInternal(child)
	return child.taskImpl
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
	return newMatchTask(op, stream, iterator).invoke()
}
