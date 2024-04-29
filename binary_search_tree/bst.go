package main

import "fmt"

type Node struct {
	Value int
	Left  *Node
	Right *Node
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

func (tree *Node) Remove(value int) *Node {
	tree.remove(value, nil)
	return tree
}

func (tree *Node) max() int {
	for tree.Right != nil {
		tree = tree.Right
	}
	return tree.Value
}
func (tree *Node) min() int {
	for tree.Left != nil {
		tree = tree.Left
	}
	return tree.Value
}

func (tree *Node) remove(value int, parent *Node) {
	// base case
	if tree == nil {
		return
	}
	if value < tree.Value {
		tree.Left.remove(value, tree)
	} else if value > tree.Value {
		tree.Right.remove(value, tree)
	} else if value == tree.Value {
		// no children
		if tree.Left == nil && tree.Right == nil {
			if parent == nil {
				return
			} else if parent.Left == tree {
				parent.Left = nil
			} else {
				parent.Right = nil
			}
			// one child
		} else if tree.Left != nil && tree.Right == nil {
			if parent == nil {
				tree.Value = tree.Left.max()
				tree.Left.remove(tree.Value, tree)
			} else if parent.Left == tree {
				parent.Left = tree.Left
			} else {
				parent.Right = tree.Left
			}
		} else if tree.Right != nil && tree.Left == nil {
			if parent == nil {
				tree.Value = tree.Right.min()
				tree.Right.remove(tree.Value, tree)
			} else if parent.Left == tree {
				parent.Left = tree.Right
			} else {
				parent.Right = tree.Right
			}
			// two children
		} else {
			tree.Value = tree.Right.min()
			tree.Right.remove(tree.Value, tree)
		}
	}
}

func main() {
	bst := &Node{Value: 12}
	bst.Insert(56)
	bst.Insert(66)
	bst.Insert(68)
	bst.Insert(67)
	bst.Insert(87)
	fmt.Println("Print BST")
	showInOrder(bst)
	fmt.Println("Search 67: ", bst.Search(67))
	fmt.Println("Remove 67", bst.Remove(67))
	fmt.Println("Search 67 again: ", bst.Search(67))
	fmt.Println("Print BST again")
	showInOrder(bst)
}
