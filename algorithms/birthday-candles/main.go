package main

import "fmt"

// https://www.hackerrank.com/challenges/birthday-cake-candles/problem
//
// You are in charge of the cake for a child's birthday. You have decided the cake will have one candle for each year
// of their total age. They will only be able to blow out the tallest of the candles. Count how many candles are tallest.
//
// Example
// candles = [4,4,1,3]
//
// The maximum height candles are 4 units high. There are 2 of them, so return 2.
//
// Function Description
//
// Complete the function birthdayCakeCandles in the editor below.
//
// birthdayCakeCandles has the following parameter(s):
// int candles[n]: the candle heights
//
// Returns
// int: the number of candles that are tallest
//
// Input Format
// The first line contains a single integer, n, the size of candles.
// The second line contains n space-separated integers, where each integer i describes the height of candles[i].
//
// Constraints
// 1 <=n<=10^5
// 1 <=candles[i] <=10^7
//
// Sample Input 0
// 4
// 3 2 1 3
//
// Sample Output 0
// 2
//
// Explanation 0
// Candle heights are [3, 2, 1, 3]. The tallest candles are 3 units, and there are 2 of them.

func main() {
	data := []int{3, 2, 1, 3}

	output := solution1(data)
	fmt.Println(output)

	output = solution2(data)
	fmt.Println(output)
}

// time complexity O(n) space complexity O(1)
func solution2(candles []int) int {
	max := 0
	count := 0

	for _, height := range candles {
		if height > max {
			max = height
			count = 1
		} else if height == max {
			count++
		}
	}

	return count
}

// time complexity O(n) space complexity O(n)
func solution1(candles []int) int {
	tallestCandle := 0
	cake := make(map[int]int)

	for _, candle := range candles {
		cake[candle] = cake[candle] + 1

		if candle > tallestCandle {
			tallestCandle = candle
		}
	}

	return cake[tallestCandle]
}
