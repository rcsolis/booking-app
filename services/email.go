package services

import (
	"fmt"
	"time"
)

/**
* function to send an email to the user
 */
func SendTicket(userTickets uint, name string, lastName string, email string, total float64, currency string) {
	time.Sleep(time.Second * 10)
	ticket := fmt.Sprintf("%v tickets for %v %v sending to %v for total of %v %v\n", userTickets, name, lastName, email, total, currency)
	fmt.Println("·············")
	fmt.Printf("Sending ticket to %v ",  ticket)
	fmt.Println("·············")
}