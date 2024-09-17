package main

import "fmt"

// interfaces are designed to define a set of one or more method signatures (i.e., the behavior)
// any type that implements those methods satisfies the interface.

// it defines the "behavior" that a type must follow
type DetailsPrinter interface{
	printDetails()
}

type Employee struct{
	id int
	name string
	email string
}

func (e Employee) printDetails(){
	fmt.Printf("----------\nEmployee details\nid: %d\nname: %s\nemail: %s\n", e.id, e.name, e.email)
}

type Department struct{
	id int
	name string
	HoD string
	total_employee int
}

func (d Department) printDetails(){
	fmt.Printf("----------\nDepartment details\nid: %d\nname: %s\nHoD: %s\ntotal employees: %d\n", d.id, d.name, d.HoD, d.total_employee)
}

// both Employee & Department satisfies the interface "DetailsPrinter" as they implement the methods defined in the interface

// Function that accepts any type that implements DetailsPrinter interface
func printDetails(s DetailsPrinter){
	s.printDetails()
}

func main(){
	e1 := Employee{1, "harry", "harry@mail.com"}
	printDetails(e1)

	d1 := Department{101, "marketing", "john doe", 30}
	printDetails(d1)
}