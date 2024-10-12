package main

import (
	"cmp"
	"fmt"
)

func main() {
	data := []int{1, 5, 2, 10, 3, 2}
	BubbleSort[int](data)
	fmt.Println(data)

	data = []int{99, 44, 6, 2, 1, 5, 63, 87, 283, 4, 0}
	BubbleSort[int](data)
	fmt.Println(data)

	dataStr := []string{"andre", "joao", "maria", "predro", "marco"}
	BubbleSort[string](dataStr)
	fmt.Println(dataStr)
}

func BubbleSort[T cmp.Ordered](values []T) {
	n := len(values) - 1
	for i := range n {
		swapped := false

		for j := range n - i {
			if values[j] > values[j+1] {
				values[j], values[j+1] = values[j+1], values[j]
				swapped = true
			}
		}

		if !swapped {
			break
		}
	}
}
