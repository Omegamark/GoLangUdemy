package main

import "fmt"

func main() {

	myGreeting := map[int]string{
		0: "Good morning!",
		1: "您好!",
		2: "Buenos dias!",
		3: "Bongiorno!",
	}
	for key, val := range myGreeting {
		fmt.Println(key, " - ", val)
	}
}
