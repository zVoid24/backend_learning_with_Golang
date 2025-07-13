package main

import "fmt"

type Student struct {
	Name   string
	Roll   int
	Height float64
}

func print(students *[2]Student) {
	for i := 0; i < 2; i++ {
		fmt.Println("Information of student ", i+1)
		fmt.Println("Name: ", students[i].Name)
		fmt.Println("Roll: ", students[i].Roll)
		fmt.Println("Height: ", students[i].Height)
	}
	fmt.Printf("Address of stored data: %p\n", students)
}

func main() {

	var students [2]Student
	for i := 1; i < 3; i++ {
		fmt.Print("Enter the name of student", i, ":")
		fmt.Scanln(&students[i-1].Name)
		fmt.Print("Enter the roll of student", i, ":")
		fmt.Scanln(&students[i-1].Roll)
		fmt.Print("Enter the height of student", i, ":")
		fmt.Scanln(&students[i-1].Height)
	}
	print(&students)
	fmt.Printf("Address of stored data %p", &students)

}
