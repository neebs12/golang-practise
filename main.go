package main

import "fmt"

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

// LinkedList prepend method (adds node at the start of linked list, given data)

// LinkedList search method (searches linked list, given data) - returns bool (true if found, false if not)

// LinkedList delete method (deletes node from linked list, given data)

// func Init, initializes LinkedList

func main() {
	node := new(Node)
	fmt.Println(node)
}
