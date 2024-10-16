package main

import "fmt"

func main() {
	fmt.Println(rob1([]int{1, 2, 3, 1}))
	fmt.Println(rob1([]int{2, 7, 9, 3, 1}))
	fmt.Println()
	fmt.Println(rob2([]int{1, 2, 3, 1}))
	fmt.Println(rob2([]int{2, 7, 9, 3, 1}))
}

// Time O(n), Space O(n)
func rob1(nums []int) int {
	n := len(nums)
	switch n {
	case 0:
		return 0
	case 1:
		return nums[0]
	}

	dp := make([]int, n)
	dp[0] = nums[0]
	dp[1] = max(nums[0], nums[1])

	for i := 2; i < n; i++ {
		dp[i] = max(dp[i-1], dp[i-2]+nums[i])
	}

	return dp[n-1]
}

// Time O(n), Space O(1)
func rob2(nums []int) int {
	n := len(nums)
	switch n {
	case 0:
		return 0
	case 1:
		return nums[0]
	}

	prev2, prev1 := 0, nums[0]
	for i := 1; i < n; i++ {
		current := max(prev1, prev2+nums[i])
		prev2 = prev1
		prev1 = current
	}

	return prev1
}
