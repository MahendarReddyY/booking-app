package main

import (
	"booking-app/message"
	"fmt"
	"sync"
)

var ConferenceName = "Go Conference"

const ConferenceTickets = 50

var RemainingTicekts = 50
var bookingList = make([]info, 0)

type info struct {
	firstName string
	lastName  string
	email     string
	tickets   int
}

var firstName string
var lastName string
var email string
var tickets int

var wg = sync.WaitGroup{}

func main() {

	mesg()

	//for RemainingTicekts > 0 && len(bookingList) < 50 {

	userinput()

	isvalidName, isvalidEmail, isvalidTickets := message.Evaluated(firstName, lastName, email, tickets, RemainingTicekts)

	if isvalidEmail && isvalidName && isvalidTickets {

		booking()
		wg.Add(1)
		go sendingpacket()
		firstname()

		if RemainingTicekts == 0 {
			fmt.Println("The tickets are sold out. Come back for the next year")
			//break
		}
	} else {
		validcheck(isvalidEmail, isvalidName, isvalidTickets)
	}
	wg.Wait()

}

func firstname() {
	firstNames := []string{}
	for _, bookingList := range bookingList {
		firstNames = append(firstNames, bookingList.firstName)
	}
}

func userinput() {

	fmt.Println("Enter your firstName:")
	fmt.Scanln(&firstName)

	fmt.Println("Enter your lastName:")
	fmt.Scanln(&lastName)

	fmt.Println("Enter your email:")
	fmt.Scanln(&email)

	fmt.Println("Enter number of tickets:")
	fmt.Scanln(&tickets)
}

func booking() {
	RemainingTicekts = RemainingTicekts - tickets
	var userData = info{
		firstName: firstName,
		lastName:  lastName,
		email:     email,
		tickets:   tickets,
	}

	bookingList = append(bookingList, userData)

	fmt.Println("this is the list of bookings:", bookingList)

	fmt.Printf("Thank you %v %v for booking  %v tickets and we will send a confirmation to your %v input email.The remaning ticktes in this event are %v. \n", firstName, lastName, tickets, email, RemainingTicekts)

}

func validcheck(isvalidEmail bool, isvalidName bool, isvalidTickets bool) {
	if !isvalidEmail {
		fmt.Println("your email address must contains @ and .com signs ")
	}
	if !isvalidName {
		fmt.Println("your name is too short ")
	}
	if !isvalidTickets {
		fmt.Println("enter required tickets values from 1 to 50 ")
	}
}
