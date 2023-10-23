package main

func maxArea(height []int) int {
	i, j, res := 0, len(height)-1, 0
	s := 0
	for i < j {
		if height[i] < height[j] {
			s = height[i] * (j - i)
			res = max(res, s)
			i++
		} else {
			s = height[j] * (j - i)
			res = max(res, s)
			j--
		}
	}
	return res

}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func main() {
	maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7})
}
