package main

type Node1 struct {
	Val int
	Left *Node1
	Right *Node1
	Next *Node1
}

func main()  {

}

func connect(root *Node1) *Node1 {
	if root == nil {
		return nil
	}
	connectTwoNode(root.Left,root.Right)
	return root
}

func connectTwoNode(node1,node2 *Node1)  {
	if node1 == nil || node2 == nil {
		return
	}
	// 将传入的两个节点连接
	node1.Next = node2
	// 连接相同父节点的两个子节点
	connectTwoNode(node1.Left,node1.Right)
	connectTwoNode(node2.Left,node2.Right)
	// 连接跨越父节点的两个子节点
	connectTwoNode(node1.Right,node2.Left)
}