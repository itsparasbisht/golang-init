package main

import "fmt"

func main() {
	counter1 := counter()
	fmt.Println(counter1())
	fmt.Println(counter1())
	fmt.Println(counter1())

	counter2 := counter()
	fmt.Println(counter2())
}

func counter() func() int {
	i := 0

	return func() int{
		i++
		return i
	}
}