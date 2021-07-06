<?php

// 652. Find Duplicate Subtrees
// Given the root of a binary tree, return all duplicate subtrees.
//
//For each kind of duplicate subtrees, you only need to return the root node of any one of them.
//
//Two trees are duplicate if they have the same structure with the same node values.

// 给定一棵二叉树，返回所有重复的子树。对于同一类的重复子树，你只需要返回其中任意一棵的根结点即可。
//
//两棵树重复是指它们具有相同的结构以及相同的结点值。
//   Definition for a binary tree node.
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
     * @return TreeNode[]
     */
    function findDuplicateSubtrees($root) {
        $tmp = $res = [];
        $this->traverse($root,$tmp,$res);
        return $res;
    }

    function traverse($root,&$tmp,&$res){
        if(empty($root)){
            return '#';
        }
        // 递归将节点拼接成字符串形式，存入hash中
        $left = $this->traverse($root->left,$tmp,$res);
        $right = $this->traverse($root->right,$tmp,$res);
        $str = $left .','.$right.','.$root->val;
        $freq = isset($tmp[$str]) ? $tmp[$str] : 0;
        // 如果是第一次加入，则将该节点放入队列
        if($freq == 1){
            $res[] = $root;
        }
        // 如果不存在，则将某个节点的字符串加入数组
        $tmp[$str] = ++$freq;
        return $str;
    }

    function node(){
        $node = new TreeNode(1);
        $node->left = new TreeNode(2);
        $node->right = new TreeNode(3);
        $node->left->left = new TreeNode(4);
        $node->right->left = new TreeNode(2);
        $node->right->left->left = new TreeNode(4);
        $node->right->right = new TreeNode(4);
        return $node;
    }
}

$obj = new Solution();
$node = $obj->node();
$obj->findDuplicateSubtrees($node);

