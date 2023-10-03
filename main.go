/*
	We must first initialize the project!!
	cmd: go init booking-app
*/

package main

import (
	"booking-app/helper" // booking app is the path from the go.mod file
	"fmt"
	"strings"
)

// Package level variable so its only visbile the main package
var conferenceName string = "Golang Conference"
var remainingTickets uint = 50 //unit indicates there can only be positive numbers
var bookings = []string{}

const conferenceTickets int = 50

func main() {

	greetUsers()

	// Loop with a conditional check
	// for remainingTickets > 0 && len(bookings) < 50
	// for true is infiite loop is the same as for
	for {
		firstName, lastName, email, userTickets := getUserInput()

		// The func returns 3 values
		isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidTicketNumber {
			bookTicket(userTickets, firstName, lastName, email)
			fmt.Printf("These are the first names of the attendess: %v.\n", getsFirstNames())

		} else if !isValidName || !isValidEmail || !isValidTicketNumber {

			// We want to check all the values, hence the use of multiple ifs
			if len(firstName) < 2 {
				fmt.Printf("Your first name was invalid input, %v.\n", firstName)
			}
			if len(lastName) < 2 {
				fmt.Printf("Your last name was invalid input, %v.\n", lastName)
			}
			if !isValidEmail {
				fmt.Printf("Your email was invalid, %v.\n", email)
			}
			if !isValidTicketNumber {
				fmt.Printf("The number of tickets was invalid, %v.\n", userTickets)
			}

		} // end of if-else statement

		// How many tickets are remainging
		if remainingTickets == 0 {
			// end the program
			fmt.Printf("The %v is fully booked. Come back next year!\n", conferenceName)
			break
		}

	} // end of main loop

	/*
		city := "london"

		switch city {
		case "new york":
			// code block
		case "london", "berlin":
			// same logic for london and berlin
		case "las vegas", "phoenix":
			// code block
		case "mexico city":
			// code block
		default:
			// handle incorrect value entered
			fmt.Println("No valid city selected!")
		}
	*/
} // end of main func

func greetUsers() {
	fmt.Printf("Welcome to the %v booking application.\n", conferenceName)
	fmt.Printf("There are a total of %v tickets.\n", conferenceTickets)
	fmt.Printf("There are %v tickets remaining.\n", remainingTickets)
	fmt.Println("Get your tickets here to attend!")
}

// the []string is the output and indicating its type
func getsFirstNames() []string {
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
	//fmt.Printf("These are first names of bookings are: %v\n", firstNames)
	// Instead of printing the data.  We want the main func to print the dat
	return firstNames
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Println("Enter your first name:")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name:")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email:")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets:")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	// Reduce the number of remaining tickets
	remainingTickets = remainingTickets - userTickets

	// bookings[0] = firstName+" "+lastName // adding to an array
	bookings = append(bookings, firstName+" "+lastName)

	fmt.Printf("Thank you %v %v for booking %v tickets.\n", firstName, lastName, userTickets)
	fmt.Printf("You will receive an email notification at %v.\n", email)
	fmt.Printf("%v tickets remaining for the %v.\n", remainingTickets, conferenceName)
}
