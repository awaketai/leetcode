<?php

//剑指 Offer 57. 和为s的两个数字

// 输入一个递增排序的数组和一个数字s，在数组中查找两个数，使得它们的和正好是s。如果有多对数字的和等于s，则输出任意一对即可。

class Solution {

    /**
     * @param Integer[] $nums
     * @param Integer $target
     * @return Integer[]
     */
    function twoSum($nums, $target) {
        $len = count($nums);
        $mem = [];
        // 递增的数列，把$target右边的数列排除掉
        // 目标值减去当前存储的值，如果等于其中一个元素，则返回
        for($i = 0;$i < $len;$i++){
            if(isset($mem[$target - $nums[$i]])){
                return [$nums[$i],$mem[$target - $nums[$i]]];
            }
            $mem[$nums[$i]] = $nums[$i];
        }
        return [];
    }
}

$nums1 = [2,7,11,15]; $target1 = 9;
$nums2 = [10,26,30,31,47,60]; $target2 = 40;
$nums3 = [14,15,16,22,53,60];$target3 = 76;
$obj = new Solution();

$ret = $obj->twoSum($nums3,$target3);
var_dump($ret);
