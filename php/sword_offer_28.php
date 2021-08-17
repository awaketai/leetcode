<?php

// 剑指 Offer 28. 对称的二叉树

// 请实现一个函数，用来判断一棵二叉树是不是对称的。如果一棵二叉树和它的镜像一样，那么它是对称的。
//
//例如，二叉树 [1,2,2,3,4,4,3] 是对称的。
//
//    1
//   / \
//  2   2
// / \ / \
//3  4 4  3
//但是下面这个 [1,2,2,null,3,null,3] 则不是镜像对称的:
//
//    1
//   / \
//  2   2
//   \   \
//   3    3
//


//  Definition for a binary tree node.
  class TreeNode {
      public $val = null;
      public $left = null;
      public $right = null;
      function __construct($value) { $this->val = $value; }
  }

class Solution {

    /**
     * @param TreeNode $root
     * @return Boolean
     */
    function isSymmetric($root) {
        if($root == null){
            return true;
        }
        return $this->traverse($root->left,$root->right);
    }

    function traverse($left,$right){
        if($left == null && $right == null){
            return true;
        }
        if($left == null || $right == null){
            return false;
        }
        if($left->val != $right->val){
            return false;
        }
        return $this->traverse($left->left,$right->right) &&
            $this->traverse($left->right,$right->left);
    }
}


