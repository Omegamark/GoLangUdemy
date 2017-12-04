package main

import "fmt"

func hello() {
	fmt.Print("hello ")
}

func world() {
	fmt.Println("world")
}

func main() {
	// defer will hold execution of the deferred function until the all other code within the closure has run and just before the function exits.
	// in this case, the world() function defers until just before the closure for main finishes execution.
	defer world()
	hello()
}
