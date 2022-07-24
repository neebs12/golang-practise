package main

import "fmt"

// key has the current value
// -- left Node key has value less than current key
// -- right Node key has value greater than current key
// -- if Node either/both references other Nodes, is a parent
// -- if Node neither refereces other Nodes, is a leaf
type Node struct {
	Key   int
	Left  *Node
	Right *Node
}

// Binary tree structure
// -- points to a current root node
type BinaryTree struct {
	Root *Node
}

// BinaryTree, Insert, inserts a value at the binary tree
// func (b *BinaryTree) Insert(i int) {}

// BinaryTree, Search, searches for a value in the binary tree
// func (b *BinaryTree) Search(i int) bool {}

// Init, initializes the Binary Tree with a Node, taking in a value
func Init(k int) *BinaryTree {
	tree := &BinaryTree{
		&Node{Key: k},
	}
	return tree
}

func main() {
	tree := Init(100)
	fmt.Println(tree)
}
