<?php

//96. Unique Binary Search Trees

//Given an integer n, return the number of structurally unique BST's (binary search trees) which has exactly n nodes of unique values from 1 to n.

class Solution {

    /**
     * @param Integer $n
     * @return Integer
     */
    function numTrees($n) {
//        $mem = [];
//        return $this->cal(1,$n,$mem);
        $arr = [];
        $arr[0] = $arr[1] = 1;
        for($i = 2;$i <= $n;$i++){
            $arr[$i] = 0;
            for($j = 1;$j <= $i;$j++){
                $arr[$i] += $arr[$j - 1] * $arr[$i - $j];
            }
        }
        return $arr[$n];
    }

    // 以有序数列的每个值为根节点，进行左右子树构建

    /**
     * 执行时间太长
     * @param $low
     * @param $high
     * @param $mem
     * @return float|int|mixed
     */
    function cal1($low,$high,$mem){
        if($low > $high){
            // 节点为null，值为1
            return 1;
        }
        $key = $low .'-'.$high;
        if(!empty($mem[$key])){
            return $mem[$key];
        }
        $res = 0;
        // $i:根节点
        for($i = $low;$i <= $high;$i++){
            $left = $this->cal1($low,$i - 1,$mem);
            $right = $this->cal1($i + 1,$high,$mem);
            $res += $left * $right;
        }
        $mem[$key] = $res;
        return $res;
    }
}

$obj = new Solution();
echo $obj->numTrees(16);