<?php

// 剑指 Offer 17. 打印从1到最大的n位数

// 输入数字 n，按顺序打印出从 1 到最大的 n 位十进制数。比如输入 3，则打印出 1、2、3 一直到最大的 3 位数 999。

// end = 10^n-1
// 考虑大数越界的情况下的打印
// 1.无论是Int/long任意类型，数字取值都是有范围的，因此大数表示用string

// 递归生成全排列：先固定高位，向低位递归
class Solution {

    /**
     * @param Integer $n
     * @return Integer[]
     */
    function printNumbers($n) {
        $res = [];
        $num = [];
        $fc = function ($x) use ($n,$res,$num){
            if($x == $n){
                return ;
            }
            for($i = 1;$i <= $n;$i++){
                $num[$x] = 1;
            }
        };
    }
}

$arr = ['abc','der',3];


array_walk($arr,function ($v,$k){
    var_dump($v,$k);
    return ucfirst($v);
});

array_map(function ($v){

},$arr);

var_dump($arr);

$s_a = 'abc';
$s_b = $s_a;
$s_a = 'bcd';

echo $s_b;

echo PHP_EOL;

$s_a = 'abc';
$s_b = &$s_a;
$s_a = 'bcd';
echo $s_b;
