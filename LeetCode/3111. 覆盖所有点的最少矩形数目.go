package LeetCode

import "sort"

func minRectanglesToCoverPoints(points [][]int, w int) int {
	sort.Slice(points, func(i, j int) bool {
		return points[i][0] < points[j][0]
	})
	res := 0
	bound := -1
	for _, p := range points {
		if p[0] > bound {
			bound = p[0] + w
			res++
		}
	}
	return res
}
