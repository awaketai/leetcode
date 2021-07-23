<?php

// 95. Unique Binary Search Trees II
// Given an integer n, return all the structurally unique BST's (binary search trees), which has exactly n nodes of unique values from 1 to n. Return the answer in any order.


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
     * @param Integer $n
     * @return TreeNode[]
     */
    function generateTrees($n) {
        $ret =  $this->build(1,$n);
        print_r($ret);
    }

    function build($low,$high){
        if($low > $high){
            return [new TreeNode()];
        }
        $node = [];
        // 穷举根节点所有情况
        for($i = $low;$i <= $high;$i++){
            // 左右子树所有情况
            $leftNode = $this->build($low,$i - 1);
            $rightNode = $this->build($i + 1,$high);
            // 补充根节点
            foreach ($leftNode as $lv){
                foreach ($rightNode as $rv){
                    $root = new TreeNode($i); // 根节点
                    $root->left = $lv;
                    $root->right = $rv;
                    $node[] = $root;
                }
            }
        }
        return $node;
    }
}

$obj = new Solution();
$obj->generateTrees(3);