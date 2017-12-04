package main

import "fmt"

func main() {
	fmt.Println(greet("Jane ", "Doe "))
}

// Multiple returns
// Returns can be separated by a comma.
// If the arguments are all of the same type, you only need to add the type at the end of the comma separated list.

func greet(fname, lname string) (string, string) {
	return fmt.Sprint(fname, lname), fmt.Sprint(lname, fname)
}
