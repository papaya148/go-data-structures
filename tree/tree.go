package tree

type Tree[T any] interface {
	Insert(value T)
}
