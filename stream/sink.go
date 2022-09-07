package stream

type sink interface {
	Begin(size int)
	Accept(any)
	End()
	CancelationRequested() bool
}
