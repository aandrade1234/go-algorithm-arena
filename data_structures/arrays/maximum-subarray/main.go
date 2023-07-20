package main

import (
	"fmt"
	"math"
)

// https://leetcode.com/problems/maximum-subarray/description/
//
// Given an integer array nums, find the subarray with the largest sum, and return its sum.
//
// Example 1:
// Input: nums = [-2,1,-3,4,-1,2,1,-5,4]
// Output: 6
// Explanation: The subarray [4,-1,2,1] has the largest sum 6.
//
// Example 2:
// Input: nums = [1]
// Output: 1
// Explanation: The subarray [1] has the largest sum 1.
//
// Example 3:
// Input: nums = [5,4,-1,7,8]
// Output: 23
// Explanation: The subarray [5,4,-1,7,8] has the largest sum 23.
//
// Constraints:
// 1 <= nums.length <= 10^5
// -10^4 <= nums[i] <= 10^4
//
// Follow up: If you have figured out the O(n) solution, try coding another solution using the divide and conquer
// approach, which is more subtle.
func main() {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}

	output := solution1(nums)
	fmt.Println(output)

	output = solution2(nums)
	fmt.Println(output)
}

// solution1 time complexity O(n) space complexity O(1)
func solution1(nums []int) int {
	maxSoFar := nums[0]
	maxEndingHere := nums[0]

	for i := 1; i < len(nums); i++ {
		maxEndingHere = max(nums[i], maxEndingHere+nums[i])
		maxSoFar = max(maxSoFar, maxEndingHere)
	}

	return maxSoFar
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// solution2 time complexity O(n^3) space complexity O(1)
func solution2(nums []int) int {
	maxSum := math.MinInt64
	n := len(nums)
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			currentSum := 0
			for k := i; k <= j; k++ {
				currentSum += nums[k]
			}
			if currentSum > maxSum {
				maxSum = currentSum
			}
		}
	}
	return maxSum
}
