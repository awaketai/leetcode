<?php

// 917. Reverse Only Letters

// Given a string s, reverse the string according to the following rules:
//
//All the characters that are not English letters remain in the same position.
//All the English letters (lowercase or uppercase) should be reversed.
//Return s after reversing it.

// ascii A-Z 65-90
// a-z 97-122

// 从前往后找，如果是字母，则找倒序的第一个字符

class Solution {

    /**
     * @param String $s
     * @return String
     */
    function reverseOnlyLetters($s) {
        $len = strlen($s);
        $end = $len-1;
        $arr = [];
        for($i = 0;$i < $len;$i++){
            if($this->isLetter(ord($s[$i]))){
                while (!$this->isLetter(ord($s[$end]))){
                    $end--;
                }
                $arr[] = $s[$end];
                $end--;
            }else{
                $arr[] = $s[$i];
            }
        }
        return implode("",$arr);
    }

    function isLetter($i){
        if((65 <= $i && $i <= 90) || (97 <= $i && $i <= 122)){
            return true;
        }
        return false;
    }
}

$s1 = "a-bC-dEf-ghIj";
$s2 = "Test1ng-Leet=code-Q!";
$obj = new Solution();
$ret = $obj->reverseOnlyLetters($s2);
echo $ret;
