<?php

// 剑指 Offer 06. 从尾到头打印链表
//输入一个链表的头节点，从尾到头反过来返回每个节点的值（用数组返回）。

// 输入：head = [1,3,2]
// 输出：[2,3,1]


//Definition for a singly-linked list.
class ListNode {
    public $val = 0;
    public $next = null;
    function __construct($val) { $this->val = $val; }
}

class Solution {

    /**
     * @param ListNode $head
     * @return Integer[]
     */
    function reversePrint($head) {
        $ret = [];
        while ($head != null){
            $ret[] = $head->val;
            $head = $head->next;
        }
        return array_reverse($ret);
    }

    function node(){
        $node = new ListNode(1);
        $node->next = new ListNode(3);
        $node->next->next = new ListNode(2);
        return $node;
    }
}

$obj = new Solution();
$node = $obj->node();
$ret = $obj->reversePrint($node);
var_dump($ret);
