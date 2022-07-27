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
// -- the return is a *node type, however this is SOLELY for facilitating the recursive algorithm
// -- *node that is returned is NOT actually used
func Delete(dataToDelete int, currNode *node) *node {
	// base case, if the current node is nil, return nil
	if currNode == nil {
		return nil

		// if less, travese left
		// -- subproblem recur with Delete
		// -- this evokes no change, unless the following subproblem (following recur) actually involves a Deletion (else if == triggered clause)
		// -- therefore, with subproblem recur call, it is important that the actual deletion tree change is captured via: `currNode.leftChild = ...`
	} else if currNode.data > dataToDelete {
		currNode.leftChild = Delete(dataToDelete, currNode.leftChild)
		return currNode

		// if more, travese right
		// -- subproblem recur with Delete
		// -- saa for description
	} else if currNode.data < dataToDelete {
		currNode.rightChild = Delete(dataToDelete, currNode.rightChild)
		return currNode

		// equal,
		// -- keep in mind that here, currNode is the NODE TO DELETE!
		// -- actual deletion, sub problem with Delete no longer required
		// -- imagine, we have arrived at the node to be deleted, consider multiple cases (three)
		// -- no children (no left, or right) -- case 1
		// -- one child (left OR right) -- case 2
		// -- both children (left and right)
	} else if currNode.data == dataToDelete {

		// case 1 & case 2 (left absent, right can be absent or present)
		// -- returns .rightChild, nil or *node
		// -- by returning, recall what Delete subprob-call is returning to!!!
		// -- this means that either .leftChild or .rightChild of the parent is being truly overridden!!
		if currNode.leftChild == nil {
			return currNode.rightChild

			// case 2 (left present, right is absent)
			// -- return the .leftChild which is *node
			// -- note that the return here will override whichever Delete subprob was called on
		} else if currNode.rightChild == nil {
			return currNode.leftChild

			// case 3 (left and right are present)
			// -- returns whichever Lift call deems appropriate
			// -- remember that if the node to be deleted has two children, we go immediately on the rhs child and
		} else {
			currNode.rightChild = Lift(currNode.rightChild, currNode)
			return currNode
		}
	}

	return nil
}

func Lift(currNode, nodeToDelete *node) *node {
	// keep recur down the leftChildren until we hit a node with a nil left
	// this is easier to understand (ish)
	if currNode.leftChild != nil {
		currNode.leftChild = Lift(currNode.leftChild, nodeToDelete) // parent
		return currNode

		// here, we hit a node that no longer has a .leftChild
		// -- therefoer at this location, the successornNode IS currNode
		// -- we then place the currNode's data within nodeToDelete
		// -- the return is the currNode.rightChild
		// -- think of the prev-call the .rightChild is then placed at the .leftChild of the subprob call of lift!
	} else {
		nodeToDelete.data = currNode.data
		// <-- why assign to nodeToDelete.data?
		// keep in mind that here, nodeToDelete is at the 'correct' location in the tree that references two children (therefore cannot be moved)
		// therefore, by reassg' nodeToDelete.data, we are replacing its data with the successor node's

		// then, now that the successor node's data is preserved, we can override the successor node's location by overriding the parent's .leftChild
		return currNode.rightChild
	}
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
	fmt.Println(150, Search(150, b.root))
	fmt.Println(201, Search(201, b.root))
	Delete(150, b.root)
	fmt.Println(150, Search(150, b.root))
}
