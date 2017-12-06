package main

import "fmt"

func main() {
	student := make([]string, 35)
	students := make([][]string, 35)
	// Since the slice was instantiated with a length/capacity it is possible to insert values directly as seen directly below.
	// student[0] = "Mark"
	// By appending to student, which is already instantiated with 35 empty indexes, "Mark" will be appended to the 36th place.
	student = append(student, "Mark")
	fmt.Println(student)
	fmt.Println(students)
}
