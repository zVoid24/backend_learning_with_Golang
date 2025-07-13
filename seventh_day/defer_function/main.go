package main

import "fmt"

func deferFucntion(a int) (result int) {
	result = a
	fmt.Println("First ", result)
	x := func(result int) {
		result = result + 10
		fmt.Println("Second ", result)
	}

	defer x(result)

	y := func() {
		result = result + 69
		fmt.Println("Third ", result)
	}

	defer y()
	result += 1
	return
}

func deferFucntion2(a int) int {
	result := a
	fmt.Println("First ", result)
	x := func(result int) {
		result = result + 10
		fmt.Println("Second ", result)
	}

	defer x(result)

	y := func() {
		result = result + 69
		fmt.Println("Third ", result)
	}

	defer y()
	result += 1
	return result
}

func main() {
	a := deferFucntion(10)
	fmt.Println(a)
	b := deferFucntion2(10)
	fmt.Println(b)
}
