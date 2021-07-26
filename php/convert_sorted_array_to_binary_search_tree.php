<?php

// 108.Convert Sorted Array to Binary Search Tree

// Given an integer array nums where the elements are sorted in ascending order, convert it to a height-balanced binary search tree.
//
//A height-balanced binary tree is a binary tree in which the depth of the two subtrees of every node never differs by more than one.

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
     * @param Integer[] $nums
     * @return TreeNode
     */
    function sortedArrayToBST($nums) {

        return $this->helper($nums,0,count($nums) - 1);
    }

    function helper($nums,$left,$right){
        if($left > $right){
            return null;
        }
        // 二叉树为有序数列，且为高度平衡二叉搜索树，选取中位数为根节点
        $mid = floor(($left + $right ) / 2);
        $root = new TreeNode($nums[$mid]);
        $root->left = $this->helper($nums,$left,$mid - 1);
        $root->right = $this->helper($nums,$mid + 1,$right);
        return $root;
    }

    function test($nums){
        // 生成的二叉树
        $root = $this->sortedArrayToBST($nums);
        var_dump($root);
        $mem = [];
        $this->traverse($root,$mem);
        // 生成的二叉树和原有的值是否相等
        foreach ($mem as $ko => $item) {
            if($nums[$ko] != $item){
                echo sprintf("test failed:element value:%d,target value:%d",$item,$nums[$ko]);
                return;
            }
        }
        echo "test success";
    }

    function traverse($root,&$mem){
        if($root == null){
            return null;
        }
        // 中序遍历
        $this->traverse($root->left,$mem);
        $mem[] = $root->val;
        $this->traverse($root->right,$mem);
        return $mem;
    }
}
$nums = [-10,-3,0,5,9];
$nums = [1,3];
$obj = new Solution();
$obj->test($nums);