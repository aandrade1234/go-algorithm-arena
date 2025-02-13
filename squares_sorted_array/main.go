package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(squareSortedArray([]int{-4, -1, 0, 3, 10}))
	fmt.Println(squareSortedArray([]int{-7, -3, 2, 3, 11}))
}

func squareSortedArray(nums []int) []int {
	squares := make([]int, len(nums))

	n := len(nums)
	left := 0
	right := len(nums) - 1
	for i := n - 1; i >= 0; i-- {
		square := 0
		if math.Abs(float64(nums[left])) >= math.Abs(float64(nums[right])) {
			square = nums[left]
			left++
		} else {
			square = nums[right]
			right--
		}
		squares[i] = square * square
	}

	return squares
}
