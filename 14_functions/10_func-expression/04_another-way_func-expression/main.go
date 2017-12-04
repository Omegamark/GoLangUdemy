package main

import "fmt"

func makeGreeter() func() string {
	// Here we are RETURNING an anonymous function of type string, which then returns the string Hello World
	return func() string {
		return "Hello World"
	}
}

func main() {
	greet := makeGreeter()
	fmt.Println(greet())
}
