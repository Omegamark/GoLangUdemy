package main

import "fmt"

func main() {
	fmt.Println(greet("Jane ", "Doe"))
}

// Named return value.
// s string - is the name and type to which the return will be stored.
func greet(fname string, lname string) (s string) {
	s = fmt.Sprint(fname, lname)
	return
}
