package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(minPathSum1([][]int{
		{1, 3, 1},
		{1, 5, 1},
		{4, 2, 1},
	}))
	fmt.Println(minPathSum1([][]int{
		{1, 2, 3},
		{4, 5, 6},
	}))
	fmt.Println()
	fmt.Println(minPathSum2([][]int{
		{1, 3, 1},
		{1, 5, 1},
		{4, 2, 1},
	}))
	fmt.Println(minPathSum2([][]int{
		{1, 2, 3},
		{4, 5, 6},
	}))
	fmt.Println()
	fmt.Println(minPathSum3([][]int{
		{1, 3, 1},
		{1, 5, 1},
		{4, 2, 1},
	}))
	fmt.Println(minPathSum3([][]int{
		{1, 2, 3},
		{4, 5, 6},
	}))
	fmt.Println()
	fmt.Println(minPathSum4([][]int{
		{1, 3, 1},
		{1, 5, 1},
		{4, 2, 1},
	}))
	fmt.Println(minPathSum4([][]int{
		{1, 2, 3},
		{4, 5, 6},
	}))
}

// time complexity O(mn) space complexity O(1)
func minPathSum1(grid [][]int) int {
	ni := len(grid)
	nj := len(grid[0])

	for i := ni - 1; i >= 0; i-- {
		for j := nj - 1; j >= 0; j-- {
			switch {
			case i < ni-1 && j < nj-1:
				grid[i][j] += min(grid[i+1][j], grid[i][j+1])
			case i < ni-1 && j == nj-1:
				grid[i][j] += grid[i+1][j]
			case j < nj-1 && i == ni-1:
				grid[i][j] += grid[i][j+1]
			}
		}
	}

	return grid[0][0]
}

// time complexity O(mn) space complexity O(n)
func minPathSum2(grid [][]int) int {
	ni := len(grid)
	nj := len(grid[0])

	dp := make([]int, nj)

	for i := ni - 1; i >= 0; i-- {
		for j := nj - 1; j >= 0; j-- {
			switch {
			case i < ni-1 && j < nj-1:
				dp[j] = grid[i][j] + min(dp[j], dp[j+1])
			case i < ni-1 && j == nj-1:
				dp[j] = grid[i][j] + dp[j]
			case j < nj-1 && i == ni-1:
				dp[j] = grid[i][j] + dp[j+1]
			case i == ni-1 && j == nj-1:
				dp[j] = grid[i][j]
			}
		}
	}

	return dp[0]
}

// time complexity O(mn) space complexity O(mn)
func minPathSum3(grid [][]int) int {
	ni := len(grid)
	nj := len(grid[0])

	dp := make([][]int, ni)
	for i := range dp {
		dp[i] = make([]int, nj)
	}

	for i := ni - 1; i >= 0; i-- {
		for j := nj - 1; j >= 0; j-- {
			switch {
			case i < ni-1 && j < nj-1:
				dp[i][j] = grid[i][j] + min(dp[i+1][j], dp[i][j+1])
			case i < ni-1 && j == nj-1:
				dp[i][j] = grid[i][j] + dp[i+1][j]
			case j < nj-1 && i == ni-1:
				dp[i][j] = grid[i][j] + dp[i][j+1]
			case i == ni-1 && j == nj-1:
				dp[i][j] = grid[i][j]
			}
		}
	}

	return dp[0][0]
}

// time complexity O(2^m+n) space complexity O(m+n)
func minPathSum4(grid [][]int) int {
	return calculate(0, 0, grid)
}

func calculate(i, j int, grid [][]int) int {
	switch {
	case i == len(grid) || j == len(grid[0]):
		return math.MaxInt
	case i == len(grid)-1 && j == len(grid[0])-1:
		return grid[i][j]
	default:
		return grid[i][j] + min(calculate(i+1, j, grid), calculate(i, j+1, grid))
	}
}
