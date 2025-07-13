package main

import "fmt"

func sum(numbers ...int) int {
	var sum int = 0
	for i := 0; i < len(numbers); i++ {
		sum += numbers[i]
	}
	return sum
}

func main() {
	var n int
	var x int
	var s []int
	fmt.Print("How many numbers do you want to input?: ")
	fmt.Scanln(&n)
	for i := 0; i < n; i++ {
		fmt.Scanln(&x)
		s = append(s, x)
	}
	y := sum(s...)
	fmt.Println("Sum: ", y)
}
