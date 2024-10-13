package main

import (
	"cmp"
	"fmt"
)

func main() {
	fmt.Println(MergeSort[int]([]int{1, 5, 2, 10, 3, 2}))
	fmt.Println(MergeSort[int]([]int{99, 44, 6, 2, 1, 5, 63, 87, 283, 4, 0}))
	fmt.Println(MergeSort[string]([]string{"andre", "joao", "maria", "pedro", "marco"}))
}

func MergeSort[T cmp.Ordered](values []T) []T {
	if len(values) == 1 {
		return values
	}

	mid := len(values) / 2
	left := MergeSort[T](values[:mid])
	right := MergeSort[T](values[mid:])

	return merge[T](left, right)
}

func merge[T cmp.Ordered](left, right []T) []T {
	result := make([]T, 0, len(left)+len(right))
	i, j := 0, 0

	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	result = append(result, left[i:]...)
	result = append(result, right[j:]...)

	return result
}
