package main

import "fmt"

type student struct{
	name string
	branch string
	rollno int
}

type teacher struct{
	name string
	subject string
	details student   // defining structure of type student inside teacher structure
}

func main(){

	s1 := student{"Himanshu","cse",1603210078}
	s2 := student{"Harshit","12th",0}

	t1 := teacher{"shalini","maths",s1}
	t2 := teacher{"geeta","primary",s2}

	fmt.Println("Teacher 1 details: ",t1)
	fmt.Println("Teacher 2 details: ",t2)

	// accessing data of student using teacher

	fmt.Println("name of the first student ", t1.details.name)
	fmt.Println("name of the second student ", s2.name)
}