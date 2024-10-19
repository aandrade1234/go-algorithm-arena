package main

import (
	"fmt"
	"slices"
)

func main() {
	fmt.Println(longestSubarray1([]int{8, 2, 4, 7}, 4))
	fmt.Println(longestSubarray1([]int{10, 1, 2, 4, 7, 2}, 5))
	fmt.Println(longestSubarray1([]int{4, 2, 2, 2, 4, 4, 2, 2}, 0))
	fmt.Println()
	fmt.Println(longestSubarray2([]int{8, 2, 4, 7}, 4))
	fmt.Println(longestSubarray2([]int{10, 1, 2, 4, 7, 2}, 5))
	fmt.Println(longestSubarray2([]int{4, 2, 2, 2, 4, 4, 2, 2}, 0))
	fmt.Println()
	fmt.Println(longestSubarray3([]int{8, 2, 4, 7}, 4))
	fmt.Println(longestSubarray3([]int{10, 1, 2, 4, 7, 2}, 5))
	fmt.Println(longestSubarray3([]int{4, 2, 2, 2, 4, 4, 2, 2}, 0))
}

func longestSubarray1(nums []int, limit int) int {
	var left, result int

	maxDeque, minDeque := []int{}, []int{}

	for right := range nums {
		for len(maxDeque) > 0 && nums[maxDeque[len(maxDeque)-1]] <= nums[right] {
			maxDeque = maxDeque[:len(maxDeque)-1]
		}

		maxDeque = append(maxDeque, right)

		for len(minDeque) > 0 && nums[minDeque[len(minDeque)-1]] >= nums[right] {
			minDeque = minDeque[:len(minDeque)-1]
		}

		minDeque = append(minDeque, right)

		for nums[maxDeque[0]]-nums[minDeque[0]] > limit {
			left++
			if maxDeque[0] < left {
				maxDeque = maxDeque[1:]
			}

			if minDeque[0] < left {
				minDeque = minDeque[1:]
			}
		}

		if right-left+1 > result {
			result = right - left + 1
		}
	}

	return result
}

// O(n^2)
func longestSubarray2(nums []int, limit int) int {
	left, result := 0, 0

	for right := range nums {
		maxVal := slices.Max(nums[left : right+1])
		minVal := slices.Min(nums[left : right+1])

		for maxVal-minVal > limit {
			left++

			maxVal = slices.Max(nums[left : right+1])
			minVal = slices.Min(nums[left : right+1])
		}

		result = max(result, right-left+1)
	}

	return result
}

// Brute force O(n^3) -> 2 nested looping + max/min
func longestSubarray3(nums []int, limit int) int {
	result := 0

	for left := range nums {
		for right := left; right < len(nums); right++ {
			maxVal := slices.Max(nums[left : right+1])
			minVal := slices.Min(nums[left : right+1])
			diff := abs(maxVal - minVal)

			if diff <= limit {
				localSize := right - left + 1
				if localSize > result {
					result = localSize
				}
			}
		}
	}

	return result
}

func abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}
