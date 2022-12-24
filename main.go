package main

import "fmt"

func main() {

	var conferenceName = "Go Conference"
	const conferenceTickets int = 50
	var remainingTickets = 50

	/*Assigning datatypes to variables ensures that we protect our variables from getting a value they're not supposed to get. */
	fmt.Printf("Welcome to the %v. We have a total of %v tickets, and %v ticktes remain so far.\n", conferenceName, conferenceTickets, remainingTickets)
	fmt.Printf("conferenceName is of %T datatype, and conferenceTickets is of %T datatype.\n", conferenceName, conferenceTickets)

	//ASKING FOR USER INPUT

	//& is called a pointer.
	// A pointer is a variable that points  to the memory address of another variable.
	//pointers are also called special variables

	fmt.Println(conferenceName)
	//line above prints the value of the variable, while the line below prints the memory address of the variable.
	//fmt.Println(&conferenceName)

	var firstName string
	var lastName string
	var email string
	var userTickets int

	fmt.Println("Enter your first name to proceed")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email address: ")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets you want: ")
	fmt.Scan(&userTickets)

	remainingTickets = remainingTickets - userTickets

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v.\n", firstName, lastName, userTickets, email)
	fmt.Printf("We have %v remaining tickets for the %v.\n", remainingTickets, conferenceName)
}
