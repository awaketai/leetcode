package main

import "fmt"

//剑指 Offer 57. 和为s的两个数字

// 输入一个递增排序的数组和一个数字s，在数组中查找两个数，使得它们的和正好是s。如果有多对数字的和等于s，则输出任意一对即可。

var nums1 = []int{2,7,11,15}
var target1 = 9
var nums2 = []int{10,26,30,31,47,60}
var target2 = 40
var nums3 = []int{14,15,16,22,53,60}
var target3 = 76

func main()  {
	ret := twoSum(nums2,target2)
	fmt.Println(ret)
}

// 目标值挨个减去数列中的值，直到找出差值等于数列中值的元素
// 将当前元素以[元素值] -> 元素值方式存储到map中
// 然后目标值减去当前循环的值，如果在map中存在
// 则返回当前循环元素和map中存在的值
func twoSum(nums []int, target int) []int {
	length := len(nums)
	mem := make(map[int]int)
	for i := 0;i < length;i++ {
		if v,ok := mem[target - nums[i]];ok {
			return []int{nums[i],v}
		}
		mem[nums[i]] = nums[i]
	}
	return []int{}
}

