package main

import "fmt"

func main() {

	//Arrays have a fixed length(like in C)
	//to declare an array
	var myIntArray = [4]int{0, 1, 2, 3}
	// another way to declare it : var myIntArray[4]int=[4]int{0,1,2,3}

	//when printing the int array by using array,len(array)
	//it will show me #1 the array inside brackets and #2 length of array outside
	fmt.Println(myIntArray, len(myIntArray))

	//MAKING A STRING ARRAY
	names := [3]string{"seba", "erica", "cosmos"}
	fmt.Println(names, len(names))

	fmt.Println("")
	fmt.Println("")
	//=============================\\
	//Let's move on to slices
	//to declare a slice, we just leave the size
	//condition undeclared
	var scores = []int{1, 2, 3, 4, 5, 6}
	fmt.Println(scores)
	//to change an element of array or slice
	scores[2] = 80
	//If i want to append
	scores = append(scores, 85)
	fmt.Println(scores, len(scores))

	//I can create a slice from an array
	rangeOne := names[0:2]
	rangeTwo := names[1:]

	fmt.Println(rangeOne, rangeTwo)
}
