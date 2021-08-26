package main

import (
	"fmt"
	"math"
)

// 剑指 Offer 55 - I. 二叉树的深度

// 输入一棵二叉树的根节点，求该树的深度。从根节点到叶节点依次经过的节点（含根、叶节点）形成树的一条路径，最长路径的长度为树的深度。


  //Definition for a binary tree node.
  type TreeNode struct {
      Val int
      Left *TreeNode
      Right *TreeNode
  }

func main()  {
	ret := maxDepth(node())
	fmt.Println(ret)
}
  // 树深度=max(左子树深度，右子树深度)+1
  // 后序遍历
func maxDepth(root *TreeNode) int {

	if root == nil {
		return 0
	}
	return int(math.Max(float64(maxDepth(root.Left)),float64(maxDepth(root.Right)))) + 1
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