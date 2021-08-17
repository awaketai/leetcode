package main

import (
	"golang.org/x/text/unicode/bidi"
	"google.golang.org/protobuf/reflect/protoreflect"
	"image"
)

// 剑指 Offer 27. 二叉树的镜像

// 请完成一个函数，输入一个二叉树，该函数输出它的镜像。
//
//例如输入：
//
//     4
//   /   \
//  2     7
// / \   / \
//1   3 6   9
//镜像输出：
//
//     4
//   /   \
//  7     2
// / \   / \
//9   6 3   1
//


//Definition for a binary tree node.
type TreeNode struct {
  Val int
  Left *TreeNode
  Right *TreeNode
}

// 递归的对树进行便来，从叶子节点开始翻转得到镜像
// 如果当前遍历到的节点root的左右子节点都已经翻转得到镜像
// 那么只需要交换两颗子树的位置，既可以得到以root为根节点的整颗子树的镜像

func mirrorTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	left := mirrorTree(root.Left)
	right := mirrorTree(root.Right)
	root.Left = right
	root.Right = left
	return root
}