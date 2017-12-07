package main

import "fmt"

func main() {

	myGreeting := map[string]string{
		"zero":  "Good morning",
		"one":   "您好",
		"two":   "Bon Dia",
		"three": "Bongiorno",
	}
	fmt.Println(myGreeting)
	delete(myGreeting, "two")
	fmt.Println(myGreeting)
}
