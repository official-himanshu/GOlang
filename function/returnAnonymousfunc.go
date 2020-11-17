package main	

import "fmt"

// return type of sum() function is func() and return type of func() function return a+b
func sum() func(a, b  int) int {
	myf := func(a , b int) int {
		return a+b
	}

	return myf
}

func main(){

	//calling sum function and store return value in a var
	value := sum()

	// calling anonymous function from return value of sum() function
	fmt.Println(value(10,20))
}