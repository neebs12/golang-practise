package main

import "fmt"

type BinarySearchTree struct {
	root *node
}

// Insertion - Insert
// -- putting new values in bst
// -- assume that the tree has already been initilialized
// -- and that currNode != nil at all cases
func Insert(data int, currNode *node) {
	newNode := &node{data: data}
	currData := currNode.data
	if currData < data {
		// the current data is greater than currNode, see right node
		rightNode := currNode.rightChild
		if rightNode == nil {
			// if nil, insert the newNode to .rightChild
			currNode.rightChild = newNode
			return
		} else {
			// else, is not nil, therefore need to traverse (recur)
			// -- subproblem
			Insert(data, rightNode)
		}
	} else if currData > data {
		leftNode := currNode.leftChild
		if leftNode == nil {
			currNode.leftChild = newNode
			return
		} else {
			// -- subproblem at left node
			Insert(data, leftNode)
		}
	}
	// else, here found a (==) which is NOT inserted
}

// Search
// -- looking for new values in bst
func Search(data int, currNode *node) bool {
	currData := currNode.data
	if currData < data {
		// traverse right
		rightNode := currNode.rightChild
		if rightNode == nil {
			return false
		}

		return Search(data, rightNode)

	} else if currData > data {
		// traverse left
		leftNode := currNode.leftChild
		if leftNode == nil {
			return false
		}

		return Search(data, leftNode)

	} else {
		// found (==), therefore return node
		return true
	}
}

// Deletion
// -- deleting from bst (omg)

type node struct {
	data       int // or any type, as long as they are comparable
	leftChild  *node
	rightChild *node
}

func Init(initialData int) *BinarySearchTree {
	b := &BinarySearchTree{
		root: &node{
			data: initialData,
		},
	}
	return b
}

func main() {
	b := Init(100)
	Insert(150, b.root)
	Insert(125, b.root)
	Insert(175, b.root)
	Insert(200, b.root)
	Insert(112, b.root)
	Insert(135, b.root)
	Insert(50, b.root)
	Insert(25, b.root)
	Insert(12, b.root)
	Insert(35, b.root)
	Insert(75, b.root)
	Insert(62, b.root)
	Insert(82, b.root)
	fmt.Println(82, Search(82, b.root))
	fmt.Println(112, Search(112, b.root))
	fmt.Println(201, Search(201, b.root))
}
