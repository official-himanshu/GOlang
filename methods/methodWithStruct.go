package main

import "fmt"

// structure of type author
type author struct{
	name string
	branch string
	age int
}

// method to print author information
func ( a author ) show(){
	fmt.Println( "Author name : ", a.name )
	fmt.Println( "Author branch : ", a.branch )
	fmt.Println( "Author age : ", a.age )
}

func main(){

	// creating struct type variable just like object in OOPS
	// initializing the structure
	value := author{
		name: "Himanshu",
		branch: "CSE",
		age: 21,
	}

	// calling method show
	value.show()
}