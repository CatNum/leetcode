package main

import "fmt"

func addToArrayForm1(num []int, k int) (ans []int) {
	//低位相加
	for i := len(num) - 1; i >= 0; i-- {
		sum := num[i] + k%10
		// k = k/10要在 k++ 之前，因为这样才能在进位上+1
		k /= 10
		if sum >= 10 {
			//k++是因为k此时的低位会加到进位上
			k++
			sum -= 10
		}
		//创建一个数组从低位到高位存放结果
		ans = append(ans, sum)
	}
	//加完了之后 k还存在，即若数组的位数小于数字的位数的情况
	for ; k > 0; k /= 10 {
		ans = append(ans, k%10)
	}
	reverse(ans)
	return
}

func reverse(num []int) {
	for i, n := 0, len(num); i < n/2; i++ {
		num[i], num[n-1-i] = num[n-1-i], num[i]
	}
}

//思路：取余从低位到高位加，逻辑太复杂，没写出来
//涉及的问题：
//	1.进位
//	2.越界，包括进位越界和num为 0，k为 23 的情况的越界
//官方题解从低位到高位加然后将结果放到数组中再翻转就避免了越界的问题
func addToArrayForm(num []int, k int) []int {
	i := len(num) - 1
	for k > 0 {
		n := k % 10
		num[i] = num[i] + n
		j := i
		//如果发生了进位
		for num[j] >= 10 {
			if j > 0 {
				num[j] = num[j] % 10
				num[j-1] = num[j-1] + 1
				j--
			} else if j == 0 { //进位造成越界，最多越界一次
				num[j] = num[j] % 10
				nums := []int{1}
				num = append(nums, num...)
			}
		}
		k = k / 10
		i--
	}
	return num
}

func main() {
	nums := []int{2, 7, 4}
	n := 181
	num := addToArrayForm1(nums, n)
	fmt.Println(num)
}

