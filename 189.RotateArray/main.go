package main

import "fmt"

//自己写的
func rotate(nums []int, k int) {
	//这里直接使用
	//原因：append在len超过cap的时候 会进行扩容
	//nums = append(nums[len(nums)-k%len(nums):],nums[:len(nums)-k%len(nums)]...)
	//fmt.Println(len(nums),cap(nums),nums)      //这里输出cap(nums)=8 证明nums发生了扩容，所以发生了改变
	//或者
	//原因：slice是引用类型，直接使用 = 赋值会改变nums的值，即实参传过来的地址，但是形参不是实参，所以只改变了形参
	//而函数外的实参没有改变
	//nums1 := append(nums[len(nums)-k%len(nums):],nums[:len(nums)-k%len(nums)]...)
	//nums = nums1
	//fmt.Println(&nums[0])
	//会造成nums的地址发生改变。。。但是不知道为什么会发生改变
	nums1 := append(nums[len(nums)-k%len(nums):], nums[:len(nums)-k%len(nums)]...)
	for i, i2 := range nums1 {
		nums[i] = i2
	}
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	n := 3
	fmt.Println(&nums[0])
	rotate(nums, n)
	fmt.Println(&nums[0])
	fmt.Println(nums)

}
