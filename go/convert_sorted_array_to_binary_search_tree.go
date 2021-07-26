package main

import (
	"leet_code/tree"
	"log"
)

// 108.Convert Sorted Array to Binary Search Tree

// Given an integer array nums where the elements are sorted in ascending order, convert it to a height-balanced binary search tree.
//
//A height-balanced binary tree is a binary tree in which the depth of the two subtrees of every node never differs by more than one.


// 二叉搜索树中序遍历是升序序列，题目给出的是升序序列，是中序遍历结果
// 高度平衡的二叉搜索树，可以选择中间数据作为根节点，这样分给左右子树的数字个数相同或相差1
// 使得保持树平衡


var arr = []int{-10,-3,0,5,9}

var arr1 = []int{1,3}
func main()  {
	testSortedArrayToBST(arr1)
	//sortedArrayToBST(arr)
}


func sortedArrayToBST(nums []int) *tree.TreeNode {

	return helper(nums,0,len(nums) - 1)
}

func helper(nums []int,left,right int) *tree.TreeNode {
	if left > right {
		return nil
	}
	// 选取中位数
	mid :=  (left + right) / 2
	root := &tree.TreeNode{Val: nums[mid]}
	root.Left = helper(nums,left,mid - 1)
	root.Right = helper(nums,mid + 1,right)
	return root
}


func testSortedArrayToBST(nums []int) {
	root := sortedArrayToBST(nums)
	// 中序遍历应该是有序的
	mem := make([]int,0)
	ret := testTraverse6(root,&mem)
	// 和原数组元素是否相等
	for i,v := range *ret {
		if nums[i] != v {
			log.Printf("test error,element value :%d,target value:%d",v,nums[i])
			break
		}
	}
	log.Println("test success")
}


func testTraverse6(root *tree.TreeNode,ret *[]int) *[]int {
	if root == nil {
		return nil
	}
	testTraverse6(root.Left,ret)
	*ret = append(*ret,root.Val)
	testTraverse6(root.Right,ret)
	return ret
}