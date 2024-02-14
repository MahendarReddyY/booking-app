package main

import (
	"fmt"
	"time"
)

func sendingpacket() {
	time.Sleep(10 * time.Second)
	var tick = fmt.Sprintf("%v tickets of %v %v are send to this ", tickets, firstName, lastName)
	fmt.Println("########################")
	fmt.Printf("sending ticket to:%v \n %v\n email address %v\n", firstName, tick, email)
	fmt.Println("########################")
	fmt.Println("Thanks for ordering your tickets for events")
	wg.Done()
}
