
package main

import (
	"fmt"
)

// 剑指 Offer 26. 树的子结构

// 输入两棵二叉树A和B，判断B是不是A的子结构。(约定空树不是任意一个树的子结构)
//
//B是A的子结构， 即 A中有出现和B相同的结构和节点值。


  //Definition for a binary tree node.
type TreeNode struct {
  Val int
  Left *TreeNode
  Right *TreeNode
}

func main()  {
	ret := isSubStructure(node11(),node22())
	fmt.Println(ret)
}
// 若树B是树A的子集，则满足以下情况
// 1.以A节点为根的子树包含B节点
// 2.树B是A树左子树的子结构
// 3.树B是A树右子树的子结构


// 遍历A中每个节点root，先判断A中以节点root为根节点的树是否包含树B
// 终止条件
// 1.节点B已经遍历完，说明树B已经匹配成功
// 2.节点A遍历完成，匹配失败
// 3.节点A和节点B值不同，匹配失败
func isSubStructure(A *TreeNode, B *TreeNode) bool {

	if A == nil || B == nil {
		return false
	}
	return traverse(A,B) ||
		isSubStructure(A.Left,B) ||
		isSubStructure(A.Right,B)
}

func traverse(A *TreeNode,B *TreeNode) bool {
	if B == nil {
		// B已经访问完，匹配完成
		return true
	}
	// 如果A树已经遍历完，未完成匹配
	if A == nil {
		return false
	}
	// 如果A节点值和B节点值不同，未完成匹配
	if A.Val != B.Val {
		return false
	}
	return traverse(A.Left,B.Left) && traverse(A.Right,B.Right)
}

func node11() *TreeNode {
	return &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 4,
			Left: &TreeNode{
				Val: 1,
			},
			Right: &TreeNode{
				Val: 2,
			},
		},
		Right: &TreeNode{
			Val: 5,
		},
	}
}

func node22() *TreeNode {
	return &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val: 1,
		},
	}
}