package main

import "fmt"

func main() {
	a := 43

	fmt.Println(a)  // 43
	fmt.Println(&a) // 0x.........

	var b *int = &a
	fmt.Println(b)  // 0x......
	fmt.Println(*b) // 43

	*b = 42        // b says, "the value at this address, changes to 42"
	fmt.Println(a) // 42

	// this is useful
	// we can pass a memorty address instead of a bunch of values (wee can pass a reference)
	// and then we can still change the value of whatever is stored at that memory address
	// this makes our programs more performant
	// we don't have to pass around large amounts of data
	// we only have to pass around addresses

	// everything is PASS BY VALUE in go, btw
	// when we pass a memory address, we are passing a value.

}
