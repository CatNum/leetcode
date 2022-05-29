package main

import "fmt"

func sort(nums []int) []int {
	//退出条件
	if len(nums) < 2 {
		return nums
	}
	//选取第一个作为基准
	temp := nums[0]
	left, right := make([]int, 0), make([]int, 0)
	//从 nums[1] 开始进行比较
	for i := 1; i < len(nums); i++ {
		if nums[i] > temp {
			right = append(right, nums[i])
		} else {
			left = append(left, nums[i])
		}
	}
	//对左边和右边也进行这样的操作
	left = sort(left)
	right = sort(right)
	return append(append(left, temp), right...)
}

func main() {
	var nums = []int{1,5,8,3,60,9,58,1,22,44,58}
	nums1 := sort(nums)
	fmt.Println("排序前：",nums)
	fmt.Println("排序后：",nums1)
}
