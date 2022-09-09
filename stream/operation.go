package stream

type operation[R any] interface {
	evaluateSequential(s *stream, iterator Iterator) R
	evaluateParallel(s *stream, iterator Iterator) R
}
