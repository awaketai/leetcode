<?php

// 剑指 Offer 22. 链表中倒数第k个节点

// 输入一个链表，输出该链表中倒数第k个节点。为了符合大多数人的习惯，本题从1开始计数，即链表的尾节点是倒数第1个节点。
//
//例如，一个链表有 6 个节点，从头节点开始，它们的值依次是 1、2、3、4、5、6。这个链表的倒数第 3 个节点是值为 4 的节点。
//



//  Definition for a singly-linked list.
  class ListNode {
      public $val = 0;
      public $next = null;
      function __construct($val) { $this->val = $val; }
  }

  // 双指针法
// 前后指针相距k，等到前指针走到链表结尾，则后指针所在的位置就是要返回的链表

// 1.初始化前、后指针，都指向head
// 2.前指针先走k步，构建前后指针距离
// 3.双指针共同移动，每轮都向前走一步，知道前指针跳出链表
class Solution {

    /**
     * @param ListNode $head
     * @param Integer $k
     * @return ListNode
     */
    function getKthFromEnd($head, $k) {
        // 1.初始化前后指针
        $advance = $latter = $head;
        // 2.前指针先移动k步
        for($i = 1;$i <= $k;$i++){
            $advance = $advance->next;
        }
        // 前后指针共同移动
        while ($advance != null){
            $advance = $advance->next;
            $latter = $latter->next;
        }
        return $latter;
    }

    function node(){
        $node = new ListNode(1);
        $node->next = new ListNode(2);
        $node->next->next = new ListNode(3);
        $node->next->next->next = new ListNode(4);
        $node->next->next->next->next = new ListNode(5);
        return $node;

    }
}

$obj = new Solution();
$obj->getKthFromEnd($obj->node(),2);