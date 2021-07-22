<?php

// 98. Validate Binary Search Tree
// Given the root of a binary tree, determine if it is a valid binary search tree (BST).
//
//A valid BST is defined as follows:
//
//The left subtree of a node contains only nodes with keys less than the node's key.
//The right subtree of a node contains only nodes with keys greater than the node's key.
//Both the left and right subtrees must also be binary search trees.


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

// 若左子树不为空，则左子树上所有节点的值均小于根节点值
// 若右子树不为空，则右子树上所有节点值均大于根节点值
// 左右子树也要为二叉搜索树
// 有问题，待解决
class Solution {

    /**
     * @param TreeNode $root
     * @return Boolean
     */
    function isValidBST($root) {
        if($root == null){
            return true;
        }
        return $this->traverse($root,null,null);
    }

    /**
     * @param $root TreeNode
     * @param $max TreeNode
     * @param $min TreeNode
     */
    function traverse($root,$max,$min){
        if($root == null){
            return true;
        }
        var_dump($max,$min,$root->val);
        if($max && $root->val >= $max->val){
            return false;
        }
        if($min && $root->val <= $min->val){
            return false;
        }
        return $this->traverse($root->left,$min,$root) && $this->traverse($root->right,$root,$max);
    }

    function node1(){
        $node = new TreeNode(10);
        $node->left = new TreeNode(5);
        $node->right = new TreeNode(15);
        $node->right->left = new TreeNode(6);
        $node->right->right = new TreeNode(20);
        return $node;
    }

    function node2(){
        $node = new TreeNode(4);
        $node->left = new TreeNode(2);
        $node->right = new TreeNode(5);
        $node->left->left = new TreeNode(1);
        $node->left->right = new TreeNode(3);
        $node->right->right = new TreeNode(6);
        return $node;
    }

    function node3(){
        $node = new TreeNode(2);
        $node->left = new TreeNode(2);
        $node->right = new TreeNode(2);
        return $node;
    }
}

$obj = new Solution();
$node = $obj->node2();

var_dump($obj->isValidBST($node));