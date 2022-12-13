package diff

type Entry[T any] struct {
	Index      int
	Element    T
	IsAddition bool
}

type Diff[T any] struct {
	Entries []Entry[T]
}

type DiffVisitor[T any] interface {
	HandleAdd(index int, element T)
	HandleMove(oldIndex int, newIndex int, element T)
	HandleRemove(index int, element T)
	HandleReplace(index int, oldElement T, newElement T)
}

func (d Diff[T]) Accept(v DiffVisitor[T]) {

}

func Compute[T any](oldArray []T, newArray []T) Diff[T] {
	return Diff[T]{}
}
