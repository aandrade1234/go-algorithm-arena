package main

import "fmt"

type tuple struct {
	i int
	j int
}

const value = 2

var directions = []tuple{
	{1, 0},
	{-1, 0},
	{0, 1},
	{0, -1},
}

func main() {
	fmt.Println(shortestBridge([][]int{
		{0, 1},
		{1, 0},
	}))
	fmt.Println(shortestBridge([][]int{
		{0, 1, 0},
		{0, 0, 0},
		{0, 0, 1},
	}))
	fmt.Println(shortestBridge([][]int{
		{1, 1, 1, 1, 1},
		{1, 0, 0, 0, 1},
		{1, 0, 1, 0, 1},
		{1, 0, 0, 0, 1},
		{1, 1, 1, 1, 1},
	}))
}

func shortestBridge(grid [][]int) int {
	i, j := findIsland(grid)

	queue := make([]tuple, 0)
	fillIsland(i, j, grid, &queue)

	return growIsland(grid, queue)
}

func growIsland(grid [][]int, queue []tuple) int {
	distance := 0
	for len(queue) > 0 {
		newQueue := make([]tuple, 0)
		for _, item := range queue {
			for _, dir := range directions {
				i, j := item.i+dir.i, item.j+dir.j
				if i >= 0 && i < len(grid) && j >= 0 && j < len(grid[0]) {
					if grid[i][j] == 0 {
						grid[i][j] = 2
						newQueue = append(newQueue, tuple{i: i, j: j})
					} else if grid[i][j] == 1 {
						return distance
					}
				}
			}
		}
		queue = newQueue
		distance++
	}
	return -1
}

func fillIsland(i, j int, grid [][]int, queue *[]tuple) {
	q := []tuple{{i, j}}
	grid[i][j] = 2

	for len(q) > 0 {
		current := q[0]
		q = q[1:]

		*queue = append(*queue, current)

		for _, dir := range directions {
			x, y := current.i+dir.i, current.j+dir.j
			if x >= 0 && x < len(grid) && y >= 0 && y < len(grid[0]) && grid[x][y] == 1 {
				grid[x][y] = 2
				q = append(q, tuple{i: x, j: y})
			}
		}
	}
}

func findIsland(grid [][]int) (i int, j int) {
	for i = range grid {
		for j = range grid[i] {
			if grid[i][j] == 1 {
				return i, j
			}
		}
	}

	return -1, -1
}
