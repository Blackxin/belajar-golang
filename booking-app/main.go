package main

import (
	"fmt"
	"strings"
)

func main() {
	conferenceName := "Go Conference"
	const conferenceTicket uint8 = 50
	var remainingTicket uint8 = 50

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available\n", conferenceTicket, remainingTicket)
	fmt.Printf("Get your tickets here to attend\n\n")

	var bookings []string

	for remainingTicket > 0 {

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

		isValidName := len(firstName) >= 2 && len(lastName) >= 2
		isValidEmail := strings.Contains(email, "@")
		isValidTicketNumber := userTickets > 0 && userTickets <= remainingTicket

		if isValidName && isValidEmail && isValidTicketNumber {
			remainingTicket -= userTickets
			bookings = append(bookings, firstName+" "+lastName)

			fmt.Printf("\nThank you %v %v for booking %v ticket(s), you will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
			fmt.Printf("Remaining ticket(s) are %v\n", remainingTicket)

			firstNames := []string{}
			for _, booking := range bookings {
				var names = strings.Fields(booking)
				var firstName = names[0]
				firstNames = append(firstNames, firstName)
			}

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
}
