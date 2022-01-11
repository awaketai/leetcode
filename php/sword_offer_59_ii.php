<?php

// 剑指 Offer 59 - II. 队列的最大值

// 请定义一个队列并实现函数 max_value 得到队列里的最大值，要求函数max_value、push_back 和 pop_front 的均摊时间复杂度都是O(1)。
//
//若队列为空，pop_front 和 max_value 需要返回 -1
class MaxQueue {

    private $stack1 = [];
    private $stack2 = []; // 最大值队列

    /**
     */
    function __construct() {

    }

    /**
     * @return Integer
     */
    function max_value() {

        if(empty($this->stack2)){
            return -1;
        }

        return $this->stack2[0];
    }

    /**
     * @param Integer $value
     * @return NULL
     */
    function push_back($value) {
        // 1.正常队列添加
        array_push($this->stack1,$value);
        // 2.最大值队列增加：单调递减队列
        while(!empty($this->stack2)){
            $end = array_pop($this->stack2);
            if($end > $value){
                array_push($this->stack2,$end);
                break;
            }
        }
        array_push($this->stack2,$value);

        return null;
    }

    /**
     * @return Integer
     */
    function pop_front() {
        if(empty($this->stack1)){
            return -1;
        }

        $front = array_shift($this->stack1);
        $max = $this->stack2[0];
        if($max == $front){
            array_shift($this->stack2);
        }

        return $front;
    }
}

/**
 * Your MaxQueue object will be instantiated and called as such:
 * $obj = MaxQueue();
 * $ret_1 = $obj->max_value();
 * $obj->push_back($value);
 * $ret_3 = $obj->pop_front();
 */
