package main

import (
	"fmt"
	"strconv"
)

func main() {

	users := make(map[int]string)

	for i := range 10 {
		users[i] = "user-" + strconv.Itoa(i)
	}

	fmt.Print(users)

	for key, value := range users{
		fmt.Printf("key: %d | value: %s\n", key, value)
	}
}