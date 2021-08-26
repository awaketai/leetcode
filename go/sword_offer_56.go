package main

import "fmt"

// 剑指 Offer 56 - II. 数组中数字出现的次数 II

// 在一个数组 nums 中除一个数字只出现一次之外，其他数字都出现了三次。请找出那个只出现一次的数字。

// 解法：
// https://leetcode-cn.com/problems/shu-zu-zhong-shu-zi-chu-xian-de-ci-shu-ii-lcof/solution/mian-shi-ti-56-ii-shu-zu-zhong-shu-zi-chu-xian-d-4/

var nums1 = []int{3,4,3,3}; var k1 = 4;
var nums2 = []int{9,1,7,9,7,9,7}; var k2 = 1;
func main()  {
	ret := singleNumber(nums2)
	fmt.Println(ret)
}

// 对于任意数字二进制位x，有：
// 异或运算：x ^ 0 = x​ ， x ^ 1 = ~x
// 与运算：x & 0 = 0 ， x & 1 = x

// 这里采用哈希表法
func singleNumber(nums []int) int {
	var mem = make(map[int]int)
	for _,v := range nums {
		if _,ok := mem[v];ok {
			mem[v]++
		}else{
			mem[v] = 1
		}
	}
	for i,v := range mem {
		if v == 1 {
			return i
		}
	}
	return 0
}