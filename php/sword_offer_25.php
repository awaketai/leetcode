<?php

// 剑指 Offer 25. 合并两个排序的链表

// 输入两个递增排序的链表，合并这两个链表并使新链表中的节点仍然是递增排序的。



//  Definition for a singly-linked list.
  class ListNode {
      public $val = 0;
      public $next = null;
      function __construct($val) { $this->val = $val; }
  }

class Solution {

    /**
     * @param ListNode $l1
     * @param ListNode $l2
     * @return ListNode
     */
    function mergeTwoLists($l1, $l2) {
        $node = new ListNode(0);
        $current = $node;
        while ($l1 != null && $l2 != null){
            if($l1->val > $l2->val){
                $current->next = $l2;
                $l2 = $l2->next;
            }else{
                $current->next = $l1;
                $l1 = $l1->next;
            }
            $current = $current->next;
        }
        if($l1){
            $current->next = $l1;
        }
        if($l2){
            $current->next = $l2;
        }
        return $node->next;
    }

    function listToSlice($l){
        $s = [];
        while ($l != null){
            $s[] = $l->val;
            $l = $l->next;
        }
        return $s;
    }

    function node1(){
        $node = new ListNode(1);
        $node->next = new ListNode(2);
        $node->next->next = new ListNode(4);
        return $node;
    }

    function node2(){
        $node = new ListNode(1);
        $node->next = new ListNode(3);
        $node->next->next = new ListNode(4);
        return $node;
    }
}

$obj = new Solution();
$ret = $obj->mergeTwoLists($obj->node1(),$obj->node2());
$ret = $obj->listToSlice($ret);
var_dump($ret);