package stream

import (
	"runtime"
	"sync/atomic"

	"github.com/chebyrash/promise"
)

type task[T any] struct {
	parent        *task[T]
	targetSize    int
	result        *atomic.Value
	iterator      Iterator
	computeResult func(*task[T]) T
	defaultResult func(*task[T]) T
	canceled      bool
	leftChild     *task[T]
	rightChild    *task[T]
}

func newRootTask[T any](iterator Iterator, computeResult func(*task[T]) T, defaultResult func(*task[T]) T) *task[T] {
	t := &task[T]{
		result:        &atomic.Value{},
		iterator:      iterator,
		targetSize:    0,
		computeResult: computeResult,
	}
	return t
}

func newChildTask[T any](parent *task[T], iterator Iterator) *task[T] {
	t := &task[T]{
		parent:        parent,
		targetSize:    parent.targetSize,
		result:        parent.result,
		computeResult: parent.computeResult,
		defaultResult: parent.defaultResult,
		iterator:      iterator,
	}
	return t
}

func (t *task[T]) Fork() *promise.Promise[T] {
	return promise.New[T](t.computeTask)
}

func (t *task[T]) Invoke() T {
	result, _ := t.Fork().Await()
	return result
}

func (t *task[T]) setResult(value T) {
	if any(value) != nil {
		t.result.Store(value)
	}
}

func (t *task[T]) computeTask(resolve func(T), reject func(error)) {
	var rs, ls Iterator
	rs = t.iterator
	sizeEstimate := rs.EstimateSize()
	sizeThreshold := t.getTargetSize(sizeEstimate)
	forkRight := true
	children := []*promise.Promise[T]{}
	var result T
	for v := t.result.Load(); v == nil; {
		if t.taskCanceled() {
			result = t.defaultResult(t)
			break
		}
		if sizeEstimate <= sizeThreshold {
			result = t.computeResult(t)
			break
		}
		if ls = rs.TrySplit(); ls == nil {
			result = t.computeResult(t)
			break
		}
		t.leftChild = newChildTask(t, ls)
		t.rightChild = newChildTask(t, rs)
		var taskToFork *task[T]
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
		promise := taskToFork.Fork()
		children = append(children, promise)
		sizeEstimate = rs.EstimateSize()
	}
	// set shared result
	if t.parent == nil {
		t.result.Store(result)
	}

	// set local result
	resolve(result)

	// wait for children completion
	promise.All[T](children...).Await()
}

func (t *task[T]) suggestTargetSize(sizeEstimate int) int {
	if est := sizeEstimate / (runtime.NumCPU() << 2); est > 0 {
		return est
	}
	return 1
}

func (t *task[T]) getTargetSize(sizeEstimate int) int {
	if t.targetSize == 0 {
		t.targetSize = t.suggestTargetSize(sizeEstimate)
	}
	return t.targetSize
}

func (t *task[T]) taskCanceled() bool {
	canceled := t.canceled
	if !canceled {
		for p := t.parent; !canceled && p != nil; p = p.parent {
			canceled = p.canceled
		}
	}
	return canceled
}

func (t *task[T]) cancel() {
	t.canceled = true
}

func (t *task[T]) cancelLaterTasks() {
	for node, p := t, t.parent; p != nil; p = p.parent {
		if p.leftChild == node {
			if rightSibling := p.rightChild; !rightSibling.canceled {
				rightSibling.cancel()
			}
		}
		node = p
	}
}
