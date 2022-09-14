package stream

import "math"

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

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func calcSize(size, skip, limit int) int {
	if size >= 0 {
		return max(0, min(size-skip, limit))
	}
	return -1
}

func newSliceStream(down *stream, skip int, limit int) *stream {
	normalizedLimit := limit
	if normalizedLimit < 0 {
		normalizedLimit = math.MaxInt
	}
	return newStream(down, func(s sink) sink {
		n := skip
		m := normalizedLimit
		return newChainedSink(s,
			begin(func(size int) {
				s.Begin(calcSize(size, skip, m))
			}),
			accept(func(a any) {
				if n == 0 {
					if m > 0 {
						m--
						s.Accept(a)
					}
				} else {
					n--
				}
			}),
			cancelationRequested(func() bool {
				return m == 0 || s.CancelationRequested()
			}),
		)
	})
}
