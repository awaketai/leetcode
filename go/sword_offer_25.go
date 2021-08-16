package main

import (
	"fmt"
)

// 剑指 Offer 25. 合并两个排序的链表

// 输入两个递增排序的链表，合并这两个链表并使新链表中的节点仍然是递增排序的。


  //Definition for singly-linked list.
type ListNode struct {
  Val int
  Next *ListNode
}


func main()  {
	ret := mergeTwoLists(node11(),node22())
	fmt.Println(ret)
	listToSlice(ret)
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	node := &ListNode{
		Val: 0,
	}
	current := node
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			current.Next = l1
			l1 = l1.Next
		}else{
			current.Next = l2
			l2 = l2.Next
		}
		current = current.Next
	}
	// 把没合并完的直接合并
	if l1 != nil {
		current.Next = l1
	}
	if l2 != nil {
		current.Next = l2
	}
	return node.Next
}

func listToSlice(l *ListNode)  {
	var s = make([]int,0)
	for l != nil {
		s = append(s,l.Val)
		l = l.Next
	}
	fmt.Println(s)
}

func node11() *ListNode {
	node := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 4,
			},
		},
	}
	return node
}

func node22() *ListNode {
	node := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val: 4,
			},
		},
	}
	return node
}


