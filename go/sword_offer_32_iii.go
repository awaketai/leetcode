package main

import "fmt"

// 剑指 Offer 32 - III. 从上到下打印二叉树 III

// 请实现一个函数按照之字形顺序打印二叉树，即第一行按照从左到右的顺序打印，第二层按照从右到左的顺序打印，第三行再按照从左到右的顺序打印，其他行以此类推。

//Definition for a binary tree node.
type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func main()  {
	ret := levelOrder(node())
	fmt.Println(ret)
}

// 利用双端队列的两端皆可添加元素的特性
// 设双端队列tmp
// 奇数层：添加至tmp尾部
// 偶数层：添加至tmp头部
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	var res = make([][]int,0)
	var queue = make([]*TreeNode,0)
	queue = append(queue,root)
	for len(queue) > 0 {
		tmp := make([]int,0)
		for i := len(queue);i > 0;i-- {
			node := queue[0]
			queue = queue[1:]
			// 如果是偶数层，添加到队列头部
			if len(res) % 2 == 0 {
				tmp = append(tmp,node.Val)
			}else{
				// 如果是奇数层，添加到队列尾部
				tmp = append([]int{node.Val},tmp...)
			}
			// 左右节点放入队列中
			if node.Left != nil {
				queue = append(queue,node.Left)
			}
			if node.Right != nil {
				queue = append(queue,node.Right)
			}
		}
		res = append(res,tmp)
	}
	return res
}


func node() *TreeNode {
	return &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 9,
		},
		Right: &TreeNode{
			Val: 20,
			Left: &TreeNode{
				Val: 15,
			},
			Right: &TreeNode{
				Val: 7,
			},
		},
	}
}
