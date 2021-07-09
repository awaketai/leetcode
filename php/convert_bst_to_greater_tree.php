<?php


// 538. Convert BST to Greater Tree
// Given the root of a Binary Search Tree (BST), convert it to a Greater Tree such that every key of the original BST is changed to the original key plus sum of all keys greater than the original key in BST.
// Definition for a binary tree node.
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
    function convertBST($root) {
        $sum = 0;
        $this->traverse($root,$sum);
        return $root;
    }

    function traverse($root,&$sum){
        if(empty($root)){
            return null;
        }
        // 将中序遍历反过来，由于已经为降序排列，因此每个节点的值是当前节点 + 之前所有节点值
        $this->traverse($root->right,$sum);
        $sum += $root->val;
        $root->val = $sum;
        $this->traverse($root->left,$sum);
    }

    function node(){
        $node = new TreeNode(4);
        $node->left = new TreeNode(1);
        $node->left->left = new TreeNode(0);
        $node->left->right = new TreeNode(2);
        $node->left->right->right = new TreeNode(3);
        $node->right = new TreeNode(6);
        $node->right->left = new TreeNode(5);
        $node->right->right = new TreeNode(7);
        $node->right->right->right = new TreeNode(8);
        return $node;
    }
}
$obj = new Solution();
$node = $obj->node();
$obj->convertBST($node);