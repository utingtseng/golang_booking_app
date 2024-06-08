package main

import "fmt"

//define entry point
func main() {
	var conferenceName = "Go Conference"
	const conferenceTickets = 50
	// prevent negative numbers
	var remainingTickets uint = 50
	//slices are dynamic arrays without a predefined size
	var bookings []string

	fmt.Printf("Conference ticket is %T\n", conferenceName)

	fmt.Printf("Welcome to the %v booking system.\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and of which %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend!")

	var firstName string
	var lastName string
	var email string
	var userTickets uint

	//pointer (&) points to the memory address
	fmt.Println("Enter your first name:")
	fmt.Scan(&firstName)
	fmt.Println("Enter your last name:")
	fmt.Scan(&lastName)
	fmt.Println("Enter your email:")
	fmt.Scan(&email)
	fmt.Println("Enter number of tickets:")
	fmt.Scan(&userTickets)

	remainingTickets = remainingTickets - userTickets
	bookings = append(bookings, firstName+" "+lastName)

	fmt.Printf("Thank you %v %v! You have booked %v tickets. You will receive your confirmation tickets at %v shortly.\n", firstName, lastName, userTickets, email)
	fmt.Printf("We now have %v tickets left.\n", remainingTickets)
	fmt.Printf("These are all the bookings so far %v", bookings)

}
