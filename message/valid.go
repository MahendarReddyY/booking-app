package message

import (
	"strings"
)

func Evaluated(firstName string, lastName string, email string, tickets int, RemainingTicekts int) (bool, bool, bool) {
	isvalidName := len(firstName) >= 2 && len(lastName) >= 2
	isvalidEmail := strings.Contains(email, "@")
	isvalidTickets := tickets > 0 && tickets <= RemainingTicekts
	return isvalidName, isvalidEmail, isvalidTickets
}
