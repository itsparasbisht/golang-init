package main

import (
	"fmt"
	"reflect"
)

// structs are used to define custom data types
type User struct {
	id    int
	name  string
	email string
}

func main() {
	user1 := User{1, "harry", "harry@mail.com"}
	// user1 := user{id:1, name: "harry", email: "harry@mail.com"}  // this also works

	// structs are mutable
	user1.id = 2

	fmt.Println(user1)
	
	t := reflect.TypeOf(user1)
	fmt.Println(t)

	// zero value initialization
	var user2 User
	fmt.Println(user2)

	// passing a struct to a func will create a copy
	updateUser(user1)
	fmt.Println(user1)

	// Pointers allow you to pass structs efficiently and modify them without copying the entire object.
	updateUser2(&user1)
	fmt.Println(user1)

	// anonymous structs, they're helpful for simple, temporary structures
	address := struct {
		city string
		state string
	}{"Dehradun", "Uttarakhand"}

	fmt.Println(address)

	user4 := User{1, "john", "john@mail.com"}
	user5 := User{1, "john", "john@mail.com"}
	fmt.Println(user4 == user5)   // true

	// embedding structs
	type Employee struct{
		department string
		user_details User
	}

	e1 := Employee{department: "sales", user_details: User{id: 10, name: "jane", email: "jane@mail.com"}}
	fmt.Println(e1)
}

func updateUser(u User){
	u.id = 5    // this will not mutate the original struct, because it's values are copied into u
	fmt.Println(u)
}

func updateUser2(u *User){
	u.id = 5    // this will mutate the original struct, because the reference is passed here
}