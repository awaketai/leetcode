<?php

// 剑指 Offer 45. 把数组排成最小的数

// 输入一个非负整数数组，把数组里所有数字拼接起来排成一个数，打印能拼接出的所有数字中最小的一个。

// 示例 1:
//
//输入: [10,2]
//输出: "102"

// 自定义排序规则 +:连接符
// 若 x+y > y+x 则 x > y
// 若 x+y < y+x 则 x < y
// 若x < y：排序后x应在y左边，否则反之
// 依据此思路，使用排序算法进行排序

// 推理此法结果准确性在剑指offer书中
class Solution {
    /**
     * @param Integer[] $nums
     * @return String
     */
    function minNumber($nums) {
        $str = '';
        $len = count($nums);
        $this->_quickSort($nums,0,$len-1);
        foreach ($nums as $vo){
            $str .= $vo;
        }
        return $str;
    }

    function _quickSort(&$nums,$left,$right){
        if($left < $right){
            $pivot = $this->doQuickSort($nums,$left,$right);
            $this->_quickSort($nums,$left,$pivot -1 );
            $this->_quickSort($nums,$pivot + 1,$right);
        }
    }

    function doQuickSort(&$nums,$left,$right){
        $pivot = $nums[$left]; // 定义基准元素
        while ($left < $right){
            // 将小于基准元素的放左边，大于基准元素的放右边
            while ($left < $right && strval($pivot.$nums[$right]) <= strval($nums[$right].$pivot)){
                $right--;
            }
            $nums[$left] = $nums[$right];
            while ($left < $right && strval($nums[$left].$pivot) <= strval($pivot.$nums[$left])){
                $left++;
            }
            $nums[$right] = $nums[$left];
        }
        $nums[$left] = $pivot;
        return $left;
    }

    /**
     * 快速排序
     * @param $arr
     * @param $left
     * @param $right
     * @return mixed
     */
    function quickSort(&$arr,$left,$right){
        if($left >= $right){
            return $arr;
        }
        $i = $left;$j = $right;
        $tmp = $arr[$i];
        while ($i < $j){
            while($arr[$j].$arr[$left] >= $arr[$left].$arr[$j] && $i < $j){
                $j--;
            }
            while($arr[$i].$arr[$left] <= $arr[$left].$arr[$i] && $i < $j){
                $i++;
            }
            $tmp = $arr[$i];
            $arr[$i] = $arr[$j];
            $arr[$j] = $tmp;
        }
        $arr[$i] = $arr[$left];
        $arr[$left] = $tmp;
        $this->quickSort($arr,$left,$i-1);
        $this->quickSort($arr,$i+1,$right);
    }
}

$nums1 = [10,2];
$nums2 = [3,30,34,5,9];
$nums3 = [1,2,3,1];

$obj = new Solution();
$ret = $obj->minNumber($nums1);
if($ret != '102'){
    die("error: expected:102 actual:".$ret);
}
$ret = $obj->minNumber($nums2);
if($ret != '3033459'){
    die("error: expected:3033459 actual:".$ret);
}
$ret = $obj->minNumber($nums3);
if($ret != '1123'){
    die("error: expected:1123 actual:".$ret);
}

echo "tet success!";
