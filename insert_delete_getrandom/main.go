package main

import (
	"fmt"
	"math/rand"
)

type RandomizedSet struct {
	hashmap map[int]int
	values  []int
}

func main() {
	rSet := Constructor()
	fmt.Println(rSet.Insert(1))
	fmt.Println(rSet.Remove(2))
	fmt.Println(rSet.Insert(2))
	fmt.Println(rSet.GetRandom())
	fmt.Println(rSet.Remove(1))
	fmt.Println(rSet.Insert(2))
	fmt.Println(rSet.GetRandom())
}

func Constructor() RandomizedSet {
	return RandomizedSet{
		hashmap: make(map[int]int),
		values:  make([]int, 0),
	}
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, exist := this.hashmap[val]; exist {
		return false
	}

	this.hashmap[val] = len(this.values)
	this.values = append(this.values, val)

	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	idx, exist := this.hashmap[val]
	if !exist {
		return false
	}

	lastVal := this.values[len(this.values)-1]
	this.values[idx] = lastVal

	this.hashmap[lastVal] = idx

	this.values = this.values[:len(this.values)-1]
	delete(this.hashmap, val)

	return true
}

func (this *RandomizedSet) GetRandom() int {
	return this.values[rand.Intn(len(this.values))]
}
