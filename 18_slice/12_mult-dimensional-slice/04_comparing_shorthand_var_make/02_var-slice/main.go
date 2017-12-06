package main

import "fmt"

func main() {
	var student []string
	var students [][]string
	// again, slices are instantiated with no value, length, nor capacity. So trying to manually insert a value at index 0 will throw an error since there is NOT a 0 index yet.
	// student[0] = "Mark"
	// See above and shorthand-slice
	student = append(student, "Mark")
	fmt.Println(student)
	fmt.Println(students)
}
