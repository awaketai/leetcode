package main

import (
	"fmt"
	"math"
	"strconv"
)

// 剑指 Offer 67. 把字符串转换成整数

// 写一个函数 StrToInt，实现把字符串转换成整数这个功能。不能使用 atoi 或者其他类似的库函数。
//
//
//
//首先，该函数会根据需要丢弃无用的开头空格字符，直到寻找到第一个非空格的字符为止。
//
//当我们寻找到的第一个非空字符为正或者负号时，则将该符号与之后面尽可能多的连续数字组合起来，作为该整数的正负号；假如第一个非空字符是数字，则直接将其与之后连续的数字字符组合起来，形成整数。
//
//该字符串除了有效的整数部分之后也可能会存在多余的字符，这些字符可以被忽略，它们对于函数不应该造成影响。
//
//注意：假如该字符串中的第一个非空格字符不是一个有效整数字符、字符串为空或字符串仅包含空白字符时，则你的函数不需要进行转换。
//
//在任何情况下，若函数不能进行有效的转换时，请返回 0。
//
//说明：
//
//假设我们的环境只能存储 32 位大小的有符号整数，那么其数值范围为 [−231,  231 − 1]。如果数值超过这个范围，请返回  INT_MAX (231 − 1) 或 INT_MIN (−231) 。

// 有限状态自动机
// 自动机：
// 一共枚举有0，1，2，3四种状态
// 0:起始状态
// 1:符号状态(+/-)
// 2:数字状态
// 3:终止状态

// 转移条件：4种
// 空格  +/-  数字位 其他
func main()  {
	//str := "42"
	//str := "  -42"
	//str := "4193 with words"
	//str := "words and 987"
	str := "-91283472332"
	ret := strToInt(str)
	fmt.Println(ret)
}

func strToInt(str string) int {
	states := []map[rune]int{
		{' ':0,'x':3,'s':1,'d':2}, //0
		{'d':2,'x':3,' ':3,'s':3}, //1
		{'d':2,'x':3,' ':3,'s':3}, //2
		{'d':3,'x':3,' ':3,'s':3}, //3
	}

	state,num,sign := 0,0,1 // 当前状态，记录数，符号
	var t rune // 当前的字符类型
	for _,c := range str {
		if c == ' ' {
			// 空格
			t = ' '
		}else if c == '+' || c == '-' {
			// 符号
			t = 's'
		}else if c >= '0' && c <= '9' {
			// 数字
			t = 'd'
			sInt,_ := strconv.Atoi(string(c))
			num = num * 10 + sInt
			fmt.Println("num-1:",num)
			if num >= math.MaxInt32 && sign == 1 {
				return math.MaxInt32
			}
			if num >= math.MaxInt32 && sign == -1 {
				return math.MinInt32
			}
		}else{
			// 终止态
			t = 'x'
		}

		state = states[state][t]
		if state == 3 {
			// 结束
			break
		}

		if c == '-'{
			sign = -1
		}
	}
	return sign * num
}