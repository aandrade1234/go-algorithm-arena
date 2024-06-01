package main

func twoSum(nums []int, target int) []int {
	cMap := make(map[int]int)

	for i, num := range nums {
		complement := target - num
		if _, found := cMap[complement]; found {
			return []int{cMap[complement], i}
		}

		cMap[num] = i
	}

	return []int{}
}
