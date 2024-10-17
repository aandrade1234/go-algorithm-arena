package main

import "fmt"

func main() {
	fmt.Println(climbStairs1(2))
	fmt.Println(climbStairs1(3))
	fmt.Println()
	fmt.Println(climbStairs2(2))
	fmt.Println(climbStairs2(3))
}

// time O(n) space O(n)
func climbStairs1(n int) int {
	if n <= 2 {
		return n
	}

	dp := make([]int, n)
	dp[0] = 1
	dp[1] = 2

	for i := 2; i < n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}

	return dp[n-1]
}

// time O(n) space O(1)
func climbStairs2(n int) int {
	if n <= 2 {
		return n
	}

	one := 2
	two := 1

	for i := 3; i <= n; i++ {
		one += two
		two = one
	}

	return one
}
