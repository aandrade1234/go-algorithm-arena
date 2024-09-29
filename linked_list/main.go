package main

import (
	"errors"
	"fmt"
	"strings"
)

// LinkedList represents a singly linked list.
type LinkedList[T any] struct {
	head *node[T]
	tail *node[T]
	size int
}

// node represents a node in the linked list.
type node[T any] struct {
	value T
	next  *node[T]
}

var ErrOutOfBounds = errors.New("index out of bounds")

func main() {
	list := New[float64]()
	list.Append(1, 2, 3)
	list.Prepend(0)
	_ = list.Insert(2, 1.5)
	fmt.Println(list) // Output: [0, 1, 1.5, 2, 3]

	value, _ := list.Remove(2)
	fmt.Printf("Removed: %v\n", value) // Output: Removed: 1.5
	fmt.Println(list)                  // Output: [0, 1, 2, 3]

	value, _ = list.Get(2)
	fmt.Printf("Value at index 2: %v\n", value) // Output: Value at index 2: 2

	fmt.Printf("Size: %d\n", list.Size()) // Output: Size: 4
}

// New creates and returns a new empty LinkedList.
func New[T any]() *LinkedList[T] {
	return &LinkedList[T]{}
}

// IsEmpty returns true if the list is empty.
func (ll *LinkedList[T]) IsEmpty() bool {
	return ll.size == 0
}

// Size returns the number of elements in the list.
func (ll *LinkedList[T]) Size() int {
	return ll.size
}

// Append adds elements to the end of the list.
func (ll *LinkedList[T]) Append(values ...T) {
	for _, value := range values {
		newNode := &node[T]{value: value}

		if ll.IsEmpty() {
			ll.head = newNode
			ll.tail = newNode
		} else {
			ll.tail.next = newNode
			ll.tail = newNode
		}
		ll.size++
	}
}

// Prepend adds elements to the beginning of the list.
func (ll *LinkedList[T]) Prepend(values ...T) {
	for i := len(values) - 1; i >= 0; i-- {
		newNode := &node[T]{value: values[i], next: ll.head}
		ll.head = newNode
		if ll.tail == nil {
			ll.tail = newNode
		}
		ll.size++
	}
}

// Insert adds elements at the specified index.
func (ll *LinkedList[T]) Insert(index int, values ...T) error {
	if index < 0 || index >= ll.size {
		return ErrOutOfBounds
	}

	if index == 0 {
		ll.Prepend(values...)
		return nil
	} else if index == ll.size {
		ll.Append(values...)
		return nil
	}

	prev := ll.head
	for range index - 1 {
		prev = prev.next
	}

	for _, value := range values {
		newNode := &node[T]{value: value, next: prev.next}
		prev.next = newNode
		ll.size++
		prev = newNode
	}

	return nil
}

// Remove removes the element at the specified index.
func (ll *LinkedList[T]) Remove(index int) (T, error) {
	var zero T
	if index < 0 || index >= ll.size {
		return zero, ErrOutOfBounds
	}

	var removedValue T
	if index == 0 {
		removedValue = ll.head.value
		ll.head = ll.head.next
		if ll.head == nil {
			ll.tail = nil
		}
	} else {
		prev := ll.head
		for range index - 1 {
			prev = prev.next
		}
		removedValue = prev.next.value
		prev.next = prev.next.next
		if index == ll.size-1 {
			ll.tail = prev
		}
	}

	ll.size--

	return removedValue, nil
}

// Get returns the element at the specified index.
func (ll *LinkedList[T]) Get(index int) (T, error) {
	var zero T
	if index < 0 || index >= ll.size {
		return zero, ErrOutOfBounds
	}

	current := ll.head
	for i := 0; i < index; i++ {
		current = current.next
	}

	return current.value, nil
}

// String returns a string representation of the list.
func (ll *LinkedList[T]) String() string {
	var sb strings.Builder
	sb.WriteString("[")
	current := ll.head
	for current != nil {
		sb.WriteString(fmt.Sprintf("%v", current.value))
		if current.next != nil {
			sb.WriteString(", ")
		}
		current = current.next
	}
	sb.WriteString("]")

	return sb.String()
}
