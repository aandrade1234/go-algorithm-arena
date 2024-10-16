package main

import "fmt"

func main() {
	fmt.Println(maxProfit1([]int{7, 1, 5, 3, 6, 4}))
	fmt.Println(maxProfit1([]int{7, 6, 4, 3, 1}))
	fmt.Println(maxProfit1([]int{2, 1, 2, 1, 0, 1, 2}))
	fmt.Println()
	fmt.Println(maxProfit2([]int{7, 1, 5, 3, 6, 4}))
	fmt.Println(maxProfit2([]int{7, 6, 4, 3, 1}))
	fmt.Println(maxProfit2([]int{2, 1, 2, 1, 0, 1, 2}))
	fmt.Println()
	fmt.Println(maxProfit3([]int{7, 1, 5, 3, 6, 4}))
	fmt.Println(maxProfit3([]int{7, 6, 4, 3, 1}))
	fmt.Println(maxProfit3([]int{2, 1, 2, 1, 0, 1, 2}))
}

// time O(n) space O(1)
func maxProfit1(prices []int) int {
	n := len(prices)
	if n == 0 {
		return 0
	}

	profit := 0
	buy := -1

	for i := 1; i < n; i++ {
		if prices[i] > prices[i-1] && (buy > prices[i-1] || buy == -1) {
			buy = prices[i-1]
		}

		if prices[i] > buy && buy >= 0 {
			profit = max(profit, prices[i]-buy)
		}
	}

	return profit
}

// time O(n) space O(1)
func maxProfit2(prices []int) int {
	n := len(prices)
	if n == 0 {
		return 0
	}

	minPrice := prices[0]
	maxProfit := 0

	for i := 1; i < n; i++ {
		if prices[i] < minPrice {
			minPrice = prices[i]
		}

		maxProfit = max(maxProfit, prices[i]-minPrice)
	}

	return maxProfit
}

// time O(n) space O(n)
func maxProfit3(prices []int) int {
	n := len(prices)
	if n == 0 {
		return 0
	}

	dp := make([]int, n)
	minPrice := prices[0]

	for i := 1; i < n; i++ {
		dp[i] = max(dp[i-1], prices[i]-minPrice)
		minPrice = min(minPrice, prices[i])
	}

	return dp[n-1]
}
