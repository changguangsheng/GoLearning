package Leetcode

func maxPoint(point [][]int, i int) int {
	Count := 0
	n := len(point)
	slopes := make(map[float64]int)
	slope := make(map[int]int)
	for j := 0; i != j && j < n; j++ {
		dx := float64(point[i][0] - point[j][0])
		dy := float64(point[i][1] - point[j][1])

		if dy == 0 {
			slope[point[i][1]]++
			if slope[point[i][1]] > Count {
				Count = slope[point[i][1]]
			}
			continue
		}

		angle := float64(dx / dy)

		slopes[angle]++
		if slopes[angle] > Count {
			Count = slopes[angle]
		}
	}
	return Count + 1
}

func maxPoints(points [][]int) int {
	ans := 1
	Numpoint := len(points)
	if Numpoint <= 2 {
		return Numpoint
	}
	for i := 0; i < Numpoint; i++ {
		num := maxPoint(points, i)
		// fmt.Println(num)
		if ans < num {
			ans = num
		}
	}
	return ans
}
