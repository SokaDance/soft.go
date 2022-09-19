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

type findTask struct {
	*taskImpl[optional.Optional[any]]
	stream        *stream
	mustFindFirst bool
}

func newFindTask(stream *stream, it Iterator, mustFindFirst bool) *findTask {
	t := &findTask{
		taskImpl:      newRootTask[optional.Optional[any]](it),
		stream:        stream,
		mustFindFirst: mustFindFirst,
	}
	t.setInternal(t)
	return t
}

func (t *findTask) computeResult() optional.Optional[any] {
	result := wrapAndCopyInto(t.stream, newFindSink(), t.iterator()).get()
	if result != nil {
		if !t.mustFindFirst {
			t.setSharedResult(result)
		} else {
			t.foundResult(result)
		}
	}
	return result
}

func (t *findTask) foundResult(result optional.Optional[any]) {
	if t.isLeftMost() {
		t.setSharedResult(result)
	} else {
		t.cancelLaterTasks()
	}
}

func (t *findTask) defaultResult() optional.Optional[any] {
	return optional.Empty[any]()
}

func (t *findTask) newChildTask(it Iterator) *taskImpl[optional.Optional[any]] {
	child := &findTask{
		taskImpl:      newChildTask[optional.Optional[any]](t.taskImpl, it),
		stream:        t.stream,
		mustFindFirst: t.mustFindFirst,
	}
	child.setInternal(child)
	return child.taskImpl
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
	return newFindTask(stream, iterator, op.mustFindFirst).invoke()
}
