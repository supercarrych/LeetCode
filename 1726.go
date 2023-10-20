package main

import "fmt"

func tupleSameProduct(nums []int) int {
	res := 0
	//cut := len(nums) / 2
	//fontNums := nums[:cut]
	//backNums := nums[cut:]
	tupleSame := map[int]int{}
	for i, v := range nums {
		for j := i + 1; j < len(nums); j++ {
			fmt.Println(v * nums[j])
			if _, ok := tupleSame[v*nums[j]]; ok {
				tupleSame[v*nums[j]]++
			} else {
				tupleSame[v*nums[j]] = 1
			}
		}
	}

	for _, value := range tupleSame {
		if value >= 2 {
			res += value * (value - 1) / 2
		}
	}
	fmt.Println(res)
	return res * 8
}

func main() {
	nums := []int{2, 3, 4, 6, 8, 12}
	tupleSameProduct(nums)
}
