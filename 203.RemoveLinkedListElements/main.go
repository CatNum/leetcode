package main

import "fmt"

//203. 移除链表元素
//Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}
//自己写的迭代写法
func removeElements(head *ListNode, val int)  *ListNode {
	//需要返回的新头结点  作用：表示一个链表
	newHead := &ListNode{Next: head}
	//需要操作的当前结点  作用：用于遍历，标识当前结点
	temp := newHead
	for temp.Next != nil {
		if temp.Next.Val == val {
			temp.Next = temp.Next.Next
		}else {
			temp = temp.Next
		}
	}

	return newHead.Next
}
//官方题解迭代解法
func removeElements1(head *ListNode, val int) *ListNode {
	dummyHead := &ListNode{Next: head}
	for tmp := dummyHead; tmp.Next != nil; {
		if tmp.Next.Val == val {
			tmp.Next = tmp.Next.Next
		} else {
			tmp = tmp.Next
		}
	}
	return dummyHead.Next
}
//官方写的递归写法
func removeElements2(head *ListNode, val int) *ListNode {
	//退出条件
	if head == nil {
		return head
	}
	head.Next = removeElements(head.Next, val)
	//head.Val == val，删除head
	if head.Val == val {
		return head.Next
	}
	//head.Val ！= val，不删除head
	return head
}

func main() {
	var l0 ListNode = ListNode{Val: 0}
	var l1 ListNode = ListNode{Val: 1}
	var l2 ListNode = ListNode{Val: 2}
	var l3 ListNode = ListNode{Val: 3}
	var l4 ListNode = ListNode{Val: 2}
	var l5 ListNode = ListNode{Val: 4}
	l0.Next = &l1
	l1.Next = &l2
	l2.Next = &l3
	l3.Next = &l4
	l4.Next = &l5
	new1 := removeElements(&l0,2)
	for new1 != nil{
		fmt.Println(new1.Val)
		new1 =new1.Next
	}

}
