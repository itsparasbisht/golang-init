package main

import "fmt"

const PI = 3.14

func main() {
	var username string = "harry"
	fmt.Printf("username: %s \nvar type: %T \n", username, username)

	var isAdmin bool = false
	fmt.Printf("is admin: %t \nvar type: %T \n", isAdmin, isAdmin)

	var userID int = 100
	fmt.Printf("user ID: %d \nvar type: %T \n", userID, userID)

	var roi float32 = 10.5
	fmt.Printf("roi: %f \nvar type: %T \n", roi, roi)

	// 	Inside a function, the := short assignment statement can be used in place of a var declaration with implicit type.
	// Outside a function, every statement begins with a keyword (var, func, and so on) and so the := construct is not available.
	email := "harry@mail.com"
	fmt.Printf("email: %s \nvar type: %T \n", email, email)

	fmt.Println(PI)
}