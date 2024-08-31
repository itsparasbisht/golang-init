package main

import (
	"fmt"
	"time"
)

func main() {
	// normal for loop
	for i := 0; i <= 10; i++ {
		fmt.Println(i)
	}

	// for loop with range
	for i:= range 5{
		fmt.Println(i)
	}

	// if else block
	role := "ADMIN"

	if role == "ADMIN" {
		fmt.Println("access granted")
	} else{
		fmt.Println("unauthorized")
	}

	// switch case block
	weekday := time.Now().Weekday()

	switch weekday {
	case time.Saturday, time.Sunday:
		fmt.Println("it's weekend!")
	case time.Monday:
		fmt.Println("a lot is pending in to do")
	default:
		fmt.Println("whatever")
	}

	// check the variable type
	varType := func(i any){
		switch i.(type){
		case int:
			fmt.Println("integer")
		case string:
			fmt.Println("string")
		case bool:
			fmt.Println("boolean")
		case float32, float64:
			fmt.Println("float")
		default:
			fmt.Printf("%T", i)
		}
	}

	varType("GO")
	varType(time.Saturday)
}