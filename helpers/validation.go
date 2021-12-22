package helpers

/*
Imports
*/
import (
	"booking-app/globals"
	"errors"
	"fmt"
	"strings"
)

/**
* function to validate user data
 */
func ValidateUserData(userName string, lastName string, email string, userCity string, userTickets uint) (uint, string, error) {
    // Local variables 
    var isValidEmail bool = len(email) > 6 && strings.Contains(email, "@") && strings.Contains(email, ".")
    var isValidData bool = len(userName) >= 2 && len(lastName) >= 2
    var isValidUserTickets bool = userTickets > 0 && userTickets <= globals.RemainingTickets
    var priceIndex uint = 0
    var currency string = "USD"
    var err error = nil

    // Validate user data
    if !isValidData || !isValidEmail {
        fmt.Println("*** User data not valid.")
        err = errors.New("user data not valid")
    }else if !isValidUserTickets {  
        fmt.Println("*** Number of tickets not valid.")
        err = errors.New("number of tickets not valid")
    }else{
        // Validate user city
        switch userCity{
            case "New York", "New York City", "NYC":
                currency = "USD"
                priceIndex = 0
            case "Hong Kong", "Hong Kong City", "HK":
                currency = "HKD"
                priceIndex = 1
            case "London", "London City", "LONDON":
                currency = "GBP"
                priceIndex = 2
            case "Paris", "Paris City", "PARIS":
                currency = "EUR" 
                priceIndex = 3
            case "Mexico City", "Mexico", "MX":
                currency = "MXN"
                priceIndex = 4
            default :
                fmt.Println("** User city not valid.")
                err = errors.New("** User city not valid")
        }
    } 
    
    return priceIndex, currency, err
}