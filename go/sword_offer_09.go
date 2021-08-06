package main

import "container/list"

// 剑指 Offer 09. 用两个栈实现队列

// 用两个栈实现一个队列。队列的声明如下，请实现它的两个函数 appendTail 和 deleteHead ，分别完成在队列尾部插入整数和在队列头部删除整数的功能。(若队列中没有元素，deleteHead 操作返回 -1 )
//

// 第一个栈的底部是最后插入的元素,顶部是待删除的元素
//
type CQueue struct {

	stack1,stack2 *list.List
}

func Constructor() CQueue {
	return CQueue{
		stack1: list.New(),
		stack2: list.New(),
	}
}

// 队尾插入元素
func (this *CQueue) AppendTail(value int)  {

	this.stack1.PushBack(value)
}

// 先插入的先删除
// 1.当栈2不为空，2中仍有已完成倒序的元素，直接返回栈2的元素
// 2.当栈1为空，两个栈都为空，返回-1
// 3.栈1不为空，将栈1元素全部移到栈2中，实现元素倒序，并返回栈2的栈顶元素
func (this *CQueue) DeleteHead() int {
	if this.stack2.Len() == 0 {
		for this.stack1.Len() > 0 {
			this.stack2.PushBack(this.stack1.Remove(this.stack1.Back()))
		}
	}
	if this.stack2.Len() != 0 {
		e := this.stack2.Back()
		this.stack2.Remove(e)
		return e.Value.(int)
	}
	return -1
}
