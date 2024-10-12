package main

import (
	"errors"
	"fmt"
)

var ErrNegativeInput = errors.New("fibonacci is not defined for negative numbers")

func main() {

	fmt.Println(findFibonacciInteractive(-1))
	fmt.Println(findFibonacciInteractive(0))
	fmt.Println(findFibonacciInteractive(1))
	fmt.Println(findFibonacciInteractive(10))
	fmt.Println()
	fmt.Println(findFibonacciRecursive(-1))
	fmt.Println(findFibonacciRecursive(0))
	fmt.Println(findFibonacciRecursive(1))
	fmt.Println(findFibonacciRecursive(10))
}

// findFibonacciInteractive - Big O in time O(n)
func findFibonacciInteractive(value int) (int, error) {
	switch {
	case value < 0:
		return -1, ErrNegativeInput
	case value < 2:
		return value, nil
	}

	prev, current := 0, 1
	for i := 2; i <= value; i++ {
		prev, current = current, prev+current
	}

	return current, nil
}

// findFibonacciInteractive - Big O in time O(2^n)
func findFibonacciRecursive(value int) (int, error) {
	switch {
	case value < 0:
		return -1, ErrNegativeInput
	case value < 2:
		return value, nil
	default:
		v1, _ := findFibonacciRecursive(value - 1)
		v2, _ := findFibonacciRecursive(value - 2)
		return v1 + v2, nil
	}
}
