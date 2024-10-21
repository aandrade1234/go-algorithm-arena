package main

import "fmt"

func main() {
	fmt.Println(subarraysDivByK([]int{4, 5, 0, -2, -3, 1}, 5))
	fmt.Println(subarraysDivByK([]int{5}, 9))
}

func subarraysDivByK(nums []int, k int) int {
	remainders := make(map[int]int)
	remainders[0] = 1

	counter := 0
	sum := 0

	for _, num := range nums {
		sum += num
		remainder := (sum%k + k) % k
		counter += remainders[remainder]
		remainders[remainder]++
	}

	return counter
}
