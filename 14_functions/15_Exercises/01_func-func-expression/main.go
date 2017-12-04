// Basic func & func expression
package main

import "fmt"

func main() {
	isEven(8)
}

func isEven(z int) (int, bool) {
	var r1 int
	var r2 bool
	r1 = z / 2
	if z%2 == 0 {
		r2 = true
	}
	fmt.Println(r1, r2)
	return r1, r2
}
