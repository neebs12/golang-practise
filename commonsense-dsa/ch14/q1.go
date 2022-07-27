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

func (l *LinkedList) Print() {
	tmpSli := make([]string, 0)
	for currNode := l.head; currNode != nil; {
		tmpSli = append(tmpSli, currNode.data)
		currNode = currNode.next
	}
	fmt.Printf(strings.Join(tmpSli, " <-- ") + "\n")
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
}
