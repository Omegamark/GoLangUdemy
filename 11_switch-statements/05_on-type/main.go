package main

import "fmt"

// switch on types
// -- normally we switch on the value of a variable
// -- Go allows you to switch on type of variable

type contact struct {
	greeting string
	name     string
}

func SwitchOnType(x interface{}) {
	switch x.(type) { // this is an assertion; asserting, "x is of this type!"
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	case contact:
		fmt.Println("contact")
	default:
		fmt.Println("Type unknown.")
	}
}

func main() {
	SwitchOnType(7)
	SwitchOnType("Grant")
	var t = contact{"Good to see you,", "Joe"}
	SwitchOnType(t)
	SwitchOnType(t.greeting)
	SwitchOnType(t.name)
}
