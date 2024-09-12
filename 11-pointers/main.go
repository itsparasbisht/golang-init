package main

import (
	"fmt"
)

// all basic data types are pass by value
// int, float, bool, string, arrays, structs

// pass by reference
// slices, maps, channels, interfaces

func handleUsers(users [3]string) {
	users[2] = "john"
}

func mutateUsers(users *[3]string){
	users[2] = "john"    // No explicit dereferencing needed for maps & slices
}

func main() {
	users := [3]string{"harry", "louis", "johnny"}

	users2 := [3]string{"harry", "louis", "johnny"}

	handleUsers(users)
	
	mutateUsers(&users2)  // passing the pointer - memeory address of users2 array

	fmt.Println(users)
	fmt.Println(users2)

	firstUserAddress := &users[0]
	fmt.Println(firstUserAddress)  // prints address of first user in users array

	*firstUserAddress = "mike"   // use pointer to modify the array element
	fmt.Println(users)
	fmt.Printf("%T", firstUserAddress)
}