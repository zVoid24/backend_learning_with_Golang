package main

import (
	"fmt"
	"hello/mathlib"
)

func main() {
	var (
		a   int
		b   int
		sum int
		mul int
		sub int
		div float32
	)
	fmt.Print("Enter a = ")
	fmt.Scanln(&a)
	fmt.Print("Enter b = ")
	fmt.Scanln(&b)
	sum = mathlib.Add(a, b)
	mul = mathlib.Mul(a, b)
	sub = mathlib.Sub(a, b)
	div = mathlib.Div(float32(a), float32(b))
	fmt.Println("Sum:", sum)
	fmt.Println("Sub:", sub)
	fmt.Println("Mul:", mul)
	fmt.Println("Div:", div)
}
