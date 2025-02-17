package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(minSubArrayLen(7, []int{2, 3, 1, 2, 4, 3}))
	fmt.Println(minSubArrayLen(4, []int{1, 4, 4}))
	fmt.Println(minSubArrayLen(11, []int{1, 1, 1, 1, 1, 1, 1, 1}))
}

// Time complexity O(n), space complexity O(1).
func minSubArrayLen(target int, nums []int) int {
	left := 0
	sum := 0
	minLength := math.MaxInt

	for right := range len(nums) {
		sum += nums[right]

		for sum >= target {
			minLength = min(minLength, right-left+1)
			sum -= nums[left]
			left++
		}
	}

	if minLength == math.MaxInt {
		return 0
	}

	return minLength
}
