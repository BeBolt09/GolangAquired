package main

import (
	"fmt"
)

//Maps in Go are a collection of unordered key-value pairs. (Think Dictionaries in Python)
//Keys must be unique and hashable, while values can be any type.
//To access a value, you simply use the key to index the map.

func main() {
	//Creating a map with MakeFunc and calling it score
	score := make(map[string]int)

	//adding my first key value pair to my map
	score["Sebastien"] = 73
	fmt.Println(score)

	score["Ruben"] = 12
	score["Erica"] = 95
	score["Robert"] = 80
	//printing my updated map
	fmt.Println(score)

	//creating a way to retrieve only sebastien's score
	getSebaScore := score["Sebastien"]
	fmt.Println(getSebaScore)

	//how to ddelete a key value pair
	delete(score, "Robert")
	fmt.Println(score)

	//This loop takes the range of my map("score") and
	//prints each key value pair inside the map.
	for k, v := range score {
		fmt.Printf("Score of %s is %v\n", k, v)
	}
}
