package main

import "fmt"

type data int

func (d1 data) add(d2 data) data{
	return d1+d2
}

// if we try to run this below code
// it will through an error

// func (d1 int) add( d2 int) int {
// 	return d1+d2
// }


// main function

func main(){
	d1 := data(23)
	d2 := data(13)

	res := d1.add(d2)
	fmt.Println(res)
}