/*
	We must first initialize the project!!
	cmd: go init booking-app
*/

package main

import (
	"fmt"
)

func main() {
	var conferenceName = "Golang Conference"
	const conferenceTickets = 50
	var remainingTickets uint = 50 //unit indicates there can only be positive numbers
	//var bookings [50]string // we cannot mix the data types in the array
	bookings := []string{}

	// Use Println to print then move to new line
	//fmt.Println("Welcome to", conferenceName, "booking application")
	fmt.Printf("Welcome to the %v booking application.\n", conferenceName)
	fmt.Printf("There are  a total of %v tickets.\n", conferenceTickets)
	fmt.Printf("There are %v tickets remaining.\n", remainingTickets)
	fmt.Println("Get your tickets here to attend!")

	for {
		var firstName string
		var lastName string
		var email string
		var userTickets uint

		// ask for user input
		fmt.Println("Enter your first name:")
		fmt.Scan(&firstName)

		fmt.Println("Enter your last name:")
		fmt.Scan(&lastName)

		fmt.Println("Enter your email:")
		fmt.Scan(&email)

		fmt.Println("Enter number of tickets:")
		fmt.Scan(&userTickets)

		remainingTickets = remainingTickets - userTickets
		// bookings[0] = firstName+" "+lastName // adding to an array
		bookings = append(bookings, firstName+" "+lastName)

		fmt.Printf("The whole array: %v\n", bookings)
		fmt.Printf("The first value: %v\n", bookings[0])
		fmt.Printf("The array type: %T\n", bookings)
		fmt.Printf("The array length: %v\n", len(bookings))

		fmt.Printf("Thank you %v %v for booking %v tickets.\n", firstName, lastName, userTickets)
		fmt.Printf("You will receive an email notification at %v.\n", email)
		fmt.Printf("%v tickers remaining for %v.\n", remainingTickets, conferenceName)

		fmt.Printf("These are all the bookings %v\n", bookings)
	}

}
