package main

import (
	"fmt"
	"strings"
)

// In the separate channel HashTable, this consists of an array

// This is the ArraySize
const ArraySize = 7

// HashTable
type HashTable struct {
	array [ArraySize]*bucket // bucket is a linked list
}

// INSERT
// SEARCH
// DELETE

// bucket
type bucket struct {
	head *node
}

// insert (works)
// -- this inserts new nodes in to the bucket, from the start (more effecient)
func (b *bucket) insert(k string) {
	// GCL if exists
	// do a search first, if true, return prematurely
	if b.search(k) {
		return
	}

	head := b.head
	newNode := node{key: k}

	// GCL, if the bucket has no other items to begin with -- ie head == nil
	if head == nil {
		b.head = &newNode
		return
	} else {
		// head is not nil and is pointing to an existing node
		nodeToBeMoved := b.head // is *node type
		b.head = &newNode
		b.head.next = nodeToBeMoved
		return
	}
}

// search (works)
func (b *bucket) search(k string) bool {
	currNode := b.head // is *node type

	for currNode != nil {
		if currNode.key == k {
			return true
		}
		// incr forward
		currNode = currNode.next
	}
	return false
}

// delete (works)
func (b *bucket) delete(k string) {
	// deletion happens from the prespective of the prev
	prevNode := b.head // is *node

	// GCL if the node is the head of the bucket
	if prevNode.key == k {
		nextNode := prevNode.next // is *node
		b.head = nextNode         // prevNode referece is dropped from b.head
		return
	}

	for prevNode.next != nil {
		currNode := prevNode.next
		if currNode.key == k {
			nextNode := currNode.next // is *node
			prevNode.next = nextNode
			return
		}
		// increment forward for linked list
		prevNode = prevNode.next
	}
}

// print (for testing purposes) (works)
// will use a pointer receiver for consistency
func (b *bucket) Print() {
	sli := make([]string, 0)
	for currNode := b.head; currNode != nil; currNode = currNode.next {
		val := currNode.key
		sli = append(sli, val)
	}
	fmt.Println(strings.Join(sli, " <--- "))
}

// node
type node struct {
	key  string
	next *node
}

// hash (works)
// -- this is the hashing function
func hashIndex(str string) int {
	// case insensitive
	str = strings.ToUpper(str)
	var summation int
	for _, rr := range str {
		summation += int(rr)
	}
	return summation % ArraySize
}

// Init
// this initializes the state of the programme by adding buckets within the HashTable
// -- here, the .array of the Hashtable will contain a bucket. Where each .head of each bucket is pointing to nil
func Init() *HashTable {
	table := new(HashTable)
	for i := 0; i < ArraySize; i += 1 {
		table.array[i] = new(bucket)
	}
	return table
}

func main() {
	// fmt.Println(hashIndex("jason"))
	// fmt.Println(hashIndex("jam"))
	// fmt.Println(hashIndex("jimmy"))
	// fmt.Println(hashIndex("jodie"))
	// fmt.Println(hashIndex("graeme"))
	// fmt.Println(hashIndex("isaac"))
	// fmt.Println(hashIndex("sotelo"))

	// table := Init()
	// // then we have to fill with a group of buckets
	// fmt.Println(table)
	b := new(bucket)
	b.insert("jason")
	b.insert("jam")
	b.insert("jimmy")
	b.insert("joseph")
	b.insert("geraldine")
	fmt.Println("jason", b.search("jason"))
	fmt.Println("jemimah", b.search("jemimah"))
	fmt.Println("jimmy", b.search("jimmy"))
	b.Print()
	b.delete("geraldine")
	b.delete("jason")
	b.delete("jimmy")
	b.Print()
}
