<?php

function binarySearch(array $arr,$target){
    $start = 0;
    $end = count($arr) - 1;
    while ($start <= $end){
        $mid = $start + (floor(($end - $start) / 2));
        if($arr[$mid] == $target){
            return $mid;
        }
        if ($arr[$mid] < $target){
            // 又半边
            $start = $mid + 1;
        }else{
            $end = $mid - 1;
        }
    }
    return -1;
}

function main(){
    $arr = [];
    for($i = 0;$i < 1000;$i++){
        $arr[$i] = $i;
    }
    print_r(binarySearch($arr,1450));
}
main();