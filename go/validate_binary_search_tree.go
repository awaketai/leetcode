package main

import (
	"fmt"
	"leet_code/tree"
)

// 98. Validate Binary Search Tree
// Given the root of a binary tree, determine if it is a valid binary search tree (BST).
//
//A valid BST is defined as follows:
//
//The left subtree of a node contains only nodes with keys less than the node's key.
//The right subtree of a node contains only nodes with keys greater than the node's key.
//Both the left and right subtrees must also be binary search trees.

// 每个节点值，root的整个左子树都要小于root.Val，整个右子树都要大于root.Val
func isValidBST(root *tree.TreeNode) bool {
	if root == nil {
		return true
	}
	return traverse3(root,nil,nil)
}

// 1.每个节点最大值max，min：每个节点最小值
// 有问题，待修改...
func traverse3(root,min,max *tree.TreeNode) bool {
	if root == nil {
		return true
	}
	if min != nil && root.Val <= min.Val {
		return false
	}
	if max != nil && root.Val >= max.Val {
		return false
	}
	return traverse3(root.Left,min,root) && traverse3(root.Right,root,max)
}

// 2.使用中序遍历：二叉树中序遍历是升序的
// 需要优化...
func inTraverse(root *tree.TreeNode,pre *[]int) bool {
	if root == nil {
		return true
	}
	if !inTraverse(root.Left,pre) {
		return false
	}
	// 当前值是否大于前面所有值
	for _,v := range *pre {
		if root.Val <= v {
			return false
		}
	}
	*pre = append(*pre,root.Val)
	return inTraverse(root.Right,pre)
}

func main() {
	node := node6()
	pre := make([]int,0)
	//var pre int
	ret := inTraverse(node,&pre)
	//fmt.Println(qu)
	fmt.Println(ret)
}

func node4() *tree.TreeNode {
	node := &tree.TreeNode{
		Val: 10,
		Left: &tree.TreeNode{
			Val: 5,
		},
		Right: &tree.TreeNode{
			Val: 15,
			Left: &tree.TreeNode{
				Val: 6,
			},
			Right: &tree.TreeNode{
				Val: 20,
			},
		},
	}
	return node
}

func node5() *tree.TreeNode {
	node := &tree.TreeNode{
		Val: 4,
		Left: &tree.TreeNode{
			Val: 2,
			Left: &tree.TreeNode{
				Val: 1,
			},
			Right: &tree.TreeNode{
				Val: 3,
			},
		},
		Right: &tree.TreeNode{
			Val: 5,
			Right: &tree.TreeNode{
				Val: 6,
			},
		},
	}
	return node
}

func node6() *tree.TreeNode {
	return &tree.TreeNode{
		Val: 2,
		Left: &tree.TreeNode{
			Val: 2,
		},
		Right: &tree.TreeNode{
			Val: 2,
		},
	}
}