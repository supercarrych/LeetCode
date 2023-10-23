package main

import "sort"

func maxSatisfaction(satisfaction []int) int {
	sort.Slice(satisfaction, func(i, j int) bool {
		return satisfaction[i] > satisfaction[j]
	})
	presum, ans := 0, 0
	for _, num := range satisfaction {
		if presum+num > 0 {
			presum += num
			ans += presum
		} else {
			break
		}
	}
	return ans
}
