<?php

// 剑指 Offer 39. 数组中出现次数超过一半的数字

// 数组中有一个数字出现的次数超过数组长度的一半，请找出这个数字。
//
//
//
//你可以假设数组是非空的，并且给定的数组总是存在多数元素。


class Solution {

    /**
     * @param Integer[] $nums
     * @return Integer
     */
    function majorityElement($nums) {
        $mem = [];
        $len = count($nums);
        for($i = 0;$i < $len;$i++){
            if(isset($mem[$nums[$i]])){
                $mem[$nums[$i]]++;
            }else{
                $mem[$nums[$i]] = 1;
            }
        }
        foreach ($mem as $ko => $vo){
            if($vo > $len / 2){
                return $ko;
            }
        }
    }

    // 方法二：排序
    // 如果将数组 nums 中的所有元素按照单调递增或单调递减的顺序排序，那么下标为 n/2的元素一定是众数
    /**
     * @param Integer[] $nums
     * @return Integer
     */
    function majorityElement2($nums)
    {
        sort($nums);
        return $nums[count($nums) / 2];
    }

    // 思路来自
    // https://leetcode-cn.com/problems/shu-zu-zhong-chu-xian-ci-shu-chao-guo-yi-ban-de-shu-zi-lcof/solution/mian-shi-ti-39-shu-zu-zhong-chu-xian-ci-shu-chao-3/
    // 方法三：摩尔投票法：
    // 记输入数组nums的众数为x，数组长度为n
    // 推论1：若记众数的票数为+1,非众数的票数为-1,则一定有所有数字的票数和>0
    // 推论2：若数组的前a个数字的票数和=0，则数组剩余(n-1)个数字的票数和仍然一定>0
    // 即后(n-a)个数字的众数仍为x
    /**
     * @param Integer[] $nums
     * @return Integer
     */
    function majorityElement3($nums)
    {
        // 1.假设初始众数为第0个元素，票数为0
        // 当票数为0时，假设当前数字为众数
        $len = count($nums);
        $votes = 0;
        $x = 0;
        for($i = 0;$i < $len;$i++){
            // 当票数为0时，指定新的众数
            if($votes == 0){
                $x = $nums[$i];
            }
            // 如果当前元素是指定的众数，则票数+1
            // 否则票数-1
            if($nums[$i] == $x){
                $votes++;
            }else{
                $votes--;
            }
        }
        // 验证是否为众数
        $count = 0;
        foreach ($nums as $vo){
            if($vo == $x){
                $count++;
            }
        }
        if($count > $len / 2){
            return $x;
        }
        return 0;
    }
}

$nums1 = [1, 2, 3, 2, 2, 2, 5, 4, 2];
$obj = new Solution();
$ret = $obj->majorityElement3($nums1);
var_dump($ret);
