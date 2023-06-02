package foo

import "fmt"

func DoCalculations() {
	fmt.Println(add(1, 2))
	fmt.Println(add(4, 2))
	fmt.Println(add(6, 1))
}

func add(a, b int) int {
	return a + b
}
