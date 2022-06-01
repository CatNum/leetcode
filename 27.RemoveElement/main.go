package main

import "fmt"

// 27. 移除元素
// 暴力解法
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

// 双指针解法：right【函数中的v】指向当前将要处理的元素，left指向下一个将要赋值的元素
// v相当于一个指针，left相当于一个指针
func removeElement1(nums []int, val int) int {
	left := 0
	for _, v := range nums {
		if v != val {
			nums[left] = v
			left++
		}
	}
	return left
}

// 左右双指针，如果左指针等于val，左右指针调换数据，右指针左移
// 如果左指针不等于val，左指针右移
// 注意：注意数组大小为1的情况
func removeElement2 (nums []int, val int) (num int) {
	// 左右双指针
	left, right := 0, len(nums)-1
	for left != right+1{
		if nums[left] == val {
			nums[left], nums[right] = nums[right], nums[left]
			// 针对数组大小为 1 的情况
			if right-1 < 0 {
				num = 0
				break
			}
			right--
			num = right + 1
		} else {
			// 针对数组大小为 1 的情况
			if left+1 > len(nums) {
				num = 1
				break
			}
			left++
			num = left
		}
	}
	return num
}

func main() {
	nums := []int{ 2}
	count := removeElement2(nums, 3)
	fmt.Println(count)
}
