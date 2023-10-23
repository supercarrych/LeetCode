package main

import (
	"container/heap"
	"fmt"
	"sort"
)

type hp struct{ sort.IntSlice }

func (h hp) Less(i, j int) bool { return h.IntSlice[i] > h.IntSlice[j] } // 最大堆
func (hp) Push(any)             {}
func (hp) Pop() (_ any)         { return }

func maxKelements(nums []int, k int) int64 {
	var score int64
	h := hp{nums}
	heap.Init(&h)

	for ; 0 < k; k-- {
		score += int64(h.IntSlice[0])
		h.IntSlice[0] = (h.IntSlice[0] + 2) / 3
		heap.Fix(&h, 0)
	}
	return score
}

func main() {
	nums := []int{1, 10, 3, 3, 3}
	score := maxKelements(nums, 3)
	fmt.Println(score)
}
