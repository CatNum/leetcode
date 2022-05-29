package main

import "fmt"
//1.两数之和
func twoSum(nums []int, target int) []int {
	map1 := make(map[int]int)
	for i, num := range nums {
		num1, ok := map1[target-num]
		//如果存在map中某个值的下标 = target-num
		if ok {
			return []int{i, num1}
		}
		//将数组中的数据存入map
		map1[num] = i
	}
	return nil
}

func main(){
	nums := []int{2,7,11,15}
	fmt.Println(twoSum(nums,9))
}