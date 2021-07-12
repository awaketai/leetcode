<?php

// 701. Insert into a Binary Search Tree

// You are given the root node of a binary search tree (BST) and a value to insert into the tree. Return the root node of the BST after the insertion. It is guaranteed that the new value does not exist in the original BST.
//
//Notice that there may exist multiple valid ways for the insertion, as long as the tree remains a BST after insertion. You can return any of them.

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

class Solution {

    /**
     * @param TreeNode $root
     * @param Integer $val
     * @return TreeNode
     */
    function insertIntoBST($root, $val) {
        if($root == null){
            return new TreeNode($val);
        }
        if($root->val == $val){
            return $root;
        }
        if($root->val > $val){
            // 给左边增加值
            $root->left = $this->insertIntoBST($root->left,$val);
        }else{
            // 给右边增加值
            $root->right = $this->insertIntoBST($root->right,$val);
        }
        return $root;
    }
}