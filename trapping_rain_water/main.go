package main

import "fmt"

func main() {

	fmt.Println(trap1([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}))
	fmt.Println(trap1([]int{4, 2, 0, 3, 2, 5}))
	fmt.Println()
	fmt.Println(trap2([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}))
	fmt.Println(trap2([]int{4, 2, 0, 3, 2, 5}))

}

func trap1(height []int) int {
	water := 0

	leftMax := make([]int, len(height))
	rightMax := make([]int, len(height))

	leftMax[0] = height[0]
	for i := 1; i < len(height); i++ {
		leftMax[i] = max(leftMax[i-1], height[i])
	}

	rightMax[len(height)-1] = height[len(height)-1]
	for i := len(height) - 2; i >= 0; i-- {
		rightMax[i] = max(rightMax[i+1], height[i])
	}

	for i := 0; i < len(height); i++ {
		water += min(leftMax[i], rightMax[i]) - height[i]
	}

	return water
}

func trap2(height []int) int {
	if len(height) == 0 {
		return 0
	}

	left, right := 0, len(height)-1
	leftMax, rightMax := 0, 0
	water := 0

	for left < right {
		if height[left] < height[right] {
			if height[left] >= leftMax {
				leftMax = height[left]
			} else {
				water += leftMax - height[left]
			}

			left++
		} else {
			if height[right] >= rightMax {
				rightMax = height[right]
			} else {
				water += rightMax - height[right]
			}

			right--
		}
	}

	return water
}
