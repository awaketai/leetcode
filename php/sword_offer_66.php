<?php

// 剑指 Offer 66. 构建乘积数组

// 给定一个数组 A[0,1,…,n-1]，请构建一个数组 B[0,1,…,n-1]，其中 B[i] 的值是数组 A 中除了下标 i 以外的元素的积, 即 B[i]=A[0]×A[1]×…×A[i-1]×A[i+1]×…×A[n-1]。不能使用除法。
class Solution {
    /**
     * @param Integer[] $a
     * @return Integer[]
     */
    function constructArr($a){
        // b[i] = 左边所有元素乘积 * 右边所有元素乘积
        $len = count($a);
        if($len == 0){
            return [];
        }

        // 左边所有元素乘积
        $leftMulti[0] = 1; // 左边界，默认值1
        for($i = 1;$i < $len;$i++){
            $leftMulti[$i] = $leftMulti[$i - 1] * $a[$i - 1];
        }

        print_r($leftMulti);

        // 右边所有元素乘积 len=5
        $rightMulti[$len - 1] = 1; // 右边界最大值
        for($j = $len - 2;$j >= 0;$j--){
            $rightMulti[$j] = $rightMulti[$j + 1] * $a[$j + 1];
        }

        print_r($rightMulti);

        $b = [];
        for($i = 0;$i < $len;$i++){
            $b[$i] = $leftMulti[$i] * $rightMulti[$i];
        }

        return $b;
    }

    /**
     * 改进版
     * @param Integer[] $a
     * @return Integer[]
     */
    function constructArr1($a){
        // b[i] = 左边所有元素乘积 * 右边所有元素乘积
        $len = count($a);
        if($len == 0){
            return [];
        }

        // 左边所有元素乘积
        $leftMulti[0] = 1; // 左边界，默认值1
        for($i = 1;$i < $len;$i++){
            $leftMulti[$i] = $leftMulti[$i - 1] * $a[$i - 1];
        }

        // 右边所有元素乘积 len=5
        $right = 1;
        for($j = $len - 1;$j >= 0;$j--){
            $leftMulti[$j] *= $right;
            $right *= $a[$j];
        }

        return $leftMulti;
    }
}

$a = [1,2,3,4,5];
$obj = new Solution();
$ret = $obj->constructArr1($a);
print_r($ret);