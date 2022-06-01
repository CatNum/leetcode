package main

import "fmt"

func main() {
	nums := []int{5,6,4,3,2,1}
	mpSort(nums)
	fmt.Println(nums)
}

//func sort(nums []int) {
//	for i,k:=0,0;k<len(nums);k++ {
//		for j := i; j < len(nums)-i-1; j++ {
//			if nums[j] > nums[j+1]{
//				nums[j],nums[j+1] = nums[j+1],nums[j]
//			}
//		}
//	}
//}
//
//func main() {
//	nums := []int{5,6,4,3,2,1}
//	sort(nums)
//	fmt.Println(nums)
//}

//二叉树 最大高度

//type TreeNode struct {
//	Left *TreeNode
//	Right *TreeNode
//	Val int
//}
//func high(tree *TreeNode) int{
//	var hight int
//	sum := 0
//	if tree.Left == nil && tree.Right == nil{
//		sum =sum+1
//		return high()
//	}else if tree.Left != nil {
//		return high(tree.Left)
//		sum++
//	}else if tree.Right != nil{
//		return high(tree.Right)
//		sum++
//	}
//
//}
