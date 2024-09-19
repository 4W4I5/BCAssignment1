// Task 2 array of structs
// Here, you can see the two different structs. The company struct uses the employee struct as an array. Write a program that do the following:
// Add three employee using the employee struct (e.g., emp1 := employee{"Amir", 80000, "Full-Stack Developer"}
// Create an array of ‘emplys’ by add the above three records (emplys:=[]employee{ …})
//  create a company struct and add values to it (e.g., {"Tetra", emplys})
// Print company details

package main

import (
	"fmt"
)

type employee struct {
	Name   string
	Salary int
	Role   string
}

type company struct {
	Name     string
	Employes []employee
}

func printEmployee(e employee) {
	fmt.Println("Name: ", e.Name)
	fmt.Println("Salary: ", e.Salary)
	fmt.Println("Role: ", e.Role)
	fmt.Println()
}

func printCompany(c company) {
	fmt.Println("Company Name: ", c.Name)
	fmt.Println("Employes: ")
	for _, e := range c.Employes {
		printEmployee(e)
	}
}

func main() {
	// Create an array of employees
	employees := []employee{
		employee{"Amir", 80000, "Full-Stack Developer"},
		employee{"Ali", 70000, "Front-End Developer"},
		employee{"Ahmed", 60000, "Back-End Developer"},
	}

	// Create a company struct
	company := company{"Tetra", employees}

	// Print the company details
	printCompany(company)
}
