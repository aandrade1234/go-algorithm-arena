package main

import (
	"cmp"
	"fmt"
)

func main() {
	fmt.Println(QuickSort1[int]([]int{1, 5, 2, 10, 3, 2}))
	fmt.Println(QuickSort1[int]([]int{99, 44, 6, 2, 1, 5, 63, 87, 283, 4, 0}))
	fmt.Println(QuickSort1[string]([]string{"andre", "joao", "maria", "pedro", "marco"}))
	fmt.Println()
	fmt.Println(QuickSort2[int]([]int{1, 5, 2, 10, 3, 2}))
	fmt.Println(QuickSort2[int]([]int{99, 44, 6, 2, 1, 5, 63, 87, 283, 4, 0}))
	fmt.Println(QuickSort2[string]([]string{"andre", "joao", "maria", "pedro", "marco"}))
	fmt.Println()
}

func QuickSort1[T cmp.Ordered](values []T) []T {
	n := len(values)
	if n <= 1 {
		return values
	}

	key := values[n-1]
	left := make([]T, 0, n)
	right := make([]T, 0, n)

	for i := range n - 1 {
		if values[i] > key {
			right = append(right, values[i])
		} else {
			left = append(left, values[i])
		}
	}

	left = QuickSort1[T](left)
	right = QuickSort1[T](right)

	result := make([]T, 0, n)
	result = append(result, left...)
	result = append(result, key)
	result = append(result, right...)

	return result
}

func QuickSort2[T cmp.Ordered](values []T) []T {
	if len(values) <= 1 {
		return values
	}

	quickSortHelper(values, 0, len(values)-1)

	return values
}

func quickSortHelper[T cmp.Ordered](values []T, low, high int) {
	if low < high {
		pivotIndex := partition(values, low, high)
		quickSortHelper[T](values, low, pivotIndex-1)
		quickSortHelper[T](values, pivotIndex+1, high)
	}
}

func partition[T cmp.Ordered](values []T, low, high int) int {
	mid := low + (high-low)/2
	values[mid], values[high] = values[high], values[mid]
	pivot := values[high]

	i := low - 1
	for j := low; j < high; j++ {
		if values[j] <= pivot {
			i++
			values[i], values[j] = values[j], values[i]
		}
	}

	values[i+1], values[high] = values[high], values[i+1]

	return i + 1
}
