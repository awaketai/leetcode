<?php


//  Definition for a binary tree node.
//230. Kth Smallest Element in a BST
//Given the root of a binary search tree, and an integer k, return the kth (1-indexed) smallest element in the tree.
//Input: root = [3,1,4,null,2], k = 1
//Output: 1
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

// 1.二叉搜索树中序遍历结果是升序，因此使用递归，时间复杂度O(N)

class Solution {

    /**
     * @param TreeNode $root
     * @param Integer $k
     * @return Integer
     */
    function kthSmallest($root, $k) {
        $rank = $ret = 0;
        $this->traverse($root,$k,$rank,$ret);
        var_dump($ret);
    }

    function traverse($root,$k,&$rank,&$ret){
        if(empty($root)){
            return;
        }
        // 中序遍历
        $this->traverse($root->left,$k,$rank,$ret);
        $rank++;
        if($rank == $k){
            $ret = $root->val;
            return;
        }
        $this->traverse($root->right,$k,$rank,$ret);
        return;
    }

    function node(){
        $node = new TreeNode(5);
        $node->left = new TreeNode(3);
        $node->left->left = new TreeNode(2);
        $node->left->left->left = new TreeNode(1);
        $node->right = new TreeNode(6);
        $node->left->right = new TreeNode(4);
        return $node;
    }
}
$obj = new Solution();
$node = $obj->node();
$obj->kthSmallest($node,5);

