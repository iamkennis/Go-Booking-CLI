package main

import (
	"fmt"
	"sync"
	"time"
)

var conferenceName = "Go conference"

const conferenceTicket int = 50

var remainingTicket uint
var bookings = make([]UserData, 0)
var NigeriaNoTicket uint = 50
var NewYorkNoTicket uint = 50
var BerlinNoTicket uint = 50
var ParisNoTicket uint = 50

type UserData struct {
	firstName  string
	lastName   string
	email      string
	city       string
	userTicket uint
}

type CityTickets struct {
	Nigeria string
	London  string
	NewYork string
	Berlin  string
	Paris   string
}

var wg = sync.WaitGroup{}

func main() {

	greetUser()

	firstName, lastName, email, city, userTicket := getUserInput()

	isValidName, isValidEmail, isValidTicketNumber := isValidInput(firstName, lastName, email, userTicket, city, remainingTicket)

	if isValidName && isValidEmail && isValidTicketNumber {
		bookingTicket(userTicket, firstName, lastName, city, email)
		wg.Add(1)
		go sendTicket(userTicket, firstName, lastName, email)

		fullNames := getFullName()
		fmt.Printf("The first names of our bookings are:%v\n", fullNames)

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

		// if !isValidCity {
		// 	fmt.Println("city you entered is not allowed to book ticket")
		// }
	}

	switch city {
	case "Nigeria":
		remainingTicket = NigeriaNoTicket - userTicket
		fmt.Printf("Ticket booked for %v and %v ticket remains\n", city, remainingTicket)
	case "NewYork":
		remainingTicket = NewYorkNoTicket - userTicket
		fmt.Printf("Ticket booked for %v and %v ticket remains\n", city, remainingTicket)
	case "Berlin":
		remainingTicket = BerlinNoTicket - userTicket
		fmt.Printf("Ticket booked for %v and %v ticket remains\n", city, remainingTicket)
	case "Paris":
		remainingTicket = ParisNoTicket - userTicket
		fmt.Printf("Ticket booked for %v and %v ticket remains\n", city, remainingTicket)
	default:
	}

	wg.Wait()
}

func greetUser() {
	fmt.Printf("Welcome to %v booking application.\n", conferenceName)
	fmt.Printf("We have the total of %v tickets and %v are still available.\n", conferenceTicket, remainingTicket)
	fmt.Println("Get your conference tickets")
}

func getFullName() []string {
	fullNames := []string{}
	for _, booking := range bookings {
		fullNames = append(fullNames, booking.firstName+" "+booking.lastName)
	}

	return fullNames
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
	// remainingTicket = ticket - userTicket

	userData := UserData{
		firstName:  firstName,
		lastName:   lastName,
		email:      email,
		city:       city,
		userTicket: userTicket,
	}

	bookings = append(bookings, userData)
	fmt.Printf("List of bookings %v \n", bookings)
	fmt.Printf("Thank you %v %v for booking %v tickets from %v. You will receive a confirmation email at %v \n ", firstName, lastName, userTicket, city, email)
	fmt.Printf("Thank you for coming %v\n", conferenceName)
}

func sendTicket(userTicket uint, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second)
	ticket := fmt.Sprintf("%v tickets for %v %v", userTicket, firstName, lastName)
	fmt.Printf("##########################\n")
	fmt.Printf("Sending ticket:\n %v to email address %v\n", ticket, email)
	fmt.Printf("##########################")

	wg.Done()
}
