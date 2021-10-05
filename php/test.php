<?php

//[-3,-2,-1,1,2,3,4,6] => [[-3,1,2],[-2,-1,3]]
//找出有序数组中随机3个数和为0的所有情况

function sumZero($nums){
    $len = count($nums);
    $ret = [];
    for($i = 0;$i < $len;$i++){
        $pivot = $nums[$i];
        $start = $i + 1;
        if($start >= $len){
            break;
        }
        $end = $len - 1;
        while ($start <= $end){
            $sum = $pivot + $nums[$start] + $nums[$end];
            if($sum == 0){
                // array push
                $ret[] = [$pivot,$nums[$start],$nums[$end]];
                $start++;
                $end--;
            }elseif($sum > 0){
                $end--;
            }elseif($sum < 0){
                $start++;
            }
        }
    }
    return $ret;
}

$nums = [-3,-2,-1,1,2,3,4,6];

$ret = sumZero($nums);

print_r($ret);