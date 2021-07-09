package main

import (
	"fmt"
	"leet_code/tree"
)

// 450. Delete Node in a BST
// Given a root node reference of a BST and a key, delete the node with the given key in the BST. Return the root node reference (possibly updated) of the BST.
//
//Basically, the deletion can be divided into two stages:
//
//Search for a node to remove.
//If the node is found, delete the node.

func main()  {
	node := node3()
	root := deleteNode(node,3)
	fmt.Println(root)
}

// 三种情况
// 1.节点A恰好是末端节点，两个子节点都为空，可以直接删除
// 2.节点A只有一个非空子节点，那么要让它的孩子接替自己的位置
// 3.节点A有两个子节点，不能破坏bst的性质，A必须找到左子树中最大的那个节点，或者右子树中最小的那个节点来接替自己

func deleteNode(root *tree.TreeNode, key int) *tree.TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == key {
		// 1.节点是末端节点，直接删除
		// 2.节点只有一个非空子节点、则让自己的孩子接替自己位置
		if root.Left == nil {
			return root.Right
		}
		if root.Right == nil {
			return root.Left
		}
		// 3.节点有两个子节点、找到右子树的最小那个节点来接替自己
		minNode := getMin(root.Right)
		root.Val = minNode.Val
		// 将右子节点置为nil
		root.Right = deleteNode(root.Right,minNode.Val)
	} else if root.Val > key {
		// 去左边找
		root.Left = deleteNode(root.Left,key)
	} else if root.Val < key {
		// 去右边找
		root.Right = deleteNode(root.Right,key)
	}
	return root
}

// 获取bst最小节点值：最左边就是最小值
func getMin(root *tree.TreeNode) *tree.TreeNode {
	 for root.Left != nil {
	 	root = root.Left
	 }
	 return root
}

func node3() *tree.TreeNode {
	node := &tree.TreeNode{
		Val: 5,
		Left: &tree.TreeNode{
			Val: 3,
			Left: &tree.TreeNode{
				Val: 2,
			},
			Right: &tree.TreeNode{
				Val: 4,
			},
		},
		Right: &tree.TreeNode{
			Val: 6,
			Right: &tree.TreeNode{
				Val: 7,
			},
		},
	}
	return node
}