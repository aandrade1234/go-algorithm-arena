package main

import "fmt"

func main() {
	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
	fmt.Println(maxArea([]int{1, 1}))
}

// Time complexity O(n), space complexity O(1).
func maxArea(height []int) int {
	left := 0
	right := len(height) - 1
	maxArea := 0

	for left < right {
		area := min(height[left], height[right]) * (right - left)
		maxArea = max(maxArea, area)

		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}

	return maxArea
}
