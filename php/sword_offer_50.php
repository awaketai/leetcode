<?php

//剑指 Offer 50. 第一个只出现一次的字符

// 在字符串 s 中找出第一个只出现一次的字符。如果没有，返回一个单空格。 s 只包含小写字母。
//
//示例:
// s = "abaccdeff"
//返回 "b"
//
//s = ""
//返回 " "

class Solution {

    /**
     * @param String $s
     * @return String
     */
    function firstUniqChar($s) {
        $len = strlen($s);
        if($len == 0){
            return ' ';
        }
        $mem = [];
        for($i = 0;$i < $len;$i++){
            if(isset($mem[$s[$i]])){
                $mem[$s[$i]]++;
            }else{
                $mem[$s[$i]] = 1;
            }
        }
        foreach ($mem as $ko => $vo){
            if($vo == 1){
                return $ko;
            }
        }
        return ' ';
    }
}

$s = 'abaccdeff';
$obj = new Solution();
$obj->firstUniqChar($s);