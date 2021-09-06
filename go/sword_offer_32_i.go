package main

import (
	"fmt"
)

// 剑指 Offer 32 - I. 从上到下打印二叉树

// 从上到下打印出二叉树的每个节点，同一层的节点按照从左到右的顺序打印。


// 从上至下打印：二叉树的广度优先搜索(BFS)
// BFS通常借助队列的先入先出特性实现


  //Definition for a binary tree node.
type TreeNode struct {
  Val int
  Left *TreeNode
  Right *TreeNode
}

func main()  {
	ret := levelOrder(node())
	fmt.Println(ret)
}

func levelOrder(root *TreeNode) []int {
	var res = make([]int,0)
	if root == nil {
		return res
	}
	// 包含根节点的队列queue
	queue := make([]*TreeNode,0)
	queue = append(queue,root)
	for len(queue) > 0 {
		// 首元素出队
		node := queue[0]
		res = append(res,node.Val)
		queue = queue[1:]
		// 如果左节点有值，则将左节点加入队列
		if node.Left != nil {
			queue = append(queue,node.Left)
		}
		// 如果右节点有值，则将有节点加入队列
		if node.Right != nil {
			queue = append(queue,node.Right)
		}
	}
	return res
}

func node() *TreeNode {
	return &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 9,
		},
		Right: &TreeNode{
			Val: 20,
			Left: &TreeNode{
				Val: 15,
			},
			Right: &TreeNode{
				Val: 7,
			},
		},
	}
}