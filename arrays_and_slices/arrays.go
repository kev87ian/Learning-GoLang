package main

import "fmt"

func main() {

	//arrays
	var arrayBookings [50]string
	var name1 string
	var name2 string
	fmt.Scan(&name1)
	fmt.Scan(&name2)
	//to add an item to array, we use indices

	arrayBookings[0] = name1 + " " + name2

	fmt.Printf("The whole array: %v \n", arrayBookings)
	fmt.Printf("The first item in the array: %v \n", arrayBookings[0])

	fmt.Printf("The array type: %T \n", arrayBookings)
	fmt.Printf("The length of the array: %v \n", len(arrayBookings))

	//slice
	var bookings []string
	var firstName string
	var lastName string
	fmt.Scan(&firstName)
	fmt.Scan(&lastName)

	bookings = append(bookings, firstName+" "+lastName)

	fmt.Printf("The whole slice: %v \n", bookings)
	fmt.Printf("The first item in the slice: %v \n", bookings[0])

	fmt.Printf("The slice type: %T \n", bookings)
	fmt.Printf("The length of the slice: %v \n", len(bookings))
}
