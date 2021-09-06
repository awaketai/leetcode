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

    // 剑指 Offer 32 - II. 从上到下打印二叉树 II
    function t($root){
        if($root == null){
            return [];
        }
        $res = [];
        $queue = [$root];
        while (count($queue) > 0){
            $tmp = [];
            for($i = count($queue);$i > 0;$i--){
                $node = $queue[0];
                $queue = array_slice($queue,1);
                $tmp[] = $node->val;
                if($node->left){
                    $queue[] = $node->left;
                }
                if($node->right){
                    $queue[] = $node->right;
                }
            }
            $res[] = $tmp;
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
print_r($ret);
