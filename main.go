package main

import (
	"booking-app/helper"
	"fmt"
	"strings"
)

var conferenceName = "Go conference"

const conferenceTicket int = 50

var remainingTicket uint = 50
var bookings = []string{}

func main() {

	greetUser()

	for {

		firstName, lastName, email, city, userTicket := getUserInput()

		isValidName, isValidEmail, isValidTicketNumber, isValidCity := helper.IsValidInput(firstName, lastName, email, userTicket, city, remainingTicket)

		if isValidName && isValidEmail && isValidTicketNumber && isValidCity {
			bookingTicket(userTicket, firstName, lastName, city, email)

			firstNames := getFirstName()
			fmt.Printf("The first names of our bookings are:%v\n", firstNames)

		} else {
			if remainingTicket == 0 {
				fmt.Println("Conference tickets sold out, come back next year")
			}

			if !isValidName {
				fmt.Println("first name or last name is too short")
			}
			if !isValidEmail {
				fmt.Println("email address is invaild")
			}
			if !isValidTicketNumber {
				fmt.Println("number of ticket you entered is invaild")
			}

			if !isValidCity {
				fmt.Println("city you entered is not allowed to book ticket")
			}
		}

	}
}

func greetUser() {
	fmt.Printf("Welcome to %v booking application.\n", conferenceName)
	fmt.Printf("We have the total of %v tickets and %v are still available.\n", conferenceTicket, remainingTicket)
	fmt.Println("Get your conference tickets")
}

func getFirstName() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		names := strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}

	return firstNames
}

func getUserInput() (string, string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var city string
	var userTicket uint

	fmt.Println("Enter your first name:")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name:")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email:")
	fmt.Scan(&email)

	fmt.Println("Enter your city:")
	fmt.Scan(&city)

	fmt.Println("Enter number of ticket:")
	fmt.Scan(&userTicket)

	return firstName, lastName, email, city, userTicket
}

func bookingTicket(userTicket uint, firstName string, lastName string, city string, email string) {
	remainingTicket = remainingTicket - userTicket
	bookings = append(bookings, firstName+" "+lastName)
	fmt.Printf("Booking name list %v\n", bookings)
	fmt.Printf("Thank you %v %v for booking %v tickets from %v. You will receive a confirmation email at %v \n ", firstName, lastName, userTicket, city, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTicket, conferenceName)
}
