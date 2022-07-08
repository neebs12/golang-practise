package main

import "fmt"

func doubleFirstSliElm(s []int) {
	s[0] = 100
}

func doubleMpPropByPropName(myMap map[string]int, name string) {
	myMap[name] = myMap[name] * 2
}

func main() {
	sli := []int{11, 22, 33, 44}
	fmt.Println(sli)
	doubleFirstSliElm(sli)
	fmt.Println(sli) // <--- works

	mp := map[string]int{
		"apple":  50,
		"orange": 500,
	}

	fmt.Println(mp)
	doubleMpPropByPropName(mp, "apple")
	fmt.Println(mp)

}
