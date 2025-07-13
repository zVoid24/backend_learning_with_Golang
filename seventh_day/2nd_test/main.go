package main

import "fmt"

func main() {
	var x []int
	x = append(x, 1)
	x = append(x, 5)
	x = append(x, 6)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))
	var y = x
	x = append(x, 4)
	fmt.Println(x)
	y = append(y, 7)
	fmt.Println(y)
	fmt.Println(len(y))
	fmt.Println(cap(y))
	x[2] = 101
	fmt.Println(x)
	fmt.Println(y)
}
