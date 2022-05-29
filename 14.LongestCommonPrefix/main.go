package main

import "fmt"
//自己编写
func longestCommonPrefix(strs []string) string {
	//选取第一个作为公共前缀
	temp := strs[0]
	for i := 1; i < len(strs); i++ {
		//如果有一个字符串为空，则返回空
		if len(strs[i]) == 0 {
			temp = ""
		}
		//遍历temp和一个子串并进行各个字符的比较
		//三种情况
		//1.某个字符不等，切割
		//2.temp遍历完了，但strs[i]还没有，则temp不变
		//3.temp没遍历完，strs[i]遍历完了，则temp=strs[i]
		for n := 0; n < len(temp) && n < len(strs[i]); n++ {
			if temp[n] != strs[i][n] {
				temp = temp[:n]
			} else if n == len(temp)-1 && n != len(strs[i])-1 {
				temp = temp
			} else if n == len(strs[i])-1 && n != len(temp)-1 {
				temp = strs[i]
			}
		}
	}
	return temp
}

func main() {
	strs := []string{"flower", "flow", "flight"}
	re := longestCommonPrefix(strs)
	fmt.Println(re)
}
