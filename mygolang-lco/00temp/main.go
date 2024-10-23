package main

import (
"fmt"
"time"
)

func main() {
	fmt.Println("Welcome to time study in golang")
		presenttime := time.Now()
		fmt.Println("Present time is: ", presenttime)
		fmt.Println(presenttime.Format("01-02-2006 Monday")) // yahi value rehti format ke liye
		// golang ke nasheeeeeeee

		createdDate := time.Date(2020,time.December,20,12,11,0,0,time.UTC)
		fmt.Println("createdDate: ", createdDate)
		fmt.Println("createdDate: ", createdDate.Format("01-02-2006 Monday"))
}
