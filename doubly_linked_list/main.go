package main

import (
	"errors"
	"fmt"
	"strings"
)

// DoublyLinkedList represents a doubly linked list.
type DoublyLinkedList[T any] struct {
	size int
	head *node[T]
	tail *node[T]
}

// node represents a doubly linked list node.
type node[T any] struct {
	value    *T
	next     *node[T]
	previous *node[T]
}

var ErrOutOfBounds = errors.New("index out of bounds")

func main() {
	list := New[float64]()
	list.Append(1, 2, 3)
	list.Prepend(0)
	_ = list.Insert(2, 1.5)
	fmt.Println(list) // Output: [0, 1, 1.5, 2, 3]

	value, _ := list.Remove(2)
	fmt.Printf("Removed: %v\n", *value) // Output: Removed: 1.5
	fmt.Println(list)                   // Output: [0, 1, 2, 3]

	value, _ = list.Get(2)
	fmt.Printf("Value at index 2: %v\n", *value) // Output: Value at index 2: 2

	fmt.Printf("Size: %d\n", list.Size()) // Output: Size: 4
}

// New creates and returns a new empty DoublyLinkedList.
func New[T any]() *DoublyLinkedList[T] {
	return &DoublyLinkedList[T]{}
}

// IsEmpty returns true if the list is empty.
func (dl *DoublyLinkedList[T]) IsEmpty() bool {
	return dl.size == 0
}

// Size returns the number of elements in the list.
func (dl *DoublyLinkedList[T]) Size() int {
	return dl.size
}

// Append adds elements to the end of the list.
func (dl *DoublyLinkedList[T]) Append(values ...T) {
	for _, value := range values {
		item := &node[T]{
			value:    &value,
			next:     nil,
			previous: dl.tail,
		}

		if dl.tail == nil {
			dl.head = item
			dl.tail = item
		} else {
			dl.tail.next = item
			dl.tail = item
		}

		dl.size++
	}
}

// Prepend adds elements to the beginning of the list.
func (dl *DoublyLinkedList[T]) Prepend(values ...T) {
	var previousNode *node[T]

	for i := len(values) - 1; i >= 0; i-- {
		newNode := &node[T]{
			value:    &values[i],
			next:     dl.head,
			previous: previousNode,
		}

		dl.head = newNode
		if dl.tail == nil {
			dl.tail = newNode
		}

		previousNode = newNode
		dl.size++
	}
}

// Insert adds elements at the specified index.
func (dl *DoublyLinkedList[T]) Insert(index int, values ...T) error {
	if index < 0 || index > dl.size {
		return ErrOutOfBounds
	}

	if index == 0 {
		dl.Prepend(values...)

		return nil
	} else if index == dl.size {
		dl.Append(values...)

		return nil
	}

	current := dl.head
	for range index - 1 {
		current = current.next
	}

	for _, value := range values {
		newNode := &node[T]{
			value:    &value,
			next:     current.next,
			previous: current,
		}

		if current.next != nil {
			current.next.previous = newNode
		}

		current.next = newNode
		current = newNode
		dl.size++
	}

	return nil
}

// Remove removes the element at the specified index.
func (dl *DoublyLinkedList[T]) Remove(index int) (*T, error) {
	if index < 0 || index >= dl.size {
		return nil, ErrOutOfBounds
	}

	var removedNode *node[T]

	// remove first element
	if index == 0 {
		removedNode = dl.head

		dl.head = dl.head.next
		if dl.head != nil {
			dl.head.previous = nil
		} else {
			dl.tail = nil
		}
	} else {
		current := dl.head
		for range index {
			current = current.next
		}

		removedNode = current
		previous := current.previous
		next := current.next

		// remove last item
		if next == nil {
			dl.tail = previous
			previous.next = nil
		} else {
			previous.next = next
			next.previous = previous
		}
	}

	removedNode.next = nil
	removedNode.previous = nil
	dl.size--

	return removedNode.value, nil
}

// Get returns the element at the specified index.
func (dl *DoublyLinkedList[T]) Get(index int) (*T, error) {
	if index < 0 || index >= dl.size {
		return nil, ErrOutOfBounds
	}

	var current *node[T]
	if index < dl.size/2 {
		current = dl.head
		for range index {
			current = current.next
		}
	} else {
		current = dl.tail
		for i := dl.size - 1; i > index; i-- {
			current = current.previous
		}
	}

	return current.value, nil
}

// String returns a string representation of the list.
func (dl *DoublyLinkedList[T]) String() string {
	var sb strings.Builder

	sb.WriteString("[")

	current := dl.head
	for current != nil {
		sb.WriteString(fmt.Sprintf("%v", *current.value))

		if current.next != nil {
			sb.WriteString(", ")
		}

		current = current.next
	}
	sb.WriteString("]")

	return sb.String()
}
