<?php

// 226.Given the root of a binary tree, invert the tree, and return its root.
// 翻转一棵二叉树。
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
     * @param TreeNode $root
     * @return TreeNode
     */
    function invertTree($root) {
        if($root == null){
            return null;
        }
        $right = $this->invertTree($root->right);
        $left = $this->invertTree($root->left);
        $root->left = $right;
        $root->right = $left;
        return $root;
    }
}