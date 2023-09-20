package main

import "fmt"

func main() {
	fmt.Println("")
	fmt.Println("Let's Dive into Functions !")
	fmt.Println("Functions are very simple, We're actually inside a function right now !")
	fmt.Println("")

	//this is using the concat function i typed down below
	fmt.Println(concat("seba is", " having a great day"))
	//this is using the sub function built below
	fmt.Println(sub(5, 23))
}

func sub(x int, y int) int {
	return x - y
}

func concat(s1 string, s2 string) string {
	return s1 + s2
}
