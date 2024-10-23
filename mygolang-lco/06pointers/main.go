package main

import "fmt"

func main() {
	fmt.Println("Let's study pointers in golang")

	//var ptr *int // pointer variable

	//fmt.Println("Value of ptr is: ", ptr) // nil

	num := 23
	var ptr = &num // reference

	fmt.Println("value 1: ", ptr)  // address of num
	fmt.Println("value 2: ", *ptr) // 23

}
