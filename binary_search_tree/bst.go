package main

import "fmt"

type TreeNode struct {
	Value int
	Left  *TreeNode
	Right *TreeNode
}

func (node *TreeNode) Insert(v int) {
	if node.Value < v {
		// move right
		if node.Right == nil {
			node.Right = &TreeNode{Value: v}
		} else {
			node.Right.Insert(v)
		}
	} else if node.Value > v {
		// move left
		if node.Left == nil {
			node.Left = &TreeNode{Value: v}
		} else {
			node.Left.Insert(v)
		}
	}
}

func (node *TreeNode) Search(v int) bool {
	if node == nil {
		return false
	}
	if node.Value < v {
		// move right
		return node.Right.Search(v)
	} else if node.Value > v {
		// move left
		return node.Left.Search(v)
	}
	return true
}

func (node *TreeNode) Remove(value int) *TreeNode {
	node.remove(value, nil)
	return node
}

func (node *TreeNode) max() int {
	for node.Right != nil {
		node = node.Right
	}
	return node.Value
}
func (node *TreeNode) min() int {
	for node.Left != nil {
		node = node.Left
	}
	return node.Value
}

func (node *TreeNode) remove(value int, parent *TreeNode) {
	// base case
	if node == nil {
		return
	}
	if value < node.Value {
		node.Left.remove(value, node)
	} else if value > node.Value {
		node.Right.remove(value, node)
	} else if value == node.Value {
		// no children
		if node.Left == nil && node.Right == nil {
			if parent == nil {
				return
			} else if parent.Left == node {
				parent.Left = nil
			} else {
				parent.Right = nil
			}
			// one child
		} else if node.Left != nil && node.Right == nil {
			if parent == nil {
				node.Value = node.Left.max()
				node.Left.remove(node.Value, node)
			} else if parent.Left == node {
				parent.Left = node.Left
			} else {
				parent.Right = node.Left
			}
		} else if node.Right != nil && node.Left == nil {
			if parent == nil {
				node.Value = node.Right.min()
				node.Right.remove(node.Value, node)
			} else if parent.Left == node {
				parent.Left = node.Right
			} else {
				parent.Right = node.Right
			}
			// two children
		} else {
			node.Value = node.Right.min()
			node.Right.remove(node.Value, node)
		}
	}
}

func main() {
	bst := &TreeNode{Value: 102}
	bst.Insert(56)
	bst.Insert(66)
	bst.Insert(68)
	bst.Insert(67)
	bst.Insert(87)
	fmt.Println("Print BST (Inorder)")
	showInOrder(bst)
	fmt.Println("Print BST (Preorder)")
	showPreOrder(bst)
	fmt.Println("Print BST (Postorder)")
	showPostOrder(bst)
	fmt.Println("Search 67: ", bst.Search(67))
	fmt.Println("Remove 67", bst.Remove(67))
	fmt.Println("Search 67 again: ", bst.Search(67))
	fmt.Println("Print BST again(Inorder)")
	showInOrder(bst)
	fmt.Println("Print BST again(Preorder)")
	showPreOrder(bst)
	fmt.Println("Print BST again(Postorder)")
	showPostOrder(bst)
}
