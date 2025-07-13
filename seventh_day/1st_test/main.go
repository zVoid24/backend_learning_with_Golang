package main

import "fmt"

func main() {
	arr := [5]string{"My", "name", "is", "Zahidul", "Islam"}
	s := arr[3:5]
	s2 := s[0:1]
	s3 := []int{1, 2, 5} //slice literal
	//s3 = append(s3, 6)
	s4 := make([]int, 3) //with make function (length and array)
	s4[0] = 1
	s4[1] = 2
	s4[2] = 69
	s5 := make([]int, 3, 5) //with make function (length and array and capacity)
	var s6 []int            //nil slice
	fmt.Println("Value of array: ", arr, "Value of slice of array: ", s, "Length of slice: ", len(s), "Capacity of slice: ", cap(s))
	fmt.Println("Value of array: ", arr, "Value of slice of array: ", s, "Value of slice of slice: ", s2, "Length of slice of slice: ", len(s2), "Capacity of slice of slice: ", cap(s2))
	fmt.Println("Value of slice literal: ", s3, "Length of slice literal: ", len(s3), "Capacity of slice literal: ", cap(s3))
	fmt.Println("Value of slice with make: ", s4, "Length of slice with make: ", len(s4), "Capacity of slice of make: ", cap(s4))
	fmt.Println("Value of slice with make: ", s5, "Length of slice with make: ", len(s5), "Capacity of slice of make: ", cap(s5))
	fmt.Println("Value of nil slice : ", s6, "Length of nil slice: ", len(s6), "Capacity of nil slice: ", cap(s6))
	s6 = append(s6, 5, 6, 7)
	fmt.Println("Value of nil slice : ", s6, "Length of nil slice: ", len(s6), "Capacity of nil slice: ", cap(s6))
}
