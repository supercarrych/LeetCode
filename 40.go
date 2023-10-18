package main

func combinationSum2(candidates []int, target int) [][]int {
	var mp map[int]int
	var res [][]int

}

func f(candidates []int, target int, mp map[int]int) []int {

	for v := range candidates {
		if _, ok := mp[target-v]; ok {

		}
	}
}
