package main

import "fmt"

func add(a, b int) int {
	return a + b
}
func main() {
	fmt.Println("Hello, Go!")
	fmt.Println("Результат работы задачи: 5 + 3 =", add(5, 3))
}
