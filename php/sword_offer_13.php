<?php

// 剑指 Offer 13. 机器人的运动范围
// 地上有一个m行n列的方格，从坐标 [0,0] 到坐标 [m-1,n-1] 。
// 一个机器人从坐标 [0, 0] 的格子开始移动，它每次可以向左、右、上、下移动一格（不能移动到方格外），也不能进入行坐标和列坐标的数位之和大于k的格子。
// 例如，当k为18时，机器人能够进入方格 [35, 37] ，因为3+5+3+7=18。但它不能进入方格 [35, 38]，因为3+5+3+8=19。请问该机器人能够到达多少个格子？
//
class Solution {

    /**
     * @param Integer $m
     * @param Integer $n
     * @param Integer $k
     * @return Integer
     */
    function movingCount($m, $n, $k) {
        if($k == 0){
            return 1;
        }
        $r = $c = $ret = 0;
        for($r = 0;$r < $m;$r++){
            for($c = 0;$c < $n;$c++){

            }
        }
    }

    function inArea($m,$n,$r,$c){

        return 0 <= $r && $r < $m && 0 <= $c && $c < $n;
    }

    function getSum($k){
        $ret = 0;
        while ($k != 0){
            $ret += $k % 10;
            $k /= 10;
        }
    }
}
$obj = new Solution();
$obj->movingCount(12,13,11);
