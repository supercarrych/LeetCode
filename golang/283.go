package main

func main() {
	nums := []int{0, 0, 1}
	moveZeroes(nums)
}

func moveZeroes(nums []int) {
	zeroNums := []int{}
	i := 0
	for range nums {
		if nums[i] == 0 {
			//入队
			zeroNums = append(zeroNums, nums[i])
			//使用切片将元素从原数组中删除
			nums = append(nums[:i], nums[i+1:]...)
		} else {
			i++
		}
	}
	nums = append(nums, zeroNums...)
}
