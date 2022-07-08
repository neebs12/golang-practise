package main

import "fmt"

type Point struct {
	x float32
	y float32
}

type Circle struct {
	radius float32
	center Point
}

type Circle2 struct {
	radius float32
	center *Point
}

type Circle3 struct {
	radius float32
	*Point
}

func changeXValue(pt Point) {
	// <--- pt creates a copy
	pt.x = 9.001
	fmt.Println(pt)
}

func changeXValue2(pt *Point) {
	(*pt).x = 23.001 // valid
	pt.x = 25.251    // valid - shorthand
}

func main() {
	var pt1 Point = Point{11.1, 4.8}
	fmt.Println(pt1)
	pt1.x = 100
	fmt.Println(pt1)
	changeXValue(pt1)
	fmt.Println(pt1) // <-- not changed

	fmt.Println("-------------------")

	var pt2 *Point = &Point{1.1, 2.1}
	fmt.Println(pt2, *pt2)
	pt2.x = 99.1
	// <--- didnt need to deref for structs!
	fmt.Println(pt2, *pt2)
	changeXValue2(pt2)
	fmt.Println(pt2, *pt2)

	fmt.Println("-------------------")

	var myCircle *Circle = &Circle{4.2, Point{11.1, 11.2}}
	fmt.Println(myCircle)
	fmt.Println(myCircle, myCircle.radius, myCircle.center)
	fmt.Println(myCircle.center.x, myCircle.center.y)
	changeXValue(myCircle.center)
	fmt.Println(myCircle)

	fmt.Println("-------------------")

	var myCircle2 *Circle2 = &Circle2{4.2001, &Point{11.1001, 11.2001}}
	fmt.Println(myCircle2)
	fmt.Println(myCircle2, myCircle2.radius, myCircle2.center)
	changeXValue2(myCircle2.center)
	fmt.Println(myCircle2.center) // <--- changes!!

	fmt.Println("-------------------")

	myCircle3 := &Circle3{9.001, &Point{4.002, 5.001}}
	fmt.Println(myCircle3, myCircle3.x, myCircle3.y)
	// changeXValue2(myCircle3)
	// <--- error: mismatched argument types
	// <--- from *Circle(arg) to *Point(param def)
}
