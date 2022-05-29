package main

import "fmt"

//Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
	slow,fast := head,head
	for fast != nil && fast.Next != nil{
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

func main() {
	var l0 ListNode = ListNode{Val: 0}
	var l1 ListNode = ListNode{Val: 1}
	var l2 ListNode = ListNode{Val: 2}
	var l3 ListNode = ListNode{Val: 3}
	var l4 ListNode = ListNode{Val: 4}
	var l5 ListNode = ListNode{Val: 5}
	l0.Next = &l1
	l1.Next = &l2
	l2.Next = &l3
	l3.Next = &l4
	l4.Next = &l5
	head := middleNode(&l0)
	fmt.Println(head.Val)
}
