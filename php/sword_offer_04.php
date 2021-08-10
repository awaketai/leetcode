<?php

//剑指 Offer 04. 二维数组中的查找

// 在一个 n * m 的二维数组中，每一行都按照从左到右递增的顺序排序，每一列都按照从上到下递增的顺序排序。请完成一个高效的函数，输入这样的一个二维数组和一个整数，判断数组中是否含有该整数。
//

class Solution {

    /**
     * 方法一：对于每一行，使用二分查找的方式
     * @param Integer[][] $matrix
     * @param Integer $target
     * @return Boolean
     */
    function findNumberIn2DArray($matrix, $target) {
        // 对每一行，按照二分法查找
        $lenRow = count($matrix);
        if($lenRow == 0){
            return false;
        }
        $lenColumn = count($matrix[0]);
        for($i = 0;$i < $lenRow;$i++){
            $start = 0;
            $end = $lenColumn - 1;
            while ($start <= $end){
                $mid = $start + floor( ($end - $start) / 2);
                var_dump($mid);
                if($matrix[$i][$mid] == $target){
                    return true;
                }
                if($matrix[$i][$mid] < $target){
                    // 从右边找
                    $start = $mid + 1;
                }else{
                    $end = $mid - 1;
                }
            }
        }
        return false;
    }

    // 方法二：以坐下角为坐标原点，移动坐标的方式进行查找"
    function findNumberIn2DArray2($matrix, $target){
        $lenRow = count($matrix);
        if($lenRow == 0){
            return false;
        }
        $lenColumn = count($matrix[0]);
        $x = $lenRow - 1;
        $y = 0;
        while ($x >= 0 && $y < $lenColumn){
            if($matrix[$x][$y] == $target){
                return true;
            }
            if($matrix[$x][$y] < $target){
                // 当前元素小于目标元素，往右侧找
                $y++;
            }else{
                // 当前元素大于目标元素，往上侧找
                $x--;
            }
        }
        return false;
    }
}

$f1 = [
    [1,   4,  7, 11, 15],
    [2,   5,  8, 12, 19],
    [3,   6,  9, 16, 22],
    [10, 13, 14, 17, 24],
    [18, 21, 23, 26, 30]
];

$obj = new Solution();
//$ret = $obj->findNumberIn2DArray($f1,20);
$ret = $obj->findNumberIn2DArray2($f1,5);
var_dump($ret);