package main

import (
	"bufio"
	"fmt"
	"os"
	// "github.com/go-delve/delve/pkg/dwarf/reader"
)

func main() {
	wel := "Welcome to user input sesh"
	fmt.Println(wel)

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter a number: ")

	// comma ok || err ok syntax 
	input, _ := reader.ReadString('\n') // tab tak input jab tak \n nahi milta
	// _ is error part  yaha kaam nhi uska isliye underscore
	// agar input ka kaam nhi toh _,err := reader.ReadString('\n') likhna padega 
	fmt.Println("You entered:", input) 
	fmt.Printf("type of input is %T", input)
}
	