package main

import "fmt"

// 剑指 Offer 42. 连续子数组的最大和

// 输入一个整型数组，数组中的一个或连续多个整数组成一个子数组。求所有子数组的和的最大值。
//
//要求时间复杂度为O(n)。

// 输入: nums = [-2,1,-3,4,-1,2,1,-5,4]
//输出: 6
//解释: 连续子数组 [4,-1,2,1] 的和最大，为 6。

func main()  {
	//nums := []int{-2,1,-3,4,-1,2,1,-5,4}
	nums := []int{2,8,1,5,9}
	ret := maxSubArray2(nums)
	fmt.Println(ret)
}
// 方法一：
func maxSubArray1(nums []int) int {
	sum,maxSum := nums[0],nums[0]
	length := len(nums)
	for i := 1;i < length;i++ {
		if sum <= 0 {
			sum = nums[i]
		} else {
			sum += nums[i]
		}
		if sum > maxSum {
			maxSum = sum
		}
	}
	return maxSum
}

// 方法二：动态规划
func maxSubArray2(nums []int) int {
	length := len(nums)
	dp := make([]int,0)
	dp[0] = nums[0]
	ans := dp[0]
	for i := 1;i < length;i++ {
		// 当以i-1个数字结尾的数组和小于0，则最大和为i所在值
		if dp[i - 1] < 0 {
			dp[i] = nums[i]
		}else{
			dp[i] = dp[i-1] + nums[i]
		}
		if dp[i] > ans {
			ans = dp[i]
		}
	}
	return ans
}
