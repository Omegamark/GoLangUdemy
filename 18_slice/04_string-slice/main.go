package main

import "fmt"

func main() {

	greeting := []string{
		"Good morning",
		"Bonjour",
		"您好",
		"Ohaiyo",
		"Selemat pagi",
		"Gutten morgen",
	}

	for i, currentEntry := range greeting {
		fmt.Println(i, currentEntry)
	}

	for j := 0; j < len(greeting); j++ {
		fmt.Println(greeting[j])
	}

	// currentEntry is only available within the scope of the for loop.
	// fmt.Println(currentEntry)

}
