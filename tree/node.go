package tree

import (
	"cmp"

	"github.com/papaya147/go-data-structures/list"
)

type Node[T cmp.Ordered] struct {
	Value T
	Left  *Node[T]
	Right *Node[T]
}

func NewNode[T cmp.Ordered](value T) *Node[T] {
	return &Node[T]{
		Value: value,
		Left:  nil,
		Right: nil,
	}
}

func (n *Node[T]) Preorder() list.List[T] {
	l := list.NewLinkedList[T]()
	preorder(n, l)
	return l
}

func preorder[T cmp.Ordered](root *Node[T], ll list.List[T]) {
	if root == nil {
		return
	}

	ll.Append(root.Value)
	preorder(root.Left, ll)
	preorder(root.Right, ll)
}

func (n *Node[T]) Inorder() list.List[T] {
	l := list.NewLinkedList[T]()
	inorder(n, l)
	return l
}

func inorder[T cmp.Ordered](root *Node[T], ll list.List[T]) {
	if root == nil {
		return
	}

	inorder(root.Left, ll)
	ll.Append(root.Value)
	inorder(root.Right, ll)
}

func (n *Node[T]) Postorder() list.List[T] {
	l := list.NewLinkedList[T]()
	postorder(n, l)
	return l
}

func postorder[T cmp.Ordered](root *Node[T], ll list.List[T]) {
	if root == nil {
		return
	}

	postorder(root.Left, ll)
	postorder(root.Right, ll)
	ll.Append(root.Value)
}
