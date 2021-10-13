package main

import (
	"fmt"
	"math"
	"sort"
)

// 16.3Sum Closest

// Given an array nums of n integers and an integer target, find three integers in nums such that the sum is closest to target. Return the sum of the three integers. You may assume that each input would have exactly one solution.

func main1()  {
	//var nums = []int{-1,2,1,-4}
	//var target1 = 1
	var nums1 = []int{1,1,1,2,3}
	var target2 = 7
	ret := threeSumClosest(nums1,target2)
	fmt.Println(ret)

}
// 1.和最接近：和减去目标值最小
// 2.穷举
// 最接近：差值的绝对值最小
func threeSumClosest(nums []int, target int) int {
	// 1.排序
	sort.Ints(nums)
	// 初始化结果值
	length := len(nums)
	res := nums[0] + nums[1] + nums[2]
	for i := 0;i < length-2;i++ {
		start,end := i+1,length-1
		for start < end {
			sum := nums[start] + nums[end] + nums[i]
			if math.Abs(float64(target-sum)) < math.Abs(float64(target-res)) {
				res = sum
			}
			if sum > target {
				end--
			}else if sum < target {
				start++
			}else{
				return res
			}
		}
	}
	return res
}

