package main

import (
	"cmp"
	"fmt"
)

func main() {
	data := []int{1, 5, 2, 10, 3, 2}
	SelectionSort[int](data)
	fmt.Println(data)

	data = []int{99, 44, 6, 2, 1, 5, 63, 87, 283, 4, 0}
	SelectionSort[int](data)
	fmt.Println(data)

	dataStr := []string{"andre", "joao", "maria", "predro", "marco"}
	SelectionSort[string](dataStr)
	fmt.Println(dataStr)
}

func SelectionSort[T cmp.Ordered](values []T) {
	n := len(values)

	if n == 0 {
		return
	}

	for i := range n - 1 {
		minIndex := i
		for j := i; j < n; j++ {
			if values[j] < values[minIndex] {
				minIndex = j
			}
		}
		if minIndex != i {
			values[i], values[minIndex] = values[minIndex], values[i]
		}
	}
}
