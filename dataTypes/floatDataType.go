package main

import "fmt"

func main(){

	var A1 float32 = 30.22
	var A2 float64 = 26.33

	// here if we declare variable dynamically so it takes base64 values
	a:= 33.33
	b:= 21.2

	var c float64 = a+b
	d := a+b

	fmt.Println(A1, A2, A2-float64(A1))
	fmt.Println(c, d)
}