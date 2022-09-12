package stream

type forEachSink struct {
	*terminalSink
}

func newForEachSink(action func(any)) *forEachSink {
	return &forEachSink{terminalSink: newTerminalSink(action)}
}

func (s *forEachSink) get() any {
	return nil
}

type forEachOperation struct {
	action func(any)
}

func newForEachOperation(action func(any)) *forEachOperation {
	return &forEachOperation{action: action}
}

func (op *forEachOperation) evaluateSequential(stream *stream, iterator Iterator) any {
	return wrapAndCopyInto(stream, newForEachSink(op.action), iterator).get()
}

func (op *forEachOperation) evaluateParallel(stream *stream, iterator Iterator) any {
	return false
}
