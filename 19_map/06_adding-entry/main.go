package main

import "fmt"

func main() {

	myGreeting := map[string]string{
		"Tim":   "Good morning",
		"Jenny": "Bonjour",
	}

	myGreeting["Harleen"] = "Hiddy ho"

	fmt.Println(myGreeting)
}
