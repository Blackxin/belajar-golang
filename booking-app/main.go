package main

import "fmt"

func main() {
	conferenceName := "Go Conference"
	const conferenceTicket uint8 = 50
	var remainingTicket uint8 = 50

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available\n", conferenceTicket, remainingTicket)
	fmt.Printf("Get your tickets here to attend\n\n")

	var firstName string
	var lastName string
	var email string
	var userTickets uint8
	// ask users for their personal informations
	fmt.Print("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Print("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Print("Enter your email: ")
	fmt.Scan(&email)

	fmt.Print("Enter number of ticket(s): ")
	fmt.Scan(&userTickets)

	remainingTicket -= userTickets

	fmt.Printf("\nThank you %v %v for booking %v ticket(s), you will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("Remaining ticket(s) are %v", remainingTicket)
}
