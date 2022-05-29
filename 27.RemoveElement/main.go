package main
//27. 移除元素
func removeElement(nums []int, val int) int {
	for i := 0; i < len(nums); i++ {
		if nums[i] == val {
			//处理最后一个等于 val的情况
			if i == len(nums)-1 {
				nums = nums[:i]
				break
			}
			nums = append(nums[:i], nums[i+1:]...)
			// i-- 是因为 nums 切割，造成了下标对应的值发生了改变
			i--
		}
	}
	return len(nums)
}
//双指针解法：right指向当前将要处理的元素，left指向下一个将要赋值的元素
//v相当于一个指针，left相当于一个指针
func removeElement1(nums []int, val int) int {
	left := 0
	for _,v := range nums{
		if v != val {
			nums[left] = v
			left++
		}
	}
	return left
}