package main

import (
	"fmt"
	"strings"
)

type LinkedList struct {
	head *node
}

func (l *LinkedList) AddNode(s string) {
	currNode := l.head
	newNode := &node{data: s}
	if currNode == nil {
		// there is no current node
		l.head = newNode
		return
	}

	// nodes exists, simply add at the start
	newNode.next = l.head
	l.head = newNode
	return
}

func (l *LinkedList) ReturnLastElement() *node {
	currNode := l.head
	if currNode == nil {
		return nil
	}

	// just before the nil, stops RIGHT at the node itself
	for currNode.next != nil {
		currNode = currNode.next
	}

	return currNode
}

func (l *LinkedList) Print() {
	tmpSli := make([]string, 0)
	for currNode := l.head; currNode != nil; {
		tmpSli = append(tmpSli, currNode.data)
		currNode = currNode.next
	}
	fmt.Printf(strings.Join(tmpSli, " <-- ") + "\n")
}

func (l *LinkedList) ReverseList() {
	listOfCurrentNodes := make([]*node, 0)
	for currentNode := l.head; currentNode != nil; {
		listOfCurrentNodes = append(listOfCurrentNodes, currentNode)
		currentNode = currentNode.next
	}

	l.head = listOfCurrentNodes[len(listOfCurrentNodes)-1]
	currentNode := l.head

	for i := len(listOfCurrentNodes) - 2; i >= 0; i -= 1 {
		// assign the .next to the appr element
		currentNode.next = listOfCurrentNodes[i]
		// then move the current node accordingly
		currentNode = listOfCurrentNodes[i]
	}
	// the for the last node, nil the .next reference
	currentNode.next = nil
	return
}

func (l *LinkedList) DeleteANode() {
	// note, this assumes populated linked list!
	secondToHeadNode := l.head.next
	extractNextData := secondToHeadNode.next.data
	secondToHeadNode.data = extractNextData
	// skip over!
	secondToHeadNode.next = secondToHeadNode.next.next
}

type node struct {
	data string
	next *node
}

func Init() *LinkedList {
	l := &LinkedList{}
	return l
}

func main() {
	list := Init()
	list.AddNode("Jason")
	list.AddNode("Jemimah")
	list.AddNode("Janina")
	list.AddNode("Jaci")
	list.Print()
	fmt.Println("The data of the last node is:", list.ReturnLastElement().data)
	list.ReverseList()
	list.Print()
	list.DeleteANode()
	list.Print()
}
