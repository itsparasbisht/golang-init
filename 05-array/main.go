package main

import (
	"fmt"
)

func main() {
	// size of an array has to defined during the array declaration
	var roles [2]string
	roles[0] = "admin"
	roles[1] = "user"
	
	fmt.Println(roles)

	// array declaration & initialization
	var users = [3]string{"harry", "john", "jane"}

	fmt.Println(users)
	fmt.Println(len(users))

	var scores = [5]float32{}
	fmt.Println(scores)   // [0 0 0 0 0]

	// looping in an array
	for i := 0; i<len(roles); i++{
		fmt.Println(roles[i])
	}

	for index, value := range users{
		fmt.Printf("%d -> %s \n", index, value)
	}

	// arrays are fixed size
	// arrays in Go are designed to hold elements of a single data type
	// arrays are passed by value

}