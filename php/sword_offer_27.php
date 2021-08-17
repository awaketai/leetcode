<?php

// 剑指 Offer 27. 二叉树的镜像

// 请完成一个函数，输入一个二叉树，该函数输出它的镜像。
//
//例如输入：
//
//     4
//   /   \
//  2     7
// / \   / \
//1   3 6   9
//镜像输出：
//
//     4
//   /   \
//  7     2
// / \   / \
//9   6 3   1
//


//  Definition for a binary tree node.
  class TreeNode {
      public $val = null;
      public $left = null;
      public $right = null;
      function __construct($value) { $this->val = $value; }
  }
  // // 递归的对树进行便来，从叶子节点开始翻转得到镜像
//// 如果当前遍历到的节点root的左右子节点都已经翻转得到镜像
//// 那么只需要交换两颗子树的位置，既可以得到以root为根节点的整颗子树的镜像
class Solution {

    /**
     * @param TreeNode $root
     * @return TreeNode
     */
    function mirrorTree($root) {
        if($root == null){
            return null;
        }
        $left = $this->mirrorTree($root->left);
        $right = $this->mirrorTree($root->right);
        $root->left = $right;
        $root->right = $left;
        return $root;
    }
}
