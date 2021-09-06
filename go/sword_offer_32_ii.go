package main

import (
	"fmt"
)

// 剑指 Offer 32 - II. 从上到下打印二叉树 II

// 从上到下按层打印二叉树，同一层的节点按从左到右的顺序打印，每一层打印到一行。


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

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	var res = make([][]int,0)
	// 将节点放入队列中
	queue := make([]*TreeNode,0)
	queue = append(queue,root)
	for len(queue) > 0 {
		var tmp = make([]int,0)
		// 为了这行加入队列
		for i := len(queue);i > 0;i-- {
			node := queue[0]
			queue = queue[1:]
			tmp = append(tmp,node.Val)
			// 左右节点放入队列中
			if node.Left != nil {
				queue = append(queue,node.Left)
			}
			if node.Right != nil {
				queue = append(queue,node.Right)
			}
		}
		res = append(res,tmp)
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
