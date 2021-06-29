package main

import "fmt"
import "leet_code/tree"

func main()  {
	node := listNode(testCase1)
	ret := invertTree(node)
	fmt.Println(ret.Left.Val)
}

func invertTree(root *tree.TreeNode) *tree.TreeNode {
	if root == nil {
		return nil
	}
	right := invertTree(root.Right)
	left := invertTree(root.Left)
	root.Right = left
	root.Left = right
	return root
}

var testCase1 = []int{4,2,7,1,3,6,9}

// 平衡二叉树
func listNode(arr []int) *tree.TreeNode {
	l := len(arr)
	node := tree.TreeNode{Val: arr[0]}
	for i := 1;i < l;i++ {
		subNode := tree.TreeNode{}
		subNode.Val = i
		if i % 2 == 0 {

			node.Left = &subNode
		}else{
			node.Right = &subNode
		}
	}
	return &node
}

