package stream

import "github.com/chebyrash/promise"

type task[T any] interface {
	fork() *promise.Promise[T]
	invoke() T
	cancel()
}
