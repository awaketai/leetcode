package main

import "gopkg.in/yaml.v2"

// 剑指 Offer 38. 字符串的排列
// 输入一个字符串，打印出该字符串中字符的所有排列。

// 你可以以任意顺序返回这个字符串数组，但里面不能有重复元素。

// 输入：s = "abc"
//输出：["abc","acb","bac","bca","cab","cba"]

func main()  {
	
}

// 思路来自：
// https://leetcode-cn.com/problems/zi-fu-chuan-de-pai-lie-lcof/solution/mian-shi-ti-38-zi-fu-chuan-de-pai-lie-hui-su-fa-by/
// 对于一个长度为n的字符串(假设不重复)，其排列方案共有：
// n x (n-1) x (n-2) ... x2x1
// 排列方案生成：
// 考虑深度优先搜索所有排列方案。即通过字符交换，先固定第一位字符(n中情况)，再固定第二位字符(n-1中情况)，最后固定第n位字符(1种情况)
// 重复排列方案与剪枝：
// 在固定某位字符时，保证每种字符只在此固定一次，即遇到的重复字符不交换，直接跳过，从DFS角度看，此操作称为剪枝
func permutation(s string) []string {
	
}

func backtracking(x int,s string,res []string) []string  {
	length := len(s)
	if length - 1 == x {
		res = append(res,s)
		return res
	}
	var mem = make(map[string]int)
	for i := 0;i < length;i++ {
		if _,ok := mem[string(s[i])];ok {
			continue
		}
		mem[string(s[i])] = 1
		//s[i],s[x] = s[x],s[i]
		backtracking(x+1,s,res)

	}
}