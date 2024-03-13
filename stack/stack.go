package stack

import "github.com/papaya147/go-data-structures/list"

type Stack[T any] struct {
	list list.List[T]
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{
		list: list.NewLinkedList[T](),
	}
}

func (s *Stack[T]) Push(value T) {
	s.list.Append(value)
}

func (s *Stack[T]) Pop() T {
	return s.list.Delete(s.list.Length() - 1)
}

func (s *Stack[T]) Peek() T {
	return s.list.TailValue()
}

func (s *Stack[T]) ToString() string {
	return s.list.ToString()
}

func (s *Stack[T]) Length() uint {
	return s.list.Length()
}
