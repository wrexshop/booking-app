/*
	We must first initialize the project!!
	cmd: go init booking-app
*/

package main

import (
	"booking-app/helper" // booking app is the path from the go.mod file
	"fmt"
	"sync"
	"time"
)

// Package level variable so its only visbile the main package
var conferenceName string = "Golang Conference"
var remainingTickets uint = 50 //unit indicates there can only be positive numbers
// var bookings = make([]map[string]string, 0) // This is a list of maps. because it a list we need to set a initial size. can be zero
var bookings = make([]UserData, 0) // This is a list containing data based on the User Data struct

const conferenceTickets int = 50

// type = a custom type of data
// similar to classes in other programming languages
// but it cannot be inherited
type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

// Tell the main thread to wait for all actions to complete before exiting
// By default the main func will exit
var wg = sync.WaitGroup{}

func main() {

	greetUsers()

	firstName, lastName, email, userTickets := getUserInput()

	// The func returns 3 values
	isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

	if isValidName && isValidEmail && isValidTicketNumber {
		bookTicket(userTickets, firstName, lastName, email)

		wg.Add(1)
		// Concurrency! Simply by adding "go"
		go sendTicket(userTickets, firstName, lastName, email)

		fmt.Printf("These are the first names of the attendees: %v.\n", getsFirstNames())

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
	}
	// Tell the main func to wait for the concurrent action to complete
	wg.Wait()

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
		firstNames = append(firstNames, booking.firstName)
	}

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

	/* Using a map for a user data. This will be [email:x, firstName:x, lastName:x, numberOfTickets:x]
	var userData = make(map[string]string)
	// add data to the map
	userData["firstName"] = firstName
	userData["lastName"] = lastName
	userData["email"] = email
	// since the is configured to take strings, we must convert the integer to str
	// the second parameter for FormatUint is base 10 / decimal
	userData["numberOfTickets"] = strconv.FormatUint(uint64(userTickets), 10)

	// bookings[0] = firstName+" "+lastName // adding to an array
	// bookings = append(bookings, firstName+" "+lastName) // adding to a slice
	bookings = append(bookings, userData) // adding the map to the list
	*/

	// Using a struct to assign user data
	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}

	bookings = append(bookings, userData)
	fmt.Printf("List of bookings: %v.\n", bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets.\n", firstName, lastName, userTickets)
	fmt.Printf("You will receive an email notification at %v.\n", email)
	fmt.Printf("%v tickets remaining for the %v.\n", remainingTickets, conferenceName)
}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second)
	// we can store the string into a variable using Sprintf
	var message = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("##########################")
	fmt.Printf("Sending ticket:\n%v\nTo email address: %v.\n", message, email)
	fmt.Println("##########################")
	// The tell main func that this func has completed its action
	wg.Done()
}
