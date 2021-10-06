<?php

// 剑指 Offer 57 - II. 和为s的连续正数序列

// 输入一个正整数 target ，输出所有和为 target 的连续正整数序列（至少含有两个数）。
//
//序列内的数字由小到大排列，不同序列按照首个数字从小到大排列。

class Solution {

    /**
     * @param Integer $target
     * @return Integer[][]
     */
    function findContinuousSequence($target) {
        $left = 1; // 滑动窗口左边界
        $right = 2; // 滑动窗口右边界
        $sum = 3; // 滑动窗口内所有数字和
        $res = [];
        while ($left < $right){
            if($sum == $target){
                // 将当前窗口内所有数字加入结果集中
                $cur = [];
                for($i = $left;$i <= $right;$i++){
                    $cur[] = $i;
                }
                $res[] = $cur;
                // 将左边界和减掉，然后左边界向右移动，继续寻找下一个符合的结果
                $sum -= $left;
                $left++;
            }elseif($sum > $target){
                // 如果当前窗口内所有数字和大于target，则缩小窗口，左边界右移
                $sum -= $left;
                $left++;
            }else{
                // 如果当前窗口所有数字和小于target，则扩大窗口，右边界右移
                $right++;
                $sum += $right;
            }
        }
        return $res;
    }
}

$obj = new Solution();
$ret = $obj->findContinuousSequence(15);
print_r($ret);
