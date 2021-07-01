package main

import (
	"fmt"
	"leet_code/tree"
)

// 106. Construct Binary Tree from Inorder and Postorder Traversal
//Given two integer arrays inorder and postorder where inorder is the inorder traversal of a binary tree and postorder is the postorder traversal of the same tree, construct and return the binary tree.

// 中序遍历：左 -> 根 -> 右
// 后序遍历：左 -> 右 -> 根

var inorder1 = []int{9,3,15,20,7}
var inorder2 = []int{5,2,6,4,7,1,8,3,9}
var postorder1 = []int{9,15,7,20,3}
var postorder2 = []int{5,6,7,4,2,8,9,3,1}



// [5,6,7,4,2,8,9,3,1]
// [5,2,6,4,7,1,8,3,9]
func main(){
	fmt.Println(buildTree1(inorder2,postorder2))
}

func buildTree1(inorder []int, postorder []int) *tree.TreeNode {
	if len(inorder) == 0 || len(postorder) == 0 {
		return nil
	}
	// 确定根节点值：后序遍历最后值
	rootVal := postorder[len(postorder) - 1]
	// 中序遍历和后序遍历，左子树和右子树长度相同
	// 确定跟节点在中序遍历中位置
	rootIndex := 0
	for i := 0;i < len(inorder);i++ {
		if rootVal == inorder[i] {
			rootIndex = i
			break
		}
	}
	node := &tree.TreeNode{Val: rootVal}

	leftLen := len(inorder[:rootIndex])
	node.Left = buildTree1(inorder[:rootIndex],postorder[:leftLen])
	// 总长度-左子树长度-根节点长度
	node.Right = buildTree1(inorder[rootIndex+1:],postorder[leftLen:(len(postorder) - 1)])
	return node
}

