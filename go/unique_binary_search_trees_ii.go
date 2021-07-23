package main

import (
	"fmt"
	"leet_code/tree"
)

// 95. Unique Binary Search Trees II
// Given an integer n, return all the structurally unique BST's (binary search trees), which has exactly n nodes of unique values from 1 to n. Return the answer in any order.

func main()  {
	generateTrees(8)
}

func generateTrees(n int) []*tree.TreeNode {
	res := build(1,n)
	fmt.Println(len(res))
	//fmt.Println(res)
	//for _,v := range res {
		//fmt.Printf("node :%v",v)
		//fmt.Println()
	//}
	return res
}

// 1.穷举root节点的所有可能
// 2.递归构造出左右子树的所有合法bst
// 3.给root节点穷举所有左右子树的组合

func build(low,high int) []*tree.TreeNode{
	if low > high {
		return []*tree.TreeNode{nil}
	}
	var node []*tree.TreeNode
	// 1.穷举root节点所有可能
	for i := low;i <= high;i++ {
		// 2.递归构造出所有左右子树的所有合法bst
		leftNode := build(low,i - 1)
		rightNode := build(i + 1,high)
		// 给root节点穷举所有左右子树结合
		for _,lv := range leftNode {
			for _,rv := range rightNode {
				root := &tree.TreeNode{Val: i} // 根节点
				root.Left = lv
				root.Right = rv
				//fmt.Println("root",root)
				node = append(node,root)
			}
		}
	}
	return node
}