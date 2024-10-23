package main

import "fmt"

const LoginToken string = "123456" // public

func main() {
	var username string = "tarun"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n",username)

	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T \n",isLoggedIn)

	var val uint8 = 255
	fmt.Println(val)
	fmt.Printf("Variable is of type: %T \n",val)

	var nextguy int
	fmt.Println(nextguy)

	// implicit type
	var something = "lexer sambhal leta ye type"
	fmt.Println(something)

	// no var type 
	// walrus operator - can't be used main ke bahar (globally nhi work krega)
	noofuser := 3000
	fmt.Println(noofuser)

		fmt.Println("Login Token is: ",LoginToken)
	
}

