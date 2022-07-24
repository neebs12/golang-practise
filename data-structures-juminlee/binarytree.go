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
func (n *Node) Insert(i int) {
	// this section inserts values in the binary tree
	// insertion begins at Root, then recursively tails down to the leaves
	// where a Right or Left reference is inserted

	// super local search guard will be here
	if n.Key == i {
		return // terminate, no insertion required
	}

	currKey := n.Key
	if currKey < i {
		// traversal to right
		rightNode := n.Right
		if rightNode == nil {
			n.Right = &Node{Key: i}
		} else {
			rightNode.Insert(i)
		}
		return
	} else if currKey > i {
		// traversal to left
		leftNode := n.Left
		if leftNode == nil {
			n.Left = &Node{Key: i}
		} else {
			leftNode.Insert(i)
		}
		return
	}
}

// BinaryTree, Search, searches for a value in the binary tree
func (n *Node) Search(i int) bool {

	if n == nil {
		return false // none is found
	}

	if n.Key == i {
		return true // a node is found!
	}

	currKey := n.Key
	if currKey < i {
		// traverse right
		rightNode := n.Right
		return rightNode.Search(i)
	} else if currKey > i {
		// traverse left
		leftNode := n.Left
		return leftNode.Search(i)
	}
	return false
}

// Init, initializes the Binary Tree with a Node, taking in a value
func Init(k int) *BinaryTree {
	tree := &BinaryTree{
		&Node{Key: k},
	}
	return tree
}

func main() {
	tree := Init(100)
	tree.Root.Insert(101)
	tree.Root.Insert(99)
	tree.Root.Insert(97)
	tree.Root.Insert(98)
	tree.Root.Insert(73)
	tree.Root.Insert(200)
	tree.Root.Insert(250)
	tree.Root.Insert(150)
	tree.Root.Insert(400)
	fmt.Println(tree)
	fmt.Println(400, tree.Root.Search(400))
	fmt.Println(97, tree.Root.Search(97))
	fmt.Println(49, tree.Root.Search(49))
	fmt.Println(96, tree.Root.Search(96))
}
