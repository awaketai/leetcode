package main

import "container/list"

// 剑指 Offer 59 - II. 队列的最大值

// 请定义一个队列并实现函数 max_value 得到队列里的最大值，要求函数max_value、push_back 和 pop_front 的均摊时间复杂度都是O(1)。
//
//若队列为空，pop_front 和 max_value 需要返回 -1

type MaxQueue struct {
	stack1 list.List
	stack2 list.List // 单调递减队列
}


func Constructor() MaxQueue {

	return MaxQueue{
		stack1: list.List{},
		stack2: list.List{},
	}
}


func (this *MaxQueue) Max_value() int {

	if this.stack2.Len() == 0 {
		return -1
	}

	return this.stack2.Front().Value.(int)
}


func (this *MaxQueue) Push_back(value int)  {
	// 1.正常队列数据添加
	this.stack1.PushBack(value)
	// 2.单调递减队列
	// 移除比当前元素小的所有元素
	for this.stack2.Len() != 0 && this.stack2.Back().Value.(int) < value {
		this.stack2.Remove(this.stack2.Back())
	}
	// 将当前元素添加到队列尾部
	this.stack2.PushBack(value)
}


func (this *MaxQueue) Pop_front() int {
	if this.stack1.Len() == 0 {
		return -1
	}

	// 1.正常队列元素弹出
	val := this.stack1.Front().Value.(int)
	this.stack1.Remove(this.stack1.Front())
	// 2.单调递减队列中元素
	max := this.stack2.Front().Value.(int)
	if max == val {
		this.stack2.Remove(this.stack2.Front())
	}

	return val
}


/**
 * Your MaxQueue object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Max_value();
 * obj.Push_back(value);
 * param_3 := obj.Pop_front();
 */

//func main()  {
//	obj := Constructor()
//	param1 := obj.Max_value()
//	obj.Push_back(1)
//}