package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to my pizza app")

	fmt.Println("Please enter the rating of this pizza: ")

	reader := bufio.NewReader(os.Stdin)
	input,_ := reader.ReadString('\n')
	fmt.Println("Thanks for rating, ", input)

	// conversions 
	numRating,err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	if err != nil {
		fmt.Println("Error: ", err)
		// panic(err)// panic is a function that stops the execution of the program
	}else{
		fmt.Println("added 1 to rating, ", numRating+1)
	}


}