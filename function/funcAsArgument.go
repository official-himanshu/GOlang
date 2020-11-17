package main

import "fmt"

func add(a func(x, y int) int){
	fmt.Println(a (10,20))
}
func main(){
	value := func(x, y int) int{
		return x+y
	}
	add(value)

}