package helper

import "strings"

func ValidateUserInput(firstName string, lastName string, email string, ticket uint, remeaningTicket uint) (bool, bool, bool) {

	isValidName := len(firstName) >= 2 && len(lastName) >= 2 || len(lastName) == 0
	isValidEmail := strings.Contains(email, "@")
	isValidTicket := remeaningTicket > 0 && ticket <= remeaningTicket

	return isValidName, isValidEmail, isValidTicket
}
