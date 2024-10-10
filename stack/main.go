package main

import (
	"errors"
	"fmt"
)

type Stack[T any] struct {
	top *node[T]
}

type node[T any] struct {
	value    T
	previous *node[T]
}

var ErrStackUnderFlow = errors.New("stack underflow")

func main() {
	stack := New[int]()
	stack.Push(10)
	stack.Push(20)
	stack.Push(30)

	v, _ := stack.Peek()
	fmt.Println(v)

	v, _ = stack.Pop()
	fmt.Println(v)
	v, _ = stack.Pop()
	fmt.Println(v)
	v, _ = stack.Pop()
	fmt.Println(v)

	v, err := stack.Pop()
	fmt.Println(err)
}

func New[T any]() *Stack[T] {
	return &Stack[T]{}
}

func (s *Stack[T]) Peek() (T, error) {
	if s.top == nil {
		var zero T
		return zero, ErrStackUnderFlow
	}

	return s.top.value, nil
}

func (s *Stack[T]) Push(value T) {
	node := &node[T]{
		value:    value,
		previous: s.top,
	}

	s.top = node
}

func (s *Stack[T]) Pop() (T, error) {
	if s.top == nil {
		var zero T
		return zero, ErrStackUnderFlow
	}

	val := s.top.value
	s.top = s.top.previous

	return val, nil
}
