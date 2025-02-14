package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(squareSortedArray1([]int{-4, -1, 0, 3, 10}))
	fmt.Println(squareSortedArray1([]int{-7, -3, 2, 3, 11}))
	fmt.Println()
	fmt.Println(squareSortedArray2([]int{-4, -1, 0, 3, 10}))
	fmt.Println(squareSortedArray2([]int{-7, -3, 2, 3, 11}))
}

func squareSortedArray1(nums []int) []int {
	n := len(nums)
	left := 0
	right := n - 1
	result := make([]int, n)

	for i := n - 1; i >= 0; i-- {
		var square int
		if math.Abs(float64(nums[left])) >= math.Abs(float64(nums[right])) {
			square = nums[left]
			left++
		} else {
			square = nums[right]
			right--
		}

		result[i] = square * square
	}

	return result
}

func squareSortedArray2(nums []int) []int {
	n := len(nums)
	left := 0
	right := n - 1
	i := n - 1
	result := make([]int, n)

	for left <= right {
		leftSquare := nums[left] * nums[left]
		rightSquare := nums[right] * nums[right]

		if leftSquare > rightSquare {
			result[i] = leftSquare
			left++
		} else {
			result[i] = rightSquare
			right--
		}

		i--
	}

	return result
}
