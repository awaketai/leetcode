<?php

class Solution {

    /**
     * @param String $s
     * @return String
     */
    function replaceSpace($s) {
        $str = "";
        $len = strlen($s);
        for($i = 0;$i<$len;$i++){
            if($s[$i] == " "){
                $str .= "%20";
            }else{
                $str .= $s[$i];
            }
        }
        return $str;
    }
}

$s = "We are happy.";
$obj = new Solution();
$ret = $obj->replaceSpace($s);
var_dump($ret);

