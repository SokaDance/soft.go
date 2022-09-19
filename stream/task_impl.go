package stream

import (
	"runtime"
	"sync/atomic"

	"github.com/chebyrash/promise"
)

type taskInternal[T any] interface {
	computeResult() T
	defaultResult() T
	newChildTask(Iterator) *taskImpl[T]
	onCompletion(result T)
}

type taskImpl[T any] struct {
	internal     taskInternal[T]
	parent       *taskImpl[T]
	targetSize   int
	sharedResult *atomic.Value
	it           Iterator
	canceled     bool
	leftChild    *taskImpl[T]
	rightChild   *taskImpl[T]
}

func newRootTask[T any](iterator Iterator) *taskImpl[T] {
	t := &taskImpl[T]{
		sharedResult: &atomic.Value{},
		it:           iterator,
		targetSize:   0,
	}
	return t
}

func newChildTask[T any](parent *taskImpl[T], iterator Iterator) *taskImpl[T] {
	t := &taskImpl[T]{
		parent:       parent,
		targetSize:   parent.targetSize,
		sharedResult: parent.sharedResult,
		it:           iterator,
	}
	return t
}

func (t *taskImpl[T]) asInternal() taskInternal[T] {
	return t.internal
}

func (t *taskImpl[T]) setInternal(internal taskInternal[T]) {
	t.internal = internal
}

func (t *taskImpl[T]) fork() *promise.Promise[T] {
	return promise.Then[T](
		promise.New[T](t.computeTask),
		func(result T) T {
			t.asInternal().onCompletion(result)
			return result
		},
	)
}

func (t *taskImpl[T]) invoke() T {
	result, _ := t.fork().Await()
	return result
}

func (t *taskImpl[T]) iterator() Iterator {
	return t.it
}

func (t *taskImpl[T]) setSharedResult(value T) {
	if any(value) != nil {
		t.sharedResult.Store(value)
	}
}

func (t *taskImpl[T]) computeTask(resolve func(T), reject func(error)) {
	var rs, ls Iterator
	rs = t.it
	sizeEstimate := rs.EstimateSize()
	sizeThreshold := t.getTargetSize(sizeEstimate)
	forkRight := true
	children := []*promise.Promise[T]{}
	var result T
	for v := t.sharedResult.Load(); v == nil; {
		internal := t.asInternal()
		if t.taskCanceled() {
			result = internal.defaultResult()
			break
		}
		if sizeEstimate <= sizeThreshold {
			result = internal.computeResult()
			break
		}
		if ls = rs.TrySplit(); ls == nil {
			result = internal.computeResult()
			break
		}
		t.leftChild = internal.newChildTask(ls)
		t.rightChild = internal.newChildTask(rs)
		var taskToFork *taskImpl[T]
		if forkRight {
			forkRight = false
			rs = ls
			t = t.leftChild
			taskToFork = t.rightChild
		} else {
			forkRight = true
			t = t.rightChild
			taskToFork = t.leftChild
		}
		promise := taskToFork.fork()
		children = append(children, promise)
		sizeEstimate = rs.EstimateSize()
	}
	// set shared result
	if t.parent == nil {
		t.sharedResult.Store(result)
	}

	// set local result
	resolve(result)

	// wait for children completion
	promise.All[T](children...).Await()
}

func (t *taskImpl[T]) suggestTargetSize(sizeEstimate int) int {
	if est := sizeEstimate / (runtime.NumCPU() << 2); est > 0 {
		return est
	}
	return 1
}

func (t *taskImpl[T]) getTargetSize(sizeEstimate int) int {
	if t.targetSize == 0 {
		t.targetSize = t.suggestTargetSize(sizeEstimate)
	}
	return t.targetSize
}

func (t *taskImpl[T]) taskCanceled() bool {
	canceled := t.canceled
	if !canceled {
		for p := t.parent; !canceled && p != nil; p = p.parent {
			canceled = p.canceled
		}
	}
	return canceled
}

func (t *taskImpl[T]) cancel() {
	t.canceled = true
}

func (t *taskImpl[T]) cancelLaterTasks() {
	for node, p := t, t.parent; p != nil; p = p.parent {
		if p.leftChild == node {
			if rightSibling := p.rightChild; !rightSibling.canceled {
				rightSibling.cancel()
			}
		}
		node = p
	}
}

func (t *taskImpl[T]) isLeftMost() bool {
	n := t
	for n != nil {
		p := n.parent
		if p != nil && p.leftChild != n {
			return false
		}
		n = p
	}
	return true
}

func (t *taskImpl[T]) onCompletion(result T) {
}
