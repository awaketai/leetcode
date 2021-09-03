<?php

// 剑指 Offer 32 - I. 从上到下打印二叉树

// 从上到下打印出二叉树的每个节点，同一层的节点按照从左到右的顺序打印。


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
     * @return Integer[]
     */
    function levelOrder($root) {
        if($root == null){
            return [];
        }
        $res = [];
        $queue = [$root];
        while (count($queue) > 0){
            $node = $queue[0];
            $queue = array_slice($queue,1);
            $res[] = $node->val;
            if($node->left){
                $queue[] = $node->left;
            }
            if($node->right){
                $queue[] = $node->right;
            }
        }
        return $res;
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
$ret = $obj->levelOrder($obj->node());
var_dump($ret);
