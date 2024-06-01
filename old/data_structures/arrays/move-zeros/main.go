package main

import "fmt"

// https://leetcode.com/problems/move-zeroes/description/
//
// Given an integer array nums, move all 0's to the end of it while maintaining the relative order of the non-zero elements.
//
// Note that you must do this in-place without making a copy of the array.
//
// Example 1:
// Input: nums = [0,1,0,3,12]
// Output: [1,3,12,0,0]
//
// Example 2:
// Input: nums = [0]
// Output: [0]
//
// Constraints:
// 1 <= nums.length <= 10^4
// -2^31 <= nums[i] <= 2^31 - 1
//
// Follow up: Could you minimize the total number of operations done?

func main() {
	nums := []int{0, 1, 0, 3, 12}

	output := solution1(nums)
	fmt.Println(output)
}

// solution1 time complexity of O(n) and space complexity O(1)
func solution1(nums []int) []int {
	pos := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[i], nums[pos] = nums[pos], nums[i]
			pos++
		}
	}

	return nums
}
