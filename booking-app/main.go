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
	var bookings []string

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
	// bookings[0] = firstName + " " + lastName
	bookings = append(bookings, firstName+" "+lastName)

	// fmt.Printf("\nThe whole slice: %v\n", bookings)
	// fmt.Printf("The first value: %v\n", bookings[0])
	// fmt.Printf("Slice type: %T\n", bookings)
	// fmt.Printf("Slice length: %v\n", len(bookings))

	fmt.Printf("\nThank you %v %v for booking %v ticket(s), you will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("Remaining ticket(s) are %v\n", remainingTicket)

	fmt.Printf("These are all our bookings: %v", bookings)
}
