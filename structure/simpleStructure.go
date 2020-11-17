package main

import "fmt"

type address struct{
	line string
	street string
	area string
	city string
	pincode int
}

func main(){

	var a address

	fmt.Println(a)

	a1 := address{"house no:23","budh bazar road","teacher colony","gajraula", 244235}
	a2 := address{"house no:26","budh bazar road","teacher colony","gajraula", 244234}

	fmt.Println("Address1:" ,a1)
	fmt.Println("Address2:", a2)
}