package main

import (
	"fmt"
	"slices"
)

func main() {
	// slices are dynamic arrays
	// declaration
	var users []string

	// uninitialized slices are nil
	fmt.Println(users == nil)

	// declaring a slice with an initial size
	var scores = make([]int, 2, 5)  // type, length, capacity
	fmt.Println(scores)    // [0 0]

	// adding elements to our slice
	users = append(users, "harry")
	users = append(users, "john")
	users = append(users, "jane")
	users = append(users, "julie")
	fmt.Println(users)
	
	// copying slices
	var users2 = make([]string, len(users))
	copy(users2, users)
	fmt.Println(users2)

	// slice operator
	var users3 = users2[0:2]
	users3 = append(users3, "paul")
	fmt.Println(users3, len(users3), cap(users3))

	// comparing slice
	fmt.Println(slices.Equal(users, users2))

}