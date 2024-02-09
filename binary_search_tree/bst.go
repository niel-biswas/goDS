package main

import "fmt"

type Node struct {
	Key   int
	Left  *Node
	Right *Node
}

func NewNode(key int) *Node {
	return &Node{Key: key}
}

func (n *Node) Insert(k int) {
	if n.Key < k {
		// move right
		if n.Right == nil {
			n.Right = &Node{Key: k}
		} else {
			n.Right.Insert(k)
		}
	} else if n.Key > k {
		// move left
		if n.Left == nil {
			n.Left = &Node{Key: k}
		} else {
			n.Left.Insert(k)
		}
	}
}

func (n *Node) Search(k int) bool {
	if n == nil {
		return false
	}
	if n.Key < k {
		// move right
		return n.Right.Search(k)
	} else if n.Key > k {
		// move left
		return n.Left.Search(k)
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
