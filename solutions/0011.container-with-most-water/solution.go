package solution

func maxArea(height []int) int {
	var (
		x       = 0
		y       = len(height) - 1
		maxArea = 0
	)

	for x < y {
		area := min(height[x], height[y]) * (y - x)
		if maxArea < area {
			maxArea = area
		}

		if height[x] <= height[y] {
			x++
		} else {
			y--
		}
	}

	return maxArea
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}
