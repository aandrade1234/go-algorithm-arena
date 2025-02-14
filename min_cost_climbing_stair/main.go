package main

import "fmt"

func main() {
	fmt.Println(minCostClimbingStairs([]int{10, 15, 20}))
	fmt.Println(minCostClimbingStairs([]int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}))
}

// minCostClimbingStairs time complexity O(n) space complexity O(1).
func minCostClimbingStairs(cost []int) int {
	if len(cost) == 1 {
		return cost[0]
	}

	prev, curr := 0, 0

	for i := 2; i <= len(cost); i++ {
		oneStep := curr + cost[i-1]
		towSteps := prev + cost[i-2]
		prev, curr = curr, min(oneStep, towSteps)
	}

	return curr
}
