package globals

/* Global variables */
const ConferenceTickets int = 50
var EventName string = "Go Conference"
var TicketPrice =[5]float64{100.00, 200.00, 300.00, 400.00, 500.00}
var RemainingTickets uint = uint(ConferenceTickets)
//var bookingsUser []map[string]string = make([]map[string]string, 0)
var BookingsUser = make([]UserData, 0)
/* User data type */
type UserData struct {
	Name string
	LastName string
	Email string
	City string
	NumberOfTickets uint
	Price float64
	Currency string
	Total float64
}
/**
 * function to create a new user
 */
func NewUser(name string, lastName string, email string, city string, numberOfTickets uint, price float64, currency string, total float64) UserData {
	return UserData{
		Name: name,
		LastName: lastName,
		Email: email,
		City: city,
		NumberOfTickets: numberOfTickets,
		Price: price,
		Currency: currency,
		Total: total,
	}
}
