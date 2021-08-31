<?php

// 剑指 Offer 34. 二叉树中和为某一值的路径

// 输入一棵二叉树和一个整数，打印出二叉树中节点值的和为输入整数的所有路径。从树的根节点开始往下一直到叶节点所经过的节点形成一条路径。


//  Definition for a binary tree node.
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

// 未完
class Solution {

    /**
     * @param TreeNode $root
     * @param Integer $target
     * @return Integer[][]
     */
    function pathSum($root, $target) {
        if($root == null){
            return [];
        }
    }

    function dfs($root,$target,&$ret){
        if($root == null){
            return;
        }
        $ret[] = $root->val;
        $this->dfs($root->left,$target,$ret);
        $this->dfs($root->right,$target,$ret);

    }
}
