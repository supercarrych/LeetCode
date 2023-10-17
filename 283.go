package main

import "fmt"

func main() {
	nums := []int{0, 0, 1}
	moveZeroes(nums)
}

func moveZeroes(nums []int) {
	zeroNums := []int{}
	i := 0
	for true {
		if nums[=i] {

		}
		num := nums[i]
		//使用切片将元素从原数组中删除
		nums = append(nums[:i], nums[i+1:]...)
		fmt.Printf("out %v\n", nums)
		//入队
		zeroNums = append(zeroNums, num)
		fmt.Printf("in %v\n", zeroNums)

	}

	nums = append(nums, zeroNums...)
	fmt.Println(nums)

}
