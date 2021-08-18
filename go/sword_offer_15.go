package main

// 剑指 Offer 15. 二进制中1的个数
// 编写一个函数，输入是一个无符号整数（以二进制串的形式），返回其二进制表达式中数字位数为 '1' 的个数（也被称为 汉明重量).）。

//     
//
//    提示：
//
//    请注意，在某些语言（如 Java）中，没有无符号整数类型。在这种情况下，输入和输出都将被指定为有符号整数类型，
// 并且不应影响您的实现，因为无论整数是有符号的还是无符号的，其内部的二进制表示形式都是相同的。
//    在 Java 中，编译器使用 二进制补码 记法来表示有符号整数。因此，在上面的 示例 3 中，输入表示有符号整数 -3。

import (
	"fmt"
)

// 根据与运算定义，设二进制数字n，则有
// 若n&1 = 0 则n最右一位为0
// 若n&1 = 1 则n最右一位为1

var u1 = uint32(00000000000000000000000000001011)
var u2 = uint32(00000000000000000000000010000000)
//var u3 = uint32(11111111111111111111111111111101)
func main()  {
	ret := hammingWeight2(u2)
	fmt.Println(ret)
}
// 方法一：
// 根据以上特点，做以下循环判断
// 1.判断n最右一位是否为1，根据结果计数
// 2.将n右移一位(无符号右移)
func hammingWeight(num uint32) int {
	var res = 0
	for num != 0 {
		if (num & 1) == 1 {
			res++
		}
		num >>= 1
	}
	return res
}

// 10101000 2^3+2^5+2^7 168
// 10101000 2^0+2^1+2^2+2^5+2^7 167
// 10101000 & 10101000 (n&(n-1)) -> 10100000

// 0011 - 1 -> 0010 2
// 0011 & 0010 -> 0010


// 方法二：n&(n-1)
// (n-1)：二进制数字n最右边的1变成0，此1右边的0都变成1
// n&(n-1)：二进制数字n最右边的1变成0，其余不变
// 循环消去最右边的1，当n=0时跳出
func hammingWeight2(num uint32) int {
	var res = 0
	for num != 0 {
		res += 1
		num &= num - 1
	}
	return res
}

