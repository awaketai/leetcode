<?php

// 剑指 Offer 55 - I. 二叉树的深度

// 输入一棵二叉树的根节点，求该树的深度。从根节点到叶节点依次经过的节点（含根、叶节点）形成树的一条路径，最长路径的长度为树的深度。


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
     * @return Integer
     */
    function maxDepth($root) {
        if($root == null){
            return 0;
        }
        return max($this->maxDepth($root->left),$this->maxDepth($root->right)) + 1;

    }

    function node(){
        $node = new TreeNode(3);
        $node->left = new TreeNode(9);
        $node->right = new TreeNode(20);
        $node->right->left = new TreeNode(15);
        $node->right->right = new TreeNode(7);
        return $node;
    }
}

$obj = new Solution();
  $ret = $obj->maxDepth($obj->node());
  var_dump($ret);
