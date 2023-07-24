package algorithms

import (
	"sort"

	"github.com/aandrade1234/go-algorithm-arena/challenges"
)

// https://www.hackerrank.com/challenges/mini-max-sum/problem
//
// Given five positive integers, find the minimum and maximum values that can be calculated by summing exactly four of
// the five integers. Then print the respective minimum and maximum values as a single line of two space-separated
// long integers.
//
// Example
// arr = [1,3,5,7,9]
//
// The minimum sum is 1 + 3 + 5 + 7 = 16 and the maximum sum is 3 + 5 + 7 + 9 = 24. The function prints 16 24
//
// Function Description
//
// Complete the miniMaxSum function in the editor below.
//
// miniMaxSum has the following parameter(s):
// arr: an array of 5 integers
//
// Print
// Print two space-separated integers on one line: the minimum sum and the maximum sum of 4 of 5 elements.
//
// Input Format
// A single line of five space-separated integers.
//
// Constraints
// 1 <= arr[i] <= 10^9
//
// Output Format
// Print two space-separated long integers denoting the respective minimum and maximum values that can be calculated by
// summing exactly four of the five integers. (The output can be greater than a 32 bit integer.)
//
// Sample Input
// 1 2 3 4 5
// Sample Output
// 10 14

type MinMaxSum struct {
	challenges.BaseChallenge
}

func (*MinMaxSum) GetName() string {
	return "min-max-sum(source: hackerrank)"
}

func (b *MinMaxSum) Execute() error {
	return b.BaseExecute("solution 1", b.GetInputs(), b.GetOutputs(), b.solution1)
}

// time complexity of O(n) space complexity of O(1)
func (*MinMaxSum) solution1(inputs any) any {
	values := inputs.([]int)

	sort.Ints(values)
	min := 0
	max := 0
	for i := 0; i < len(values)-1; i++ {
		min += values[i]
		max += values[i+1]
	}
	return [2]int{min, max}
}

func (*MinMaxSum) GetInputs() []any {
	return []any{
		[]int{1, 3, 5, 7, 9},
		[]int{1, 2, 3, 4, 5},
		[]int{256741038, 623958417, 467905213, 714532089, 938071625},
		[]int{396285104, 573261094, 759641832, 819230764, 364801279},
		[]int{140638725, 436257910, 953274816, 734065819, 362748590},
		[]int{769082435, 210437958, 673982045, 375809214, 380564127},
		[]int{156873294, 719583602, 581240736, 605827319, 895647130},
		[]int{426980153, 354802167, 142980735, 968217435, 734892650},
		[]int{942381765, 627450398, 954173620, 583762094, 236817490},
		[]int{539674108, 549382170, 270968351, 746219035, 140597628},
		[]int{254961783, 604179258, 462517083, 967304281, 860273491},
		[]int{501893267, 649027153, 379408215, 452968170, 487530619},
		[]int{140537896, 243908675, 670291834, 923018467, 520718469},
		[]int{793810624, 895642170, 685903712, 623789054, 468592370},
		[]int{5, 5, 5, 5, 5},
		[]int{7, 69, 2, 221, 8974},
	}
}

func (*MinMaxSum) GetOutputs() []any {
	return []any{
		[2]int{16, 24},
		[2]int{10, 14},
		[2]int{2063136757, 2744467344},
		[2]int{2093989309, 2548418794},
		[2]int{1673711044, 2486347135},
		[2]int{1640793344, 2199437821},
		[2]int{2063524951, 2802298787},
		[2]int{1659655705, 2484892405},
		[2]int{2390411747, 3107767877},
		[2]int{1500622257, 2106243664},
		[2]int{2181931615, 2894274113},
		[2]int{1821800271, 2091419209},
		[2]int{1575456874, 2357937445},
		[2]int{2572095760, 2999145560},
		[2]int{20, 20},
		[2]int{299, 9271},
	}
}
