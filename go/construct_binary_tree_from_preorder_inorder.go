package main

import (
	"fmt"
	"leet_code/tree"
)

var preorder = []int{3,9,20,15,7}
var inorder = []int{9,3,15,20,7}

// 105. Construct Binary Tree from Preorder and Inorder Traversal
// Given two integer arrays preorder and inorder where preorder is the preorder traversal of a binary tree and inorder is the inorder traversal of the same tree, construct and return the binary tree.
// 根据一棵树的前序遍历与中序遍历构造二叉树。
func main()  {
	ret := buildTree(preorder,inorder)
	fmt.Println(ret)
}

// 根据一棵树的前序遍历与中序遍历构造二叉树
// 注意：可以假设树中没有重复的元素
// 例如，给出
// 前序遍历 preorder = [3,9,20,15,7] 根 -> 左 -> 右
// 中序遍历 inorder = [9,3,15,20,7] 左 -> 根 -> 右
func buildTree(preorder []int, inorder []int) *tree.TreeNode {

	// 在中序遍历中获取到根节点的位置，那么就可以确定左子树和右子树的节点数目
	// 同一颗子树的前序遍历和中序遍历长度是相同的
	if len(preorder) == 0 || len(inorder) == 0{
		return nil
	}
	// 根节点：前序遍历第0元素
	rootVal := preorder[0]
	// 确定中序遍历中根节点位置
	rootIndex := 0
	for i := 0;i < len(inorder);i++ {
		if inorder[i] == rootVal {
			rootIndex = i
			break
		}
	}
	// 以中序遍历结果确定，左子树和右子树节点数据
	node := &tree.TreeNode{Val: rootVal}
	// 确定左右子树长度
	leftLen := len(inorder[:rootIndex])
	//rightLen := len(inorder[rootIndex:])

	node.Left = buildTree(preorder[1:leftLen+1],inorder[:rootIndex])
	node.Right = buildTree(preorder[(1 + leftLen):],inorder[rootIndex + 1:])
	return node
}
