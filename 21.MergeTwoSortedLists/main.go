package main

import "fmt"

//Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	//创建一个结点表示新链表
	newHead := &ListNode{}
	//创建一个结点用于链表的遍历和增加
	curr := newHead
	for list1 != nil || list2 != nil {
		//如果其中一个链表空了
		if list1 == nil {
			curr.Next = list2
			break
		}
		if list2 == nil {
			curr.Next = list1
			break
		}
		//取两个链表头结点小的放入新链表
		if list1.Val > list2.Val {
			curr.Next = list2
			list2 = list2.Next
		} else {
			curr.Next = list1
			list1 = list1.Next
		}
		curr = curr.Next
	}
	return newHead.Next
}

func main() {
	var l0 ListNode = ListNode{Val: 0}
	var l1 ListNode = ListNode{Val: 1}
	var l2 ListNode = ListNode{Val: 2}
	var l3 ListNode = ListNode{Val: 3}
	var l4 ListNode = ListNode{Val: 4}
	var l5 ListNode = ListNode{Val: 5}
	var l6 ListNode = ListNode{Val: 6}
	var l7 ListNode = ListNode{Val: 7}
	l0.Next = &l2
	l1.Next = &l3
	l2.Next = &l4
	l3.Next = &l5
	l4.Next = &l6
	l5.Next = &l7
	sum := mergeTwoLists(&l0, &l1)
	for sum != nil {
		fmt.Println(sum.Val)
		sum = sum.Next
	}
}
