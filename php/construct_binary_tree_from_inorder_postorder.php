<?php

// 106. Construct Binary Tree from Inorder and Postorder Traversal
// 106. 从中序与后序遍历序列构造二叉树


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

// 中序遍历：左 -> 根 -> 右
// 后序遍历：左 -> 右 -> 根
$inorder1 = [9,3,15,20,7];
$postorder1 = [9,15,7,20,3];

$inorder2 = [5,2,6,4,7,1,8,3,9];
$postorder2 = [5,6,7,4,2,8,9,3,1];


class Solution {

    /**
     * @param Integer[] $inorder
     * @param Integer[] $postorder
     * @return TreeNode
     */
    function buildTree($inorder, $postorder) {
        if(empty($inorder) || empty($postorder)){
            return null;
        }
        // 确定根节点值
        $rootVal = $postorder[count($postorder) - 1];
        // 根节点值在中序遍历中位置
        $inorderLen = count($inorder);
        $rootIndex = 0;
        for($i = 0;$i < $inorderLen;$i++){
            if($inorder[$i] == $rootVal){
                $rootIndex = $i;
                break;
            }
        }
        // 中序遍历左子树长度和后序遍历左子树长度相同，右子树同理
        // 左子树长度
        $node = new TreeNode($rootVal);
        $node->left = $this->buildTree(array_slice($inorder,0,$rootIndex),array_slice($postorder,0,$rootIndex));
        $node->right = $this->buildTree(array_slice($inorder,$rootIndex + 1),array_slice($postorder,$rootIndex,(count($postorder) - $rootIndex - 1)));
        return $node;
    }
}

$obj = new Solution();
print_r($obj->buildTree($inorder1,$postorder1));