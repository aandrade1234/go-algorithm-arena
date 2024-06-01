package main

import "fmt"

func main() {

	fmt.Println(containsDuplicate([]int{1, 2, 3, 1}))
	fmt.Println(containsDuplicate([]int{1, 2, 3, 4}))
	fmt.Println(containsDuplicate([]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}))

}

func containsDuplicate(nums []int) bool {
	container := make(map[int]struct{}, len(nums))
	for _, num := range nums {
		if _, found := container[num]; found {
			return true
		} else {
			container[num] = struct{}{}
		}
	}
	return false
}
