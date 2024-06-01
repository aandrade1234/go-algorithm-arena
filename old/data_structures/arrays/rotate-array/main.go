package main

import "fmt"

// https://leetcode.com/problems/rotate-array/description/
//
// Given an integer array nums, rotate the array to the right by k steps, where k is non-negative.
//
// Example 1:
// Input: nums = [1,2,3,4,5,6,7], k = 3
// Output:       [5,6,7,1,2,3,4]
// Explanation:
// rotate 1 steps to the right: [7,1,2,3,4,5,6]
// rotate 2 steps to the right: [6,7,1,2,3,4,5]
// rotate 3 steps to the right: [5,6,7,1,2,3,4]
//
// Example 2:
// Input: nums = [-1,-100,3,99], k = 2
// Output: [3,99,-1,-100]
// Explanation:
// rotate 1 steps to the right: [99,-1,-100,3]
// rotate 2 steps to the right: [3,99,-1,-100]
//
// Constraints:
// 1 <= nums.length <= 10^5
// -2^31 <= nums[i] <= 2^31 - 1
// 0 <= k <= 10^5
//
// Follow up:
//
// Try to come up with as many solutions as you can. There are at least three different ways to solve this problem.
// Could you do it in-place with O(1) extra space?
func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	k := 3

	output := solution1(nums, k)
	fmt.Println(output)

	output = solution2(nums, k)
	fmt.Println(output)
}

// solution1 time complexity O(n) space complexity O(n)
func solution1(nums []int, k int) []int {
	n := len(nums)
	output := make([]int, n)
	for i := range nums {
		output[(i+k)%n] = nums[i]
	}
	return output
}

// solution1 time complexity O(n) space complexity O(1)
func solution2(nums []int, k int) []int {
	k %= len(nums)
	reverse(nums, 0, len(nums)-1)
	reverse(nums, 0, k-1)
	reverse(nums, k, len(nums)-1)

	return nums
}

func reverse(nums []int, start, end int) {
	for start < end {
		nums[start], nums[end] = nums[end], nums[start]
		start++
		end--
	}
}
