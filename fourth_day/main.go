package main

import "fmt"

func call() func(a int, b int) int {
	return add
}

func add(a int, b int) int {
	return a + b
}

func highOrgerFunction(x int, y int, op func(p int, q int) int) int {
	return op(x, y)
}

func main() {
	var (
		m   int                    = 5
		n   int                    = 7
		sum func(a int, b int) int = call()
	)
	fmt.Println(highOrgerFunction(m, n, sum))
}
