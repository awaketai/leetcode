<?php

// 剑指 Offer 26. 树的子结构

// 输入两棵二叉树A和B，判断B是不是A的子结构。(约定空树不是任意一个树的子结构)
//
//B是A的子结构， 即 A中有出现和B相同的结构和节点值。


//  Definition for a binary tree node.
class TreeNode {
  public $val = null;
  public $left = null;
  public $right = null;
  function __construct($value) { $this->val = $value; }
}

class Solution {

    /**
     * @param TreeNode $A
     * @param TreeNode $B
     * @return Boolean
     */
    function isSubStructure($A, $B) {
        if($A == null || $B == null){
            return false;
        }
        // l 1
        // 1.从根节点判断
        // 2.从A的左右子节点判断
        return $this->traverse($A,$B) ||
            $this->isSubStructure($A->left,$B) ||
            $this->isSubStructure($A->right,$B);
    }

    function traverse($A, $B){
        // 如果B为空，说明已经访问完了，确定是A的子级
        if($B == null){
            return true;
        }
        // 如果B不为空，A为空，或者这两个节点值不同，说明B不是A的子集
        if($A == null || $A->val != $B->val){
            return false;
        }
        return $this->traverse($A->left,$B->left) && $this->traverse($A->right,$B->right);
    }

    function nodeA(){
        $node = new TreeNode(3);
        $node->left = new TreeNode(4);
        $node->right = new TreeNode(5);
        $node->left->left = new TreeNode(1);
        $node->left->right = new TreeNode(2);
        return $node;
    }

    function nodeB(){
        $node = new TreeNode(4);
        $node->right = new TreeNode(2);
        return $node;
    }
}

$obj = new Solution();
$ret = $obj->isSubStructure($obj->nodeA(),$obj->nodeB());
var_dump($ret);