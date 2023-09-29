package main

import "fmt"

func main() {
	//	LET'S LEARN ABOUT INTERFACES
	//	What's an interface?
	//	--> An interface in Go is a named set of method signatures.
	//	 A value of an interface type can hold any value that implements those methods.
	//	 Interfaces are used to decouple code that depends on specific implementations
	//	 from the code that provides those implementations.
	var doggy Animal = &Dog{name: "cosmos"}

	fmt.Println(doggy.Speak())

	var kitty Animal = &Cat{name: "zeus"}

	fmt.Println(kitty.Speak())

}

type Animal interface {
	Speak() string
}

type Cat struct {
	name string
}

func (d *Cat) Speak() string {
	return "meaw!"
}

type Dog struct {
	name string
}

func (d *Dog) Speak() string {
	return "Woof!"
}
