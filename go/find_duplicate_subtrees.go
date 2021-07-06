package main

import (
	"fmt"
	"leet_code/tree"
	"strconv"
)

// 652. Find Duplicate Subtrees
// Given the root of a binary tree, return all duplicate subtrees.
//
//For each kind of duplicate subtrees, you only need to return the root node of any one of them.
//
//Two trees are duplicate if they have the same structure with the same node values.

// 给定一棵二叉树，返回所有重复的子树。对于同一类的重复子树，你只需要返回其中任意一棵的根结点即可。
//
// 两棵树重复是指它们具有相同的结构以及相同的结点值。

// Input: root = [1,2,3,4,null,2,4,null,null,4]
// Output: [[2,4],[4]]

// 1.以我为根节点的这颗二叉树是什么样
// 2.以其他节点为根的子树是什么样
// 使用后序遍历

//var mem = make(map[string]int)
//var res []*tree.TreeNode

func main()  {
	node := rootNode()
	findDuplicateSubtrees(node)
}

func findDuplicateSubtrees(root *tree.TreeNode) []*tree.TreeNode {
	var mem = make(map[string]int)
	var res []*tree.TreeNode
	traverse(root,mem,&res)
	return res
}

func traverse(node *tree.TreeNode,mem map[string]int,res *[]*tree.TreeNode) string {
	// 如果是空节点，使用#号表示
	if node == nil {
		return "#"
	}
	// 将左右子树序列化成字符串
	left := traverse(node.Left,mem,res)
	right := traverse(node.Right,mem,res)
	// 后序遍历，用,号分割每个二叉树节点值
	// 为了描述二叉树样子，什么排序不重要
	// 1.以我为根节点的二叉树长什么样？
	subTree := left + "," + right + "," + strconv.Itoa(node.Val)
	// 2.以其他节点为根的子树是什么样？
	freq := mem[subTree]
	// 只增加第一次值到
	if freq == 1 {
		fmt.Println("node:",node)
		*res = append(*res,node)
	}
	mem[subTree] = freq + 1
	return subTree
}

func rootNode() *tree.TreeNode {
	root := &tree.TreeNode{
		Val: 1,
		Left: &tree.TreeNode{
			Val: 2,
			Left: &tree.TreeNode{
				Val: 4,
			},
		},
		Right: &tree.TreeNode{
			Val: 3,
			Left: &tree.TreeNode{
				Val: 2,
				Left: &tree.TreeNode{
					Val: 4,
				},
			},
			Right: &tree.TreeNode{
				Val: 4,
			},
		},
	}
	return root
}