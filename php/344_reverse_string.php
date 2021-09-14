<?php

// 344. Reverse String

// Write a function that reverses a string. The input string is given as an array of characters s.

// Example 1:
//
//Input: s = ["h","e","l","l","o"]
//Output: ["o","l","l","e","h"]
//Example 2:
//
//Input: s = ["H","a","n","n","a","h"]
//Output: ["h","a","n","n","a","H"]


class Solution {

    /**
     * @param String[] $s
     * @return NULL
     */
    function reverseString(&$s) {
        $start = 0;
        $len = count($s);
        $end = $len - 1;
        while ($start < $end){
            $tmp = $s[$start];
            $s[$start] = $s[$end];
            $s[$end] = $tmp;
            $start++;
            $end--;
        }
    }
}

$obj = new Solution();
$s1 = ["h","e","l","l","o"];
$s2 = ["H","a","n","n","a","h"];

$ret = $obj->reverseString($s2);
var_dump($s2);