package main

import "fmt"

//88. 合并两个有序数组
func merge1(nums1 []int, m int, nums2 []int, n int) {
	//逆向双指针
	for p1, p2, tail := m-1, n-1, m+n-1; p1 >= 0 || p2 >= 0; tail-- {
		var cur int
		if p1 == -1 {
			cur = nums2[p2]
			p2--
		} else if p2 == -1 {
			cur = nums1[p1]
			p1--
		} else if nums1[p1] > nums2[p2] {
			cur = nums1[p1]
			p1--
		} else {
			cur = nums2[p2]
			p2--
		}
		nums1[tail] = cur
	}
}

func merge2(nums1 []int, m int, nums2 []int, n int)  {
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
	nums1 := []int{1,2,3,0,0,0}
	nums2 := []int{2,5,6}
	fmt.Println("开始：",&nums1[0])
	merge1(nums1,3,nums2,3)
	//copy(nums1[:],nums2)
	fmt.Println(nums1,nums2)
}

