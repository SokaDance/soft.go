package diff

type Iterable[T any] interface {
	Iterator() Iterator[T]
}
