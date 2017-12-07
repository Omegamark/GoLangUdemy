package main

import "fmt"

func main() {

	myGreeting := map[int]string{
		0: "Good morning",
		1: "Bonjour",
		2: "您好",
		3: "Bon dia!",
	}
	fmt.Println(myGreeting)
	delete(myGreeting, 7)
	fmt.Println(myGreeting)
}
