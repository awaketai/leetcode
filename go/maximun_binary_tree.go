package main

import (
	"leet_code/tree"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

var nums = []int{3,2,1,6,0,5}

func main()  {
	constructMaximumBinaryTree(nums)
}

// 给定一个不含重复元素的整数数组，一个以此数组构建的最大二叉树定义如下
// 1.二叉树的根是数组中的最大元素
// 2.左子树是通过数组中最大值左边部分构造出的最大二叉树
// 3.右子树是通过数组中最大值右边部分构造出的最大二叉树
func constructMaximumBinaryTree(nums []int) *tree.TreeNode {
	// 1.找出最大值
	// 2.对最大值最边和右边分别找最大值
	length := len(nums)
	if length <= 0 {
		return nil
	}
	// 找出最大值
	var max,index = 0,0
	for i := 0;i < length;i++ {
		if nums[i] > max {
			max = nums[i]
			index = i
		}
	}
	node := &tree.TreeNode{
		Val: max,
	}
	// 构造左右子树
	left := nums[:index]
	right := nums[index+1:]
	node.Left = constructMaximumBinaryTree(left)
	node.Right = constructMaximumBinaryTree(right)

	return node
}


