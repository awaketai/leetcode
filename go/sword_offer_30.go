package main

import (
	"container/list"
	"fmt"
)

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


type MinStack struct {
	stack1 *list.List
	stack2 *list.List // 辅助栈，存储最小元素值
}

func main()  {
	test3()
	//stack := constructor()
	//stack.Push(-2)
	//stack.Push(0)
	//stack.Push(-3)

	//length := stack.stack1.Len()
	//for i := 0;i < length;i++{
	//	fr := stack.stack1.Front()
	//	fmt.Println(fr.Value)
	//	stack.stack1.Remove(fr)
	//}
	//fmt.Println("********")
	//length1 := stack.stack2.Len()
	//for i := 0;i < length1;i++{
	//	fr := stack.stack2.Front()
	//	fmt.Println(fr.Value)
	//	stack.stack2.Remove(fr)
	//}

	//fmt.Println(stack.Min())
	//stack.Pop()
	//fmt.Println(stack.Top())
	//fmt.Println(stack.Min())
}

func constructor() MinStack  {
	return MinStack{
		stack1: list.New(),
		stack2: list.New(),
	}
}

func (this *MinStack) Push(x int)  {
	this.stack1.PushBack(x)

	if this.stack2.Len() == 0 || this.stack2.Back().Value.(int) >= x {
		this.stack2.PushBack(x)
	}
}

func (this *MinStack) Pop()  {
	// 移除顶端元素
	pop := this.stack1.Back()
	this.stack1.Remove(pop)
	// 移除辅助栈中相同的元素
	el := this.stack2.Back()
	if el.Value.(int) == pop.Value.(int) {
		this.stack2.Remove(el)
	}
}

func (this *MinStack) Top() int {

	return this.stack1.Back().Value.(int)
}

func (this *MinStack) Min() int {

	return this.stack2.Back().Value.(int)
}

// ["MinStack","push","push","push","min","pop","top","min"]
//[[],[-2],[0],[-3],[],[],[],[]]
//["MinStack","push","push","push","min","pop","min"]
//[[],[0],[1],[0],[],[],[]]


// ["MinStack","push","push","push","top","pop","min","pop","min","pop","push","top","min","push","top","min","pop","min"]
//[[],[2147483646],[2147483646],[2147483647],[],[],[],[],[],[],[2147483647],[],[],[-2147483648],[],[],[],[]]

func test3()  {
	// [null,null,null,null,2147483647,null,2147483646,null,2147483646,null,null,2147483647,2147483646,null,-2147483648,-2147483648,null,2147483646] 输出
	// [null,null,null,null,2147483647,null,2147483646,null,2147483646,null,null,2147483647,2147483647,null,-2147483648,-2147483648,null,2147483647] 预期
	stack := constructor()
	stack.Push(2147483646)
	stack.Push(2147483646)
	stack.Push(2147483647)
	fmt.Println(stack.Top())
	stack.Pop()
	fmt.Println(stack.Min())
	stack.Pop()
	fmt.Println(stack.Min())
	stack.Pop()
	stack.Push(2147483647)
	fmt.Println(stack.Top())
	fmt.Println(stack.Min())
	stack.Push(-2147483648)
	fmt.Println(stack.Top())
	fmt.Println(stack.Min())
	fmt.Println(stack.Top())
	fmt.Println(stack.Min())
}