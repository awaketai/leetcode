package main

import "fmt"

// 1.Two Sum
// Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.
//
//You may assume that each input would have exactly one solution, and you may not use the same element twice.
//
//You can return the answer in any order.

var nums1 = []int{2,7,11,15}
var target1 = 9

var nums2 = []int{3,2,4}
var target2 = 6

var nums3 = []int{3,3}
var target3 = 6

var nums4 = []int{-3,4,3,90}
var target4 = 0

func main()  {
	ret := twoSum(nums4,target4)
	fmt.Println(ret)
}

func twoSum(nums []int, target int) []int {
	mem := make(map[int]int)
	for i := 0;i < len(nums);i++ {
		//
		if e,ok := mem[target - nums[i]];ok {
			return []int{e,i}
		}
		mem[nums[i]] = i
	}
	return nil
}