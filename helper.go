package main

import "strings"

func validateUserInput(firstName string, lastName string, email string, userTickets uint) (bool, bool, bool) {
	// Validate first and last name. both must be atleast 2 chars
	// both conditions must be true
	isValidName := len(firstName) >= 2 && len(lastName) >= 2

	// Validate the email entered is in a correct format
	isValidEmail := strings.Contains(email, "@")

	// Validate the number of tickets. Value must be a positive numbers
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets

	// Example of an OR conditional
	//isValidCity := city == "las vegas" || city =="fallon"
	return isValidName, isValidEmail, isValidTicketNumber
}
