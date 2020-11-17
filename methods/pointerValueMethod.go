package main

import "fmt"

type author struct{
	name string
	branch string
}

// method with pointer
// reciever of author type
func (a *author) show_1(abranch string){
	(*a).branch = abranch
} 


// method with value
// reciver of value type
func ( a author) show_2(){
	a.name = "Chaudhary"
	fmt.Println( "Author name(Before): ",a.name )
}

func main(){

	// initializing the author value

	res := author{
		name: "Himanshu",
		branch: "CSE",
	}

	fmt.Println("Branch name(Before) :", res.branch)

	//calling show_1 method
	// pointer(*) method with value

	res.show_1("ECE")
	fmt.Println("Branch name(After) :", res.branch)

	// calling show_2 method
	// value method with pointer(*)

	(&res).show_2()
	fmt.Println("Author name(After) :", res.name)
}

// This is how we can call pointer method with value and vice versa. 