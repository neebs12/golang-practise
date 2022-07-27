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

// Inorder Bst Traversal - depth first, logs in ascending order
func InOrderBSTTraversal(currNode *node) {
	if currNode == nil {
		return // na, base case
	}

	// traverse left
	InOrderBSTTraversal(currNode.leftChild)

	// then print
	fmt.Println(currNode.data)

	// then traverse right
	InOrderBSTTraversal(currNode.rightChild)
}

// Preorder bst traversal means that we are essentially logging the path of the traversal (Println prior to the subproblem calls)
func PreOrderBSTTraversal(currNode *node) {
	if currNode == nil {
		return
	}

	fmt.Println(currNode.data)

	PreOrderBSTTraversal(currNode.leftChild)
	PreOrderBSTTraversal(currNode.rightChild)
}

// Postorder bst traversal
// -- this is logging the unwinding of the stack??
func PostOrderTraversal(currNode *node) {
	if currNode == nil {
		return
	}

	PostOrderTraversal(currNode.leftChild)
	PostOrderTraversal(currNode.rightChild)
	fmt.Println(currNode.data)
}

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

	return b
}

func main() {
	b := Init(100)

	// InOrderBSTTraversal(b.root)
	// PreOrderBSTTraversal(b.root)
	PostOrderTraversal(b.root)
}
