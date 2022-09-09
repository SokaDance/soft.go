package stream

type operation interface {
	evaluateSequential(iterator Iterator)
	evaluateParallel(iterator Iterator)
}
