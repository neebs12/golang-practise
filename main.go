package main

import "fmt"

type human struct{}

func (h human) greet(str string) {
	fmt.Printf("Hello, I hope we have found ourselves in a good place in this rocky planet, here take this word: %v\n", str)
}

type modernHuman struct {
	*human
}

func (mh modernHuman) greet(flt float32) {
	fmt.Printf("Good day to you! We aren't primitives! Take this more sophisticated number! %.2f\n", flt)
}

func main() {
	h := &human{}
	h.greet("flabergasted")
	mh := &modernHuman{}
	mh.greet(2)
	mh.greet(2.2)
	// mh.greet("Hello")  <--- not allowed, not really overloaded
}
