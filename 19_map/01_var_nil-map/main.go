package main

import "fmt"

func main() {
	// This is NOT the way to create a map. Having a nil reference for a map seemingly makes it useless.
	var myGreeting map[string]string
	fmt.Println(myGreeting)
	fmt.Println(myGreeting == nil)
}

// add these lines:
/*
   myGreeting["Tim"] = "Good morning"
   myGreeting["Jenny"] = "Bonjour"
*/
// and you will get this
// panic: assignment to entry in nil map
