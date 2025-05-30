package main

import (
	"fmt"
	"practice/models" // Correct import path based on our module name
)

func main() {
	// Using the constructor function from models package
	emp1 := models.NewEmployee("Kartik", 36, 500)

	// This line would cause a compilation error if uncommented:
	// fmt.Println(emp1.name) // ERROR: emp1.name is not accessible

	// Instead, we must use the exported fields and methods:
	fmt.Println(emp1.EMPID)     // Works - EMPID is exported (starts with uppercase)
	fmt.Println(emp1.Salary)    // Works - Salary is exported (starts with uppercase)
	fmt.Println(emp1.GetName()) // Works - using accessor method for unexported field

	// Let's change the name using the setter method
	emp1.SetName("Kartik Arora")
	fmt.Println(emp1.GetName()) // Will print the updated name
}
