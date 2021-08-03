package main

import "fmt"

// 600. Non-negative Integers without Consecutive Ones

// Given a positive integer n, return the number of the integers in the range [0, n] whose binary representations do not contain consecutive ones.

func main()  {
	findIntegers(9)
}

// 这种方法数字大，执行时间超时
func findIntegers(n int) int {
	var res = n + 1
	// 将0-n数字二进制转换
	for i := 0;i <= n;i++ {
		binArr := DecToBin(i)
		fmt.Println(binArr)
		length := len(binArr)
		loop:
		for j := 0; j < length;j++ {
			if j == 0 || j == length {
				continue
			}
			if binArr[j] == binArr[j - 1] && binArr[j] == 1 {
				res--
				break loop
			}
		}
	}
	fmt.Println(res)
	return res
}

// 原理：除2取模是最低位
func DecToBin(n int) []int {
	var res []int
	if n == 0 {
		return []int{0}
	}
	for ; n > 0;n /= 2 {
		lsb := n % 2
		res = append(res,lsb)
	}
	return res
}
