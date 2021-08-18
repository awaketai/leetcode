<?php

// 剑指 Offer 21. 调整数组顺序使奇数位于偶数前面

// 输入一个整数数组，实现一个函数来调整该数组中数字的顺序，使得所有奇数位于数组的前半部分，所有偶数位于数组的后半部分。

// 方法一：双指针法
// 首位双指针left,right
// left往右移，直到值为偶数
// right往左移，直到值为奇数
// 交换left和right的值
// 直到left和right相遇

// $a & $b 按位与，把$a和$b中都为1的设为1
// $a | $b 按位或，把$a和$b中任何一个为1的位设为1
// $a ^ $b 按位异或，把$a和$b中一个为1另一个为0的位设为1
// ~$a 按位取反，把$a中为0的位设为1，反之亦然
class Solution {

    /**
     * @param Integer[] $nums
     * @return Integer[]
     */
    function exchange($nums) {
        $len = count($nums);
        if($len == 0){
            return $nums;
        }
        $left = 0;
        $right = $len - 1;
        while ($left < $right){

            // $a & $b 按位与，把$a和$b中都为1的设为1
            // 0001 0001 -> 0001
            // 0010 0001 -> 0000
            if(($nums[$left] & 1) != 0){
                // 当前位是奇数
                $left++;
                continue;
            }
            if(($nums[$right] & 1) != 1){
                // 当前位是偶数
                $right--;
                continue;
            }
            $tmp = $nums[$left];
            $nums[$left] = $nums[$right];
            $nums[$right] = $tmp;
        }
        return $nums;
    }

    // 方法二：快慢指针法
    // 定义快慢指针，fast和low，fast在前，low在后
    // fast向前搜索奇数位置，low指向下一个奇数应当存储的位置
    // fast向前移动，当搜索到奇数时，将ta和nums[low]交换，此时low向前移动一个位置
    // 重复上述操作，直到fast指向数组末尾
    function exchange2($nums){
        $len = count($nums);
        if($len == 0){
            return $nums;
        }
        $fast = $low = 0;
        while ($fast < $len){
            if($nums[$fast] & 1){
                // 如果当前位是奇数
                $tmp = $nums[$fast];
                $nums[$fast] = $nums[$low];
                $nums[$low] = $tmp;
                $low++;
            }
            $fast++;
        }
        return $nums;
    }
}

$n = [1,2,3,4];
$n1 = [];
for($i = 1;$i <= 20;$i++){
    $n1[] = $i;
}
shuffle($n1);
$obj = new Solution();
$ret = $obj->exchange($n1);
print_r($ret);

var_dump($n[0] & 1);
var_dump($n[1] & 1);