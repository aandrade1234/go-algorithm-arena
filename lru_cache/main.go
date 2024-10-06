package main

import "fmt"

func main() {
	cache := Constructor(2)
	cache.Put(1, 1)
	cache.Put(2, 2)
	fmt.Println(cache.Get(1)) // 1
	cache.Put(3, 3)
	fmt.Println(cache.Get(2)) // -1 (not found)
	cache.Put(4, 4)
	fmt.Println(cache.Get(1)) // -1 (not found)
	fmt.Println(cache.Get(3)) // 3
	fmt.Println(cache.Get(4)) // 4
}

type LRUCache struct {
	capacity int
	head     *Node
	tail     *Node

	index map[int]*Node
}

type Node struct {
	key      int
	value    int
	previous *Node
	next     *Node
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity: capacity,
		index:    make(map[int]*Node),
	}
}

func (this *LRUCache) Get(key int) int {
	node, ok := this.index[key]
	if !ok {
		return -1
	}

	this.moveTop(node)

	return node.value
}

func (this *LRUCache) Put(key int, value int) {
	node, ok := this.index[key]
	if ok {
		node.value = value

		this.moveTop(node)
	} else {
		node := &Node{
			key:   key,
			value: value,
			next:  this.head,
		}

		if this.head != nil {
			this.head.previous = node
		}

		this.head = node

		if this.tail == nil {
			this.tail = node
		}

		this.index[key] = node

		if len(this.index) > this.capacity {
			delete(this.index, this.tail.key)

			if this.tail.previous != nil {
				this.tail = this.tail.previous
				this.tail.next = nil
			} else {
				this.head = nil
				this.tail = nil
			}
		}
	}
}

func (this *LRUCache) moveTop(node *Node) {
	if node != this.head {
		node.previous.next = node.next

		if node.next == nil {
			this.tail = node.previous
		} else {
			node.next.previous = node.previous
		}

		node.previous = nil
		node.next = this.head

		if this.head != nil {
			this.head.previous = node
		}

		this.head = node
	}
}
