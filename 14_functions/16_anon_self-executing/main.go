package main

import "fmt"

func main() {
	func() {
		fmt.Println("I'm driving!")
	}()
}

// More or less the same as JavaScript IIFE.
