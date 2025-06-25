package main

import "fmt"

func main() {
	var (
		a int = 10
		b int = 20
	)
	func(a int, b int) {
		c := a + b
		fmt.Println(c)
	}(a, b)

	add := func(a int, b int) int {
		return a + b
	}
	fmt.Println(add(a, b))
}
