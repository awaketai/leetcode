<?php

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

$preorder = [3,9,20,15,7];
$inorder = [9,3,15,20,7];

class Solution {

    /**
     * @param Integer[] $preorder
     * @param Integer[] $inorder
     * @return TreeNode
     */
    function buildTree($preorder, $inorder) {
        if(empty($preorder) || empty($inorder)){
            return null;
        }
        // 根节点值$preorder[0]
        // 根节点在中序遍历中位置
        $rootIndex = 0;
        $len = count($inorder);
        for($i = 0;$i <$len;$i++){
            if($preorder[0] == $inorder[$i]){
                $rootIndex = $i;
                break;
            }
        }
        // 确定左右子树长度
        $node = new TreeNode($preorder[0]);
        $node->left = $this->buildTree(array_slice($preorder,1,$rootIndex),array_slice($inorder,0,$rootIndex));
        $node->right = $this->buildTree(array_slice($preorder,(1 + $rootIndex)),array_slice($inorder,$rootIndex + 1));
        return $node;
    }
}

$obj = new Solution();
var_dump($obj->buildTree($preorder,$inorder));
