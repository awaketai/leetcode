package main

// 剑指 Offer 28. 对称的二叉树

// 请实现一个函数，用来判断一棵二叉树是不是对称的。如果一棵二叉树和它的镜像一样，那么它是对称的。
//
//例如，二叉树 [1,2,2,3,4,4,3] 是对称的。
//
//    1
//   / \
//  2   2
// / \ / \
//3  4 4  3
//但是下面这个 [1,2,2,null,3,null,3] 则不是镜像对称的:
//
//    1
//   / \
//  2   2
//   \   \
//   3    3
//


//Definition for a binary tree node.
type TreeNode struct {
  Val int
  Left *TreeNode
  Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {

	if root == nil {
		return true
	}
	return traverse(root.Left,root.Right)
}

// 节点为空的情况：
// 1.左空、右不为空，不对称
// 2.右空、左不为空，不对称
// 3.左右为空，对称
// 4.左右不为空，判断节点值是否相等

func traverse(left,right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if left == nil || right == nil {
		return false
	}
	if left.Val != right.Val {
		return false
	}
	return traverse(left.Left,right.Right) && traverse(left.Right,right.Left)
}


