package main

import (
	"fmt"
)

type node struct {
	value int
	prev  *node
	next  *node
}

type list struct {
	head *node
	tail *node
	size int
}

func main() {
	fmt.Println(numBusesToDestination([][]int{{1, 2, 7}, {3, 6, 7}}, 1, 6))
	fmt.Println(numBusesToDestination([][]int{{7, 12}, {4, 5, 15}, {6}, {15, 19}, {9, 12, 13}}, 15, 12))
}

func numBusesToDestination(routes [][]int, source, target int) int {
	if source == target {
		return 0
	}

	// create a map[busStop] = [route index]
	stopToRoutes := make(map[int][]int)

	for i, route := range routes {
		for _, stop := range route {
			stopToRoutes[stop] = append(stopToRoutes[stop], i)
		}
	}

	visitedStops := make(map[int]bool)
	visitedStops[source] = true

	visitedRoutes := make(map[int]bool)
	buses := 0

	queue := []int{source}
	for len(queue) > 0 {
		buses++
		nextQueue := []int{}

		for _, stop := range queue {
			for _, route := range stopToRoutes[stop] {
				if visitedRoutes[route] {
					continue
				}

				visitedRoutes[route] = true

				for _, nextStop := range routes[route] {
					if nextStop == target {
						return buses
					}

					if !visitedStops[nextStop] {
						visitedStops[nextStop] = true

						nextQueue = append(nextQueue, nextStop)
					}
				}
			}
		}

		queue = nextQueue
	}

	return -1
}
