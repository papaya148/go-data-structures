package tree

import (
	"cmp"

	"github.com/papaya147/go-data-structures/list"
)

type BinarySearchTree[T cmp.Ordered] struct {
	Root *Node[T]
	Size uint
}

func NewBinarySearchTree[T cmp.Ordered]() *BinarySearchTree[T] {
	return &BinarySearchTree[T]{
		Root: nil,
		Size: 0,
	}
}

func (b *BinarySearchTree[T]) Insert(value T) {
	if b.Root == nil {
		b.Root = NewNode(value)
	}
	insertBST(b.Root, value)
}

func insertBST[T cmp.Ordered](node *Node[T], value T) {
	if node == nil {
		node = NewNode(value)
	}
	if node.Value > value {
		insertBST(node.Right, value)
	} else {
		insertBST(node.Left, value)
	}
}

func (b *BinarySearchTree[T]) Preorder() list.List[T] {
	return b.Root.Preorder()
}

func (b *BinarySearchTree[T]) Inorder() list.List[T] {
	return b.Root.Inorder()
}

func (b *BinarySearchTree[T]) Postorder() list.List[T] {
	return b.Root.Postorder()
}
