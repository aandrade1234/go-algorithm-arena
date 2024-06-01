package main

import "fmt"

// https://www.geeksforgeeks.org/find-first-repeating-element-array-integers/
//
// Given an array of integers, find the first recurring value. If there are no recurring characters, return null.
//
// examples:
// [2,5,1,2,3,5,1,2,4] returns 2
// [2,1,1,2,3,5,1,2,4] returns 1
// [2,3,4,5] returns null
func main() {
	nums := []int{2, 5, 1, 2, 3, 5, 1, 2, 4}

	output := *solution1(nums)
	fmt.Println(output)
}

// Space complexity O(n) Time complexity O(n)
func solution1(values []int) *int {
	data := make(map[int]bool)

	for _, value := range values {
		if _, exists := data[value]; exists {
			retValue := value
			return &retValue
		}
		data[value] = true
	}
	return nil
}
