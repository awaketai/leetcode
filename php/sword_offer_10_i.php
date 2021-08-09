<?php


// 剑指 Offer 10- I. 斐波那契数列

//写一个函数，输入 n ，求斐波那契（Fibonacci）数列的第 n 项（即 F(N)）。斐波那契数列的定义如下：
//
//F(0) = 0,   F(1) = 1
//F(N) = F(N - 1) + F(N - 2), 其中 N > 1.
//斐波那契数列由 0 和 1 开始，之后的斐波那契数就是由之前的两数相加而得出。
//
//答案需要取模 1e9+7（1000000007），如计算初始结果为：1000000008，请返回 1。

class Solution {

    /**
     * @param Integer $n
     * @return Integer
     */
    function fib($n) {
        if($n <= 1){
            return $n;
        }
        // 方法一
//        $mem = [];
//        return $this->dfs($n,$mem);
        // 方法二
        return $this->dp($n);
    }

    // 方法一：视n -> (1,0)为一棵二叉树，使用深度遍历方法
    // 自顶向下推导
    function dfs($n,&$mem){
        if($n <= 1){
            return $n;
        }
        if(isset($mem[$n])){
            return $mem[$n];
        }
        // 按照左侧为n-1,右侧为n-2的原则去
        $leftVal = $this->dfs($n - 1,$mem);
        $rightVal = $this->dfs($n - 2,$mem);
        $val = ($leftVal + $rightVal) % 1000000007;
        $mem[$n] = $val;
        return $val;
    }

    // 方法二：dp，动态规划，自底向上推导，n的斐波那契数列就是n-1 + n-2
    // 而0,1的斐波那契数列值是已知的，所以可以以0,1开始，推导出后面的元素的
    // 斐波那契值
    function dp($n){
        if($n <= 1){
            return $n;
        }

        // 初始化0,1元素的值
        $fib[0] = 0;
        $fib[1] = 1;
        // 以0，1开始推导后续的值
        for($i = 2;$i <= $n;$i++){
            $fib[$i] = ($fib[$i-1] + $fib[$i-2]) % 1000000007;
        }
        return $fib[$n];
    }
}

$obj = new Solution();
$ret = $obj->fib(100);
var_dump($ret);
