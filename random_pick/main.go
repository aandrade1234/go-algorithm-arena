package main

import (
	"math/rand"
	"sort"
)

type Solution struct {
	sum       int
	sumPrefix []int
}

func Constructor(w []int) Solution {
	sumPrefix := make([]int, len(w))
	sumPrefix[0] = w[0]
	for i := 1; i < len(w); i++ {
		sumPrefix[i] = sumPrefix[i-1] + w[i]
	}
	sum := sumPrefix[len(w)-1]

	solution := Solution{
		sum:       sum,
		sumPrefix: sumPrefix,
	}

	return solution
}

func (this *Solution) PickIndex() int {
	target := rand.Intn(this.sum) + 1
	return sort.Search(len(this.sumPrefix), func(i int) bool {
		return this.sumPrefix[i] >= target
	})
}
