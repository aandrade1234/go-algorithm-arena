package main

import (
	"fmt"
)

func main() {
	fmt.Println(spiralOrder(
		[][]int{
			{1, 2, 3},
			{4, 5, 6},
			{7, 8, 9},
		}))
	fmt.Println(spiralOrder(
		[][]int{
			{1, 2, 3, 4},
			{5, 6, 7, 8},
			{9, 10, 11, 12},
		}))
}

// Time complexity O(m.n), Space complexity O(1)
func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return []int{}
	}

	rows := len(matrix)
	cols := len(matrix[0])

	top, bottom := 0, rows-1
	left, right := 0, cols-1

	result := make([]int, 0, rows*cols)

	for top <= bottom && left <= right {
		for i := left; i <= right; i++ {
			result = append(result, matrix[top][i])
		}
		top++

		for i := top; i <= bottom; i++ {
			result = append(result, matrix[i][right])
		}
		right--

		if top <= bottom {
			for i := right; i >= left; i-- {
				result = append(result, matrix[bottom][i])
			}
			bottom--
		}

		if left <= right {
			for i := bottom; i >= top; i-- {
				result = append(result, matrix[i][left])
			}
			left++
		}
	}

	return result
}
