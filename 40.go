package main

import (
	"fmt"
	"sort"
)

func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	var res [][]int
	l := len(candidates)
	var path []int
	if l == 0 {
		return res
	}
	var dfs func([]int, int, int, int)
	dfs = func(candidates []int, l int, index int, target int) {

		if target == 0 {
			res = append(res, append([]int{}, path...))
			return
		}

		for i := index; i < l; i++ {
			//大剪枝
			if target < candidates[i] {
				return
			}

			//小剪枝
			if i > index && candidates[i] == candidates[i-1] {
				continue
			}

			path = append(path, candidates[i])
			fmt.Printf("递归前：%s 当前剩余：%d\n", path, target)
			dfs(candidates, l, i+1, target-candidates[i])

			path = path[:len(path)-1]
			fmt.Printf("递归后：%s  当前剩余：%d\n", path, target)
		}

	}

	dfs(candidates, l, 0, target)
	fmt.Println(res)
	return res

}

func main() {
	candidates := []int{1, 2, 2, 2, 5}
	combinationSum2(candidates, 5)
}
