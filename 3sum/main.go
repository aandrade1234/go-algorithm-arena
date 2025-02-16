package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4}))
	fmt.Println(threeSum([]int{-0, 1, 1}))
	fmt.Println(threeSum([]int{0, 0, 0}))
}

// Time complexity O(n^2), Space complexity O(1).
func threeSum(nums []int) [][]int {
	result := make([][]int, 0)
	n := len(nums)

	sort.Ints(nums)

	for i := range n - 2 {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		left := i + 1
		right := n - 1

		result = append(result, twoSum(i, left, right, nums)...)
	}

	return result
}

func twoSum(i, left, right int, nums []int) [][]int {
	result := make([][]int, 0)

	for left < right {
		sum := nums[i] + nums[left] + nums[right]

		switch {
		case sum == 0:
			result = append(result, []int{nums[i], nums[left], nums[right]})
			left++
			right--

			for left < right && nums[left] == nums[left-1] {
				left++
			}

			for left < right && nums[right] == nums[right+1] {
				right--
			}
		case sum < 0:
			left++
		default:
			right--
		}
	}

	return result
}
