<?php

// 剑指 Offer 11. 旋转数组的最小数字


//把一个数组最开始的若干个元素搬到数组的末尾，我们称之为数组的旋转。输入一个递增排序的数组的一个旋转，输出旋转数组的最小元素。例如，数组 [3,4,5,1,2] 为 [1,2,3,4,5] 的一个旋转，该数组的最小值为1。  
//

class Solution {

    /**
     * @param Integer[] $numbers
     * @return Integer
     */
    function minArray($numbers) {
        $len = count($numbers);
        if($len == 0){
            return 0;
        }
        $end = $len - 1;
        $start = 0;
        while ($start <= $end) {
            $mid = $start + floor(($end - $start) / 2);
            if($numbers[$mid] < $numbers[$end]){
                // 如果中间值小于最右侧值，则最小值在左侧
                // 为了不漏值，不-1
                $end = $mid;
            }elseif($numbers[$mid] > $numbers[$end]){
                // 如果中间值大于最右侧值，则最小值在右侧
                $start = ++$mid;
            }else{
                // 如果中间值和最右侧值相等，
                // 则移动最右侧指针，一个一个找
                $end--;
            }
        }
        return $numbers[$start];
    }
}

$a11 = [3,4,5,1,2];
$a22 = [2,2,2,0,1];
$a33 = [7,0,1,1,1,1,1,2,3,4];
$a44 = [3,4,5,6,1,2,3];

$obj = new Solution();
$ret = $obj->minArray($a44);
var_dump($ret);
