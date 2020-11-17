package main

import "fmt"

type emp struct{
	string
	int
	float32
}

func main(){

	value := emp{"himanshu",21, 15644.23}

	fmt.Print(value.string+"\n", value.int, value.float32)
}