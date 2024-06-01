package main

import "math"

func minAreaRect(points [][]int) int {
	vertices := make(map[int]map[int]struct{})

	for _, point := range points {
		x, y := point[0], point[1]
		if _, found := vertices[x]; !found {
			vertices[x] = make(map[int]struct{})
		}

		vertices[x][y] = struct{}{}
	}

	minArea := math.MaxInt

	for i := range len(points) {
		for j := range len(points) {
			x1, y1 := points[i][0], points[i][1]
			x2, y2 := points[j][0], points[j][1]

			if x1 != x2 && y1 != y2 {
				if _, found1 := vertices[x1][y2]; found1 {
					if _, found2 := vertices[x2][y1]; found2 {
						area := abs(x1-x2) * abs(y1-y2)
						if area < minArea {
							minArea = area
						}
					}
				}
			}
		}
	}

	if minArea == math.MaxInt {
		minArea = 0
	}

	return minArea
}

func abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}
