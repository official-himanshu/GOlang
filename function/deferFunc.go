package main

import "fmt"

func add(x, y int) int {
	fmt.Println(x+y)
	return 0
}

func main(){
	fmt.Println("Start")
	// when we call function as a defer so it execute at last in LIFO order.
	defer fmt.Println("End")
	defer fmt.Println("FIRST END")

	add(10, 20)
	add(2, 43)
}