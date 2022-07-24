package main

import "fmt"

type Queue struct {
	content []string
}

// Enqueue, adds an item at the end of the queue
// (back) A --> B --> C (front)
// (back) NEW --> A --> B --> C (front)
func (q *Queue) Enqueue(val string) {
	q.content = append(q.content, val)
}

// Dequeue, removes an item at the start of the queue, returns dequeued item
// (back) A --> B --> C (front)
// (back) A --> B (front), C removed and returned
func (q *Queue) Dequeue() string {
	returnVal := q.content[0]
	q.content = q.content[1:]
	return returnVal
}

// Print
func (q Queue) Print() {
	l := len(q.content)
	fmt.Printf("%s", q.content[l-1])
	for i := l - 2; i >= 0; i -= 1 {
		fmt.Printf(" --> %s", q.content[i])
	}
	fmt.Println()
}

func main() {
	queue := Queue{
		[]string{
			"Jason",
			"Kaz",
			"Josh",
		},
	}
	queue.Print()
	queue.Enqueue("JR")
	queue.Enqueue("Sam")
	queue.Enqueue("Hannah")
	queue.Print()
	queue.Dequeue()
	queue.Dequeue()
	queue.Dequeue()
	queue.Print()
}
