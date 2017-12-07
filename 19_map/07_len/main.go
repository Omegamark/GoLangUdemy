package main

import "fmt"

func main() {

	myGreeting := map[string]string{

		"Tim":   "Good morning",
		"Jenny": "Bonjour",
	}

	myGreeting["Harleen"] = "Hidy Ho!"
	myGreeting["Mark"] = "Sup"

	fmt.Println(len(myGreeting))
}
