package main

import (
	"fmt"
	"math"
	"sort"
)

// 16.3Sum Closest

// Given an array nums of n integers and an integer target, find three integers in nums such that the sum is closest to target. Return the sum of the three integers. You may assume that each input would have exactly one solution.

func main()  {
	//var nums = []int{-1,2,1,-4}
	//var target1 = 1
	var nums1 = []int{1,1,1,2,3}
	var target2 = 7
	ret := threeSumClosest(nums1,target2)
	fmt.Println(ret)

}
// 1.和最接近：和减去目标值最小
// 2.穷举
func threeSumClosest(nums []int, target int) int {
	// 1.排序
	sort.Ints(nums)
	// 初始化结果值
	length := len(nums)
	res := nums[0] + nums[1] + nums[2]
	for i := 0;i < length;i++ {
		// 设置指针
		left := i + 1
		right := length - 1
		for left < right {
			fmt.Println(nums[i],nums[left],nums[right])
			sum := nums[i] + nums[left] + nums[right]
			// 如果结果值大于当前和，则替换
			if math.Abs(float64(target - sum)) < math.Abs(float64(target - res)) {
				res = sum
			}
			if sum == target {
				return sum
			}
			if sum > target {
				right--
				// 解决元素重复nums[right]重复
				for left != right && nums[right] == nums[right + 1] {
					right--
				}
			}else {
				left++
				// 解决元素重复nums[left]重复
				for left != right && nums[left] == nums[left - 1] {
					left++
				}
			}
		}
		// 解决元素重复nums[i]重复
		for i < len(nums) - 2 && nums[i] == nums[i + 1] {
			i++
		}

	}
	return res
}