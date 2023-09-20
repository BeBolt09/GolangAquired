package main

import "fmt"

func main() {
	fmt.Println("")
	fmt.Println("Hello, world! Let's learn GO !")
	fmt.Println("Let's go over the Variables of Go in this Chapter !")
	fmt.Println("")

	//Starting with integers
	//first way to declare a variable
	var intExample int
	intExample = 14
	//Here's another way
	var myInt = 10
	// we can print them out now
	fmt.Printf("here's the first int we declared :%d and here's the second one:%d", intExample, myInt)
	fmt.Println("")

	//Let's do floats now
	var myFloat float32
	myFloat = 32.1

	//Let's do Strings
	var myString string
	myString = "seba's String"

	//Let's do bool
	var myBool bool
	myBool = true

	//let's do arrays
	myIntArray := [5]int{1, 2, 3, 4, 5}
	myStringArray := [5]string{"a", "r", "r", "a", "y"}

	//print out all the vars we created
	fmt.Println(myFloat)
	fmt.Println(myString)
	fmt.Println(myBool)
	fmt.Println(myIntArray)
	fmt.Print(myStringArray)

}
