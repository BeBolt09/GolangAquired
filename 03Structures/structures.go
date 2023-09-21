package main

import "fmt"

func main() {
	fmt.Println("")
	fmt.Println("")
	fmt.Println("Okayyy we back for our STRUCTURES !!!")

	//Let's make a bunch of stuctures and learn about
	//How they can be useful to us in some projects !
	sebastien := User{"Sebatien", "sebastiengdupont@gmail.com", true, 21}
	fmt.Println("")
	fmt.Println("Print the struct (Sebastien)")

	fmt.Println(sebastien)
	//better to way to access details about a struct
	fmt.Printf("Sebastien's details are: %+v\n", sebastien)
	//to only have some details
	fmt.Printf("Name is %v and email is %v", sebastien.Name, sebastien.Email)

}

//There is no inheritence
//First capital letter because this is kinda like a class
type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
