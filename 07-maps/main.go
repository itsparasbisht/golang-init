package main

import (
	"fmt"
	"maps"
)

func main() {
	// declaring an empty map
	fd := make(map[string]string)
	fmt.Println(fd)

	// declaring & initializing a map
	user := map[string]string{"name": "harry", "email": "harry@mail.com"}

	// adding key/value
	user["id"] = "111"

	fmt.Println(user)
	fmt.Println("user id: ", user["id"])
	fmt.Println(len(user))

	// returns zeroed value if the key doesn't exist
	fmt.Println(user["language"])  // empty string

	// to check if the key exists or not use this
	val, exists := user["language"]
	if exists == false{
		fmt.Println("language is not available!")
	}else{
		fmt.Println("language:", val)
	}

	delete(user, "id")
	fmt.Println(user)

	user2 := maps.Clone(user)
	fmt.Println(user2)

	if maps.Equal(user, user2) {
		fmt.Println("user == user2")
	}

	// clear(user)
	// fmt.Println(user)

	for i:= range maps.Keys(user){
		fmt.Println(i)
	}
}