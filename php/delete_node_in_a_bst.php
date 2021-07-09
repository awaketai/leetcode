<?php

// 450. Delete Node in a BST
// Given a root node reference of a BST and a key, delete the node with the given key in the BST. Return the root node reference (possibly updated) of the BST.
//
//Basically, the deletion can be divided into two stages:
//
//Search for a node to remove.
//If the node is found, delete the node.

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
     * @param Integer $key
     * @return TreeNode
     */
    function deleteNode($root, $key) {
        if($root == null){
            return null;
        }
        if ($root->val == $key){
            // 1.节点是末端节点，可以直接删除
            // 2.节点下只有一个子节点，则让子节点接替自己
            if($root->left == null){
                return $root->right;
            }
            if($root->right == null){
                return $root->left;
            }
            // 3.节点有两个子节点，则让右子树的最小节点(或者左子树的最大节点)来接替自己
            $minNode = $this->getMin($root->right);
            $root->val = $minNode->val;
            // 将右子树
            $root->right = $this->deleteNode($root->right,$minNode->val);

        }elseif($root->val > $key){
            // 去左边找
            $root->left = $this->deleteNode($root->left,$key);
        }elseif($root->val < $key){
            // 去右边找
            $root->right = $this->deleteNode($root->right,$key);
        }
        return $root;
    }

    /**
     * 获取bst最小节点值，左子树末端值
     * @param $root TreeNode
     * @return TreeNode
     */
    function getMin($root){
        while ($root->left != null){
            $root = $root->left;
        }
        return $root;
    }
}