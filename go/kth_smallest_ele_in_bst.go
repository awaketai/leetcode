package main

import (
	"leet_code/tree"
)

// 230. Kth Smallest Element in a BST
// Given the root of a binary search tree, and an integer k, return the kth (1-indexed) smallest element in the tree.
// Input: root = [3,1,4,null,2], k = 1
// Output: 1
// 230. 二叉搜索树中第K小的元素
// 给定一个二叉搜索树的根节点 root ，和一个整数 k ，请你设计一个算法查找其中第 k 个最小元素（从 1 开始计数）。
func main()  {
	node := node()
	kthSmallest(&node,5)
}

// 1.二叉搜索树中序遍历是升序，升序排序，然后找第k个元素
// 时间复杂度 O(N)
func kthSmallest(root *tree.TreeNode, k int) int {
	var ret,rank int
	traverse1(root,k,&ret,&rank)
	return ret
}

func traverse1(root *tree.TreeNode,k int,ret,rank *int)  {
	if root == nil {
		return
	}
	traverse1(root.Left,k,ret,rank)
	*rank++
	if k == *rank {
		*ret = root.Val
		return
	}
	traverse1(root.Right,k,ret,rank)
	return
}

func node() tree.TreeNode {
	node := tree.TreeNode{
		Val: 5,
		Left: &tree.TreeNode{
			Val: 3,
			Left: &tree.TreeNode{
				Val: 2,
				Left: &tree.TreeNode{
					Val: 1,
				},
			},
			Right: &tree.TreeNode{
				Val: 4,
			},
		},
		Right: &tree.TreeNode{
			Val: 6,
		},
	}
	return node
}