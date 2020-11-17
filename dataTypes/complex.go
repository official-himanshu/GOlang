package main

import "fmt"

func main(){

	var X complex64 = complex(2,6)
	var Y complex128 = complex(3,2)

	a := complex(5,6)
	fmt.Println(X, Y+a)
	fmt.Printf("%T "+" %T", X, Y+a)
}