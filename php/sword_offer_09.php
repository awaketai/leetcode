<?php

// 剑指 Offer 09. 用两个栈实现队列

//用两个栈实现一个队列。队列的声明如下，请实现它的两个函数 appendTail 和 deleteHead ，分别完成在队列尾部插入整数和在队列头部删除整数的功能。(若队列中没有元素，deleteHead 操作返回 -1 )

class CQueue {
    private $stack1 = [];
    private $stack2 = [];
    /**
     */
    function __construct() {
        $this->stack1 = [];
        $this->stack2 = [];
    }

    /**
     * @param Integer $value
     * @return NULL
     */
    function appendTail($value) {

        array_push($this->stack1,$value);
    }

    /**
     * @return Integer
     */
    function deleteHead() {
        if(count($this->stack2) == 0){
            while (count($this->stack1) > 0){
                $e = array_shift($this->stack1);
                array_push($this->stack2,$e);
            }
        }
        if(count($this->stack2) > 0){
            return array_shift($this->stack2);
        }
        return -1;
    }
}

/**
 * Your CQueue object will be instantiated and called as such:
 * $obj = CQueue();
 * $obj->appendTail($value);
 * $ret_2 = $obj->deleteHead();
 */

$obj = new CQueue();
//$obj->appendTail(3);
//$ret = $obj->deleteHead();
//var_dump($ret);
//$ret = $obj->deleteHead();
//var_dump($ret);

$ret = $obj->deleteHead();
var_dump($ret);
$obj->appendTail(5);
$obj->appendTail(2);
$ret = $obj->deleteHead();
var_dump($ret);
$ret = $obj->deleteHead();
var_dump($ret);



//$s = [];
//array_push($s,1);
//array_push($s,2);
//array_push($s,3);
//var_dump($s);
//
//array_pop($s);
//var_dump($s);

