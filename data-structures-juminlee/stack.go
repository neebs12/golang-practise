package main

import "fmt"

// Stack is LIFO - last in, first out
type Stack struct {
	content []string
}

// Stack has Pop, this takes out the value at the top of the stack (last value)
// -- and returns the assc. value removed
func (s *Stack) Pop() string {
	l := len(s.content)
	topOfStack := s.content[l-1]
	s.content = s.content[:l-1]
	return topOfStack
}

// Stack has PUSH, this places a value at the top of the stack
func (s *Stack) Push(val string) {
	s.content = append(s.content, val)
}

// Print, prints the stack
func (s Stack) Print() {
	l := len(s.content)
	fmt.Printf("%s", s.content[l-1])
	for i := l - 2; i >= 0; i -= 1 {
		fmt.Printf(" --> %s", s.content[i])
	}
	fmt.Println()
}

func main() {
	stack := Stack{
		[]string{
			"Jason", "JR", "Daniel",
		},
	}
	stack.Print()
	stack.Push("Sam")
	stack.Push("Josh")
	stack.Push("Kazz")
	stack.Print()
	stack.Pop()
	stack.Pop()
	stack.Pop()
	stack.Print()
}
