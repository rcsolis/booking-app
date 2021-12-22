package main

/*
Imports
*/
import (
	"booking-app/globals"
	"booking-app/greetings"
	"booking-app/services"
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

/**
* Main function
 */
func main(){
    // Greeting users  
    greetings.GreetUsers()
    
    for globals.RemainingTickets > 0 {
        var userName string
        var lastName string
        var email string
        var userCity string
        var userTickets uint

        var currency string = "USD"
        var priceIndex uint = 0

        // Get user data
        userName, lastName, email, userCity, userTickets, priceIndex, currency, e := getUserData()
    
        if e != nil {
            fmt.Println("*** Invalid data. Please try again.")
            continue 
        }
        
        // Book conference ticket
        bookConferenceTicket(userName, lastName, email, userCity, userTickets, priceIndex, currency)

        // Balance remaining tickets
        fmt.Println("We have total of", globals.ConferenceTickets,"tickets and", globals.RemainingTickets,"are still available")
    }
    // end execution
    greetings.FinishExecution()
}

/**
* function to book a conference ticket
*/
func bookConferenceTicket(userName string, lastName string, email string, userCity string, userTickets uint, priceIndex uint, currency string) {
    // Local variables
    var total float64 = float64(userTickets) * globals.TicketPrice[priceIndex]
    // Create a map for user
    var user = globals.NewUser(userName, lastName, email, userCity, userTickets, globals.TicketPrice[priceIndex], currency, total)
    // Add user to bookings
    globals.BookingsUser = append(globals.BookingsUser, user)
    // Add to wait group
    wg.Add(1)
    // Start execution of the send email function
    go func(user globals.UserData) {
        //call the send email function
        services.SendTicket(user.NumberOfTickets, user.Name, user.LastName, user.Email, user.Total, user.Currency)
        //signal the wait group that the goroutine is done
        wg.Done()
    }(user)

    // Update remaining tickets
    globals.RemainingTickets = globals.RemainingTickets - userTickets
    
    fmt.Println("Dear", userName, lastName, "your ticket for", userCity, "in", currency, "is booked. Total: ", total)

    // Wait for the send email function to finish
    wg.Wait()
}


