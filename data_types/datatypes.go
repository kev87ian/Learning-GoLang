package main

import "fmt"

func main() {

	var userName string
	var userTickets int
	//ask user their name

	userName = "Kefini"
	userTickets = 2
	fmt.Printf("User's name is %v, and he has booked %v tickets.\n", userName, userTickets)

	//printing the datatype of variables
	fmt.Printf("username is %T, userTickets is %T", userName, userTickets)
}
