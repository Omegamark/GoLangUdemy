package main

import "fmt"

func main() {
	var student []string
	var students [][]string
	fmt.Println(student)
	fmt.Println(students)
	// Look at shorthand vs. var implimentation. var instantiates with a value of 'nil'. Whereas := instantiates with a value. ie: '' for a string or 0 for int.
	fmt.Println(student == nil)
}
