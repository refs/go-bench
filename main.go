package main

import (
	"fmt"
	"performance-go/pointers"
)

// The following example illustrate the cost of indirection of [interface / struct] method sets.

// Talker defines something that talks.
type Talker interface {
	Talk(string)
}

// Lion implements the Talker interface.
type Lion struct{}

// Talk implements the Talker interface.
func (l Lion) Talk(message string) {
	// print(message)
}

func sayInterface(t Talker, message string) {
	t.Talk(message)
}

func sayStruct(l Lion, message string) {
	l.Talk(message)
}

func main() {
	a := pointers.P{
		A: 42,
	}
	fmt.Printf("%p\n", &a)
	pointers.PMem(&a)
	fmt.Printf("%v\n", a.A) // 21, modified inside pointers.PMem
}
