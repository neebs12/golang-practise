package main

import "fmt"

var _ = fmt.Println // for debugging

func RecurPrint(nested interface{}) {
	switch nested := nested.(type) {
	case int: // if we encounter an integer
		fmt.Println(nested)
		return
	case []int, []interface{}:
		for _, val := range nested {

		}
		// default:
		// 	for _, val := range []interface{}(nested) {
		// 		RecurPrint(val)
		// 		return
		// 	}
	}
}

func main() {
	// myVal := []interface{}{
	// 	1, 2, 3,
	// 	[]int{4, 5, 6},
	// 	7,
	// 	[]interface{}
	// }

}
