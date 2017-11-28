package main

import "fmt"

var x int = 4

func main() {
	fmt.Println(x)
	foo()
}

func foo() {
	fmt.Println(x)
}
