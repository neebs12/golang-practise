package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
)

func main() {
  // try to get name, and add a greeting to it
  scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Please type your name here: ")
	scanner.Scan()
	input := scanner.Text()
	fmt.Printf("Good morning %v!!\n", input)
	
	fmt.Print("What year were you born? ")
	scanner.Scan()
	secondInput, _ := strconv.ParseInt(scanner.Text(), 10, 32)
	fmt.Printf("You are %d years old!\n", 2022 - secondInput)
}
