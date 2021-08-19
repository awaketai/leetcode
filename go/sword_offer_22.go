package main

// 剑指 Offer 22. 链表中倒数第k个节点

// 输入一个链表，输出该链表中倒数第k个节点。为了符合大多数人的习惯，本题从1开始计数，即链表的尾节点是倒数第1个节点。
//
//例如，一个链表有 6 个节点，从头节点开始，它们的值依次是 1、2、3、4、5、6。这个链表的倒数第 3 个节点是值为 4 的节点。
//

  //Definition for singly-linked list.
  type ListNode struct {
      Val int
      Next *ListNode
  }

// 双指针法：
// 1.初始化前后指针，都指向head
// 2.前指针先走k步
// 3.前后指针都向前移动，知道前指针先跳出，后指针就是需要的链表
func getKthFromEnd(head *ListNode, k int) *ListNode {
	advance,latter := head,head
	// 前指针先移动k步
	for i := 1;i <= k;i++ {
		advance = advance.Next
	}
	// 前后指针共同移动
	for advance != nil {
		advance = advance.Next
		latter = latter.Next
	}
	return latter
}
