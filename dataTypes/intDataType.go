package main

import "fmt"

func main(){
	var X1 int8 = 127
	var X2 uint8 = 255
	// signed integer value to print data ranges from -128 to 127
	fmt.Println(X1+1, X1-1)
	// signed integer value to print data ranges from 0 to 255
	fmt.Println(X2+1, X2-1)

	// signed integer value to print data ranges from -32768 to 32767
	var Y1 int16 = 32767
	// signed integer value to print data ranges from 0 to 65535
	var Y2 uint16 = 65535
	fmt.Println(Y1+1, Y1-1, Y2+1, Y2-1)

	// signed integer value to print data ranges from -(2^32)/2 to ((2^32)/2)-1
	var Z1 int32 = 2147483647
	// signed integer value to print data ranges from 0 to (2^32)-1
	var Z2 uint32 = 4294967295

	fmt.Println(Z1+1, Z1-1, Z2+1, Z2-1)

	var A uint = 4294967295
	fmt.Println(A+2, A)
}