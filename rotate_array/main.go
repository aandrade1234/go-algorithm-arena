package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	rotate(nums, 3)
	fmt.Println(nums)

	nums = []int{-1, -100, 3, 99}
	rotate(nums, 2)
	fmt.Println(nums)
}

func rotate(nums []int, k int) {
	n := len(nums)
	k %= n // if k > len(nums)

	reverse(nums, 0, n-1) // Step 1: reverse whole array [7 6 5 4 3 2 1]
	reverse(nums, 0, k-1) // Step 2: reverse first k elements [5 6 7 4 3 2 1]
	reverse(nums, k, n-1) // Step 3: reverse last n-k elements [5 6 7 1 2 3 4]
}

func reverse(nums []int, start, end int) {
	for start < end {
		nums[start], nums[end] = nums[end], nums[start]
		start++
		end--
	}
}
