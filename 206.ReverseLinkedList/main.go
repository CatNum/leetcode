package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList2(head *ListNode)( *ListNode)  {
	return
}
//迭代实现反转链表
//想象成两个链表，第一个链表是反转之后的结点，第二个是还未反转的结点
//prev指向第一个链表的头结点，next指向第二个链表的头结点
//可以想象为两个链表，一个头结点为prev，一个头结点为next，curr为反转时候用的结点
func reverseList(head *ListNode) *ListNode {
	curr := head
	var prev *ListNode
	for curr != nil {
		//先将未反转的结点拿下来
		next := curr.Next
		//当前结点指向前一个链表的头结点
		curr.Next = prev
		//将第一个链表的头结点（prev）指向当前节点（当前节点即新加入第一个链表的结点，即头结点）
		prev = curr
		//当前节点指向第二个链表的头结点
		curr = next
	}
	return prev
}

// 递归实现反转链表
//递归函数之前是递的过程的操作。递归函数之后是归的过程的操作
func reverseList1(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var p = reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	//此处的p一直是反转后链表的头结点，即未反转链表的最后一个结点
	return p
}

func main() {
	head := new(ListNode)
	head.Val = 1
	ln2 := new(ListNode)
	ln2.Val = 2
	ln3 := new(ListNode)
	ln3.Val = 3
	ln4 := new(ListNode)
	ln4.Val = 4
	ln5 := new(ListNode)
	ln5.Val = 5
	head.Next = ln2
	ln2.Next = ln3
	ln3.Next = ln4
	ln4.Next = ln5
	//为了遍历的时候保护head不变设置cur
	cur := head
	for cur != nil {
		fmt.Println(cur)
		cur = cur.Next
	}
	newhead := reverseList(head)
	for newhead != nil {
		fmt.Println(newhead)
		newhead = newhead.Next
	}
}
