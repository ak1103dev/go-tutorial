package main

import "fmt"

func main() {
	var conferenceName = "Go Conference"
	const conferenceTickets = 50
	var remainingTickets uint = 50
	var bookings []string

	fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName is %T\n", conferenceTickets, remainingTickets, conferenceName)

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still avaliable.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	var firstName string
	var lastName string
	var userTickets uint

	fmt.Print("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Print("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Print("Enter number of tickets: ")
	fmt.Scan(&userTickets)

	remainingTickets = remainingTickets - userTickets
	bookings = append(bookings, firstName + " " + lastName)

	fmt.Printf("The whole slice: %v\n", bookings)
	fmt.Printf("The first value: %v\n", bookings[0])
	fmt.Printf("Slice Type: %T\n", bookings)
	fmt.Printf("Slice length: %v\n", len(bookings))

	fmt.Printf("User %v %v booked %v tickets.\n", firstName, lastName, userTickets)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

	fmt.Printf("These are all our bookins: %v\n", bookings)
}