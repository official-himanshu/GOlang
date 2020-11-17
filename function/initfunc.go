package main

import "fmt"

var pass string

func init(){
	fmt.Print("Enter password: ")
	fmt.Scanln(&pass)
}

func print(s string){
	fmt.Println("welcome", s)
}

func main(){

	if( pass == "himanshu"){
		print(pass)
	}else{
		fmt.Println("Please enter valid password")
	}
}



