package main

func max(i int, j int) int {
	if i > j {
		return i
	}
	return j
}

func trap(height []int) int {
	left := 0
	right := 1
	ans := 0
	if len(height) <= 1 {
		return ans
	}

	for i := 0; i < len(height); i++ {
		left = i

		res := 0
		for height[right] < height[left] {
			maxN := max(height[left], height[right])
			if right == len(height)-1 && height[right] < maxN {
				res = 0
				left++
				right = left
				break
			} else if height[right] >= maxN {
				right = len(height) - 1
				break
			}

			res += maxN - height[right]

			right++
		}
		if right-left != 1 && left == i {
			ans += res
			i = right - 1
		}
		if right == len(height)-1 {
			break
		}
		right++

	}
	return ans
}
