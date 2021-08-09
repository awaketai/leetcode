package main

import (
	"fmt"
	"time"
)

// 剑指 Offer 10- I. 斐波那契数列

// 写一个函数，输入 n ，求斐波那契（Fibonacci）数列的第 n 项（即 F(N)）。斐波那契数列的定义如下：
// F(0) = 0,   F(1) = 1
// F(N) = F(N - 1) + F(N - 2), 其中 N > 1.

//斐波那契数列由 0 和 1 开始，之后的斐波那契数就是由之前的两数相加而得出。
//
//答案需要取模 1e9+7（1000000007），如计算初始结果为：1000000008，请返回 1。

// DP(Dynamic Programming)动态规划
//
func main()  {
	start := time.Now()
	ret := fib(100)
	fmt.Println(ret)
	cost := time.Since(start)
	fmt.Printf("cost=[%s]",cost)
}

// 1.递归大量重复计算会超时
// 2.考虑int范围
// 回溯：通过枚举的方式得到n的斐波那契数列
// 转化为二叉树的遍历，使用深度优先遍历方法
//                 6
//         5                 4
//     4       3         3       2
//   3      2     2   2     1     1  1   0
// 2   1  1  0  1  0
func fib(n int) int {
	if n <=1 {
		return n
	}
	var mem = make(map[int]int)
	return dfs2(n,mem)
}

// 递归在于时间复杂度太高，因此借助于外部存储，使用空间换取时间
func dfs2(n int,mem map[int]int) int {
	if n <= 1 {
		return n
	}
	if v,ok := mem[n];ok {
		return v
	}
	leftVal := dfs2(n-1,mem)
	rightVal := dfs2(n-2,mem)
	val := (leftVal + rightVal) % 1000000007
	// 将已有的值放入map中，避免重复计算
	mem[n] = val
	return val
}

// 自底而上方法dp,以f(0),f(1)往上推导
func dp(n int) int {

	if n <= 1{
		return n
	}
	var mem = make(map[int]int)
	// 初始化0，1斐波那契值
	mem[0] = 0
	mem[1] = 1
	// 自底向上推导，斐波那契数列的后一个值为前两个数值之和
	for i := 2;i <= n;i++ {
		mem[i] = (mem[i - 1] + mem[i - 2]) % 1000000007
	}
	// 反推导后，数组最后一个值就是当前数字的斐波那契数列
	return mem[n]
}