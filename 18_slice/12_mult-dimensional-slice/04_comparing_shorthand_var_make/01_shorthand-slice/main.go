package main

import "fmt"

func main() {
	student := []string{}
	students := [][]string{}
	// Since the slice was instantiated without a value, length, nor capacity, no indexs yet exist.
	// student[0] = "Mark"
	// As a result, append must be used.
	student = append(student, "Mark")
	fmt.Println(student)
	fmt.Println(students)
}
