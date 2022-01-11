package main

// 剑指 Offer 64. 求1+2+…+n
// 求 1+2+...+n ，要求不能使用乘除法、for、while、if、else、switch、case等关键字及条件判断语句（A?B:C）。

func sumNums(n int) int {
	res := 0
	sum(n,&res)
	return res
}

func sum(n int,res *int) bool {
	*res += n
	return (n > 1) && sum(n - 1,res)
}
