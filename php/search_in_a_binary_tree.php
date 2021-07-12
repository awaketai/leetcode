<?php


// 700. Search in a Binary Search Tree
// You are given the root of a binary search tree (BST) and an integer val.

//   Find the node in the BST that the node's value equals val and return the subtree rooted with that node.
//  If such a node does not exist, return null.

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
     * @param Integer $val
     * @return TreeNode
     */
    function searchBST($root, $val) {
        if($root == null){
            return null;
        }
        if($root->val == $val){
            return $root;
        }
        if($root->val > $val){
            // 去左边找
            return $this->searchBST($root->left,$val);
        }else{
            // 去右边找
            return $this->searchBST($root->right,$val);
        }
    }
}