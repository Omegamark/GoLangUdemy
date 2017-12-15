package main

import (
	"fmt"
	"sort"
)

type people []string

// These 3 methods are needed for the Interface interface from the sort package.
// The methods are being attached to the user defined type of "people"
func (p people) Len() int           { return len(p) }
func (p people) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p people) Less(i, j int) bool { return p[i] < p[j] }

func main() {
	studyGroup := people{"Zeno", "John", "Al", "Jenny"}

	fmt.Println(studyGroup)
	sort.Sort(studyGroup)
	fmt.Println(studyGroup)
}
