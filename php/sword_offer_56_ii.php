<?php

// 剑指 Offer 56 - II. 数组中数字出现的次数 II

// 在一个数组 nums 中除一个数字只出现一次之外，其他数字都出现了三次。请找出那个只出现一次的数字。

class Solution {

    /**
     * @param Integer[] $nums
     * @return Integer
     */
    function singleNumber($nums) {
        // 哈希表统计
        $len = count($nums);
        $mem = [];
        for($i = 0;$i < $len;$i++){
            if(isset($mem[$nums[$i]])){
                $mem[$nums[$i]]++;
            }else{
                $mem[$nums[$i]] = 1;
            }
        }
        foreach ($mem as $ko => $vo){
            if($vo == 1){
                return $ko;
            }
        }
        return 0;
    }
    // 方法二：位运算
    // https://leetcode-cn.com/problems/shu-zu-zhong-shu-zi-chu-xian-de-ci-shu-ii-lcof/solution/3chong-jie-fa-ji-jian-cdai-ma-by-fengzil-fmlr/
    // 通过统计所有数，每个bit位出现的次数，再把每个位的 次数%3，也就算出只出现1次数的bit位
    // 最后再各个bit位拼接起来，就得到了只出现1次的数
    function singleNumber2($nums){

    }
}
$nums1 = [3,4,3,3]; $k1 = 4;
$nums2 = [9,1,7,9,7,9,7]; $k2 = 1;
$obj = new Solution();
$ret = $obj->singleNumber($nums2);
var_dump($ret);
