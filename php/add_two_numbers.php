<?php

// 2.You are given two non-empty linked lists representing two non-negative integers. The digits are stored in reverse order, and each of their nodes contains a single digit. Add the two numbers and return the sum as a linked list.
//
//You may assume the two numbers do not contain any leading zero, except the number 0 itself.


// Definition for a singly-linked list.
class ListNode {
    public $val = 0;
    public $next = null;
    function __construct($val = 0, $next = null) {
        $this->val = $val;
        $this->next = $next;
    }
}

class Solution {

    /**
     * @param ListNode $l1
     * @param ListNode $l2
     * @return ListNode
     */
    function addTwoNumbers($l1, $l2) {
        $ret = new ListNode();
        $res = $ret;
        $add = 0; // 进位数
        $carry = false;
        // 2,4,3
        // 5,6,4
        while ($l1 != null || $l2 != null){
            $val1 = $val2 = 0;
            if ($l1 != null){
                $val1 = $l1->val;
                $l1 = $l1->next;
            }
            if ($l2 != null){
                $val2 = $l2->val;
                $l2 = $l2->next;
            }

            $sum = $val1 + $val2 + $add;
            $residue = $sum % 10;// 当前位置值
            $carry = $sum >= 10 ? true : false;
            $add = $sum >= 10 ? 1 : 0;

            $ret->next = new ListNode($residue);
            $ret = $ret->next;
        }
        if ($carry) {
            $ret->next = new ListNode(1);
        }
        return $res->next;
    }
}