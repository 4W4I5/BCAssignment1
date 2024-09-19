// Task 1 - Pass Struct as Function Arguments

package main

import (
	"fmt"
)

type student struct {
	Name   string
	Age    int
	RollNo string
}

func printStudent(s student) {
	fmt.Println("Name: ", s.Name)
	fmt.Println("Age: ", s.Age)
	fmt.Println("RollNo: ", s.RollNo)
	fmt.Println()
}

func main() {
	// Create a struct of student
	student1 := student{"Awais Shahid", 23, "21i-2764"}

	// Print the student details
	printStudent(student1)

}
