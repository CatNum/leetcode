package main

import "fmt"

//88. 合并两个有序数组

func merge(nums1 []int, m int, nums2 []int, n int)  {
	//逆向双指针，利用nums1中的后半部分空余部分
	if m == 0 {
		//for i,v  := range nums2 {
		//	nums1[i] = v
		//}
		fmt.Println("之前：",&nums1[0])
		nums1 = nums2   //会改变地址
		fmt.Println(nums1)
		//nums1[0] =nums2[0]    //不会改变地址
		fmt.Println("之后：",&nums1[0])
		return
	}
	if n == 0{
		return
	}
	i := 1
	for m != 0 || n != 0{
		if m == 0{
			copy(nums1[:n],nums2[:n])
			return
		}
		if n == 0{
			return
		}
		if nums1[m-1] > nums2[n-1]{
			nums1[cap(nums1)-i] = nums1[m-1]
			m--
			i++
		}else {
			nums1[cap(nums1)-i] = nums2[n-1]
			n--
			i++
		}
	}
	return
}
func main(){
	nums1 := []int{0}
	nums2 := []int{1}
	fmt.Println("开始：",&nums1[0])
	merge(nums1,0,nums2,1)
	//copy(nums1[:],nums2)
	fmt.Println(nums1,nums2)
}