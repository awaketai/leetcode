<?php

// 654. Maximum Binary Tree
// You are given an integer array nums with no duplicates. A maximum binary tree can be built recursively from nums using the following algorithm:
//
//Create a root node whose value is the maximum value in nums.
//Recursively build the left subtree on the subarray prefix to the left of the maximum value.
//Recursively build the right subtree on the subarray suffix to the right of the maximum value.
//Return the maximum binary tree built from nums.

//Definition for a binary tree node.
  class TreeNode {
      public $val = null;
      public $left = null;
      public $right = null;
      function __construct($val = 0, $left = null, $right = null) {
          $this->val = $val;
          $this->left = $left;
          $this->right = $right;
      }
  }

class Solution {

    /**
     * @param Integer[] $nums
     * @return TreeNode
     */
    function constructMaximumBinaryTree($nums) {
        $len = count($nums);
        if($len <= 0){
            return null;
        }
        // 获取最大值以及索引
        $max = $index = 0;
        for($i = 0;$i < $len;$i++){
            if($nums[$i] > $max){
                $max = $nums[$i];
                $index = $i;
            }
        }
        // 构建节点
        $node = new TreeNode($max);
        $left = array_slice($nums,0,$index);
        $right = array_slice($nums,$index + 1);
        $node->left = $this->constructMaximumBinaryTree($left);
        $node->right = $this->constructMaximumBinaryTree($right);
        return $node;
    }
}