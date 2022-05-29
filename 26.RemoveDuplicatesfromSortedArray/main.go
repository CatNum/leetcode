package main
//26. 删除有序数组中的重复项
//双指针实现，left，right，如果nums[left] != nums[right]，就将nums[right]赋值给nums[left]
func removeDuplicates(nums []int) int {
	left := 0
	if len(nums) == 0{
		return 0
	}
	for _,v := range nums{
		if v != nums[left]{
			nums[left+1] = v
			left++
		}
	}
	return left+1
}