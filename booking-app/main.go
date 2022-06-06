package main

import (
	"booking-app/helper"
	"fmt"
	"strings"
)

const conferenceTicket uint8 = 50

var conferenceName string = "Go Conference"
var remainingTicket uint8 = 50
var bookings []string

func main() {
	greetUsers()

	for remainingTicket > 0 {
		// getting user input
		firstName, lastName, email, userTickets := getUserInput()

		isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTicket)

		if isValidName && isValidEmail && isValidTicketNumber {
			// booking tickets
			remainingTicket, bookings = bookTicket(userTickets, firstName, lastName, email)

			// call function print first names
			firstNames := printFirstNames()
			fmt.Printf("These are all our bookings: %v\n\n", firstNames)

			if remainingTicket == 0 {
				fmt.Println("Our conference is booked out. Come back next year.")
				break
			}
		} else {
			fmt.Println("")
			if !isValidName {
				fmt.Println("First name or last name you entered is too short")
			}
			if !isValidEmail {
				fmt.Println("Your email doesn't contain @ symbol")
			}
			if !isValidTicketNumber {
				fmt.Println("Number of tickets you entered is invalid")
			}
			fmt.Println("")
		}
	}

	// city := "London"

	// switch city {
	// 	case "New York":
	// 		// execute code for booking New York conference tickets
	// 	case "Singapore", "Hong Kong":
	// 		// execute code for booking Singapore & Hong Kong conference tickets
	// 	case "London", "Berlin":
	// 		// execute code for booking London & Berlin conference tickets
	// 	default:
	// 		fmt.Print("No valid city selected")
	// }
}

func greetUsers() {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available\n", conferenceTicket, remainingTicket)
	fmt.Printf("Get your tickets here to attend\n\n")
}

func printFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		var names = strings.Fields(booking)
		var firstName = names[0]
		firstNames = append(firstNames, firstName)
	}

	return firstNames
}

func getUserInput() (string, string, string, uint8) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint8

	// ask users for their personal informations
	fmt.Print("Enter your first name     : ")
	fmt.Scan(&firstName)

	fmt.Print("Enter your last name      : ")
	fmt.Scan(&lastName)

	fmt.Print("Enter your email          : ")
	fmt.Scan(&email)

	fmt.Print("Enter number of ticket(s) : ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint8, firstName string, lastName string, email string) (uint8, []string) {
	remainingTicket -= userTickets
	bookings = append(bookings, firstName+" "+lastName)

	fmt.Printf("\nThank you %v %v for booking %v ticket(s), you will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("Remaining ticket(s) are %v\n", remainingTicket)

	return remainingTicket, bookings
}
