package main

import "fmt"

// 剑指 Offer 03. 数组中重复的数字

//找出数组中重复的数字。
//
//
//在一个长度为 n 的数组 nums 里的所有数字都在 0～n-1 的范围内。数组中某些数字是重复的，但不知道有几个数字重复了，也不知道每个数字重复了几次。请找出数组中任意一个重复的数字。
//
//示例 1：
//
//输入：
//[2, 3, 1, 0, 2, 5, 3]
//输出：2 或 3
//

var n1 = []int{2, 3, 1, 0, 2, 5, 3}

func main()  {
	ret := findRepeatNumber1(n1)
	fmt.Println(ret)
}

// 借用外部存储
func findRepeatNumber(nums []int) int {
	length := len(nums)
	var mem = make(map[int]int)
	for i := 0;i < length;i++ {
		if _,ok := mem[nums[i]];ok {
			return nums[i]
		}else{
			mem[nums[i]] = 1
		}
	}
	return 0
}

// 使用当前数组，原地交换元素到自己对应位置
func findRepeatNumber1(nums []int) int {
	length := len(nums)
	for i := 0;i < length;i++ {
		// 如果当前索引位置是当前值，则跳过
		if nums[i] == i {
			continue
		}
		// 如果当前索引的值等于即将要换到数组索引位置的元素值
		// 则已经存在重复元素
		if nums[i] == nums[nums[i]]{
			return nums[i]
		}
		// 交换元素到自己对应的索引位置
		nums[i],nums[nums[i]] = nums[nums[i]],nums[i]
	}
	return 0
}


