package main

import (
	"leet_code/tree"
)

func main()  {
	node := node2()
	convertBST(node)
}

// 538. Convert BST to Greater Tree
// Given the root of a Binary Search Tree (BST), convert it to a Greater Tree such that every key of the original BST is changed to the original key plus sum of all keys greater than the original key in BST.

func convertBST(root *tree.TreeNode) *tree.TreeNode {
	var sum = 0
	traverse2(root,&sum)
	return root
}

// 中序便来的顺序反过来，某个节点的值是大于等于其的所有节点值之和
func traverse2(root *tree.TreeNode,sum *int) *tree.TreeNode {
	if root == nil {
		return nil
	}
	// 将中序遍历的顺序反过来
	// 8,7,6,5,4,3,2,1,0
	// 最终每个节点的值为：本身+前面所有节点值 (因为之前所有节点是大于当前节点值的)
	traverse2(root.Right,sum)
	*sum += root.Val
	root.Val = *sum
	traverse2(root.Left,sum)
	return root
}

func node2() *tree.TreeNode {
	node := &tree.TreeNode{
		Val: 4,
		Left: &tree.TreeNode{
			Val: 1,
			Left: &tree.TreeNode{
				Val: 0,
			},
			Right: &tree.TreeNode{
				Val: 2,
				Right: &tree.TreeNode{
					Val: 3,
				},
			},
		},
		Right: &tree.TreeNode{
			Val: 6,
			Left: &tree.TreeNode{
				Val: 5,
			},
			Right: &tree.TreeNode{
				Val: 7,
				Right: &tree.TreeNode{
					Val: 8,
				},
			},
		},
	}
	return node
}
