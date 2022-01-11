<?php

// 剑指 Offer 67. 把字符串转换成整数
// 写一个函数 StrToInt，实现把字符串转换成整数这个功能。不能使用 atoi 或者其他类似的库函数。
//
//
//
//首先，该函数会根据需要丢弃无用的开头空格字符，直到寻找到第一个非空格的字符为止。
//
//当我们寻找到的第一个非空字符为正或者负号时，则将该符号与之后面尽可能多的连续数字组合起来，作为该整数的正负号；假如第一个非空字符是数字，则直接将其与之后连续的数字字符组合起来，形成整数。
//
//该字符串除了有效的整数部分之后也可能会存在多余的字符，这些字符可以被忽略，它们对于函数不应该造成影响。
//
//注意：假如该字符串中的第一个非空格字符不是一个有效整数字符、字符串为空或字符串仅包含空白字符时，则你的函数不需要进行转换。
//
//在任何情况下，若函数不能进行有效的转换时，请返回 0。
class Solution {
    function strToInt($str){
        $intMax32 = 2**31 - 1;
        $intMin32 = -2**31;
        // 流转状态
        // 4种状态：0 起始态；1 符号态；2 数字态；3 结束态
        // 4种转移条件：空格 +/- 数字 其他
        $states = [
            [" " => 0,"s" => 1,"d" => 2,"x" => 3],
            ["d" => 2," " => 3,"s" => 3,"x" => 3],
            ["d" => 2," " => 3,"s" => 3,"x" => 3],
            [" " => 3,"s" => 3,"d" => 3,"x" => 3]
        ];
        $state = 0; // 当前状态
        $num = 0; // 记录数
        $sign = 1; // 符号
        $len = strlen($str);
        for($i = 0;$i < $len;$i++){
            if($str[$i] == " "){
                // 空格
                $t = " ";
            }elseif($str[$i] == "+" || $str[$i] == "-"){
                $t = "s";
            }elseif(is_numeric($str[$i])){
                $t = "d";
                $num = $num * 10 + $str[$i];
                if($num > $intMax32 && $sign == 1){
                    return $intMax32;
                }
                if($num > $intMax32 && $sign == -1){
                    return $intMin32;
                }
            }else{
                $t = "x";
            }

            // 当前状态
            $state = $states[$state][$t];
            if ($state == 3){
                // 结束态
                break;
            }

            if ($str[$i] == "-"){
                $sign = -1;
            }
        }
        return $num * $sign;
    }
}

$str1 = "42";
$str2 = "   -42";
$str3 = "4193 with words";
$str4 = "words and 987";
$str5 = "-91283472332";
$str6 = "23 45 89 a 123";
$str7 = "-23 45 89 a 123";
$str8 = "-2147483647";
$str9 = "+1";

$obj = new Solution();
var_dump($obj->strToInt($str1));
var_dump($obj->strToInt($str2));
var_dump($obj->strToInt($str3));
var_dump($obj->strToInt($str4));
var_dump($obj->strToInt($str5));
var_dump($obj->strToInt($str6));
var_dump($obj->strToInt($str7));
var_dump($obj->strToInt($str8));
var_dump($obj->strToInt($str9));