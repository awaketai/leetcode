package main

import (
	"fmt"
	"leet_code/tree"
)

// 剑指 Offer 07. 重建二叉树

// 输入某二叉树的前序遍历和中序遍历的结果，请构建该二叉树并返回其根节点。
//
//假设输入的前序遍历和中序遍历的结果中都不含重复的数字。


/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

var preorder5 = []int{3,9,20,15,7}
var inorder5 = []int{9,3,15,20,7}
func main()  {
	ret := buildTree5(preorder5,inorder5)
	fmt.Println(ret)
}

// 1.前序遍历第一个元素为根节点值
// 2.前序遍历左右节点长度和中序遍历左右节点长度相同
// 3.通过前序遍历根节点确定中序遍历中根节点位置
func buildTree5(preorder []int, inorder []int) *tree.TreeNode {
	if  len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}
	rootIndex := 0
	length := len(inorder)
	for i := 0;i < length;i++ {
		if inorder[i] == preorder[0] {
			rootIndex = i
			break
		}
	}
	node := &tree.TreeNode{
		Val: preorder[0],
	}
	node.Left = buildTree5(preorder[1:rootIndex + 1],inorder[:rootIndex])
	node.Right = buildTree5(preorder[rootIndex + 1:],inorder[rootIndex + 1:])
	return node
}