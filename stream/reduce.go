package stream

type countSink struct {
	*terminalSink
	count int
}

func newCountSink() *countSink {
	s := &countSink{}
	s.terminalSink = newTerminalSink(func(a any) {
		s.count++
	}, begin(func(i int) {
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
