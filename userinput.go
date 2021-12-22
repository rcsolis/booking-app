package main

/*
imports
*/
import (
	"booking-app/helpers"
	"fmt"
)

/**
* function to get user input
 */
func getUserData() (string, string, string, string, uint, uint, string, error) {
    // Local variables 
    var userName string
    var lastName string
    var email string
    var userCity string
    var userTickets uint
    var err error = nil
    // Get user input
    fmt.Println("Please enter your name: ")
    fmt.Scanln(&userName)

    fmt.Println("Please enter your last name: ")
    fmt.Scanln(&lastName)

    fmt.Println("Please enter your email address: ")
    fmt.Scanln(&email)

    fmt.Println("Please enter your city: ")
    fmt.Scanln(&userCity)

    fmt.Print("How many tickets would you like to buy? ")
    fmt.Scan(&userTickets)
    // validate user data
    priceIndex, currency, err := helpers.ValidateUserData(userName, lastName, email, userCity, userTickets)
    
    return userName, lastName, email, userCity, userTickets, priceIndex, currency, err
}