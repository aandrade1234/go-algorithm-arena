package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(FindIntersection([]string{"1, 3, 4, 7, 13", "1, 2, 4, 13, 15"}))
	fmt.Println(FindIntersection([]string{"1, 3, 9, 10, 17, 18", "1, 4, 9, 10"}))
}

func FindIntersection(strArr []string) string {
	set1 := strToSet(strArr[0])
	set2 := strToSet(strArr[1])

	intersection := make([]int, 0)

	for num := range set1 {
		if _, found := set2[num]; found {
			intersection = append(intersection, num)
		}
	}

	if len(intersection) == 0 {
		return "false"
	}

	sort.Ints(intersection)

	strResult := make([]string, len(intersection))
	for i, val := range intersection {
		strResult[i] = strconv.Itoa(val)
	}

	return strings.Join(strResult, ",")
}

func strToSet(str string) map[int]struct{} {
	nums := strings.Split(str, ",")
	set := make(map[int]struct{}, len(nums))

	for _, num := range nums {
		val, err := strconv.Atoi(strings.TrimSpace(num))
		if err == nil {
			set[val] = struct{}{}
		}
	}

	return set
}
