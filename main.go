package main

import (
	"fmt"
	"math"
)

type shape interface {
	getArea() float32
}

type shape2 interface {
	getArea() float32
	getParam() float32
}

type circle struct {
	radius float32
}

type rectangle struct {
	side1 float32
	side2 float32
}

func (c *circle) getArea() float32 {
	return math.Pi * (c.radius * c.radius)
}

func (c *circle) getParam() float32 {
	return 2 * math.Pi * c.radius
}

func (r *rectangle) getArea() float32 {
	return r.side1 * r.side2
}

func main() {
	myCircle := circle{3.5}
	myRect := rectangle{4.24, 9.85}

	myShapes := []shape{&myCircle, &myRect}

	for _, val := range myShapes {
		fmt.Println(val.getArea())
	}

	fmt.Println("------------------------")

	myCircle2 := circle{8.95}
	myShapes2 := []shape2{&myCircle, &myCircle2}
	// <--- not valid to incl myRect due to missing struct method
	for _, val := range myShapes2 {
		fmt.Println(val.getParam())
	}
}
