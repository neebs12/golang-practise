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
func (b *BinaryTree) Insert(i int, currNode *Node) {
	// this section inserts values in the binary tree
	// insertion begins at Root, then recursively tails down to the leaves
	// where a Right or Left reference is inserted

	// super local search guard will be here
	if currNode.Key == i {
		return // terminate, no insertion required
	}

	for currNode != nil {
		currKey := currNode.Key
		if currKey < i {
			// traversal to right
			rightNode := currNode.Right
			if rightNode == nil {
				currNode.Right = &Node{Key: i}
			} else {
				b.Insert(i, rightNode)
			}
			return
		} else if currKey > i {
			// traversal to left
			leftNode := currNode.Left
			if leftNode == nil {
				currNode.Left = &Node{Key: i}
			} else {
				b.Insert(i, leftNode)
			}
			return
		}
		// case for if i already exists? use Search
	}

}

// BinaryTree, Search, searches for a value in the binary tree
func (b *BinaryTree) Search(i int, currNode *Node) bool {

	if currNode == nil {
		return false // none is found
	}

	if currNode.Key == i {
		return true // a node is found!
	}

	currKey := currNode.Key
	if currKey < i {
		// traverse right
		rightNode := currNode.Right
		return b.Search(i, rightNode)
	} else if currKey > i {
		// traverse left
		leftNode := currNode.Left
		return b.Search(i, leftNode)
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
	tree.Insert(101, tree.Root)
	tree.Insert(99, tree.Root)
	tree.Insert(97, tree.Root)
	tree.Insert(98, tree.Root)
	tree.Insert(73, tree.Root)
	tree.Insert(200, tree.Root)
	tree.Insert(250, tree.Root)
	tree.Insert(150, tree.Root)
	tree.Insert(400, tree.Root)
	fmt.Println(tree)
	fmt.Println(400, tree.Search(400, tree.Root))
	fmt.Println(97, tree.Search(97, tree.Root))
	fmt.Println(49, tree.Search(49, tree.Root))
	fmt.Println(96, tree.Search(96, tree.Root))
}
