package main

import "fmt"

func main() {
	i := 0
	for {
		fmt.Println(i)
		i++
	}
}

// This is a for loop with no post, it will loop infinitely, could be used in situations where a function is acting as a listener.
