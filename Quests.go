package main

import "fmt"

func subtract(a, b int) int {
	return a - b
}
func main() {
	fmt.Println("Hello, Go!")
	fmt.Println("5 - 3 =", subtract(5, 3))
}
