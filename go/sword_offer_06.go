package main

import "fmt"

// 剑指 Offer 06. 从尾到头打印链表
//输入一个链表的头节点，从尾到头反过来返回每个节点的值（用数组返回）。

// 输入：head = [1,3,2]
// 输出：[2,3,1]



  //Definition for singly-linked list.
  type ListNode struct {
      Val int
      Next *ListNode
  }


func main()  {
	node := node7()
	ret := reversePrint(node)
	fmt.Println(ret)
}

// 先将链表的值压入数组中
// 然后将数组的值反向输出
func reversePrint(head *ListNode) []int {
	var ret = make([]int,0)
	for head != nil {
		ret = append(ret,head.Val)
		head = head.Next
	}
	res := make([]int,0)
	length := len(ret)
	for i := length-1;i >= 0;i-- {
		res = append(res,ret[i])
	}
	return res
}

func node7() *ListNode {
	node := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val: 2,
			},
		},
	}
	return node
}
