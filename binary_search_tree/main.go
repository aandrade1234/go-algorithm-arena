package main

import (
	"cmp"
	"errors"
	"fmt"
)

type BinarySearchTree[T cmp.Ordered] struct {
	root *node[T]
}

type node[T cmp.Ordered] struct {
	value T
	left  *node[T]
	right *node[T]
}

var ErrValueNotFound = errors.New("value not found")

func main() {
	bst := New[int]()
	bst.Insert(9)
	bst.Insert(4)
	bst.Insert(6)
	bst.Insert(20)
	bst.Insert(170)
	bst.Insert(15)
	bst.Insert(1)

	val, err := bst.Lookup(170)
	fmt.Println(val)

	val, err = bst.Lookup(17)
	fmt.Println(err)

	_ = bst.Remove(20)

	_, err = bst.Lookup(20)
	fmt.Println(err)
}

func New[T cmp.Ordered]() *BinarySearchTree[T] {
	return &BinarySearchTree[T]{}
}

func (bst *BinarySearchTree[T]) Insert(value T) {
	newNode := &node[T]{value: value}

	if bst.root == nil {
		bst.root = newNode
		return
	}

	currNode := bst.root

	for {
		if value < currNode.value {
			if currNode.left == nil {
				currNode.left = newNode
				return
			}
			currNode = currNode.left
		} else {
			if currNode.right == nil {
				currNode.right = newNode
				return
			}
			currNode = currNode.right
		}
	}
}

func (bst *BinarySearchTree[T]) Lookup(value T) (T, error) {
	var zero T

	currNode := bst.root
	for currNode != nil {
		switch {
		case value == currNode.value:
			return value, nil
		case value > currNode.value:
			currNode = currNode.right
		case value < currNode.value:
			currNode = currNode.left
		}
	}

	return zero, ErrValueNotFound
}

func (bst *BinarySearchTree[T]) Remove(value T) error {
	var parent *node[T]

	currNode := bst.root
	for currNode != nil && currNode.value != value {
		parent = currNode

		if value < currNode.value {
			currNode = currNode.left
		} else {
			currNode = currNode.right
		}
	}

	if currNode == nil {
		return ErrValueNotFound
	}

	switch {
	case currNode.left == nil && currNode.right == nil:
		switch {
		case parent == nil:
			bst.root = nil
		case parent.left == currNode:
			parent.left = nil
		case parent.right == currNode:
			parent.right = nil
		}
	case currNode.left == nil || currNode.right == nil:
		var child *node[T]
		if currNode.left != nil {
			child = currNode.left
		} else {
			child = currNode.right
		}

		switch {
		case parent == nil:
			bst.root = child
		case parent.left == currNode:
			parent.left = child
		case parent.right == currNode:
			parent.right = child
		}
	case currNode.left != nil && currNode.right != nil:
		successorParent := currNode

		successor := currNode.right
		for successor.left != nil {
			successorParent = successor
			successor = successor.left
		}

		currNode.value = successor.value

		if successorParent.left == successor {
			successorParent.left = successor.right
		} else {
			successorParent.right = successor.right
		}
	}

	return nil
}
