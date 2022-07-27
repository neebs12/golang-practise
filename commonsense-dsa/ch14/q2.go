package main

import (
	"fmt"
	"strings"
)

type DoublyLinkedList struct {
	head *node
	tail *node
}

// add a node - at start
func (d *DoublyLinkedList) AddNodeAtStart(s string) {
	headNode := d.head
	newNode := &node{data: s}
	if headNode == nil {
		// first node - also both head and tail are pointed at the same node if the first node!
		// keep the newNode's next and prev pointed at nil!
		d.head = newNode
		d.tail = newNode
		return
	}

	// else, there is more than one node in DoublyLinkedList
	existingHeadNode := d.head      // save reference!
	d.head = newNode                // reassg head
	newNode.next = existingHeadNode // newNode looks forward
	existingHeadNode.prev = newNode // existingHeadNode looks backwards
	return
}

// add a node - at end
func (d *DoublyLinkedList) AddNodeAtEnd(s string) {
	tailNode := d.tail
	newNode := &node{data: s}
	if tailNode == nil {
		d.head = newNode
		d.tail = newNode
		return
	}

	// else, there are existing node
	existingTailNode := d.tail      // preseve ref
	d.tail = newNode                // reassg
	newNode.prev = existingTailNode // new looking backwards
	existingTailNode.next = newNode // old looking forwards
	return
}

// print from head
func (d *DoublyLinkedList) PrintFromHead() {
	tmpSli := make([]string, 0)
	for currNode := d.head; currNode != nil; {
		tmpSli = append(tmpSli, currNode.data)
		currNode = currNode.next
	}
	fmt.Printf(strings.Join(tmpSli, " <-- ") + "\n")
}

// print from tail
func (d *DoublyLinkedList) PrintFromTail() {
	tmpSli := make([]string, 0)
	for currNode := d.tail; currNode != nil; {
		tmpSli = append(tmpSli, currNode.data)
		currNode = currNode.prev // <-- difference!
	}
	fmt.Printf(strings.Join(tmpSli, " --> ") + "\n")
}

type node struct {
	data string
	next *node
	prev *node
}

func Init() *DoublyLinkedList {
	return new(DoublyLinkedList)
}

func main() {
	d := Init()
	fmt.Println(d)
	d.AddNodeAtStart("Jemimah")
	d.AddNodeAtStart("Jason")
	d.AddNodeAtStart("Janina")
	d.AddNodeAtStart("Jaci")
	d.PrintFromHead()
	d.PrintFromTail()
	d.AddNodeAtEnd("Geraldine")
	d.PrintFromHead()
}
