package main

import (
	"cmp"
	"fmt"
)

func main() {
	data := []int{1, 5, 2, 10, 3, 2}
	InsertionSort[int](data)
	fmt.Println(data)

	data = []int{99, 44, 6, 2, 1, 5, 63, 87, 283, 4, 0}
	InsertionSort[int](data)
	fmt.Println(data)

	dataStr := []string{"andre", "joao", "maria", "predro", "marco"}
	InsertionSort[string](dataStr)
	fmt.Println(dataStr)
}

func InsertionSort[T cmp.Ordered](values []T) {
	for i := 1; i < len(values); i++ {
		if values[i] < values[i-1] {
			key := values[i]

			j := i - 1
			for j >= 0 && values[j] > key {
				values[j+1] = values[j]
				j--
			}

			values[j+1] = key
		}
	}
}
