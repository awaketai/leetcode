<?php

// 剑指 Offer 40. 最小的k个数

// 输入整数数组 arr ，找出其中最小的 k 个数。例如，输入4、5、1、6、2、7、3、8这8个数字，则最小的4个数字是1、2、3、4。

class Solution {

    /**
     * @param Integer[] $arr
     * @param Integer $k
     * @return Integer[]
     */
    function getLeastNumbers($arr, $k) {
        $len = count($arr);
        if($len < 2){
            return $arr;
        }
        // 先排序
//        $ret = $this->QuickSort($arr);
        $ret = $this->MergeSort($arr);
        return array_slice($ret,0,$k);
    }

    // 方法一：快速排序
    function QuickSort($arr){

        return $this->_quickSort($arr,0,count($arr) - 1);
    }

    function _quickSort(&$arr,$left,$right){
        if ($left < $right) {
            $index = $this->partition($arr,$left,$right);
            $this->_quickSort($arr,$left,$index - 1);
            $this->_quickSort($arr,$index + 1,$right);
        }
        return $arr;
    }

    function partition(&$arr,$left,$right){
        $pivot = $left;
        $index = $pivot + 1;
        for($i = $index;$i <= $right;$i++){
            if($arr[$i] < $arr[$pivot]){
                $this->swap($arr,$i,$index);
                $index++;
            }
        }
        $this->swap($arr,$pivot,$index-1);
        return $index - 1;
    }

    function swap(&$arr,$i,$j){
        list($arr[$i],$arr[$j]) = [$arr[$j],$arr[$i]];
        return $arr;
    }

    // 方法二：归并排序
    function MergeSort($arr){
        $len = count($arr);
        if($len < 2){
            return $arr;
        }
        $middle = floor($len / 2);
        $left = array_slice($arr,0,$middle);
        $right = array_slice($arr,$middle);
        return $this->merge($this->MergeSort($left),$this->MergeSort($right));
    }

    function merge($left,$right){
        $ret = [];
        while (count($left) != 0 && count($right) != 0){
            if($left[0] <= $right[0]){
                $ret[] = $left[0];
                $left = array_slice($left,1);
            }else{
                $ret[] = $right[0];
                $right = array_slice($right,1);
            }
        }
        // 合并剩余的元素
        while (count($left) != 0){
            $ret[] = $left[0];
            $left = array_slice($left,1);
        }
        while (count($right) != 0){
            $ret[] = $right[0];
            $right = array_slice($right,1);
        }
        return $ret;
    }
}

$a1 = [3,2,1];
$t1 = 2;
$a2 = [0,1,2,1];
$t2 = 1;
$a3 = [4,5,1,6,2,7,3,8];
$t3 = 4;

$obj = new Solution();
$ret = $obj->getLeastNumbers($a3,$t3);
print_r($ret);

