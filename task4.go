// Task 4– Creating a list of students

// Extending task 3, add an array for the subject a student is currently studying.
// Calculate hash of the block data and display it

package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strings"
)

// Student struct
type student struct {
	id      int
	Name    string
	Age     int
	RollNo  string
	Subject string
	hashVal string
}

// Print method for student
func (s student) Print() {
	fmt.Printf("%s List %d %s \n", strings.Repeat("=", 20), s.id, strings.Repeat("=", 20))
	fmt.Println("Name: ", s.Name)
	fmt.Println("Age: ", s.Age)
	fmt.Println("RollNo: ", s.RollNo)
	fmt.Println("Subject: ", s.Subject)
	fmt.Println("Hash: ", s.hashVal)
	fmt.Println()
}

// Function to add hash to student
func (s *student) addHash() {
	hash := sha256.New()
	hash.Write([]byte(s.Name))
	hash.Write([]byte(s.RollNo))
	hash.Write([]byte(s.Subject))
	s.hashVal = hex.EncodeToString(hash.Sum(nil))
}

func main() {
	// Create an array of students
	students := []student{
		student{1, "Awais Shahid", 23, "21i-2764", "Computer Science", ""},
		student{2, "Abdullah Abubaker", 21, "21i-1917", "Software Engineering", ""},
		student{3, "Muhammad Qasim", 24, "21i-1580", "Data Science", ""},
	}

	// Calculating Hash for each student
	for i := range students {
		students[i].addHash()
	}

	// Print the student details
	for _, s := range students {
		s.Print()
	}

}
