package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println()
	fmt.Println("Let's handle some Errors  >(")

	//Okay so we will attempt to convert a sting to an integer
	//if the err is nil then all went well, on the other hand if the
	//error is not nill we will return an error message

	i, err := strconv.Atoi("42b")
	if err != nil {
		fmt.Println("Coudln't convert:", err)
		return
	} else {
		fmt.Println(i)
	}

	//Things to remember
	//Use errors.New() if i want to make a custom error message.
}
