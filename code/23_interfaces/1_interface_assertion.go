package main

import "fmt"

func main() {
	// Compiler will deduce the type of 123 as int.
	var x interface{} = 123

	// Case 1:
	n, ok := x.(int)
	fmt.Println(n, ok) // 123 true
	n = x.(int)
	fmt.Println(n) // 123

	// Case 2:
	a, ok := x.(float64)
	fmt.Println(a, ok) // 0 false

	// Case 3:
	a = x.(float64) // will panic
}
