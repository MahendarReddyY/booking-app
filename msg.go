package main

import "fmt"

func mesg() {
	fmt.Println("Welcome to Our", ConferenceName, " Booking-App")
	fmt.Println("Get your tickets here to attend")

	fmt.Println("Available tickets:", ConferenceTickets, " and Remaining tickets:", RemainingTicekts)
}
