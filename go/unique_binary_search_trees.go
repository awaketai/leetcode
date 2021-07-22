package main

import (
	"fmt"
	"strconv"
)

// 96. Unique Binary Search Trees
// Given an integer n, return the number of structurally unique BST's (binary search trees) which has exactly n nodes of unique values from 1 to n.

func main()  {
	numTrees(19)
}

func numTrees(n int) int {
	//mem := make(map[string]int)
	//res := count(1,n,mem)
	//fmt.Println(res,mem)
	//return res
	res := cal(n)
	fmt.Println(res)
	return res
}

// 
func cal(n int) int {
	var arr = make([]int,n+1)
	arr[0],arr[1] = 1,1
	for i := 2;i <= n;i++ {
		for j := 1;j <=i;j++ {
			arr[i] += arr[j - 1] * arr[i - j]
		}
	}
	return arr[n]
}

// 递归执行时间太长
func count(lo,hi int,mem map[string]int) int {
	if lo > hi {
		return 1
	}
	key := strconv.Itoa(lo) + "-" + strconv.Itoa(hi)
	if mem[key] != 0 {
		return mem[key]
	}
	res := 0
	// 以每个元素为根节点
	for i := lo;i <= hi;i++ {
		left := count(lo,i - 1,mem)
		right := count(i + 1,hi,mem)
		res += left * right
	}
	mem[key] = res
	return res
}