package main

import (
	"fmt"
	"sync"
	"time"
)

const conferenceTickets = 50

var conferenceName = "Go Conference"
var remainingTickets uint = 50

//bookings array's length can be 1,2,..... 49 or 50 but let's say we know it's 50 for sure....(means 1 person->1 ticket booking)
// var bookings=[50]string{} //an array-you have to write length of the array in go lang. and make sure it's true or you wont like the consequences...
// var bookings = []string{} //a slice
//var bookings = make([]map[string]string, 0) //a slice of maps(a map list), 0=initial size
//map supports only 1 data type. all keys have to be same data type and all values have to be same data type

var bookings = make([]UserData, 0) //a slice of UserData struct list

type UserData struct {
	//struct->structure, can hold mixed data types
	//type creates a new data type with the the name of you specify(custom data type)
	//design your own data type:)
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

// wait group- main thread should wait for the launched goroutines in wait group to finish
var wg = sync.WaitGroup{}

func main() {

	greetUsers()

	for {
		firstName, lastName, email, userTickets := getUserInput()
		isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidTicketNumber {

			bookTicket(userTickets, firstName, lastName, email)
			//"go" starts a new goroutine( is a green thread which needs little memory space-small stack,same hardware context-. So it is lightweight)
			//green thread(user level) is an abstraction of an actual(OS level-kernel-hardware-heavyweight) thread
			//now we have main thread and sendTicket() thread
			wg.Add(1) //adding the (1)launched goroutine(sendTicket() thread) to wait group
			go sendTicket(userTickets, firstName, lastName, email)

			firstNames := getFirstNames()
			fmt.Printf("these are all first names: %v\n", firstNames)

			if remainingTickets == 0 {

				fmt.Println("our conference is booked out.")
				break
			}
		} else {
			fmt.Println(".......................................................")
			if !isValidName {
				fmt.Printf("first name or last name you entered %v %v is too short!\n", firstName, lastName)
			}
			if !isValidEmail {
				fmt.Printf("email address you entered %v doesn't contain @ sign!\n", email)
			}
			if !isValidTicketNumber {
				fmt.Println("number of tickets you entered is invalid!")
			}
			fmt.Println("please try again")
			fmt.Println(".......................................................")
		}

	}
	wg.Wait()
}

func greetUsers() {
	fmt.Printf("\nwelcome to %v booking application\n", conferenceName)
	fmt.Printf("we have total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("get your tickets here to attend")
}

func getFirstNames() []string {
	firstNames := []string{}
	//blank identifier _ ,to ignore a variable you don't want to use
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Println("\nenter your first name: ")
	fmt.Scan(&firstName) //pointer & to its address

	fmt.Println("enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("enter your email: ")
	fmt.Scan(&email)

	fmt.Println("enter number of tickets you want: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets
	//create a UserData struct for a user
	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}
	bookings = append(bookings, userData)
	fmt.Printf("\nlist of bookings is %v\n", bookings)

	fmt.Printf("thank you %v %v for booking %v tickets. You will receive a conformation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(15 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println()
	fmt.Println("###############################################")
	fmt.Printf("sending ticket: %v  to email address %v \n", ticket, email)
	fmt.Println("#################################################")
	fmt.Println()
	wg.Done()
}
