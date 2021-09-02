<?php

// 剑指 Offer 33. 二叉搜索树的后序遍历序列

// 输入一个整数数组，判断该数组是不是某二叉搜索树的后序遍历结果。如果是则返回 true，否则返回 false。假设输入的数组的任意两个数字都互不相同。

// 后序遍历：左子树 -> 右子树 -> 根节点
// 二叉搜索树：左子树中所有节点 < 根节点值；右子树所有节点值 > 根节点值
// 其左、右子树也分别为二叉搜索树

// [1,3,2,6,5] true
// [1,6,3,2,5] false

// 思路来自：https://leetcode-cn.com/problems/er-cha-sou-suo-shu-de-hou-xu-bian-li-xu-lie-lcof/solution/di-gui-he-zhan-liang-chong-fang-shi-jie-jue-zui-ha/

// [3，5，4，10，12，9]
// 后序遍历的最后一个数字是根节点:9
// 从前往后找到第一个比根节点大的数字[10,12]，则[10,12]是9的右子节点
// [3,5,4]都是9的左子节点
// 如果[10,12]后还有小于9的值，则不是二叉搜索树

// [3，5，13，10，12，9]
// 1.根节点是9，则右子树[13,10,12],左子树[3,5]
// 2.[13,10,12]的根节点是12，则其右子树为[13,10],10<12根节点，其不是二叉搜索树
class Solution {

    /**
     * @param Integer[] $postorder
     * @return Boolean
     */
    function verifyPostorder($postorder) {

        return $this->recur($postorder,0,count($postorder) - 1);
    }

    function recur($postorder,$left,$right){
        if($left >= $right){
            // 没有节点
            return true;
        }
        $root = $postorder[$right]; // 确定根节点值
        // 从前往后找比根节点大的第一个数组,二分法
        $point = $left;
        while ($postorder[$point] < $root){
            $point++;
        }
        // 确定到前面所有的小于根节点的值后，查看后面的元素值是否都
        // 比根节点值大
        $tmp = $point;
        while ($tmp < $right){
            if($postorder[$tmp++] < $root){
                // 如果存在比根节点值小的元素，则不是二叉搜索树
                return false;
            }
        }
        // 对左右子节点递归调用
        return $this->recur($postorder,$left,$point - 1) && $this->recur($postorder,$point,$right - 1);
    }
}

$s1 = [1,3,2,6,5];
$s2 = [1,6,3,2,5];
$s3 = [3,5,4,10,12,9];
$s4 = [3,5,13,10,12,9];
$obj = new Solution();
$ret = $obj->verifyPostorder($s4);
var_dump($ret);