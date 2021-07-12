package main

import "leet_code/tree"

// 700. Search in a Binary Search Tree
// You are given the root of a binary search tree (BST) and an integer val.

//   Find the node in the BST that the node's value equals val and return the subtree rooted with that node.
//  If such a node does not exist, return null.
func main()  {

}

func searchBST(root *tree.TreeNode, val int) *tree.TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == val {
		return root
	}
	// 二叉树搜索树左子树节点值小于右子树节点值
	if root.Val > val {
		return searchBST(root.Left,val)
	}else{
		return searchBST(root.Right,val)
	}
}
