package greetings

/*
Imports
*/
import (
	"booking-app/globals"
	"fmt"
)

/**
* function to greet users
 */
func GreetUsers(){
    fmt.Println("Welcome to", globals.EventName,"booking application")
    fmt.Println("Get your tickets here to attend.")
}

/**
* function to finish execution
*/
func FinishExecution() {
    // Local variables 
    var firstNames = []string{}
    // Get first names 
    for _, user := range globals.BookingsUser {
        firstNames = append(firstNames, user.Name)
    }

    fmt.Println("Thank you for booking with us. Your tickets are booked. ")
    fmt.Printf("Total attendes: %d Attendees %v \n", len(globals.BookingsUser), firstNames)
}