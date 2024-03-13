package list

import (
	"fmt"
	"strings"

	"github.com/papaya147/go-data-structures/node"
)

type LinkedList[T any] struct {
	Head *node.Node[T]
	Tail *node.Node[T]
	Size uint
}

func NewLinkedList[T any]() *LinkedList[T] {
	return &LinkedList[T]{
		Head: nil,
		Tail: nil,
		Size: 0,
	}
}

func (l *LinkedList[T]) Length() uint {
	return l.Size
}

func (l *LinkedList[T]) HeadValue() T {
	return l.Head.Value
}

func (l *LinkedList[T]) TailValue() T {
	return l.Tail.Value
}

func (l *LinkedList[T]) Append(value T) {
	if l.Head == nil {
		l.Head = node.NewNode(value)
		l.Tail = l.Head
	} else {
		l.Tail.Next = node.NewNode(value)
		l.Tail = l.Tail.Next
	}
	l.Size++
}

func (l *LinkedList[T]) Insert(value T, position uint) {
	if position >= l.Size {
		l.Append(value)
		return
	}
	var prev *node.Node[T]
	curr := l.Head
	for i := 0; i < int(position); i++ {
		prev = curr
		curr = curr.Next
	}
	newNode := node.NewNode(value)
	newNode.Next = curr
	if prev == nil {
		l.Head = newNode
	} else {
		prev.Next = newNode
	}
	l.Size++
}

func (l *LinkedList[T]) Delete(position uint) T {
	if position >= l.Size {
		position = l.Size - 1
	}
	var prev *node.Node[T]
	curr := l.Head
	for i := 0; i < int(position); i++ {
		prev = curr
		curr = curr.Next
	}
	value := curr.Value
	if prev == nil {
		l.Head = l.Head.Next
	} else {
		prev.Next = curr.Next
	}
	l.Size--
	return value
}

func (l *LinkedList[T]) ToString() string {
	str := make([]string, l.Size)
	index := 0
	for curr := l.Head; curr != nil; curr = curr.Next {
		str[index] = fmt.Sprintf("%v", curr.Value)
		index++
	}
	return strings.Join(str, " ")
}
