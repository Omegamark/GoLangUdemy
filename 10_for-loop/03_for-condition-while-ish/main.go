package main

import "fmt"

func main() {
	i := 0       // initializer
	for i < 10 { // loop and condition; functions like a while loop since Go does not have while loops.
		fmt.Println(i)
		i++ // post loop operation
	}
}
