package main

import "fmt"

// 剑指 Offer 39. 数组中出现次数超过一半的数字

// 数组中有一个数字出现的次数超过数组长度的一半，请找出这个数字。
//
//
//
//你可以假设数组是非空的，并且给定的数组总是存在多数元素。

// 思路来自：
// https://leetcode-cn.com/problems/shu-zu-zhong-chu-xian-ci-shu-chao-guo-yi-ban-de-shu-zi-lcof/solution/mian-shi-ti-39-shu-zu-zhong-chu-xian-ci-shu-chao-3/

func main()  {
	nums1 := []int{1, 2, 3, 2, 2, 2, 5, 4, 2}
	ret := majorityElement(nums1)
	fmt.Println(ret)
}
// 方法三：摩尔投票法：
// 记输入数组nums的众数为x，数组长度为n
// 推论1：若记众数的票数为+1,非众数的票数为-1,则一定有所有数字的票数和>0
// 推论2：若数组的前a个数字的票数和=0，则数组剩余(n-1)个数字的票数和仍然一定>0
// 即后(n-a)个数字的众数仍为x
func majorityElement(nums []int) int {
	x,votes := 0,0
	length := len(nums)
	for i := 0;i < length;i++ {
		// 1.如果票数为0，则重新选择当前元素为新的众数
		if votes == 0 {
			x = nums[i]
		}
		// 如果当前元素和指定众数相等，则票数+1
		// 否则票数-1
		if nums[i] == x {
			votes++
		}else{
			votes--
		}
	}
	// 判断x是不是众数
	count := 0
	for _,v := range nums {
		if v == x {
			count++
		}
	}
	if count > length / 2 {
		return x
	}
	return 0
}
