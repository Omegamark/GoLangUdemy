package main

import "fmt"

func main() {

	greeting := []string{
		"Good morning",
		"Bonjour",
		"dias!",
		"Bongiorno!",
		"您好",
		"Ohayo",
		"Selemat pagi",
		"Gutten morgen!", // Notice that trailing commas do not cause problems.
	}

	fmt.Print("[1:2] ")
	fmt.Println(greeting[1:2])
	fmt.Print("[:2] ")
	fmt.Println(greeting[:2])
	fmt.Print("[5:] ")
	fmt.Println(greeting[5:])
	fmt.Print("[:] ")
	fmt.Println(greeting[:])
}
