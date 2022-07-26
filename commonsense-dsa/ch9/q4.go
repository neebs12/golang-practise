package main

import (
	"fmt"
	"strings"
)

type Stack struct {
	data []string
}

func (s *Stack) Pop() string {
	l := len(s.data) - 1
	returnVal := s.data[l]
	s.data = s.data[:l]
	return returnVal
}

func (s *Stack) Push(str string) {
	s.data = append(s.data, str)
}

func (s *Stack) Length() int {
	return len(s.data)
}

func ReverseString(str string) string {
	stack := new(Stack)
	for _, rr := range str {
		localStr := string(rr)
		stack.Push(localStr)
	}

	reversedStrSli := make([]string, stack.Length())
	for stack.Length() > 0 {
		reversedStrSli = append(reversedStrSli, stack.Pop())
	}

	return strings.Join(reversedStrSli, "")
}

func main() {
	name := "Jason Aricheta"
	fmt.Println(ReverseString(name))

}
