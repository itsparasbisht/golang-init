package main

import (
	"fmt"
	"strconv"
)

func main() {
	// interpreted string
	txt := "harry\nruns🏃"
	fmt.Println(txt)

	// raw sting
	query := `harry \n runs`
	fmt.Println(query)

	// this will return the total number of bytes
	fmt.Println(len(txt))

	fmt.Println(txt[0])

	// A rune in Go is an alias for the int32 type and represents a Unicode code point
	// A rune literal in Go is a single character enclosed in single quotes
	r := 'A'
	fmt.Println(r)   // prints 65, unicode for A

	// modifying a string using rune
	txt2 := []rune(txt)
	fmt.Println((txt2))
	txt2[0] = 'B'
	str := string(txt2)
	fmt.Println(str)
	
	e := 'Ģ'
	fmt.Printf("unicode %d to binary is %b\n", e, e)

	eToBinary := fmt.Sprintf("%b", e)
	fmt.Println(eToBinary)

	r1 := rune(127757)
	r2 := string(r1)
	fmt.Println(r2)

	b := "11111010"
	num, _ := strconv.ParseInt(b, 2, 64)
	fmt.Println(num)

	
}