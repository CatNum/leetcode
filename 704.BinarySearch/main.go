package main

import "fmt"
//自己编的二分查找
func find(nums []int, v int) (n int) {
	low := 0
	high := len(nums) - 1
	flag := true
	for flag && high >= low {
		m := (low + high) / 2
		if v == nums[m] {
			n = m
			flag = false
		} else if v > nums[m] {
			low = m + 1
		} else {
			high = m - 1
		}
	}
	if flag {
		n = -1
	}
	return n
}
//力扣官方题解二分查找
func search(nums []int, target int) int {
	low, high := 0, len(nums)-1
	for low <= high {
		mid := (high-low)/2 + low
		num := nums[mid]
		if num == target {
			return mid
		} else if num > target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1
}

func main() {
	nums := []int{5}
	n := 5
	i := find(nums, n)
	fmt.Println(i)
}
