package main

import (
	"fmt"
)

func main(){

	// anonymous function
	/*func(){
		fmt.Println("Hello from anonymous function")
	}()*/

	func(ele string){
		fmt.Println("hello",ele)
	}("Himanshu")

}