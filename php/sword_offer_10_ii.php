<?php


// 剑指 Offer 10- II. 青蛙跳台阶问题

// 一只青蛙一次可以跳上1级台阶，也可以跳上2级台阶。求该青蛙跳上一个 n 级的台阶总共有多少种跳法。
//
//答案需要取模 1e9+7（1000000007），如计算初始结果为：1000000008，请返回 1。
//

// dp方法，与斐波那契数列等价，不过此处初始条件成了：
// f(0)=1,f(1)=1,f(2)=2
// 设跳上n级台阶有f(n)中跳法，在所有跳法总，青蛙的最后一步只有两种情况：跳上1级台阶或者2级台阶
// 1.当为1级台阶：剩n-1个台阶，此情况共有f(n-1)种跳法
// 2.当为2级台阶：剩n-2个台阶，此情况共有f(n-2)种跳法
// f(n)即为以上两种情况之和，即：f(n) = f(n-1) + f(n-2)


class Solution {

    /**
     * 方法一：dp
     * @param Integer $n
     * @return Integer
     */
    function numWays($n) {

        if($n == 0){
            return 1;
        }
        //dp
        $mem[0] = 1;
        $mem[1] = 1;
        for($i = 2;$i <= $n;$i++){
            $mem[$i] = ($mem[$i - 1] + $mem[$i - 2]) % 1000000007;
        }
        return $mem[$n];
    }

    function numWays2($n){
        $mem = [];
        return $this->dfs($n,$mem);
    }

    // 方法二：假设要求7的所有方法，则必须先求出6，5的所有步骤，一次类推，将其视为一颗二叉树
    // 进行深度遍历
    // 递归有可能会超时
    function dfs($n,&$mem){
        if($n <= 1){
            return 1;
        }
        if(isset($mem[$n])){
            return $mem[$n];
        }
        // 左子节点
        $leftVal = $this->dfs($n - 1,$mem);
        // 右子节点
        $rightVal = $this->dfs($n - 2,$mem);
        $val = ($leftVal + $rightVal) % 1000000007;
        $mem[$n] = $val;
        return $val;
    }
}

$obj = new Solution();
//$ret = $obj->numWays2(7);
//$ret = $obj->numWays2(2);
$ret = $obj->numWays2(100);

var_dump($ret);
