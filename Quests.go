package main

import "fmt"

func subtract(a, b int) int {
	return a - b
}
func add(a, b int) int {
	return a + b
}
func main() {
	fmt.Println("Hello, Go!")
	fmt.Println("5 - 3 =", subtract(5, 3))
	fmt.Println("Результат работы задачи: 5 + 3 =", add(5, 3))
}
