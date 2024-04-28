package main

import "fmt"

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

func NewNode(value int) *Node {
	return &Node{Value: value}
}

func (n *Node) Insert(v int) {
	if n.Value < v {
		// move right
		if n.Right == nil {
			n.Right = &Node{Value: v}
		} else {
			n.Right.Insert(v)
		}
	} else if n.Value > v {
		// move left
		if n.Left == nil {
			n.Left = &Node{Value: v}
		} else {
			n.Left.Insert(v)
		}
	}
}

func (n *Node) Search(v int) bool {
	if n == nil {
		return false
	}
	if n.Value < v {
		// move right
		return n.Right.Search(v)
	} else if n.Value > v {
		// move left
		return n.Left.Search(v)
	}
	return true
}

func main() {
	bst := NewNode(12)
	bst.Insert(56)
	bst.Insert(66)
	bst.Insert(68)
	bst.Insert(67)
	bst.Insert(87)

	fmt.Println(bst.Search(67))
}
