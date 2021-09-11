<?php

// 剑指 Offer 30. 包含min函数的栈

// 定义栈的数据结构，请在该类型中实现一个能够得到栈的最小元素的 min 函数在该栈中，调用 min、push 及 pop 的时间复杂度都是 O(1)。

// 示例:
//
//MinStack minStack = new MinStack();
//minStack.push(-2);
//minStack.push(0);
//minStack.push(-3);
//minStack.min();   --> 返回 -3.
//minStack.pop();
//minStack.top();      --> 返回 0.
//minStack.min();   --> 返回 -2.

class MinStack{

    private $stack;
    private $stack1; // 辅助栈

    public function __construct(){

        $this->stack = [];
        $this->stack1 = [];
    }

    /**
     * @param integer $x
     * @return null
     */
    public function push($x){
        $this->stack[] = $x;
        // 将最小元素添加到辅助栈栈顶
        if(count($this->stack1) == 0 || $this->stack1[count($this->stack1) - 1] >= $x){
            $this->stack1[] = $x;
        }
    }

    /**
     * @return null
     */
    public function pop(){
        $pop = array_pop($this->stack);
        if($pop == $this->stack1[count($this->stack1) - 1]){
            // 如果弹出的元素也在辅助栈中，则将辅助栈中的元素移除
            array_pop($this->stack1);
        }
    }

    /**
     * @return integer
     */
    public function top(){

        return $this->stack[count($this->stack) - 1];
    }

    /**
     * @return integer
     */
    public function min(){

        return $this->stack1[count($this->stack1) - 1];
    }

    public function println(){
        print_r($this->stack);
        print_r($this->stack1);
    }
}

$obj = new MinStack();
$obj->push(-2);
$obj->push(0);
$obj->push(-3);

$ret = $obj->min();
var_dump("min:",$ret);
$obj->pop();
$obj->println();
$ret = $obj->top();
var_dump("top:",$ret);

$ret = $obj->min();
var_dump("min:",$ret);
