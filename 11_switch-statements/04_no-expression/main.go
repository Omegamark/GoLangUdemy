package main

import "fmt"

func main() {
	myFriendsName := "Brian"

	switch {
	case len(myFriendsName) == 2:
		fmt.Println("Sup, your name is length 2")
	case myFriendsName == "Marcus", myFriendsName == "Medhi":
		fmt.Println("Your name is either Marcus or Medhi")
	default:
		fmt.Println("nothing matched; this is the default.")
	}
}
