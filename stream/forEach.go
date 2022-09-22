package stream

import (
	"github.com/chebyrash/promise"
)

type forEachSink struct {
	*terminalSink
}

func newForEachSink(action func(any)) *forEachSink {
	return &forEachSink{terminalSink: newTerminalSink(action)}
}

type forEachOperation struct {
	action func(any)
}

func newForEachOperation(action func(any)) *forEachOperation {
	return &forEachOperation{action: action}
}

type forEachTask struct {
	*taskImpl[any]
	stream *stream
	op     *forEachOperation
}

func newForEachTask(op *forEachOperation, stream *stream, iterator Iterator) *forEachTask {
	t := &forEachTask{
		taskImpl: newRootTask[any](iterator),
		stream:   stream,
		op:       op,
	}
	t.setInternal(t)
	return t
}

func (t *forEachTask) computeTask(resolve func(any), reject func(error)) {
	rightSplit := t.it
	sizeEstimate := rightSplit.EstimateSize()
	sizeThreshold := t.targetSize
	if sizeThreshold == 0 {
		t.targetSize = t.suggestTargetSize(sizeEstimate)
		sizeThreshold = t.targetSize
	}
	forkRight := false
	taskSink := newForEachSink(t.op.action)
	children := []*promise.Promise[any]{}
	for !taskSink.CancelationRequested() {
		if sizeEstimate <= sizeThreshold {
			copyInto(taskSink, rightSplit)
			break
		}

		leftSplit := rightSplit.TrySplit()
		if leftSplit == nil {
			copyInto(taskSink, rightSplit)
			break
		}

		leftTask := &forEachTask{
			taskImpl: newChildTask(t.taskImpl, leftSplit),
			stream:   t.stream,
			op:       t.op,
		}
		leftTask.setInternal(leftTask)
		var taskToFork *forEachTask
		if forkRight {
			forkRight = false
			rightSplit = leftSplit
			taskToFork = t
			t = leftTask
		} else {
			forkRight = true
			taskToFork = leftTask
		}
		promise := taskToFork.fork()
		children = append(children, promise)
		sizeEstimate = rightSplit.EstimateSize()
	}
	if len(children) > 0 {
		promise.All(children...).Await()
	}
	resolve(nil)
}

func (op *forEachOperation) evaluateSequential(stream *stream, iterator Iterator) any {
	wrapAndCopyInto(stream, newForEachSink(op.action), iterator)
	return nil
}

func (op *forEachOperation) evaluateParallel(stream *stream, iterator Iterator) any {
	return newForEachTask(op, stream, iterator).invoke()
}
