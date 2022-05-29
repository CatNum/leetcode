package main

import "fmt"

func sort(nums []int) {
	for i,k:=0,0;k<len(nums);k++ {
		for j := i; j < len(nums)-i-1; j++ {
			if nums[j] > nums[j+1]{
				nums[j],nums[j+1] = nums[j+1],nums[j]
			}
		}
	}
}

func main() {
	nums := []int{5,6,4,3,2,1}
	sort(nums)
	fmt.Println(nums)
}
