<?php
// 1.Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.
//
//You may assume that each input would have exactly one solution, and you may not use the same element twice.
//
//You can return the answer in any order.

// 给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出 和为目标值 target  的那 两个 整数，并返回它们的数组下标。
//
//你可以假设每种输入只会对应一个答案。但是，数组中同一个元素在答案里不能重复出现。
//
//你可以按任意顺序返回答案。

class Solution {

    /**
     * @param Integer[] $nums
     * @param Integer $target
     * @return Integer[]
     */
    function twoSum($nums, $target) {
        if(empty($nums) || !is_array($nums)){
            return [];
        }
        $len = count($nums);
        for ($i = 0;$i < $len;$i++){
            for ($j = $i + 1;$j < $len;$j++){
                if(($nums[$i] + $nums[$j]) == $target){
                    return [$i,$j];
                }
            }

        }
    }

    /**
     * @param Integer[] $nums
     * @param Integer $target
     * @return Integer[]
     */
    function twoSum2($nums, $target) {
        $len = count($nums);
        $mem = [];
        for($i = 0;$i < $len;$i++){
            // 当前循环索引 = 已存储元素索引
            if(isset($mem[$target - $nums[$i]])){
                return [$mem[$target - $nums[$i]],$i];
            }
            $mem[$nums[$i]] = $i;
        }
        return null;
    }
}

$nums1 = [2,7,11,15];
$target1 = 9;

$nums2 = [3,2,4];
$target2 = 6;

$nums3 = [3,3];
$target3 = 6;

$nums4 = [-3,4,3,90];
$target4 = 0;

$obj = new Solution();
$ret = $obj->twoSum2($nums2,$target2);
var_dump($ret);
