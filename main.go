package main

import (
	"fmt"
	"strings"
)

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

	for {
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

		isValidName := len(firstName) >= 2 && len(lastName) >= 2
		isValidEmail := strings.Contains(email, "@")
		isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets

		if isValidName && isValidEmail && isValidTicketNumber {
			remainingTickets = remainingTickets - userTickets
			bookings = append(bookings, firstName+" "+lastName)

			fmt.Printf("Thank you %v %v! You have booked %v tickets. You will receive your confirmation tickets at %v shortly.\n", firstName, lastName, userTickets, email)
			fmt.Printf("We now have %v tickets left.\n", remainingTickets)

			firstNames := []string{}
			for _, booking := range bookings {
				var names = strings.Fields(booking)
				firstNames = append(firstNames, names[0])
			}
			fmt.Printf("The first names of bookings are %v\n", firstNames)

			if remainingTickets == 0 {
				// end the programme
				fmt.Println("Our conference is booked out! Come back next year.")
				break
			}
		} else {
			fmt.Println("Input data is incorrect. Try again, please.")
			if !isValidName {
				fmt.Println("Both first and last name needs to be longer than 2 characters.")
			}
			if !isValidEmail {
				fmt.Println("A valid email needs to have an @")
			}
			if !isValidTicketNumber {
				fmt.Println("The number of tickets is incorrect.")
			}
			continue
		}

	}

}
