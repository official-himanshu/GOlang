package main

import "fmt"

type emp struct{
	fn, ln string
	age, salary int
}

func main(){

	// taking pointer to struct

	emp8 := &emp{"Himanshu", "Chaudhary", 21, 15000}

	//printing name and salary of emp8 with *emp8 and emp8

	fmt.Println("Full name of employee is : ",(*emp8).fn, emp8.ln)
	fmt.Println("Age and Salary employee is : ",emp8.age, (*emp8).salary)
}