<?php

//Definition for a binary tree node.
class TreeNode {
    public $val = null;
    public $left = null;
    public $right = null;
    function __construct($value) { $this->val = $value; }
}

// 剑指 Offer 07. 重建二叉树

// 输入某二叉树的前序遍历和中序遍历的结果，请构建该二叉树并返回其根节点。
//
//假设输入的前序遍历和中序遍历的结果中都不含重复的数字。


// 前序遍历从根节点开始
// preorder = [3,9,20,15,7], inorder = [9,3,15,20,7]
class Solution {

    /**
     * @param Integer[] $preorder
     * @param Integer[] $inorder
     * @return TreeNode
     */
    function buildTree($preorder, $inorder) {
        if($preorder == null || $inorder == null){
            return null;
        }

        // 确定根节点值:前序遍历第一个值 $preorder[0]
        // 确定左右子树元素
        // 确定根节点在中序遍历中位置
        $rootIndex = 0;
        $len = count($inorder);
        for($i = 0;$i < $len;$i++){
            if($inorder[$i] == $preorder[0]){
                $rootIndex = $i;
                break;
            }
        }
        // 构建左右节点
        $node = new TreeNode($preorder[0]);
        // 左节点长度：中序遍历根节点左边的所有值
        // 确定左节点
        // 前序遍历和中序遍历中左右节点长度是相等的
        $node->left = $this->buildTree(array_slice($preorder,1,$rootIndex),array_slice($inorder,0,$rootIndex));
        $node->right = $this->buildTree(array_slice($preorder,$rootIndex + 1),array_slice($inorder,1 + $rootIndex));
        return $node;
    }
}

$preorder = [3,9,20,15,7];
$inorder = [9,3,15,20,7];
$obj = new Solution();
$ret = $obj->buildTree($preorder,$inorder);
var_dump($ret);
