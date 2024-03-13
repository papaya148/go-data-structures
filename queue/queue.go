package queue

import "github.com/papaya147/go-data-structures/list"

type Queue[T any] struct {
	list list.List[T]
}

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{
		list: list.NewLinkedList[T](),
	}
}

func (q *Queue[T]) Enqueue(value T) {
	q.list.Append(value)
}

func (q *Queue[T]) Dequeue() T {
	return q.list.Delete(0)
}

func (q *Queue[T]) Peek() T {
	return q.list.HeadValue()
}

func (q Queue[T]) ToString() string {
	return q.list.ToString()
}

func (q Queue[T]) Length() uint {
	return q.list.Length()
}
