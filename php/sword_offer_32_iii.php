<?php

// 剑指 Offer 32 - III. 从上到下打印二叉树 III

// 请实现一个函数按照之字形顺序打印二叉树，即第一行按照从左到右的顺序打印，第二层按照从右到左的顺序打印，第三行再按照从左到右的顺序打印，其他行以此类推。


//  Definition for a binary tree node.
class TreeNode {
    public $val = null;
    public $left = null;
    public $right = null;
    function __construct($value) { $this->val = $value; }
}

class Solution
{

    /**
     * @param TreeNode $root
     * @return Integer[][]
     */
    function levelOrder($root)
    {
        if($root == null){
            return [];
        }
        $ret = [];
        $queue = [$root];
        while (count($queue) > 0){
            $tmp = [];
            for($i = count($queue);$i > 0;$i--){
                $node = $queue[0];
                $queue = array_slice($queue,1);
                if(count($ret) % 2 == 0){
                    // 如果偶数节点，将节点添加到tmp双端队列头部
                    $tmp[] = $node->val;
                }else{
                    // 如果是奇数节点，将节点值添加到tmp双端队列尾部
                    array_unshift($tmp,$node->val);
                }
                // 将左右节点放入队列中
                if($node->left){
                    $queue[] = $node->left;
                }
                if($node->right){
                    $queue[] = $node->right;
                }
            }
            $ret[] = $tmp;
        }
        return $ret;
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