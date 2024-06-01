package main

import "fmt"

func main() {
	fmt.Println(mergeSortedArrays([]int{0, 3, 35}, []int{4, 6, 30, 33}))
}

func mergeSortedArrays(arr1, arr2 []int) []int {
	result := make([]int, 0, len(arr1)+len(arr2))

	i1, i2 := 0, 0
	for i1 < len(arr1) && i2 < len(arr2) {
		if arr1[i1] <= arr2[i2] {
			result = append(result, arr1[i1])
			i1++
		} else {
			result = append(result, arr2[i2])
			i2++
		}
	}

	for i1 < len(arr1) {
		result = append(result, arr1[i1])
		i1++
	}

	for i2 < len(arr2) {
		result = append(result, arr2[i2])
		i2++
	}

	return result
}
