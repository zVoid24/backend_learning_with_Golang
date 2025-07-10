package main

import "fmt"

type Employee struct {
	Name        string
	Age         int
	Designation string
}

func employeeInfo(employee1 Employee) {

	fmt.Println("===== Inputted Data =====")
	fmt.Printf("Employee name: %s\nEmployee age: %d\nEmployee designation: %s", employee1.Name, employee1.Age, employee1.Designation)
}

func main() {
	var employee1 Employee
	fmt.Print("Enter employee name: ")
	fmt.Scanln(&employee1.Name)
	fmt.Print("Enter employee Age: ")
	fmt.Scanln(&employee1.Age)
	fmt.Print("Enter employee Designation: ")
	fmt.Scanln(&employee1.Designation)
	employeeInfo(employee1)

}
