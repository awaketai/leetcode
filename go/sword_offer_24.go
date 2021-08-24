package main

import (
	"fmt"
)

// 剑指 Offer 24. 反转链表

// 定义一个函数，输入一个链表的头节点，反转该链表并输出反转后链表的头节点。
//
// 
//
//示例:
//
//输入: 1->2->3->4->5->NULL
//输出: 5->4->3->2->1->NULL
//

func main()  {
	node := reverseList(node())
	listToSlice(node)
}

  //Definition for singly-linked list.
type ListNode struct {
  Val int
  Next *ListNode
}

// 双指针法
// 定义两个指针：pre/cur，pre在前，cur在后
// 每次让pre的next指向cur，实现一次局部反转
// 局部反转完成后，pre和cur同时往前移动一个位置
// 循环执行上述过程，知道pre到达链表尾部
func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	pre,cur := head,&ListNode{}
	for pre != nil {
		pre.Next = cur


	}




}

func listToSlice(head *ListNode)  {
	s := make([]int,0)
	for head != nil {
		s = append(s,head.Val)
		head = head.Next
	}
	fmt.Println(s)
}

func node() *ListNode {
	return &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 5,
					},
				},
			},
		},
	}
}