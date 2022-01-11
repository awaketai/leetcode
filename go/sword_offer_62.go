package main

import "fmt"
// 剑指 Offer 62. 圆圈中最后剩下的数字
// 0,1,···,n-1这n个数字排成一个圆圈，从数字0开始，每次从这个圆圈里删除第m个数字（删除后从下一个数字开始计数）。求出这个圆圈里剩下的最后一个数字。
//
//例如，0、1、2、3、4这5个数字组成一个圆圈，从数字0开始每次删除第3个数字，则删除的前4个数字依次是2、0、4、1，因此最后剩下的数字是3。
func main()  {
	lastRemaining(5,3)
}

// 取巧，剩余数的起始索引位置正好和数字相等
func lastRemaining(n,m int) int {
	finally := 0
	for i := 2;i <= n;i++ {
		finally = (finally + m) % i
		fmt.Println("finally:",finally)
	}
	return finally
}