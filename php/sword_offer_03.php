<?php


class Solution {

    /**
     * @param Integer[] $nums
     * @return Integer
     */
    function findRepeatNumber($nums) {
        return $this->method2($nums);
    }

    // 方法一：借用外部存储
    function method1($nums){
        $mem = [];
        $len = count($nums);
        for($i = 0;$i < $len;$i++){
            if (isset($mem[$nums[$i]])){
                return $nums[$i];
            }
            $mem[$nums[$i]] = 1;
        }
        return 0;
    }

    // 方法二：在当前数组内交换元素，将当前元素放置到当前元素索引位置
    function method2($nums){
        $len = count($nums);
        for($i = 0;$i < $len;$i++){
            if($i == $nums[$i]){
                continue;
            }
            // 如果当前索引位置已经存在值，则证明重复
            if($nums[$i] == $nums[$nums[$i]]){
                return $nums[$i];
            }
            // 将当前索引位置，放入当前索引相等的值
            $tmp = $nums[$nums[$i]];
            $nums[$nums[$i]] = $nums[$i];
            $nums[$i] = $tmp;
        }
        return 0;
    }
}

$nums = [2, 3, 1, 0, 2, 5, 3];
$obj = new Solution();
$ret = $obj->findRepeatNumber($nums);
var_dump($ret);