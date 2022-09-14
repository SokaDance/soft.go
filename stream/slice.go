package stream

type toSliceSink struct {
	*terminalSink
	slice []any
}

func newToSliceSink() *toSliceSink {
	s := &toSliceSink{slice: []any{}}
	s.terminalSink = newTerminalSink(func(a any) {
		s.slice = append(s.slice, a)
	})
	return s
}

type toSliceOperation struct {
}

func newToSliceOperation() *toSliceOperation {
	return &toSliceOperation{}
}

func (op *toSliceOperation) evaluateSequential(stream *stream, iterator Iterator) []any {
	return wrapAndCopyInto(stream, newToSliceSink(), iterator).slice
}

func (op *toSliceOperation) evaluateParallel(stream *stream, iterator Iterator) []any {
	return nil
}
