package main

// 剑指 Offer 18. 删除链表的节点

// 给定单向链表的头指针和一个要删除的节点的值，定义一个函数删除该节点。
//
//返回删除后的链表的头节点。
//
//注意：此题对比原题有改动


  //Definition for singly-linked list.
type ListNode struct {
  Val int
  Next ListNode
}

func deleteNode(head *ListNode, val int) *ListNode {
	node := head
	for head != nil {
		if head.Val == val {

		}
	}
}
