package main

import "fmt"

func main() {

	myGreeting := map[string]string{
		"Tim":   "Good morning",
		"Jenny": "Bonjour",
	}
	myGreeting["Harleen"] = "Hidy Ho!"
	fmt.Println(myGreeting)
	myGreeting["Harleen"] = "Sup"
	fmt.Println(myGreeting)
}
