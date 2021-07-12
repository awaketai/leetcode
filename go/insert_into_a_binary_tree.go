package main

import "leet_code/tree"

// 701. Insert into a Binary Search Tree

// You are given the root node of a binary search tree (BST) and a value to insert into the tree. Return the root node of the BST after the insertion. It is guaranteed that the new value does not exist in the original BST.
//
//Notice that there may exist multiple valid ways for the insertion, as long as the tree remains a BST after insertion. You can return any of them.

func main()  {

}

func insertIntoBST(root *tree.TreeNode, val int) *tree.TreeNode {
	if root == nil {
		return &tree.TreeNode{Val: val}
	}
	if root.Val == val {
		return root
	}
	if root.Val < val {
		// 去右侧
		root.Right = insertIntoBST(root.Right,val)
	}
	if root.Val > val {
		// 去左边
		root.Left = insertIntoBST(root.Left,val)
	}
	return root
}