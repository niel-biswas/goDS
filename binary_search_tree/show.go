package main

import "fmt"

// Print the tree in-order
// Traverse the left sub-tree, root, right sub-tree
func showInOrder(root *Node) {
	if root != nil {
		showInOrder(root.Left)
		fmt.Println(root.Value)
		showInOrder(root.Right)
	}
}

// Print the tree pre-order
// Traverse the root, left sub-tree, right sub-tree
func showPreOrder(root *Node) {
	if root != nil {
		fmt.Println(root.Value)
		showInOrder(root.Left)
		showInOrder(root.Right)
	}
}

// Print the tree post-order
// Traverse left sub-tree, right sub-tree, root
func showPostOrder(root *Node) {
	if root != nil {
		showInOrder(root.Left)
		showInOrder(root.Right)
		fmt.Println(root.Value)
	}
}
