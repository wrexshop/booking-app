/*
	We must first initialize the project!!
	cmd: go init booking-app
*/

package main

import (
	"fmt"
	"strings"
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

		// Handle invalid from the user input
		if userTickets > remainingTickets {
			fmt.Printf("There are only %v tickets remaining! So you cannot book %v tickets", remainingTickets, userTickets)
			// Force the user to fill out the information once again by using "continue"
			continue
		}

		remainingTickets = remainingTickets - userTickets
		// bookings[0] = firstName+" "+lastName // adding to an array
		bookings = append(bookings, firstName+" "+lastName)

		fmt.Printf("The whole array: %v\n", bookings)
		fmt.Printf("The first value: %v\n", bookings[0])
		fmt.Printf("The array type: %T\n", bookings)
		fmt.Printf("The array length: %v\n", len(bookings))

		fmt.Printf("Thank you %v %v for booking %v tickets.\n", firstName, lastName, userTickets)
		fmt.Printf("You will receive an email notification at %v.\n", email)
		fmt.Printf("%v tickets remaining for the %v.\n", remainingTickets, conferenceName)

		// Slice only for first names
		firstNames := []string{}

		// returns two values the position and thing
		// "range" is how we iterate through the array and slice
		// "_" replaced the "index" as it is not used at all.
		for _, booking := range bookings {

			// This will split the names up. So it will have two values: first and last name
			var full_name = strings.Fields(booking)

			// We will only add the first name to the firstNames slice
			firstNames = append(firstNames, full_name[0])
		}

		fmt.Printf("These are first names of bookings are: %v\n", firstNames)

		// How many tickets are remainging
		if remainingTickets == 0 {
			// end the program
			fmt.Printf("The %v is fully booked. Come back next year!\n", conferenceName)
			break
		}

	} // main for loop

} // main func
