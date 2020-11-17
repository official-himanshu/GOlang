package main

import "fmt"

func main(){

	// if you are using y but not using so compiler throw an error
	//x, y := mul_div(10, 5)

	// instead of using y use _ (blank identifier)
	x, _ := mul_div(10, 5)

	fmt.Println("mul is: ", x)
}

func mul_div(x, y int) (int, int){

	return x*y, x/y
}