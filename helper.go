package main

import "strings"

func isValidInput(firstName string, lastName string, email string, userTicket uint, city string, remainingTicket uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTicket > 0 && userTicket >= remainingTicket
	// isValidCity := city == "London" || city == "NewYork"

	return isValidName, isValidEmail, isValidTicketNumber
}
