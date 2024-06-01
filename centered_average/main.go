package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(centeredAverage([]int{1, 2, 3, 4, 100}))
	fmt.Println(centeredAverage([]int{1, 1, 5, 5, 10, 8, 7}))
	fmt.Println(centeredAverage([]int{-10, -4, -2, -4, -2}))
}

func centeredAverage(nums []int) int {
	min, max := math.MaxInt, math.MinInt
	sum := 0
	for _, n := range nums {
		sum += n

		if n < min {
			min = n
		}

		if n > max {
			max = n
		}
	}

	return (sum - min - max) / (len(nums) - 2)
}
