package main

import (
	"fmt"
	"testing"
)

type Node struct {
	Val int
	Left *Node
	Right *Node
	Next *Node
}

var root = []int{1,2,3,4,5,6,7}

func TestConnect(t *testing.T) {
	node := perfectTree(root)
	fmt.Println(node.Next.Val)
}

func main()  {
	node := perfectTree(root)
	//fmt.Printf("%v",node)
	fmt.Println(node)
}

func perfectTree(arr []int) *Node {
	var node = &Node{
		Val: 1,
		Left: &Node{
			Val: 2,
			Left: &Node{
				Val:   4,
				Left:  nil,
				Right: nil,
				Next:  nil,
			},
			Right: &Node{
				Val:   5,
				Left:  nil,
				Right: nil,
				Next:  nil,
			},
			Next: nil,
		},
		Right: &Node{
			Val: 3,
			Left: &Node{
				Val:   6,
				Left:  nil,
				Right: nil,
				Next:  nil,
			},
			Right: &Node{
				Val:   7,
				Left:  nil,
				Right: nil,
				Next:  nil,
			},
			Next: nil,
		},
		Next: nil,
	}
	return node
}


