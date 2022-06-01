package main

// mpSort 冒泡排序
func mpSort(nums []int) {
	for i,k:=0,0;k<len(nums);k++ {
		for j := i; j < len(nums)-i-1; j++ {
			if nums[j] > nums[j+1]{
				nums[j],nums[j+1] = nums[j+1],nums[j]
			}
		}
	}
}