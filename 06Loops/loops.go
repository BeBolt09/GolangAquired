package main

import (
	"fmt"
)

func main() {
	fmt.Println()
	fmt.Println("Let's Go Over Loops in here")

	//Loops are writing like this:
	//for INITIAL;CONDITIOIN;AFTER;{
	//do something
	//}

	fmt.Println("in the loop below, Print 0->9")
	for i := 0; i < 10; i++ {
		fmt.Print(i)
	} //print 0->9

	fmt.Println("Let's play the fizzbuzz game where i have to print")
	fmt.Println("number 1->100 (1 each line) But multiples of 3 are replaced")
	fmt.Println("with the word FIZZ and Multiples of 5 are replaced with the")
	fmt.Println("word BUZZ. Multiples of 3 and 5 will be FIZZBUZZ")
	fizzbuzz()
}

func fizzbuzz() {
	for i := 0; i <= 100; i++ {
		three := false
		five := false
		if i%3 == 0 {
			three = true
		}
		if i%5 == 0 {
			five = true
		}
		if three && five {
			fmt.Println("fizzbuzz")
		} else if three {
			fmt.Println("fizz")
		} else if five {
			fmt.Println("buzz")
		} else {
			fmt.Println(i)
		}
	}
}
