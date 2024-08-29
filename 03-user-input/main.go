package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("***************** WELCOME *****************")

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Please enter your name: ")
	name, _ := reader.ReadString('\n') // reader.ReadString('\n') returns string and error, we are ignoring the error (comma ok syntax)

	fmt.Print("Please enter your age: ")
	age, _ := reader.ReadString('\n')

	fmt.Printf("Name: %sAge: %s", name, age)

	intAge, error := strconv.ParseFloat(strings.TrimSpace(age), 64)

	if error != nil {
		fmt.Println(error)
	}else{
		fmt.Println("Age + 1 =", intAge + 1)
	}
}