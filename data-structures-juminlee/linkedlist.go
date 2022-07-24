package main

import (
	"fmt"
	"strings"
)

// making a linked list

// Node
// contains `key`, information inside Node
// contains `next`, references next Node
type Node struct {
	key  string
	next *Node // ptr for identical struct
}

// LinkedList
// contains `head` ref *Node
type LinkedList struct {
	head *Node
}

// LinkedList Prepend method (adds node at the start of linked list, given data)
func (l *LinkedList) Prepend(k string) {
	// is a ptr receiver as we want to mutate receiver
	// init the new Node to `newNode`
	// get a reference to the next value `nextNode`
	// reassg head to *newNode
	// reassg newNode.next to `nextNode`

	newNode := &Node{key: k} // <-- can ignore `next` field?, is default zero value?
	nextNode := l.head
	l.head = newNode
	newNode.next = nextNode
}

// LinkedList Search method (searches linked list, given data) - returns bool (true if found, false if not)
func (l LinkedList) Search(k string) bool {
	currNode := l.head
	for currNode != nil {
		if currNode.key == k {
			return true
		}
		currNode = currNode.next
	}
	return false
}

// LinkedList Delete method (deletes node from linked list, given data)
func (l *LinkedList) Delete(k string) {
	// head deletion
	prevNode := l.head
	if prevNode.key == k {
		l.head = prevNode.next
		return
	}

	for prevNode.next != nil {
		if prevNode.next.key == k {
			// deletion occurs here
			// tail deletion
			prevNode.next = prevNode.next.next // <-- chain broken
			return
		}
		prevNode = prevNode.next
	}
}

// LinkedList, Print, prints the values in the linked list in proper succession
func (l LinkedList) Print() {
	// construct a sli of strings
	sli := make([]string, 0)
	currNode := l.head
	for currNode != nil {
		sli = append(sli, currNode.key)

		// increment
		currNode = currNode.next
	}

	fmt.Println(strings.Join(sli, " <--- "))
}

// func Init, initializes LinkedList
func Init(initialValue string) *LinkedList {
	node := new(Node) // ptr to node
	node.key = initialValue
	linkedList := new(LinkedList)
	linkedList.head = node
	return linkedList
}

func main() {
	linkedList := Init("Jason")
	linkedList.Prepend("Jemimah")
	linkedList.Prepend("Joseph")
	linkedList.Prepend("Janina")
	linkedList.Prepend("Jaci")
	linkedList.Print()
	fmt.Println("Jaci", linkedList.Search("Jaci"))
	fmt.Println("Geraldine", linkedList.Search("Geraldine"))
	linkedList.Delete("Joseph") // mid deletion!
	linkedList.Print()
	linkedList.Delete("Jaci") // head deletion
	linkedList.Print()
	linkedList.Delete("Jason") // tail deletion
	linkedList.Print()
}
