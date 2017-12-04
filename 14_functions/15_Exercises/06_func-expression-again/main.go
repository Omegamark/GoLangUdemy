package main

import "fmt"

func main() {
	isEven := func(z int) (int, bool) {
		var r1 int
		var r2 bool
		r1 = z / 2
		if z%2 == 0 {
			r2 = true
		}
		return r1, r2
	}
	fmt.Println(isEven(8))
}
