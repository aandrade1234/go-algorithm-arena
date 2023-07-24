package algorithms

import (
	"errors"

	"github.com/aandrade1234/go-algorithm-arena/challenges"
)

// https://www.hackerrank.com/challenges/birthday-cake-candles/problem
//
// You are in charge of the cake for a child's birthday. You have decided the cake will have one candle for each year
// of their total age. They will only be able to blow out the tallest of the candles. Count how many candles are tallest.
//
// Example
// candles = [4, 4, 1, 3]
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
// The first line contains a single integer n, the size of candles.
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

type BirthdayCandles struct {
	challenges.BaseChallenge
}

func (*BirthdayCandles) GetName() string {
	return "birthday-candles(source: hackerrank)"
}

func (b *BirthdayCandles) Execute() error {
	err := b.BaseExecute("solution 1", b.GetInputs(), b.GetOutputs(), b.solution1)
	return errors.Join(err, b.BaseExecute("solution 2", b.GetInputs(), b.GetOutputs(), b.solution2))
}

// time complexity O(n) space complexity O(1)
func (*BirthdayCandles) solution2(inputs any) any {
	candles := inputs.([]int)

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
func (*BirthdayCandles) solution1(inputs any) any {
	candles := inputs.([]int)

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

func (*BirthdayCandles) GetInputs() []any {
	return []any{
		[]int{3, 2, 1, 3},
		[]int{18, 90, 90, 13, 90, 75, 90, 8, 90, 43},
		[]int{82, 49, 82, 82, 41, 82, 15, 63, 38, 25},
		[]int{44, 53, 31, 27, 77, 60, 66, 77, 26, 36},
	}
}

func (*BirthdayCandles) GetOutputs() []any {
	return []any{2, 5, 4, 2}
}
