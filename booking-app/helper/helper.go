package helper

import "strings"

func ValidateUserInput(firstName string, lastName string, email string, userTickets uint8, remainingTicket uint8) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTicket

	return isValidName, isValidEmail, isValidTicketNumber
}
