package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(merge([][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}))
	fmt.Println(merge([][]int{{1, 4}, {4, 5}}))
}

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	result := make([][]int, 0)
	current := intervals[0]
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] > current[1] {
			result = append(result, current)
			current = intervals[i]
		} else if intervals[i][1] > current[1] {
			current[1] = intervals[i][1]
		}
	}
	result = append(result, current)
	return result
}
