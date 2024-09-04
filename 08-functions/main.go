package main

import "fmt"

func main() {
	sum := addNum(2, 3)
	fmt.Println(sum)

	name, id := getUserDetails()
	fmt.Println(id, "->", name)

	sum = variadicSum(3,4,5,6,7,8,9)
	fmt.Println(sum)

	numSlice := []int{1, 2, 3, 4, 5}
	sum = variadicSum(numSlice...)
	fmt.Println(sum)
}

func addNum(a int, b int) int {
	return a + b
}

// multiple return values
func getUserDetails() (string, int){
	return "harry", 111
}

// Variadic Functions - functions that can accept variable number of arguments, eg. fmt.Println()
func variadicSum(nums ...int) int{
	total := 0

	for _, num:= range nums {
		total += num
	}

	return total
}