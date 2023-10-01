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
	var remainingTickets = 50

	// Use Println to print then move to new line
	//fmt.Println("Welcome to", conferenceName, "booking application")
	fmt.Printf("Welcome to the %v booking application.\n", conferenceName)
	fmt.Printf("There are  a total of %v tickets.\n", conferenceTickets)
	fmt.Printf("There are %v tickets remaining.\n", remainingTickets)
	fmt.Println("Get your tickets here to attend!")

}
