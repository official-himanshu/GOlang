package main

import "fmt"

func main(){
	// we can assign a anonymous function to a variable and can pass argumnet to a function
	value := func(a, b int){
	//var a int = 10
	//var b int = 20

	fmt.Println(a+b)
	}
	value(10,20)
}