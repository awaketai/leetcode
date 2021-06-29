package main

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}

var d1 = []int{9,9,9,9,9,9,9}
var d2 = []int{9,9,9,9};

var a1 = []int{2,4,3}
var a2 = []int{5,6,4}

var b1 = []int{0}
var b2 = []int{0}

var c1 = []int{9}
var c2 = []int{9}

func main(){
	al1 := gen(a1)
	al2 := gen(a2)
	ret := addTwoNumbers(al1,al2)
	fmt.Println(ret)

}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	ret := &ListNode{}
	res := ret
	add := 0 // 是否向前进位
	// 2,4,3
	// 5,6,4
	for l1 != nil || l2 != nil {
		val1,val2 := 0,0
		if l1 != nil {
			val1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			val2 = l2.Val
			l2 = l2.Next
		}
		sum := val1 + val2 + add
		residue := sum % 10 // 当前位置值
		if sum >= 10 {
			add = 1
		}else{
			add = 0
		}
		ret.Next = &ListNode{Val: residue}
		ret = ret.Next
	}
	if add == 1 {
		ret.Next = &ListNode{Val: 1}
	}
	return res.Next
}


func gen(arr []int) *ListNode {
	cLen := len(arr)
	//fmt.Println(cLen)
	cNode := &ListNode{}
	ret := cNode
	for i := 0;i < cLen;i++ {
		cNode.Next = &ListNode{Val: arr[i]}
		cNode = cNode.Next
	}
	return ret.Next
}
