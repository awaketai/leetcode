<?php

// 剑指 Offer 42. 连续子数组的最大和

// 输入一个整型数组，数组中的一个或连续多个整数组成一个子数组。求所有子数组的和的最大值。
//
//要求时间复杂度为O(n)。

// 1.状态定义：设动态规划列表dp，dp[i]代表以nums[i]结尾的连续子数组最大和
// 2.若dp[i-1] <=0,说明dp[i-1]对dp[i]产生负贡献，即dp[i-1]+nums[i]还不如nums[i]本身大
// 当dp[i-1] >0时，执行dp[i]=dp[i-1]+nums[i]
// 当dp[i-1]<=0时，执行dp[i]=nums[i]

// 初始状态：dp[0]=nums[0],即以nums[0]结尾的连续子数组最大和为nums[0]
// 由于dp[i]只与dp[i-1]和nums[i]有关系，因此可以将原数组nums用作dp列表
class Solution {

    /**
     *
     * @param Integer[] $nums
     * @return Integer
     */
    function maxSubArray($nums) {
        // 方法一：
        $len = count($nums);
        $sum = $maxSum = $nums[0];
        for($i = 1;$i < $len;$i++){
            if($sum <= 0){
                $sum = $nums[$i];
            }else{
                $sum += $nums[$i];
            }
            if($sum > $maxSum){
                $maxSum = $sum;
            }
        }
        return $maxSum;
    }

    // 方法二：动态规划
    function maxSubArray2($nums){
        $len = count($nums);
        $dp[0] = $nums[0];
        $sum = $dp[0];
        for($i = 1;$i < $len;$i++){
            if($dp[$i - 1] <= 0){
                $dp[$i] = $nums[$i];
            }else{
                $dp[$i] = $dp[$i - 1] + $nums[$i];
            }
            if($sum < $dp[$i]){
                $sum = $dp[$i];
            }
        }
        return $sum;
    }
}

$nums1 = [-2,1,-3,4,-1,2,1,-5,4]; $ret1 = 6;
$nums2 = [1];$ret2 = 1;
$nums3 = [-1];$ret3 = -1;
$nums4 = [1,2];$ret4 = 3;
$obj = new Solution();
$ret = $obj->maxSubArray($nums1);
var_dump($ret);

