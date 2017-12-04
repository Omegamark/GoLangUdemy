package main

import "fmt"

func main() {
	// 1, 25 are the limiters for this function. The slice here is not an array since I cannot assign a value to position [1].
	// According to Brett, the limiter will be doubled if the upper limit is exceeded. That is likely what is happening here since my string exceeds 25 characters.
	m := make([]string, 1, 25)
	fmt.Println(m) // [ ]
	changeMe(m)
	fmt.Println(m) // Mark
	fmt.Printf("%T", m)
}

func changeMe(z []string) {
	z[0] = "Mark is absolutely the greatest programmer of all time!!!!!!!!!"
	fmt.Println(z)
}
